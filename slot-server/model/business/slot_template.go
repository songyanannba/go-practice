// 自动生成模板SlotTemplate
package business

import (
	"slot-server/global"
)

// SlotTemplate 结构体
type SlotTemplate struct {
	global.GVA_MODEL
	SlotId int    `json:"slotId" form:"slotId" gorm:"column:slot_id;default:0;comment:游戏编号;size:32;"`
	Type   uint8  `json:"type" form:"type" gorm:"column:type;default:1;comment:类型;size:8;"`
	Column int    `json:"column" form:"column" gorm:"column:column;default:0;comment:列号;size:32;"`
	Layout string `json:"layout" form:"layout" gorm:"column:layout;type:text;comment:排布;"`
	GenId  int    `json:"genId" form:"genId" gorm:"column:gen_id;default:0;comment:生成编号;size:32;"`
}

// TableName SlotTemplate 表名
func (SlotTemplate) TableName() string {
	return "b_slot_template"
}
