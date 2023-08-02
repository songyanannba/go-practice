// 自动生成模板SlotTests
package business

import (
	"slot-server/global"
)

// SlotTests 结构体
type SlotTests struct {
	global.GVA_MODEL
	Type     uint8  `json:"type" form:"type" gorm:"column:type;default:1;comment:玩法类型;size:8;"`
	SlotId   uint   `json:"slotId" form:"slotId" gorm:"column:slot_id;default:0;comment:机器id;size:32;"`
	Hold     int    `json:"hold" form:"hold" gorm:"column:hold;default:0;comment:持有;size:32;"`
	Amount   int    `json:"amount" form:"amount" gorm:"column:amount;default:0;comment:押注;size:32;"`
	Win      int    `json:"win" form:"win" gorm:"column:win;default:0;comment:总赢钱;size:64;"`
	MaxNum   int    `json:"maxNum" form:"maxNum" gorm:"column:max_num;default:0;comment:最大次数;size:64;"`
	RunNum   int    `json:"runNum" form:"runNum" gorm:"column:run_num;default:0;comment:实际次数;size:64;"`
	Detail   string `json:"detail" form:"detail" gorm:"column:detail;type:longtext;comment:详情;"`
	Status   uint8  `json:"status" form:"status" gorm:"column:status;default:1;comment:状态;size:8;"`
	Bet      int    `json:"bet" form:"bet" gorm:"column:bet;default:0;comment:押注金额;size:32;"`
	Raise    int    `json:"raise" form:"raise" gorm:"column:raise;default:0;comment:加注金额;size:32;"`
	GameType int    `json:"gameType" form:"gameType" gorm:"column:game_type;default:0;comment:游戏类型;size:8;"`
	TestId   int    `json:"testId" form:"testId" gorm:"column:test_id;default:0;comment:主游戏Id;size:32;"`
	Rank     int    `json:"rank" form:"rank" gorm:"column:rank;default:0;comment:进度;size:32;"`
	GameData string `json:"gameData" form:"gameData" gorm:"column:game_data;comment:游戏数据;type:text"`
}

// TableName SlotTests 表名
func (SlotTests) TableName() string {
	return "b_slot_test"
}
