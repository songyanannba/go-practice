// 自动生成模板TxnSub
package business

import (
	"slot-server/global"
)

// TxnSub 结构体
type TxnSub struct {
	global.GVA_MODEL
	MerchantId uint  `json:"merchantId" form:"merchantId" gorm:"column:merchant_id;default:0;comment:Merchant ID;size:32;"`
	Pid        uint  `json:"pid" form:"pid" gorm:"column:pid;default:0;comment:PID;size:32;"`
	RecordId   uint  `json:"recordId" form:"recordId" gorm:"column:record_id;default:0;comment:Record ID;size:64;"`
	Type       uint8 `json:"type" form:"type" gorm:"column:type;default:1;comment:Type;size:8;"`
	Bet        int64 `json:"bet" form:"bet" gorm:"column:bet;comment:Bet;size:64;"`
	Raise      int64 `json:"raise" form:"raise" gorm:"column:raise;comment:Raise;size:64;"`
	Win        int64 `json:"win" form:"win" gorm:"column:win;comment:Win;size:64;"`
	BeforeBal  int64 `json:"beforeBal" form:"beforeBal" gorm:"column:before_bal;comment:Before Bal;size:64;"`
	ChangeBal  int64 `json:"changeBal" form:"changeBal" gorm:"column:change_bal;comment:Change Bal;size:64;"`
	AfterBal   int64 `json:"afterBal" form:"afterBal" gorm:"column:after_bal;comment:After Bal;size:64;"`
}

// TableName TxnSub 表名
func (TxnSub) TableName() string {
	return "b_txn_sub"
}
