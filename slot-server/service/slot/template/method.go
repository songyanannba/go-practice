package template

import (
	"fmt"
	"github.com/samber/lo"
	"slot-server/service/slot/base"
	"slot-server/service/slot/template/flow"
	"slot-server/utils/helper"
	"strconv"
)

// GetTagMapByName 获取当前显示窗口所有Tag名称的Map
func (s *SpinInfo) GetTagMapByName() map[string][]*base.Tag {
	tags := s.GetAllTags()
	return lo.GroupBy(tags, func(item *base.Tag) string {
		return item.Name
	})
}

// GetEmptyTags 获取窗口空白标签
func (s *SpinInfo) GetEmptyTags() []*base.Tag {
	return lo.Filter(s.GetAllTags(), func(tag *base.Tag, i int) bool {
		return tag.Name == ""
	})
}

// SetAllLocation 重新设置所有标签的位置
func (s *SpinInfo) SetAllLocation() {
	for x, tags := range s.Display {
		for y, tag := range tags {
			tag.X = x
			tag.Y = y
		}
	}
}

func (s *SpinInfo) GetWin(bet int) int {
	win := 0
	spinMul := lo.SumBy(s.SpinFlow, func(item flow.SpinFlow) float64 {
		return item.SumMul
	})
	if spinMul > 0 {
		win += int(spinMul * float64(bet))
	}

	if s.Scatter != nil && s.Scatter.Mul > 0 {
		win += int(s.Scatter.Mul * float64(bet))
	}

	sumMul := lo.SumBy(s.Multiplier, func(item *base.Tag) float64 {
		return item.Multiple
	})
	if sumMul > 0 {
		win = int(float64(win) * sumMul)
	}
	return win
}

func (s *SpinInfo) PrintTable(str string) string {
	str += ":\n"
	for _, row := range s.Display {
		for _, col := range row {
			str += fmt.Sprintf("%s\t", strconv.Itoa(col.X)+":"+strconv.Itoa(col.Y)+" "+helper.If(col.Name == "", "🀆", col.Name))
		}
		str += "\r\n"
	}
	//fmt.Println(str)
	return str + "\r\n"
}
