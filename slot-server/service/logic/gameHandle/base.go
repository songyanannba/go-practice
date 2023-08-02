package gameHandle

import (
	"fmt"
	"github.com/lonng/nano/session"
	"google.golang.org/protobuf/reflect/protoreflect"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/pbs"
	"slot-server/service/logic/upper/seamlessWallet"
	"slot-server/service/slot/component"
	"slot-server/utils/helper"
)

type Handle struct {
	Handler

	Session  *session.Session
	User     *business.User
	Merchant *business.Merchant

	MerchantReq *seamlessWallet.MerchantHandle

	SlotId      uint
	Bet         int64              // 押注
	Raise       int64              // 加注
	BuyFreeCoin int64              // 购买免费次数金额
	BuyReCoin   int64              // 购买重转次数金额
	Amount      int64              // 原本余额
	Options     []component.Option // 机台选项

	Req *pbs.Spin
	Ack protoreflect.ProtoMessage

	UserSpin *business.SlotUserSpin // 用户转统计

	Demo   bool                 // 是否是试玩
	Txn    *business.Txn        // 交易记录
	Record *business.SlotRecord // 游戏记录
	Bill   Bill                 // 账单
	Code   int                  // 错误code
}

type Bill struct {
	Currency       string // 货币
	OriginalAmount int64  // 原本余额
	TotalBet       int64  // 总押注
	TotalRaise     int64  // 总加注
	AfterBetAmount int64  // 押注后余额
	TotalWin       int64  // 总赢
	AfterAmount    int64  // 转后余额
	PlatformTxnId  string // 平台交易id
}

// AfterBet 押注后
func (b *Bill) AfterBet(afterBetAmount float64, platformTxnId string) {
	// 押注后余额
	b.AfterBetAmount = helper.Mul100(afterBetAmount)
	// 押注前余额
	b.OriginalAmount = b.AfterBetAmount + b.TotalBet
	// 平台单号
	b.PlatformTxnId = platformTxnId
}

func (h *Handle) addOption(option ...component.Option) {
	h.Options = append(h.Options, option...)
}

// MulIsExceed 倍率是否超出限制
func (h *Handle) MulIsExceed(i int, limit float64) bool {
	win := h.GetTotalWin()
	mul := float64(win) / float64(h.Bet)
	if mul <= limit {
		return false
	}
	global.GVA_LOG.Warn(fmt.Sprintf("第%d次转结果为%d,总倍数为%f,超过%.2f倍,继续转", i+1, win, mul, limit))
	return true
}

// GetAllBet 获取总押注
func (h *Handle) GetAllBet() int64 {
	return h.Bet + h.Raise + h.BuyReCoin + h.BuyFreeCoin
}

// GetAllRaise 获取总加注
func (h *Handle) GetAllRaise() int64 {
	return h.Raise + h.BuyReCoin + h.BuyFreeCoin
}
