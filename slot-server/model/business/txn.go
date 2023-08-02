// 自动生成模板Txn
package business

import (
	"slot-server/global"
	"time"
)

// Txn 结构体
type Txn struct {
	global.GVA_MODEL
	MerchantId uint   `json:"merchantId" form:"merchantId" gorm:"index;column:merchant_id;default:0;comment:Merchant ID;size:32;"`
	UserId     uint   `json:"userId" form:"userId" gorm:"column:user_id;comment:User ID;size:32;"`
	PlayerName string `json:"playerName" form:"playerName" gorm:"column:player_name;comment:Player Name;size:100;"`

	GameId   int    `json:"gameId" form:"gameId" gorm:"column:game_id;default:0;comment:Game ID;size:32;"`
	TxnId    string `json:"txnId" form:"txnId" gorm:"column:txn_id;comment:Txn ID;size:50;"`
	Currency string `json:"currency" form:"currency" gorm:"column:currency;comment:Currency;size:10;"`

	Bet   int64 `json:"bet" form:"bet" gorm:"column:bet;comment:Bet;size:64;"`
	Raise int64 `json:"raise" form:"raise" gorm:"column:raise;comment:Raise;size:64;"`
	Win   int64 `json:"win" form:"win" gorm:"column:win;comment:Win;size:64;"`

	ChangeBal int64 `json:"changeBal" form:"changeBal" gorm:"column:change_bal;comment:Change Bal;size:64;"`
	BeforeBal int64 `json:"beforeBal" form:"beforeBal" gorm:"column:before_bal;comment:Before Bal;size:64;"`
	AfterBal  int64 `json:"afterBal" form:"afterBal" gorm:"column:after_bal;comment:After Bal;size:64;"`
	RealBal   int64 `json:"realBal" form:"realBal" gorm:"column:real_bal;comment:Real Bal;size:64;"`

	PlatformTxnId string `json:"platformTxnId" form:"platformTxnId" gorm:"column:platform_txn_id;comment:Platform Txn ID;size:100;"`

	Status uint8 `json:"status" form:"status" gorm:"column:status;default:1;comment:Status;size:8;"`

	CreatedAt time.Time `gorm:"index;size:0"` // 创建时间
}

// TableName Txn 表名
func (Txn) TableName() string {
	return "b_txn"
}
