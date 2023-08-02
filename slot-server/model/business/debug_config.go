// 自动生成模板DebugConfig
package business

import (
	"slot-server/global"
)

// DebugConfig 结构体
type DebugConfig struct {
	global.GVA_MODEL
	SlotId     int    `json:"slotId" form:"slotId" gorm:"column:slot_id;default:0;comment:机台编号;size:32;"`
	PalyType   uint8  `json:"palyType" form:"palyType" gorm:"column:paly_type;default:1;comment:游戏类型;size:8;"`
	DebugType  uint8  `json:"debugType" form:"debugType" gorm:"column:debug_type;default:1;comment:对象类型;size:8;"`
	ResultData string `json:"resultData" form:"resultData" gorm:"column:result_data;comment:结果数据;type:text "`
	Start      uint8  `json:"start" form:"start" gorm:"column:start;comment:是否启用;size:255;"`
	UserId     int    `json:"userId" form:"userId" gorm:"column:user_id;default:0;comment:用户编号;size:32;"`
}

// TableName DebugConfig 表名
func (DebugConfig) TableName() string {
	return "b_debug_config"
}
