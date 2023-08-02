package unit4

import (
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"math/rand"
	"slot-server/enum"
	"slot-server/service/slot/base"
	"slot-server/service/slot/component"
	"slot-server/utils/helper"
)

// Machine 划线 + 特殊免费玩
type Machine struct {
	Spin *component.Spin
}

func NewMachine(spin *component.Spin) *Machine {
	return &Machine{Spin: spin}
}

func (m *Machine) GetSpin() *component.Spin {
	return m.Spin
}

func (m *Machine) GetSpins() []*component.Spin {
	return []*component.Spin{m.Spin}
}

func (m *Machine) Exec() error {
	if len(m.Spin.IsSetResult) > 0 {
		m.Spin.InitDataList = m.Spin.IsSetResult
	} else if len(m.Spin.InitDataList) == 0 {
		m.GetInitData()
	}
	if m.Spin.IsMustFree {
		m.MustFree()
		m.Spin.BuyFreeCoin = decimal.NewFromInt(int64(m.Spin.Bet)).Mul(decimal.NewFromFloat(m.Spin.Config.BuyFee)).IntPart()
	}
	if m.Spin.IsMustRes {
		m.MustRes()
		c, _ := component.GetSlotConfig(4, false)
		m.Spin.BuyReCoin = decimal.NewFromInt(int64(m.Spin.Bet)).Mul(decimal.NewFromFloat(c.BuyRes)).IntPart()
	}

	m.LinkFixedValue() //设置标签倍率
	m.GetResData()
	m.SumGain()
	return nil
}

func (m *Machine) MustFree() {
	// 根据滚轮配置初始化数据
	for i := 0; i < m.Spin.Config.Index; i++ {
		reel := m.Spin.Config.Reel[i]
		reelData := reel.GetReelData(1, 1)
		start := 0
		switch i {
		case 0, 2, 4:
			var keys []int
			for key, name := range reelData.Data {
				if name == "scatter" {
					keys = append(keys, key)
				}
			}
			start = keys[rand.Intn(len(keys))]
		default:
			start = rand.Intn(len(reelData.Data))
		}
		offsets := []int{-1, 0, -2}
		offset := offsets[rand.Intn(len(offsets))]
		reelTag := helper.SliceByRange(reelData.Data, start+offset, 3)
		m.Spin.AddInitData(i, reelTag)
	}
	verification := make(map[string]bool)
	for _, tag := range m.Spin.InitDataList[0] {
		verification[tag.Name] = true
	}
	for _, tag := range m.Spin.InitDataList[0] {
		for i, t := range m.Spin.InitDataList[2] {
			if tag.Name == t.Name && tag.Name != "scatter" {
				for _, tag := range m.Spin.Config.GetAllTag() {
					if verification[tag.Name] {
						continue
					} else {
						m.Spin.InitDataList[2][i] = &tag
						break
					}
				}
			}
		}
	}

}

type ResData struct {
	Names     []string
	ReelIndex int
	ColIndex2 int
}

func (m *Machine) MustRes() {
	// 根据滚轮配置初始化数据
	var taglist = make([][]ResData, m.Spin.Config.Index)
	for i := 0; i < m.Spin.Config.Index; i++ {
		reel := m.Spin.Config.Reel[i]
		reelData := reel.GetReelData(1, 1)
		ResDatas := make([]ResData, 0)
		for a := 0; a < len(reelData.Data)-2; a++ {
			names := []string{
				reelData.Data[a],
				reelData.Data[a+1],
				reelData.Data[a+2],
			}
			ResDatas = append(ResDatas, ResData{
				Names:     names,
				ReelIndex: i,
				ColIndex2: a,
			})
		}
		ResDatass := lo.Filter(ResDatas, func(item ResData, i int) bool {
			return item.Names[0] != "link_collect" && item.Names[1] != "link_collect" && item.Names[2] != "link_collect"
		})
		taglist[i] = ResDatass
	}
	options := []int{0, 1, 2}
	choices := []int{0, 2}
	choice := choices[rand.Intn(len(choices))]
	count := 0
	for i := 0; i < m.Spin.Config.Index; i++ {
		switch i {
		case choice:
			lists := lo.Filter(taglist[i], func(item ResData, i int) bool {
				return item.Names[0] == "link_coin" && item.Names[1] == "link_coin" && item.Names[2] == "link_coin"
			})
			reelTag := lists[rand.Intn(len(lists))]
			m.Spin.AddInitData(i, reelTag.Names)
		case 4:
			lists := lo.Filter(taglist[i], func(item ResData, i int) bool {
				coins := lo.Filter(item.Names, func(item string, i int) bool {
					return item == "link_coin"
				})
				return len(coins) == 3-count
			})
			reelTag := lists[rand.Intn(len(lists))]
			m.Spin.AddInitData(i, reelTag.Names)
		default:
			option := options[rand.Intn(len(options))]
			comon := helper.If(count+option > 3, 3-count, option)
			count += comon
			lists := lo.Filter(taglist[i], func(item ResData, i int) bool {
				coins := lo.Filter(item.Names, func(item string, i int) bool {
					return item == "link_coin"
				})
				return len(coins) == comon
			})
			reelTag := lists[rand.Intn(len(lists))]
			m.Spin.AddInitData(i, reelTag.Names)
		}
	}
}
func (m *Machine) ReDebugConversion() {
	var tagList [][]*base.Tag
	for _, tags := range m.Spin.InitDataList {
		for _, tag := range tags {
			var newTag = *tag
			newTag.X = tag.Y*5 + tag.X
			newTag.Y = 0
			tagList = append(tagList, []*base.Tag{
				&newTag,
			})
		}
	}
}

func (m *Machine) GetInitData() {
	if !m.Spin.IsReSpin {
		if m.Spin.IsFree {
			m.Spin.GetInitDataByReel(enum.SlotReelType2FS, 1)
		} else {
			m.Spin.GetInitDataByReel(enum.SlotReelType1Normal, 1)
		}
	} else {
		m.GetResInitData()
	}
	return
}

// GetResInitData 获取重转的初始数据
//
//	@Description:  获取重转的初始数据
//	@receiver m *Machine
func (m *Machine) GetResInitData() {
	var tagList [][]*base.Tag

	//将锁定标签按照X坐标分组
	lockTags := lo.GroupBy(m.Spin.TagsLock, func(item *base.Tag) int {
		return item.X
	})

	//设置转换成15列的初始数据
	for i := 0; i < 15; i++ {
		//如果有锁定标签则取锁定标签
		if lockTags[i] != nil {
			tagList = append(tagList, lockTags[i])
		} else {
			//从权重列表中获取标签
			var tag = *m.Spin.Config.GetTag("null")
			tag.X = i
			tagList = append(tagList, []*base.Tag{
				&tag,
			})
		}
	}
	m.Spin.InitDataList = tagList
}

// GetResData 根据坐标获取payline结果
//
//	@Description: 不是重转的情况下获取payline结果
//	@receiver m
func (m *Machine) GetResData() {
	if !m.Spin.IsReSpin {
		m.Spin.GetPaylineByCoords()
	}
}

// LinkFixedValue 设置link_coin的倍数 为link_collect汇总倍数
//
//	@Description: 根据权重设置本次转出的Coin标签倍率, 设置link_collect标签的汇总倍率
//	@receiver m
func (m *Machine) LinkFixedValue() {
	s := m.Spin
	//设置本次link_coin的倍数
	for _, tags := range s.InitDataList {
		for _, tag := range tags {
			if tag.Name == "link_coin" && !tag.ISLock {
				//获取link_coin标签的权重
				tag.Multiple = float64(m.Spin.Config.Event.M[1].(*base.ChangeTableEvent).Fetch())
			}
		}
	}
	//获取link_coin|link_collect标签
	likeList := base.GetSpecialTags(s.InitDataList, "link_coin", "link_collect")

	//获取除了本次的合计标签外的所有标签倍数总和
	coinsMulSum := lo.SumBy(likeList, func(item *base.Tag) float64 {
		if !item.ISLock && item.Name == "link_collect" {
			return 0
		}
		return item.Multiple
	})

	//记录link_collect标签的汇总倍数,设置link_collect标签的汇总倍数
	multiple := float64(0)
	for _, tags := range m.Spin.InitDataList {
		for _, tag := range tags {
			if tag.Name == "link_collect" && !tag.ISLock {
				//tag.Multiple = coinsMulSum + multiple
				//multiple += tag.Multiple
				tag.Multiple, _ = decimal.NewFromFloat(coinsMulSum).Add(decimal.NewFromFloat(multiple)).Float64()
				multiple, _ = decimal.NewFromFloat(tag.Multiple).Add(decimal.NewFromFloat(multiple)).Float64()
			}
		}
	}
}

func (m *Machine) SumGain() {
	m.GameAssembly()
	m.Spin.Gain = 0
	if m.Spin.Jackpot != nil {
		m.Spin.Gain = int(decimal.NewFromFloat(m.Spin.Jackpot.End).Mul(decimal.NewFromInt(int64(m.Spin.Bet))).IntPart())
		return
	}

	if !m.Spin.IsReSpin {
		if m.Spin.IsFree {
			m.Spin.FreeNum--
		}
		//计算PayTables的倍数
		if len(m.Spin.PayTables) > 0 {
			var mulArr []float64
			for _, table := range m.Spin.PayTables {
				mulArr = append(mulArr, table.Multiple)
			}
			mul := helper.FloatSum(mulArr...)
			m.Spin.Gain = int(decimal.NewFromInt(int64(m.Spin.Bet)).Mul(decimal.NewFromFloat(mul)).IntPart())
		}
		//计算Wild的倍数
		if len(m.Spin.SingleList) > 0 {
			if m.Spin.Gain == 0 {
				m.Spin.Gain = m.Spin.Bet
			}
		}
		for _, wild := range m.Spin.WildList {
			if wild.IsSingle {
				//m.Spin.Gain += m.Spin.Bet * wild.Multiple
				betMulMultiple := int(decimal.NewFromFloat(wild.Multiple).Mul(decimal.NewFromFloat(float64(m.Spin.Bet))).IntPart())
				m.Spin.Gain += betMulMultiple
			}
		}

	} else {
		m.Spin.ResNum--
		//最后一次重转,计算重转的倍数
		if m.Spin.ResNum == 0 {
			likeList := base.GetSpecialTags(m.Spin.InitDataList, "link_coin", "link_collect")
			MuSum := lo.SumBy(likeList, func(item *base.Tag) float64 {
				return item.Multiple
			})
			m.Spin.WinLines = append(m.Spin.WinLines, &component.Line{
				Tags: likeList,
				Win:  helper.IntMulFloatToInt(m.Spin.Bet, MuSum),
			})
			//m.Spin.Gain = MuSum * m.Spin.Bet
			m.Spin.Gain = int(decimal.NewFromFloat(MuSum).Mul(decimal.NewFromFloat(float64(m.Spin.Bet))).IntPart())
		}
	}
	m.Spin.WinLineMergeSame()
}

// GameAssembly 游戏组装
//
//	@Description: 判断是否有jackpot, 设置jackpot倍数, 设置payTable, 设置FreeSpin次数,设置ReSpin次数
//	@receiver m
func (m *Machine) GameAssembly() {
	s := m.Spin

	//如果ink_coin|link_collect标签数量大于等于15 设置jackpot倍数
	likeList := base.GetSpecialTags(s.InitDataList, "link_coin", "link_collect")
	if len(likeList) >= 15 {
		s.Jackpot = s.Config.JackpotList[0]
	}
	if !s.IsReSpin {
		// 每条划线可能匹配到多个payTable 取最大的payTable
		for _, tags := range s.ResDataList {
			for _, table := range s.Config.PayTableList {
				if ok, res := table.Match(tags); ok {
					s.PayTables = append(s.PayTables, res)
					break
				}
			}
		}

		//如果scatter标签数量大于等于3 则进入免费游戏
		scatters := lo.Filter(base.GetSpecialTags(s.InitDataList, "scatter"), func(item *base.Tag, i int) bool {
			if item.Name == "scatter" {
				return true
			}
			return false
		})
		if len(scatters) >= 3 {
			s.FreeSpinParams.Count++
		}
		s.FreeSpinParams.FreeNum += s.FreeSpinParams.Count * 10
		s.FreeNum += s.FreeSpinParams.FreeNum

		//如果link_coin|link_collect标签数量大于等于6 则进入ReSpin游戏
		if len(likeList) >= 6 {
			s.FreeSpinParams.ReNum += 6
			s.ResNum += s.FreeSpinParams.ReNum
		}

	} else {
		//出现plus_spin标签则增加重转次数
		plusSpins := base.GetSpecialTags(s.InitDataList, "plus_spin")
		s.FreeSpinParams.ReNum += len(plusSpins)
		s.ResNum += s.FreeSpinParams.ReNum
	}
}
