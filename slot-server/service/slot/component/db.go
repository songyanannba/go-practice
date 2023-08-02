package component

import (
	"github.com/samber/lo"
	enum "slot-server/enum"
	"slot-server/global"
	. "slot-server/model/business"
	"slot-server/service/slot/base"
	"slot-server/utils/helper"
)

// DbRawData 数据库原始数据
type DbRawData struct {
	Slot      *Slot
	Reel      []*SlotReelData // 滚轮
	PayTable  []*SlotPayTable // 赢钱组合
	Payline   []*SlotPayline  // 划线规格
	Symbol    []*SlotSymbol   // 图标
	Jackpot   []*Jackpot      // 奖池规则
	Event     []*SlotEvent    // 特殊事件
	Fake      []*SlotFake     // 特殊事件
	Debugs    []*DebugConfig  // 调试配置
	Templates []*SlotTemplate // 模板类型=>列号=>标签
}

// NewDbRawDataBySlotId 创建数据库原始数据集
func NewDbRawDataBySlotId(slotId uint) (rawData *DbRawData, err error) {
	rawData = &DbRawData{}
	err = global.GVA_DB.First(&rawData.Slot, "id = ?", slotId).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Find(&rawData.Reel, "slot_id = ?", rawData.Slot.ID).Error
	if err != nil {
		return
	}
	// payTable按倍数由高到低排序
	err = global.GVA_DB.Order("win_multiple desc").Find(&rawData.PayTable, "slot_id = ?", rawData.Slot.ID).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Find(&rawData.Payline, "no = ?", rawData.Slot.PaylineNo).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Find(&rawData.Symbol, "slot_id = ?", rawData.Slot.ID).Error
	if err != nil {
		return
	}
	rawData.Jackpot, err = GetJackpotListBySlot(rawData.Slot)
	if err != nil {
		return
	}
	err = global.GVA_DB.Find(&rawData.Event, "slot_id = ?", rawData.Slot.ID).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Find(&rawData.Fake, "slot_id = ?", rawData.Slot.ID).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Find(&rawData.Debugs, "slot_id = ? and start = 1", rawData.Slot.ID).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Find(&rawData.Templates, "slot_id = ?", rawData.Slot.ID).Error
	if err != nil {
		return
	}
	return
}

// NewSlotConfig 从原始数据创建slotConfig
func (d *DbRawData) NewSlotConfig() (*Config, *Config) {
	c := &Config{
		SlotId:   d.Slot.ID,
		Raise:    d.Slot.Raise,
		BetMap:   base.NewBetMap(d.Slot.BetNum),
		BuyFee:   d.Slot.BuyFreeSpin,
		BuyRes:   d.Slot.BuyReSpin,
		Status:   d.Slot.Status,
		tagMap:   map[string]*base.Tag{},
		tagIdMap: map[int]*base.Tag{},
		TopMul:   d.Slot.TopMul,
		Template: map[int]map[int][]*base.Tag{},
	}

	// 解析规格
	c.Row, c.Index = d.Payline[0].ParseSpec()

	// 解析坐标
	for _, payline := range d.Payline {
		if payline.Position != "" {
			c.Coords = append(c.Coords, ParseCoordinate(payline.Position))
		}
	}

	reel, reelDemo := helper.Apart(d.Reel, func(v *SlotReelData) bool {
		return v.Demo == enum.No
	})

	// 解析滚轮
	c.Reel = parseReelData(reel)

	// 解析图标
	for _, symbol := range d.Symbol {
		var include []string
		if symbol.IsWild == enum.Yes {
			include = symbol.ParseInclude()
		}
		if symbol.Multiple < 1 {
			symbol.Multiple = 1
		}
		tag := base.NewTag(int(symbol.ID), symbol.Name, float64(symbol.Multiple), symbol.IsSingleWin, include...)
		c.tagMap[symbol.Name] = tag
		c.tagIdMap[int(symbol.ID)] = tag
	}

	// 解析paytable
	for _, payTable := range d.PayTable {
		combine1, combine2 := payTable.ParseCombine()
		tags := c.GetTags(combine1...)
		if payTable.Type == enum.SlotPayTableType1Common {
			c.PayTableList = append(c.PayTableList, base.NewCommonPayTable(payTable.ID, tags, payTable.WinMultiple))
		} else {
			tags2 := c.GetTags(combine2...)
			c.PayTableList = append(c.PayTableList, base.NewAnyPayTable(payTable.ID, tags, payTable.CombineNum1, tags2, payTable.CombineNum2, payTable.WinMultiple))
		}
	}

	// 解析奖池
	for _, jackpot := range d.Jackpot {
		c.JackpotList = append(c.JackpotList, NewJackpotData(jackpot.ID, 0, 0, jackpot.End, jackpot.ParseCombine()))
	}
	c.place = make([]int, c.Index)
	c.freePlace = make([]int, c.Index)
	// 解析额外配置
	//xConfig, _ := GetXConfigCacheByName(enum.ConfigNameSlotDefaultTag + strconv.Itoa(int(c.SlotId)))
	//if xConfig.Status == enum.Yes {
	//	common, free, _ := strings.Cut(xConfig.Value, "&")
	//	c.place = ParseDefaultTag(common, c.Index)
	//	c.freePlace = ParseDefaultTag(free, c.Index)
	//}

	// 解析事件
	event, eventDemo := helper.Apart(d.Event, func(v *SlotEvent) bool {
		return v.Demo == enum.No
	})
	parseEventData(c, event)

	// 解析假数据
	c.Fakes = newFakes(d.Fake)
	c.Debugs = d.Debugs

	// 解析demo
	demo := *c
	// 解析滚轮
	demo.Reel = parseReelData(reelDemo)
	// 解析事件
	parseEventData(&demo, eventDemo)

	groupBy := lo.GroupBy(d.Templates, func(i *SlotTemplate) uint8 {
		return i.Type
	})
	for i, templates := range groupBy {
		c.Template[int(i)] = map[int][]*base.Tag{}
		for _, template := range templates {
			c.Template[int(i)][template.Column] = c.GetTagsLayout(template.Layout)
		}
	}
	return c, &demo
}

func parseReelData(reelData []*SlotReelData) []*Reel {
	var reels []*Reel
	// 解析滚轮
	reelMap := lo.GroupBy(reelData, func(i *SlotReelData) int {
		return i.Group
	})

	for i := 1; i <= len(reelMap); i++ {
		reels = append(reels, ParseReel(reelMap[i]))
	}
	return reels
}

func parseEventData(c *Config, eventData []*SlotEvent) {
	c.Event = base.NewEvent()
	switch c.SlotId {
	case 5:
		c.Event.Unit5NewEvent(eventData)
	case 6:
		c.Event.Unit6NewEvent(eventData)
	case 8:
		c.Event.Unit8NewEvent(eventData)
	default:
		for i, ev := range eventData {
			c.Event.Add(i+1, ev.Event1)
		}
	}
}
