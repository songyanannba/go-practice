// 自动生成模板SlotPayline
package business

import (
	"slot-server/global"
	"slot-server/utils/helper"
	"strconv"
)

// SlotPayline 结构体
type SlotPayline struct {
	global.GVA_MODEL
	No       int    `json:"no" form:"no" gorm:"column:no;index;default:0;comment:编号;size:32;"`
	Sorted   int    `json:"sorted" form:"sorted" gorm:"column:sorted;default:0;comment:排序;size:32;"`
	Size     string `json:"size" form:"size" gorm:"column:size;comment:尺寸规格;size:50;"`
	Num      int    `json:"num" form:"num" gorm:"column:num;default:0;comment:线数;size:32;"`
	Position string `json:"position" form:"position" gorm:"column:position;comment:划线坐标;size:1000;"`
	Status   uint8  `json:"status" form:"status" gorm:"column:status;default:1;comment:状态;size:8;"`
}

// TableName SlotPayline 表名
func (SlotPayline) TableName() string {
	return "b_slot_payline"
}

func (s *SlotPayline) GetNo() int {
	return s.No
}

func (s *SlotPayline) GetCondition() string {
	return "sorted = " + strconv.Itoa(s.Sorted)
}

// ParseSpec 解析规格 返回行数和列数
func (s SlotPayline) ParseSpec() (int, int) {
	arr := helper.SplitInt[int](s.Size, "*")
	return helper.SliceVal(arr, 0), helper.SliceVal(arr, 1)
}
