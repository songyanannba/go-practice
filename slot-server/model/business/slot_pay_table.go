// 自动生成模板SlotPayTable
package business

import (
	"slot-server/global"
	"strings"
)

// SlotPayTable 结构体
type SlotPayTable struct {
	global.GVA_MODEL
	SlotId      int     `json:"slotId" form:"slotId" gorm:"column:slot_id;default:0;comment:机器id;size:32;"`
	Type        uint8   `json:"type" form:"type" gorm:"column:type;default:1;comment:类型;size:8;"`
	Combine1    string  `json:"combine1" form:"combine1" gorm:"column:combine1;type:text;comment:组合1;"`
	CombineNum1 int     `json:"combineNum1" form:"combineNum1" gorm:"column:combine_num1;default:0;comment:组合1数量;size:32;"`
	Combine2    string  `json:"combine2" form:"combine2" gorm:"column:combine2;type:text;comment:组合2;"`
	CombineNum2 int     `json:"combineNum2" form:"combineNum2" gorm:"column:combine_num2;default:0;comment:组合2数量;size:32;"`
	WinMultiple float64 `json:"winMultiple" form:"winMultiple" gorm:"column:win_multiple;type:decimal(14,2);default:0;comment:赢钱倍数;"`
	Status      uint8   `json:"status" form:"status" gorm:"column:status;default:1;comment:状态;size:8;"`
}

// TableName SlotPayTable 表名
func (SlotPayTable) TableName() string {
	return "b_slot_pay_table"
}

func (s *SlotPayTable) GetCondition() string {
	return ""
}

func (s SlotPayTable) ParseCombine() ([]string, []string) {
	return strings.Split(s.Combine1, "&"), strings.Split(s.Combine2, "&")
}
