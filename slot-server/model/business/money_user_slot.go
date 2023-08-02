// 自动生成模板MoneyUserSlot
package business

import (
	"slot-server/global"
)

// MoneyUserSlot 结构体
type MoneyUserSlot struct {
	global.GVA_MODEL
	Date      string `json:"date" form:"date" gorm:"column:date;comment:日期;size:10;"`
	UserId    int    `json:"userId" form:"userId" gorm:"column:user_id;default:0;comment:用户编号;size:32;"`
	SlotId    int    `json:"slotId" form:"slotId" gorm:"column:slot_id;default:0;comment:机器编号;size:32;"`
	SumBet    int    `json:"sumBet" form:"sumBet" gorm:"column:sum_bet;default:0;comment:押注合计;size:32;"`
	SumGain   int    `json:"sumGain" form:"sumGain" gorm:"column:sum_gain;default:0;comment:赢钱合计;size:32;"`
	CountBet  int    `json:"countBet" form:"countBet" gorm:"column:count_bet;default:0;comment:押注次数;size:32;"`
	CountGain int    `json:"countGain" form:"countGain" gorm:"column:count_gain;default:0;comment:赢钱次数;size:32;"`
	CountBk   int    `json:"countBk" form:"countBk" gorm:"column:count_bk;default:0;comment:破产次数;size:32;"`
}

// TableName MoneyUserSlot 表名
func (MoneyUserSlot) TableName() string {
	return "b_money_user_slot"
}
