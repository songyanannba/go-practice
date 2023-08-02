package gameHandle

import (
	"errors"
	"github.com/lonng/nano/session"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"gorm.io/gorm"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/pbs"
	"slot-server/service/cache"
	"slot-server/service/slot/component"
	"slot-server/utils/helper"
	"time"
)

func Start(u *business.User, merchant *business.Merchant, req *pbs.Spin, ack protoreflect.ProtoMessage, handler Handler, s *session.Session) (h *Handle, err error) {
	var (
		slotId = uint(req.Opt.GameId)
		c      *component.Config
	)

	h = &Handle{
		Handler:  handler,
		User:     u,
		Merchant: merchant,
		SlotId:   slotId,
		Req:      req,
		Ack:      ack,
		Bet:      req.Opt.Bet,
		Demo:     req.Head.Demo,
		Txn:      &business.Txn{},
		Record:   &business.SlotRecord{},
		UserSpin: &business.SlotUserSpin{},
		Session:  s,
	}

	handler.SetHandle(h)

	// 获取机台配置
	c, err = component.GetSlotConfig(slotId, false)
	if err != nil {
		return
	}

	if c.Status == enum.No {
		//h.Ack.Head.Code = pbs.Code_StatusError
		err = enum.ErrNoServer
		return
	}

	global.GVA_LOG.Info("start spin",
		zap.Uint("slotId", slotId),
		zap.Uint("userId", u.ID),
		zap.String("currency", u.Currency),
		zap.Int64("bet", req.Opt.Bet),
	)

	if !c.BetMap.Check(u.Currency, h.Bet) {
		err = errors.New("bet range error")
		return
	}

	// 加注
	if req.Opt.Raise {
		h.Raise = decimal.NewFromInt(h.Bet).Mul(decimal.NewFromFloat(c.Raise)).IntPart()
		h.addOption(component.WithRaise(h.Raise))
	}
	if req.Opt.BuyFree {
		h.BuyFreeCoin = decimal.NewFromInt(h.Bet).Mul(decimal.NewFromFloat(c.BuyFee)).IntPart()
		h.addOption(component.WithIsMustFree())
	}
	if req.Opt.BuyRe {
		h.BuyReCoin = decimal.NewFromInt(h.Bet).Mul(decimal.NewFromFloat(c.BuyRes)).IntPart()
		h.addOption(component.WithIsMustRes())
	}

	h.Bill.TotalBet = h.GetAllBet()
	h.Bill.TotalRaise = h.GetAllRaise()
	h.Bill.Currency = h.User.Currency

	// demo流程
	if h.Demo {
		h.UserSpin, _ = cache.GetSlotUserSpinInfo(s.ID(), uint(req.Opt.GameId))
		// 增加游玩次数以匹配假数据
		h.addOption(component.WithPlayNum(h.UserSpin.PlayNum+1), component.WithDemo())
		err = h.Run()
		h.Ack = h.GetAck()
		// 增加游玩次数
		playNum, freeNum := h.GetPlayNum()
		_ = cache.UserSpinInfoInc(h.Session.ID(), h.SlotId, 0, freeNum, playNum, nil)
		return
	}

	// 通知商户下注并获取用户余额
	if err = h.NotifyBet(); err != nil {
		return
	}

	// 第一次转可能需要指定结果
	h.addOption(component.WithNeedSpecify(true))

	// 进行游戏逻辑 超出最大限制则重转 最多10次
	for i := 0; i < 10; i++ {
		err = h.Run()
		if err != nil {
			return h, err
		}
		if !h.MulIsExceed(i, float64(c.TopMul)) {
			break
		}
	}

	h.Bill.TotalWin = h.GetTotalWin()
	h.Bill.AfterAmount = h.Bill.AfterBetAmount + h.Bill.TotalWin
	// 保存回复
	h.Ack = h.GetAck()

	err = h.end()
	if err != nil {
		return h, err
	}

	return
}

func (h *Handle) end() (err error) {
	// 变更记录
	if err = global.GVA_DB.Transaction(h.update); err != nil {
		return err
	}

	h.Session.Set(enum.SessionDataGameRecordId, h.Record.ID)
	return
}

func (h *Handle) update(tx *gorm.DB) (err error) {
	// 创建游戏记录
	err = h.InsertRecords(tx)
	if err != nil {
		global.GVA_LOG.Error("创建游戏记录失败 err:" + err.Error())
		return err
	}

	err = h.SaveTxn(tx)
	if err != nil {
		global.GVA_LOG.Error("Save Txn err:" + err.Error())
		return err
	}
	return err
}

func (h *Handle) InsertRecords(tx *gorm.DB) error {
	var (
		ackProto, _ = proto.Marshal(h.Ack)
		spin        = h.GetSpin()
		jackpotId   uint
		jackpotMul  float64
	)
	if spin != nil && spin.Jackpot != nil {
		jackpotId = spin.Jackpot.Id
		jackpotMul = spin.Jackpot.End
	}
	record := business.SlotRecord{
		Date:       time.Now().Format("20060102"),
		MerchantId: h.Merchant.ID,
		UserId:     h.User.ID,
		Bet:        int(h.Bet),
		Raise:      h.Raise,
		SlotId:     h.SlotId,
		Gain:       int(h.Bill.TotalWin),
		IsBk:       enum.No,
		TxnId:      h.Txn.ID,
		JackpotId:  jackpotId,
		JackpotMul: jackpotMul,
		Status:     uint8(enum.CommonStatusFinish),
		Ack:        ackProto,

		Currency:  h.User.Currency,
		BeforeBal: h.Bill.OriginalAmount,
		ChangeBal: h.Bill.TotalWin - h.Bill.TotalBet,
		AfterBal:  h.Bill.AfterAmount,
	}
	err := tx.Create(&record).Error
	if err != nil {
		return err
	}

	h.Record = &record
	return nil
}

// SaveTxn 保存主注单与子注单
func (h *Handle) SaveTxn(tx *gorm.DB) (err error) {
	var subs []*business.TxnSub
	spins := h.GetSpins()
	beforeBal := h.Record.BeforeBal
	// 生产子注单
	for i := 0; i < len(spins); i++ {
		spin := spins[i]
		win := int64(spin.Gain)
		bet := helper.If(i == 0, h.Bet, 0)
		raise := helper.If(i == 0, h.Bill.TotalRaise, 0)
		totalBet := bet + raise
		afterBal := beforeBal + win - totalBet
		subs = append(subs, &business.TxnSub{
			MerchantId: h.Merchant.ID,
			RecordId:   h.Record.ID,
			Pid:        h.Txn.ID,
			Type:       uint8(spin.Type()),
			Bet:        bet,
			Raise:      raise,
			Win:        win,
			BeforeBal:  beforeBal,
			ChangeBal:  win - totalBet,
			AfterBal:   afterBal,
		})
		beforeBal = afterBal
	}
	err = tx.CreateInBatches(subs, 200).Error
	if err != nil {
		return err
	}

	h.Txn.Win = h.Bill.TotalWin
	h.Txn.ChangeBal += h.Bill.TotalWin
	h.Txn.AfterBal += h.Bill.TotalWin
	h.Txn.BeforeBal = h.Bill.AfterAmount
	h.Txn.Status = enum.TxnStatus2CompleteInProcess
	err = tx.Select("win", "change_bal", "after_bal", "before_bal", "status").Save(h.Txn).Error
	return err
}

// isBk 是否破产
func isBk(amount int64, betNum []int) int {
	if amount < int64(helper.SliceVal(betNum, 0)) {
		return enum.Yes
	}
	return enum.No
}
