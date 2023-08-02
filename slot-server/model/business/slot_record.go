package business

import (
	"slot-server/global"
	"slot-server/utils"
	"slot-server/utils/helper"
	"time"
)

// SlotRecord 结构体
type SlotRecord struct {
	global.GVA_MODEL
	Date       string `json:"date" form:"date" gorm:"index;column:date;default:'';comment:日期;size:10;"`
	MerchantId uint   `json:"merchantId" form:"merchantId" gorm:"column:merchant_id;default:0;comment:商户编号;size:32;"`
	UserId     uint   `json:"userId" form:"userId" gorm:"index;column:user_id;default:0;comment:用户编号;size:64;"`
	SlotId     uint   `json:"slotId" form:"slotId" gorm:"index;column:slot_id;default:0;comment:机器编号;size:64;"`
	TxnId      uint   `json:"txnId" form:"txnId" gorm:"column:txn_id;default:0;comment:交易编号;size:64;"`
	Bet        int    `json:"bet" form:"bet" gorm:"column:bet;default:0;comment:押注;size:64;"`
	Gain       int    `json:"gain" form:"gain" gorm:"column:gain;default:0;comment:赢钱金额;size:64;"`
	IsBk       int    `json:"isBk" form:"isBk" gorm:"column:is_bk;default:2;comment:是否破产;size:16;"`
	Raise      int64  `json:"raise" form:"raise" gorm:"column:raise;default:0;comment:加注;size:64;"`
	Status     uint8  `json:"status" form:"status" gorm:"column:status;default:1;comment:状态;size:8;"`
	PayTableId string `json:"payTableId" form:"payTableId" gorm:"column:pay_table_id;type:text;comment:赢钱组合编号;"`
	Remark     string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"` //备注
	Ack        []byte `json:"ack" form:"ack" gorm:"type:blob;column:ack;comment:回复;"`         //回复

	ChangeBal int64 `json:"changeBal" form:"changeBal" gorm:"column:change_bal;comment:Change Bal;size:64;"`
	BeforeBal int64 `json:"beforeBal" form:"beforeBal" gorm:"column:before_bal;comment:Before Bal;size:64;"`
	AfterBal  int64 `json:"afterBal" form:"afterBal" gorm:"column:after_bal;comment:After Bal;size:64;"`

	JackpotId  uint    `json:"jackpotId" form:"jackpotId" gorm:"column:jackpot_id;default:0;comment:奖池编号;size:32;"`
	JackpotMul float64 `json:"jackpotMul" form:"jackpotMul" gorm:"column:jackpot_mul;type:decimal(14,2);default:0;comment:奖池倍数;"`

	AckStr string `json:"ackStr" form:"ackStr" gorm:"-"` //回复字符串
	No     string `json:"no" form:"no" gorm:"-"`         //编号

	Currency  string    `json:"currency" form:"currency" gorm:"column:currency;comment:货币;size:10;"` //货币
	CreatedAt time.Time `gorm:"index;type:timestamp;size:0"`                                         // 创建时间
}

// TableName SlotRecord 表名
func (SlotRecord) TableName() string {
	return "b_slot_record"
}

func FmtSlotRecordNo(id uint) string {

	front := utils.Base62Encode(helper.IntReverse(int(id)))

	back := utils.Base62Encode(len(front)*3 + 1)

	return front + utils.Base62Encode(int(id)+1000000000) + back
}

func ParseSlotRecordId(no string) uint {
	if len(no) < 5 {
		return 0
	}
	back := (utils.Base62Decode(no[len(no)-1:]) - 1) / 3
	if back < 0 || back > len(no)-2 {
		return 0
	}

	front := no[:back]
	no = no[back : len(no)-1]

	n := utils.Base62Decode(no)
	if n < 1000000000 {
		return 0
	}
	res := n - 1000000000
	if utils.Base62Encode(helper.IntReverse(res)) != front {
		return 0
	}

	return uint(res)
}
