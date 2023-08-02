package component

import (
	"fmt"
	"github.com/samber/lo"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/service/slot/base"
	"slot-server/service/slot/eliminate"
	"slot-server/service/slot/template"
	"slot-server/utils/helper"
	"strings"
)

// Coordinate 坐标
type Coordinate struct {
	X int
	Y int
}

type Line struct {
	Name string
	Tags []*base.Tag
	Win  int64
}

func (l Line) String() string {
	return fmt.Sprintf("%s %v %d", l.Name, l.Tags, l.Win)
}

type Spin struct {
	*Options

	Config       *Config       `json:"-"`
	InitDataList [][]*base.Tag `json:"-"` // 初始数据
	ResDataList  [][]*base.Tag // 结果数据
	WildList     []*base.Tag   // 百搭数据
	SingleList   []*base.Tag   // 单出数据
	WinLines     []*Line       // 赢钱线

	Jackpot   *JackpotData     // 最终奖池
	PayTables []*base.PayTable // 最终payTable
	Bet       int              // 压注
	Gain      int              // 赢钱

	FreeSpinParams FreeSpinParams // 免费游戏参数

	Table    *eliminate.Table
	Id       int   // spinId
	ParentId int   // 父级ID
	PSpin    *Spin // 父级spin

	Which    int                // 选择哪个轮子
	SpinInfo *template.SpinInfo // spinInfo
}

func (s *Spin) String() string {
	res := ""
	for _, v := range lo.Slice(s.ResDataList, 0, 100) {
		res += fmt.Sprintf("%v", v)
		res += "\n"
	}
	if len(s.ResDataList) > 100 {
		res += "..."
	}

	text := fmt.Sprintf("resData: \n%v \n"+
		"winLineList: %v \n"+
		"wildList: %v \n"+
		"singleList: %v \n"+
		"jackpot: %v \n"+
		"bet: %v \n"+
		"gain: %v \n"+
		"freeSpin: %+v \n"+
		"options: %+v ", res, s.WinLines, s.WildList, s.SingleList, s.Jackpot, s.Bet, s.Gain, s.FreeSpinParams, s.Options)
	return text
}

func (s *Spin) Type() int {
	if s.IsReSpin {
		return enum.SlotSpinType3Respin
	} else if s.IsFree {
		return enum.SlotSpinType2Fs
	}
	return enum.SlotSpinType1Normal
}

func (s *Spin) GateData() string {
	str := s.String()
	//替换[]符号变成]/n[
	str = strings.Replace(str, "] [", "]\n[", -1)
	return str
}

func NewSpin(slotId uint, amount int, options ...Option) (*Spin, error) {
	s := &Spin{
		Options: &Options{},
		Bet:     amount,
		Gain:    0,
	}
	c, err := GetSlotConfig(slotId, s.Demo)
	if err != nil {
		return s, err
	}
	s.Config = c
	s.setOption(options...)
	return s, nil
}

func (s *Spin) setOption(options ...Option) {
	s.Options.Spin = s
	for _, option := range options {
		option(s.Options)
	}
}

func (s *Spin) AddInitData(x int, arr []string) {
	var tags []*base.Tag
	for y, v := range arr {
		tag := *s.Config.GetTag(v)
		tags = append(tags, &tag)
		tag.X = x
		tag.Y = y
	}
	s.InitDataList = append(s.InitDataList, tags)
}

// JackpotMatch 请保证Multiple从高到低排序 以优先匹配最高金额
func (s *Spin) JackpotMatch() *JackpotData {
	for _, tags := range s.ResDataList {
		for _, jackpot := range s.Config.JackpotList {
			if line, ok := jackpot.Match(tags); ok {
				s.Jackpot = jackpot
				// 计算赢钱
				line.Win = helper.IntMulFloatToInt(s.Bet, s.jackpotSum(jackpot))
				s.WinLines = append(s.WinLines, line)
				return jackpot
			}
		}
	}
	return nil
}

// PayTableOnceMatch 请保证PayTableList按Multiple从高到低排序 以优先匹配最高金额
func (s *Spin) PayTableOnceMatch() *base.PayTable {
	for _, tags := range s.ResDataList {
		for _, table := range s.Config.PayTableList {
			if ok, newTable := table.Match(tags); ok {
				line := &Line{
					Tags: newTable.Tags,
					Win:  helper.IntMulFloatToInt(s.Bet, newTable.Multiple),
				}
				s.WinLines = append(s.WinLines, line)

				s.PayTables = append(s.PayTables, newTable)
				return newTable
			}
		}
	}
	return nil
}

// PayTableMatch 匹配所有payTable
func (s *Spin) PayTableMatch() {
	for _, tags := range s.ResDataList {
		for _, table := range s.Config.PayTableList {
			if ok, newTable := table.Match(tags); ok {
				s.WinLines = append(s.WinLines, &Line{
					Tags: newTable.Tags,
					Win:  helper.IntMulFloatToInt(s.Bet, newTable.Multiple),
				})

				s.PayTables = append(s.PayTables, newTable)
				break
			}
		}
	}
}

func ParseDefaultTag(v string, index int) []int {
	res := make([]int, index)
	if v == "" {
		return res
	}
	arr := helper.SplitInt[int](v, ",")
	for k, vv := range arr {
		res[k] = vv
		if k >= index-1 {
			break
		}
	}
	return res
}

// SetDebugInitData 设置调试初始数据
func (s *Spin) SetDebugInitData(userId uint) {
	debugType := uint8(1)
	playType := uint8(1)
	if s.IsTest {
		debugType = 2
	}
	if s.IsReSpin {
		playType = 2
	}
	if s.IsFree {
		playType = 3
	}
	debugs := lo.Filter(s.Config.Debugs, func(item *business.DebugConfig, index int) bool {
		return item.DebugType == debugType && item.PalyType == playType && (item.UserId == int(userId) || item.UserId == 0)
	})
	if len(debugs) == 0 {
		return
	}
	userDebugs := lo.Filter(debugs, func(item *business.DebugConfig, index int) bool {
		return item.UserId == int(userId)
	})
	if len(userDebugs) > 0 {
		s.JsonToTags(userDebugs[0].ResultData)
	} else {
		s.JsonToTags(debugs[0].ResultData)
	}
}

func (s *Spin) JsonToTags(jsonStr string) {
	s.InitDataList = [][]*base.Tag{}
	var arr [][]string
	if err := global.Json.Unmarshal([]byte(jsonStr), &arr); err != nil {
		panic(err)
	}
	arr = helper.ArrVertical(arr)
	for k, v := range arr {
		s.AddInitData(k, v)
	}
}

func (s *Spin) GetDebugData() string {
	debugType := uint8(1)
	playType := uint8(1)
	if s.IsTest {
		debugType = 2
	}
	if s.IsReSpin {
		playType = 2
	}
	if s.IsFree {
		playType = 3
	}
	debugs := lo.Filter(s.Config.Debugs, func(item *business.DebugConfig, index int) bool {
		return item.DebugType == debugType && item.PalyType == playType
	})
	if len(debugs) == 0 {
		return ""
	}
	return debugs[0].ResultData
}
