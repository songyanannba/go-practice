package component

import (
	"github.com/shopspring/decimal"
	"slot-server/service/slot/base"
)

func (s *Spin) jackpotSum(jackpot *JackpotData) float64 {
	// jackpot值改为固定 只需取最终值即可
	s.Gain = int(decimal.NewFromInt(int64(s.Bet)).Mul(decimal.NewFromFloat(jackpot.End)).IntPart())
	return jackpot.End
}

type WinLineMerge struct {
	Name   string
	TagMap map[*base.Tag]struct{}
	Win    int64
}

func NewWinLineMerge(name string) *WinLineMerge {
	return &WinLineMerge{
		Name:   name,
		TagMap: map[*base.Tag]struct{}{},
	}
}

// WinLineMergeSame 将winLine中所有相同名称的标签组合并
func (s *Spin) WinLineMergeSame() {
	var (
		lineMap = map[string]*WinLineMerge{}
		newLine []*Line
	)
	// 合并相同name的线
	for _, line := range s.WinLines {
		// 获取首个不是wild的tag
		name := base.GetTagsNameByFunc(line.Tags, func(tag *base.Tag) bool {
			return !tag.IsWild
		})
		line.Name = name
		// 使用这个tag的name做为key
		merge, ok := lineMap[name]
		if !ok {
			lineMap[name] = NewWinLineMerge(name)
			merge = lineMap[name]
		}
		for _, tag := range line.Tags {
			merge.TagMap[tag] = struct{}{}
		}
		merge.Win += line.Win
		lineMap[name] = merge
	}
	// 重新生成winLine
	for name, merge := range lineMap {
		var tags []*base.Tag
		for tag := range merge.TagMap {
			tags = append(tags, tag)
		}
		newLine = append(newLine, &Line{
			Name: name,
			Tags: tags,
			Win:  merge.Win,
		})
	}
	s.WinLines = newLine
}
