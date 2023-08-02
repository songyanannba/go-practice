package eliminate

import (
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"slot-server/service/slot/base"
	"slot-server/utils/helper"
	"sort"
)

// WinMatchList 赢钱划线匹配
func (t *Table) WinMatchList(tagList [][]*base.Tag, isFree bool) ([]*base.Eliminate, float64) {
	var eliminates []*base.Eliminate
	sumMul := float64(0)
	for _, tags := range tagList {
		//mul := t.WinMatch(tags)
		//获取匹配画线的赔率
		mul := t.MatchWinMul(tags)
		doubleTags := t.DoubleProc(tags, isFree)

		for _, tag := range doubleTags {
			if tag.Multiple > 0 {
				mul = mul.Mul(decimal.NewFromFloat(tag.Multiple))
			}
		}
		//wild倍率
		for _, tag := range tags {
			if tag.IsWild {
				mul = mul.Mul(decimal.NewFromFloat(tag.Multiple))
			}
		}
		sumMul, _ = decimal.NewFromFloat(sumMul).Add(mul).Float64()
		elMul, _ := mul.Float64()

		eliminate := &base.Eliminate{
			RemoveList: doubleTags,
			Mul:        elMul,
		}

		eliminates = append(eliminates, eliminate)
	}
	return eliminates, sumMul
}

// DoubleProc FreeSpin随机翻倍标签
func (t *Table) DoubleProc(tags []*base.Tag, isFree bool) []*base.Tag {
	rTags := make([]*base.Tag, len(tags))
	for i, i2 := range tags {
		tag := i2.Copy()
		tag.Multiple = 0
		rTags[i] = tag
	}
	if !isFree {
		return rTags
	} else {
		mulNum := t.Target.MulNumEvent.Fetch()
		if mulNum >= len(rTags) {
			for c, _ := range rTags {
				rTags[c].Multiple = 2
			}
		} else {
			v := make(map[int]int)
			count := 0
			for count < mulNum {
				index := helper.RandInt(len(rTags))
				if _, ok := v[index]; !ok {
					v[index] = index
					count++
				}
			}
			for c, _ := range rTags {
				if _, ok := v[c]; ok {
					rTags[c].Multiple = 2
				} else {
					rTags[c].Multiple = 0
				}
			}
		}
	}
	return rTags
}

// WinMatch 赢钱划线匹配
func (t *Table) WinMatch(tags []*base.Tag) decimal.Decimal {
	mul := decimal.NewFromInt(0)
	payTables := lo.Filter(t.PayTableList, func(item *base.PayTable, index int) bool {
		return item.Tags[0].Name == tags[0].Name
	})
	if len(payTables) == 0 {
		return mul
	}

	sort.Slice(payTables, func(i, j int) bool {
		return len(payTables[i].Tags) > len(payTables[j].Tags)
	})

	payTable := lo.Filter(payTables, func(item *base.PayTable, index int) bool {
		return len(item.Tags) == len(tags)
	})

	if len(payTable) > 0 {
		mul = mul.Add(decimal.NewFromFloat(payTable[0].Multiple))
	} else if len(tags) > len(payTables[0].Tags) {
		mul = mul.Add(decimal.NewFromFloat(payTables[0].Multiple))
	}

	return mul
}

func (t *Table) MatchWinMul(tags []*base.Tag) decimal.Decimal {
	mul := decimal.NewFromInt(0)
	payTableListMaps := t.PayTableListMaps
	if payTableListMaps == nil || len(payTableListMaps) == 0 {
		return mul
	}
	payTableLists, ok := payTableListMaps[tags[0].Name]
	if !ok {
		return mul
	}
	// 和每种画线的标签数量进行比较 数量相等 即匹配成功
	if payTable, ok := payTableLists[len(tags)]; ok {
		mul = mul.Add(decimal.NewFromFloat(payTable.Multiple))
	} else {
		var multiple float64
		// 如果没有相等的 又有比要消除标签小的情况
		// 向下兼容 匹配最近一个
		for i := len(tags); i >= 5; i-- {
			if pv, okk := payTableLists[i]; okk {
				multiple = pv.Multiple
				break
			}
		}

		mul = mul.Add(decimal.NewFromFloat(multiple))
	}
	return mul
}

func (t *Table) WinMatchName(name string, l int) decimal.Decimal {
	mul := decimal.NewFromInt(0)
	payTables := lo.Filter(t.PayTableList, func(item *base.PayTable, index int) bool {
		return item.Tags[0].Name == name
	})
	if len(payTables) == 0 {
		return mul
	}

	sort.Slice(payTables, func(i, j int) bool {
		return len(payTables[i].Tags) > len(payTables[j].Tags)
	})

	payTable := lo.Filter(payTables, func(item *base.PayTable, index int) bool {
		return len(item.Tags) == l
	})

	if len(payTable) > 0 {
		mul = mul.Add(decimal.NewFromFloat(payTable[0].Multiple))
	} else if l > len(payTables[0].Tags) {
		mul = mul.Add(decimal.NewFromFloat(payTables[0].Multiple))
	}

	return mul
}

// WinMatchSlot8  赢钱划线匹配
func (t *Table) WinMatchSlot8(tagList [][]*base.Tag, isFree bool) ([]*base.Eliminate, float64) {
	var eliminates []*base.Eliminate
	sumMul := float64(0)
	for _, tags := range tagList {
		mul := t.MatchWinMul(tags)

		//wild倍率
		for _, tag := range tags {
			if tag.IsWild {
				mul = mul.Mul(decimal.NewFromFloat(tag.Multiple))
			}
		}

		sumMul, _ = decimal.NewFromFloat(sumMul).Add(mul).Float64()
		elMul, _ := mul.Float64()

		eliminate := &base.Eliminate{
			RemoveList: helper.CopyList(tags),
			Mul:        elMul,
		}

		eliminates = append(eliminates, eliminate)
	}
	return eliminates, sumMul
}
