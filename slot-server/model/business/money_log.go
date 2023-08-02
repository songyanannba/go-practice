// 自动生成模板MoneyLog
package business

import (
	"slot-server/enum"
	"slot-server/global"
)

// MoneyLog 结构体
type MoneyLog struct {
	global.GVA_MODEL
	Date        string `json:"date" form:"date" gorm:"index;column:date;comment:创建时间;size:10;"`
	UserId      uint   `json:"userId" form:"userId" gorm:"index;column:user_id;default:0;comment:用户编号;size:32;"`
	Action      uint   `json:"action" form:"action" gorm:"index;column:action;default:0;comment:操作类型;size:32;"`
	ActionType  uint   `json:"actionType" form:"actionType" gorm:"column:action_type;default:0;comment:子类型;size:32;"`
	CoinInitial int64  `json:"coinInitial" form:"coinInitial" gorm:"column:coin_initial;default:0;comment:初始金币;size:64;"`
	CoinChange  int64  `json:"coinChange" form:"coinChange" gorm:"column:coin_change;default:0;comment:金币变化;size:64;"`
	CoinResult  int64  `json:"coinResult" form:"coinResult" gorm:"column:coin_result;default:0;comment:金币结果;size:64;"`
	GameId      uint   `json:"gameId" form:"gameId" gorm:"index;column:game_id;default:0;comment:游戏编号;size:64;"`
	TxnId       string `json:"txnId" form:"txnId" gorm:"column:txn_id;comment:三方交易编号;size:50;"`
}

// TableName MoneyLog 表名
func (MoneyLog) TableName() string {
	return "b_money_log"
}

type MoneyLogOption func(*MoneyLog)

func MoneyLogWithSlot(slotId uint) MoneyLogOption {
	return func(l *MoneyLog) {
		l.GameId = slotId
		l.Action = enum.MoneyAction1Play
		l.ActionType = enum.MoneyType1Spin
	}
}

func MoneyLogWithRecharge() MoneyLogOption {
	return func(l *MoneyLog) {
		l.Action = enum.MoneyAction2Cash
		l.ActionType = enum.MoneyType2Recharge
	}
}

func MoneyLogWithGive() MoneyLogOption {
	return func(l *MoneyLog) {
		l.Action = enum.MoneyAction3System
		l.ActionType = enum.MoneyType3Give
	}
}

func MoneyLogWithTxnId(slotId uint, txnId string, actionType uint) MoneyLogOption {
	return func(l *MoneyLog) {
		l.TxnId = txnId
		l.GameId = slotId
		l.Action = enum.MoneyAction1Play
		l.ActionType = actionType
	}
}
