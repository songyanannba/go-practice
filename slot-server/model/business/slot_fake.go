// 自动生成模板SlotFake
package business

import (
	"slot-server/global"
	"slot-server/utils/helper"
)

// SlotFake 结构体
type SlotFake struct {
	global.GVA_MODEL
	Type     int    `json:"type" form:"type" gorm:"column:type;default:0;comment:玩法类型;size:32;"`
	SlotId   int    `json:"slotId" form:"slotId" gorm:"column:slot_id;default:0;comment:机器id;size:32;"`
	Num      uint   `json:"num" form:"num" gorm:"column:num;default:0;comment:第几次触发;size:32;"`
	Position string `json:"position" form:"position" gorm:"column:position;comment:排布位置;"`
	Which    uint8  `json:"which" form:"which" gorm:"column:which;default:1;comment:第几列数据;size:8;"`
}

// TableName SlotFake 表名
func (SlotFake) TableName() string {
	return "b_slot_fake"
}

func (f *SlotFake) ParsePosition() []int {
	return helper.SplitInt[int](f.Position, "&")
}
