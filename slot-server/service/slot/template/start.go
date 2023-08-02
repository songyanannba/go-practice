package template

import (
	"fmt"
	"github.com/samber/lo"
	"slot-server/service/slot/base"
	"slot-server/service/slot/template/flow"
	"slot-server/utils/helper"
	"sort"
)

type SpinInfo struct {
	config     *Config         // 游戏配置
	Display    [][]*base.Tag   // 游戏展示
	initWindow [][]*base.Tag   // 初始化展示
	SpinFlow   []flow.SpinFlow // 游戏流程
	startValue int             // 初始取值,从模版里面取第几行

	templateRowMap map[int]int // 模版行数映射

	//第八台游戏需要数据
	Scatter    *flow.WinLine        // scatter数量
	Multiplier []*base.Tag          // multiplier数量
	payTable   map[string][]*NumMul // 划线对应赢钱倍率

}

func (s *SpinInfo) GetInfo() string {
	str := fmt.Sprintf("初始取值:%d\n Scatter:%d个%g倍\n翻倍标签:", s.startValue, len(s.Scatter.Tags), s.Scatter.Mul)
	for _, i2 := range s.Multiplier {
		str += fmt.Sprintf("%g,", i2.Multiple)
	}
	str += "\n"
	return str
}

func (s *SpinInfo) GetStartValue() int {
	return s.startValue
}

func NewGameInfo(config *Config) *SpinInfo {
	game := &SpinInfo{
		config:         config,
		startValue:     helper.RandInt(len(config.Template)),
		templateRowMap: make(map[int]int),
		Scatter: &flow.WinLine{
			Tags: make([]*base.Tag, 0),
			Mul:  0,
		},
	}
	//templateRowMap 用于每一列的X坐标从那里取值（列上随机一个点,既是x的初始取值,每列上的x值也不一样）
	for y := 0; y < config.Col; y++ {
		game.templateRowMap[y] = helper.RandInt(len(config.Template[y]))
	}

	//坐标初始化 表里面的数据（标签）为空
	game.Display = helper.NewTable(config.Col, config.Row, func(x, y int) *base.Tag {
		return &base.Tag{X: x, Y: y, Name: ""}
	})
	game.SetAllLocation()  //重置坐标
	game.FillInitDisplay() //标签初始化，填入初始数据
	game.SetPayTable()     //结构化 赢钱标签对应的赔率和个数
	game.initWindow = helper.CopyListArr(game.Display)
	return game
}

type NumMul struct {
	Num int
	Mul float64
}

func (s *SpinInfo) Copy() *SpinInfo {
	newSpinInfo := *s
	return &newSpinInfo
}

func (s *SpinInfo) SetPayTable() {
	s.payTable = make(map[string][]*NumMul)

	payTableMap := lo.GroupBy(s.config.PayTableList, func(item *base.PayTable) string {
		return item.Tags[0].Name
	})
	for name, tables := range payTableMap {
		sort.Slice(tables, func(i, j int) bool {
			return tables[i].Num <= tables[j].Num
		})
		s.payTable[name] = lo.FilterMap(tables, func(item *base.PayTable, i int) (*NumMul, bool) {
			return &NumMul{
				Num: len(item.Tags),
				Mul: item.Multiple,
			}, true
		})
	}
}
