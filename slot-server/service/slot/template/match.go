package template

import (
	"github.com/samber/lo"
	"slot-server/service/slot/base"
	"slot-server/service/slot/template/flow"
	"slot-server/utils/helper"
)

//第八台游戏逻辑

// GetWinLine 获取一条赢钱划线
func (s *SpinInfo) GetWinLine(tags []*base.Tag) *flow.WinLine {
	if len(tags) == 0 {
		return &flow.WinLine{
			Tags: make([]*base.Tag, 0),
			Mul:  0,
		}
	}
	return &flow.WinLine{
		Tags: helper.CopyList(tags),
		Mul:  s.MatchTagsWin(tags),
	}
}

// GetWinLines 获取多条赢钱划线
func (s *SpinInfo) GetWinLines(tagList [][]*base.Tag) []*flow.WinLine {
	return lo.FilterMap(tagList, func(tags []*base.Tag, i int) (*flow.WinLine, bool) {
		line := s.GetWinLine(tags)
		if line == nil {
			return nil, false
		}
		return line, true
	})
}

// MatchTagsWin 获取一条划线的倍率
func (s *SpinInfo) MatchTagsWin(tags []*base.Tag) (mul float64) {
	nowPayTable := s.payTable[tags[0].Name]
	mul = 0.0
	for _, table := range nowPayTable {
		if len(tags) >= table.Num {
			mul = table.Mul
		} else {
			return mul
		}
	}
	return mul
}

// MatchTagListWin 获取多条划线的倍率
func (s *SpinInfo) MatchTagListWin(list [][]*base.Tag) (mul float64) {
	mul = 0.0
	for _, tags := range list {
		mul = helper.Float64Add(mul, s.MatchTagsWin(tags))
	}
	return mul
}
