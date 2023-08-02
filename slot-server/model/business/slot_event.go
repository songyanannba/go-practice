// 自动生成模板SlotEvent
package business

import (
	"slot-server/global"
)

// SlotEvent 结构体
type SlotEvent struct {
	global.GVA_MODEL
	SlotId int    `json:"slotId" form:"slotId" gorm:"column:slot_id;default:0;comment:机器id;size:32;"`
	Event1 string `json:"event1" form:"event1" gorm:"column:event1;type:text;comment:机器特殊事件;"`
	Demo   uint8  `json:"demo" form:"demo" gorm:"column:demo;default:2;comment:是否试玩;size:8;"`
}

// TableName SlotEvent 表名
func (SlotEvent) TableName() string {
	return "b_slot_event"
}
