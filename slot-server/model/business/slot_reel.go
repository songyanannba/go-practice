// 自动生成模板SlotReel
package business

import (
	"slot-server/global"
)

// SlotReel 结构体
type SlotReel struct {
	global.GVA_MODEL
	SlotId   int    `json:"slotId" form:"slotId" gorm:"column:slot_id;default:0;comment:机器id;size:32;"`
	Size     string `json:"size" form:"size" gorm:"column:size;comment:显示区域;size:50;"`
	Num      int    `json:"num" form:"num" gorm:"column:num;default:0;comment:显示几个格子;size:32;"`
	Space    int    `json:"space" form:"space" gorm:"column:space;default:0;comment:格子间距;size:32;"`
	Speed    int    `json:"speed" form:"speed" gorm:"column:speed;default:0;comment:转动速度;size:32;"`
	StopTime int    `json:"stopTime" form:"stopTime" gorm:"column:stop_time;default:0;comment:停止时间;size:32;"`
	Spring   string `json:"spring" form:"spring" gorm:"column:spring;comment:滚动回弹节奏;size:100;"`
	Demo     uint8  `json:"demo" form:"demo" gorm:"column:demo;default:2;comment:是否试玩;size:8;"`
}

// TableName SlotReel 表名
func (SlotReel) TableName() string {
	return "b_slot_reel"
}

func (s *SlotReel) GetCondition() string {
	return ""
}
