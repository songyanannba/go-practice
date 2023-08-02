// 自动生成模板Tracking
package business

import (
	"slot-server/global"
)

// Tracking 结构体
type Tracking struct {
	global.GVA_MODEL
	Date   string `json:"date" form:"date" gorm:"column:date;comment:日期;size:10;"`
	Type   uint8  `json:"type" form:"type" gorm:"column:type;default:0;comment:类型;size:8;"`
	UserId uint   `json:"userId" form:"userId" gorm:"column:user_id;default:0;comment:用户id;size:32;"`
	Num    int    `json:"num" form:"num" gorm:"column:num;default:0;comment:次数;size:32;"`
}

// TableName Tracking 表名
func (Tracking) TableName() string {
	return "b_tracking"
}
