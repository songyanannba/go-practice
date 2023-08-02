package gameHandle

import (
	"errors"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/service/logic/upper/seamlessWallet"
	"slot-server/utils"
	"slot-server/utils/helper"
	"time"
)

// NotifyBet 通知第三方押注
func (h *Handle) NotifyBet() error {
	txnId := "TX" + time.Now().Format("0601021504") + utils.RandomString(10)

	betFloat := amount2penny(h.Bill.TotalBet)

	// 创建商户请求
	h.MerchantReq = seamlessWallet.NewMerchantReq(h.Merchant, h.User.Token)

	// 从三方商户获取用户余额
	req := seamlessWallet.BetReq{
		PlayerName:    h.User.Username,
		GameId:        int(h.SlotId),
		Currency:      h.User.Currency,
		BetAmount:     betFloat,
		TransactionId: txnId,
		UpdatedTime:   time.Now().UnixMilli(),
	}
	player, code, err := h.MerchantReq.Bet(req)
	if err != nil {
		if code == enum.MerchantApiCode10AmountInsufficient {
			h.Code = enum.MerchantApiCode10AmountInsufficient
			err = errors.New("amount=" + helper.Itoa(int(helper.Mul100(player.PlayerBalance))))
		}
		return err
	}

	// 押注后操作
	h.Bill.AfterBet(player.PlayerBalance, player.AgentTxid)

	h.Txn = &business.Txn{
		PlayerName:    h.User.Username,
		MerchantId:    h.Merchant.ID,
		UserId:        h.User.ID,
		GameId:        int(h.SlotId),
		TxnId:         txnId,
		Currency:      h.User.Currency,
		Bet:           h.Bet,
		Raise:         h.Bill.TotalRaise,
		ChangeBal:     -h.Bill.TotalBet,
		BeforeBal:     h.Bill.AfterBetAmount + h.Bet,
		AfterBal:      h.Bill.AfterBetAmount,
		Status:        enum.TxnStatus1InProgress,
		PlatformTxnId: player.AgentTxid,
	}
	err = global.GVA_DB.Create(h.Txn).Error
	if err != nil {
		global.GVA_LOG.Error("创建Txn失败", zap.Error(err))
		refundErr := NotifyRefund(h.Merchant, h.Txn)
		if refundErr != nil {
			global.GVA_LOG.Error("请求退款失败", zap.Error(refundErr))
		}
		return err
	}
	return nil
}

func (h *Handle) CreateRecord() error {
	h.Record = &business.SlotRecord{
		Date:      time.Now().Format("20060102"),
		UserId:    h.User.ID,
		SlotId:    h.SlotId,
		TxnId:     0,
		Bet:       int(h.Bet),
		Raise:     h.Bill.TotalRaise,
		BeforeBal: h.Bill.OriginalAmount,
		ChangeBal: -h.Bill.TotalBet,
		AfterBal:  h.Bill.AfterBetAmount,
		Currency:  h.User.Currency,
		Status:    enum.TxnStatus1InProgress,
	}
	err := global.GVA_DB.Create(h.Record).Error
	if err != nil {
		global.GVA_LOG.Error("创建Record失败", zap.Error(err))
		return err
	}
	return nil
}

func NotifyRefund(merchant *business.Merchant, txn *business.Txn) (err error) {
	req := seamlessWallet.RefundReq{
		GameId:        txn.GameId,
		BetAmount:     helper.Div100(txn.Bet + txn.Raise),
		PlayerName:    txn.PlayerName,
		TransactionId: txn.TxnId,
		AgentTxid:     txn.PlatformTxnId,
		UpdatedTime:   time.Now().UnixMilli(),
	}
	_, _, err = seamlessWallet.NewMerchantReq(merchant, "").Refund(req)
	if err != nil {
		txn.Status = enum.TxnStatus3CancelInProcess
	} else {
		txn.Status = enum.TxnStatus5Canceled
	}
	global.GVA_DB.Select("status").Updates(txn)
	return err
}

func amount2penny(v int64) float64 {
	if v == 0 {
		return 0
	}
	i, _ := decimal.NewFromInt(v).Div(decimal.NewFromInt(100)).Float64()
	return i
}

func NotifyResult(merchant *business.Merchant, txn *business.Txn) (err error) {
	req := seamlessWallet.ResultReq{
		PlayerName:    txn.PlayerName,
		GameId:        txn.GameId,
		Currency:      txn.Currency,
		TransactionId: txn.TxnId,
		BetAmount:     helper.Div100(txn.Bet),
		WinAmount:     helper.Div100(txn.Win),
		AgentTxid:     txn.PlatformTxnId,

		UpdatedTime: time.Now().UnixMilli(),
	}
	res, _, err := seamlessWallet.NewMerchantReq(merchant, "").Result(req)
	if err != nil {
		return err
	}
	txn.Status = enum.TxnStatus4Completed
	txn.RealBal = helper.Mul100(res.PlayerBalance)
	err = global.GVA_DB.Select("status", "real_bal").Updates(txn).Error
	return err
}
