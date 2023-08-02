// 自动生成模板MoneySlot
package business

import (
	"slot-server/global"
)

// MoneySlot 结构体
type MoneySlot struct {
	global.GVA_MODEL
	Date          string  `json:"date" form:"date" gorm:"column:date;comment:日期;size:10;"`
	SlotId        uint    `json:"slotId" form:"slotId" gorm:"column:slot_id;default:0;comment:机器编号;size:32;"`
	CoinReduce    int     `json:"coinReduce" form:"coinReduce" gorm:"column:coin_reduce;default:0;comment:消耗金币;size:32;"`
	CoinIncrease  int     `json:"coinIncrease" form:"coinIncrease" gorm:"column:coin_increase;default:0;comment:产出金币;size:32;"`
	RtpRatio      float64 `json:"rtpRatio" form:"rtpRatio" gorm:"column:rtp_ratio;comment:返还比;size:10;"`
	RecentPlayers int     `json:"recentPlayers" form:"recentPlayers" gorm:"column:recent_players;default:0;comment:游玩人数;size:32;"`
	RecentSpins   int     `json:"recentSpins" form:"recentSpins" gorm:"column:recent_spins;default:0;comment:Spin次数;size:32;"`
	AvgSpins      float64 `json:"avgSpins" form:"avgSpins" gorm:"column:avg_spins;default:0;comment:平均Spin次数;size:32;"`
	BkrpPeoples   int     `json:"bkrpPeoples" form:"bkrpPeoples" gorm:"column:bkrp_peoples;default:0;comment:破产人数;size:32;"`
	BkrpTimes     int     `json:"bkrpTimes" form:"bkrpTimes" gorm:"column:bkrp_times;default:0;comment:破产次数;size:32;"`
	BkrpAddTimes  int     `json:"bkrpAddTimes" form:"bkrpAddTimes" gorm:"column:bkrp_add_times;default:0;comment:破产充值次数;size:32;"`
	BkrpAddAmount int     `json:"bkrpAddAmount" form:"bkrpAddAmount" gorm:"column:bkrp_add_amount;default:0;comment:破产充值金额;size:32;"`
	TopUpPeoples  int     `json:"topUpPeoples" form:"topUpPeoples" gorm:"column:top_up_peoples;default:0;comment:充值人数;size:32;"`
	TopUpTimes    int     `json:"topUpTimes" form:"topUpTimes" gorm:"column:top_up_times;default:0;comment:充值次数;size:32;"`
	TopUpAmount   int     `json:"topUpAmount" form:"topUpAmount" gorm:"column:top_up_amount;default:0;comment:充值金额;size:32;"`
}

// TableName MoneySlot 表名
func (MoneySlot) TableName() string {
	return "b_money_slot"
}
