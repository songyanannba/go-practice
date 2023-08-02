package business

import (
	"slot-server/global"
)

// SlotReelData 结构体
type SlotReelData struct {
	global.GVA_MODEL
	SlotId     int    `json:"slotId" form:"slotId" gorm:"column:slot_id;default:0;comment:机器编号;size:32;"`
	WeightData string `json:"weightData" form:"weightData" gorm:"column:weight_data;type:text;comment:普通标签权重;"`
	ReelData   string `json:"reelData" form:"reelData" gorm:"column:reel_data;type:text;comment:普通标签排布;"`
	FsWeight   string `json:"fsWeight" form:"fsWeight" gorm:"column:fs_weight;type:text;comment:免费转标签权重;"`
	FsReelData string `json:"fsReelData" form:"fsReelData" gorm:"column:fs_reel_data;type:text;comment:免费转标签排布;"`
	Group      int    `json:"group" form:"group" gorm:"column:group;default:0;comment:组;size:32;"`
	Which      int    `json:"which" form:"which" gorm:"column:which;default:0;comment:位置;size:32;"`
	Demo       uint8  `json:"demo" form:"demo" gorm:"column:demo;default:2;comment:是否试玩;size:8;"`
}

// TableName SlotReelData 表名
func (SlotReelData) TableName() string {
	return "b_slot_reel_data"
}
