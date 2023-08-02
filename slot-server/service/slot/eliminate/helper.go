package eliminate

import (
	"slot-server/service/slot/base"
	"slot-server/utils/helper"
)

// GetCombination
//
//	@Description: 获取组合
//	@param target 目标值
//	@return []int 组合
func GetCombination(target int) []int {
	if target < 10 {
		return []int{target}
	}
	var result []int
	sum := 0

	for sum != target {
		value := helper.RandInt(target-sum-4) + 5
		if target-sum-value < 5 {
			result = append(result, target-sum)
			break
		} else {
			result = append(result, value)
			sum += value
		}
	}
	return result
}

func SetSite(tag *base.Tag, x, y int) *base.Tag {
	newTag := tag.Copy()
	newTag.X = x
	newTag.Y = y
	return newTag
}

// CopList 将普通类型转换成值类型
func CopList(listTag [][]base.Tag) [][]*base.Tag {
	resDataList := make([][]*base.Tag, 0)
	for _, tags := range listTag {
		tagRow := make([]*base.Tag, 0)
		for _, tag := range tags {
			tagRow = append(tagRow, tag.Copy())
		}
		resDataList = append(resDataList, tagRow)
	}
	return resDataList
}

func GenericsCopList[T any](list [][]T) [][]*T {
	resDataList := make([][]*T, 0)
	for i, ts := range list {
		row := make([]*T, 0)
		for i2, t := range ts {
			copt := t
			row[i2] = &copt
		}
		resDataList[i] = row
	}
	return resDataList
}
