package game

import (
	"errors"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
	"go.uber.org/zap"
	"google.golang.org/protobuf/reflect/protoreflect"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/pbs"
	"slot-server/service/cache"
	"slot-server/service/logic"
	"slot-server/service/logic/gameHandle"
	"slot-server/service/logic/gameHandle/eliminateHandle"
	"slot-server/service/logic/gameHandle/slotHandle"
	"slot-server/service/logic/upper/seamlessWallet"
	"slot-server/utils/env"
	"slot-server/utils/helper"
	"time"
)

var (
	Component = &component.Components{}
	Service   = newService()
)

func init() {
	Component.Register(Service)
}

type GameService struct {
	component.Base
}

func newService() *GameService {
	return &GameService{}
}

type SyncMessage struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// IntoGame 进入游戏
func (gs *GameService) IntoGame(s *session.Session, req *pbs.IntoGame) error {
	ack := &pbs.IntoGameAck{}
	ack.Head = logic.NewAckHead(req.Head)
	userId := uint(0)
	currency := ""

	if req.Head.Demo {
		userId = uint(s.ID())
		currency = "USD"
	} else {
		u, _, err := logic.RequireMerchantToken(req.Head, false)
		if err != nil {
			return logic.RespError(s, req, ack, err)
		}
		userId = u.ID
		currency = u.Currency
		s.Bind(int64(userId))
	}

	err := logic.IntoGameFunc(userId, currency, req, ack)
	if err != nil {
		return logic.RespError(s, req, ack, err)
	}

	err = cache.DeleteRecordListCache(userId, req.GameId, time.Now().Format("2006-01-02 15"))
	if err != nil {
		global.GVA_LOG.Error("cache.DeleteRecordListCache err", zap.Error(err))
	}

	return s.Response(ack)
}

// Spin 转动
func (gs *GameService) Spin(s *session.Session, req *pbs.Spin) error {
	var (
		u       = &business.User{}
		m       = &business.Merchant{}
		handler gameHandle.Handler
		ack     any
		err     error
		h       *gameHandle.Handle
	)
	switch req.Opt.GameId {
	case enum.SlotId5, enum.SlotId6:
		ack = &pbs.MatchSpinAck{
			Head: logic.NewAckHead(req.Head),
		}
		handler = eliminateHandle.NewHandle()
	default:
		ack = &pbs.SpinAck{
			Head: logic.NewAckHead(req.Head),
		}
		handler = slotHandle.NewHandle()
	}

	if s.HasKey(enum.SessionDataGameRecordId) {
		return logic.RespError(s, req, ack.(logic.AckHeads), errors.New(`operation is too frequent. Please try again later`))
	}

	if !req.Head.Demo {
		u, m, err = logic.RequireMerchantToken(req.Head, true)
		if err != nil {
			return logic.RespError(s, req, ack.(logic.AckHeads), err)
		}
	}

	h, err = logic.SpinFunc(u, m, req, ack.(protoreflect.ProtoMessage), handler, s)
	if err != nil {
		if h.Code != enum.MerchantApiCode7Fail {
			err = logic.NewErr(pbs.Code_AmountInsufficientError, err)
		}
		return logic.RespError(s, req, ack.(logic.AckHeads), err)
	}

	ack = h.Ack
	if s.HasKey(enum.SessionDataTestSpinHandle) {
		s.Set(enum.SessionDataTestSpinHandle, h)
	}
	return s.Response(ack)
}

// SpinStop 转动停止
func (gs *GameService) SpinStop(s *session.Session, req *pbs.SpinStop) error {
	ack := &pbs.SpinStopAck{}
	ack.Head = logic.NewAckHead(req.Head)
	defer s.Remove(enum.SessionDataGameRecordId)
	u, m, err := logic.RequireMerchantToken(req.Head, true)
	if err != nil {
		return logic.RespError(s, req, ack, err)
	}

	err = logic.SpinStopFunc(u.ID, m, req, ack)
	if err != nil {
		return logic.RespError(s, req, ack, err)
	}

	return s.Response(ack)
}

// Recharge 充值
func (gs *GameService) Recharge(s *session.Session, req *pbs.Recharge) error {
	ack := &pbs.RechargeAck{}
	ack.Head = logic.NewAckHead(req.Head)

	if env.Mode == env.Prod {
		return logic.RespError(s, req, ack, enum.ErrSysError)
	}

	u, _, err := logic.RequireMerchantToken(req.Head, false)
	if err != nil {
		return logic.RespError(s, req, ack, err)
	}

	var moneyLog *business.MoneyLog
	moneyLog, err = cache.ChangeMoney(u.ID, req.Amount, nil, business.MoneyLogWithRecharge())
	if err != nil {
		return logic.RespError(s, req, ack, err)
	}

	ack.Amount = moneyLog.CoinResult
	return s.Response(ack)
}

// Amount 余额
func (gs *GameService) Amount(s *session.Session, req *pbs.Amount) error {
	ack := &pbs.AmountAck{}
	ack.Head = logic.NewAckHead(req.Head)
	u, m, err := logic.RequireMerchantToken(req.Head, true)
	if err != nil {
		return logic.RespError(s, req, ack, err)
	}
	playerAmount, _, err := seamlessWallet.NewMerchantReq(m, req.Head.Token).Balance(u.Username)
	if err != nil {
		return logic.RespError(s, req, ack, err)
	}
	ack.Amount = helper.Mul100(playerAmount.PlayerBalance)

	return s.Response(ack)
}

// BackendOperate game后台操作
func (gs *GameService) BackendOperate(s *session.Session, req *pbs.BackendOperate) error {
	ack := &pbs.BackendOperateAck{}
	ack.Head = logic.NewAckHead(req.Head)
	_, err := logic.ParseBackendToken(req.Head)
	if err != nil {
		ack.Head.Code = pbs.Code_TokenInvalid
		return logic.RespError(s, req, ack, err)
	}

	err = logic.HandleBackendOperate(req)
	if err != nil {
		return logic.RespError(s, req, ack, err)
	}

	return s.Response(ack)
}

func (gs *GameService) RecordList(s *session.Session, req *pbs.RecordListReq) error {
	ack := &pbs.RecordListAck{}
	ack.Head = logic.NewAckHead(req.Head)

	uid := uint(s.UID())
	if uid == 0 {
		return logic.RespError(s, req, ack, enum.ErrTokenInvalid)
	}

	//获取列表
	if err := gameHandle.RecordList(s, uid, req, ack); err != nil {
		return logic.RespError(s, req, ack, err)
	}

	return s.Response(ack)
}

func (gs *GameService) RecordMenu(s *session.Session, req *pbs.RecordMenuReq) error {
	ack := &pbs.RecordMenuAck{}
	ack.Head = logic.NewAckHead(req.Head)

	uid := uint(s.UID())
	if uid == 0 {
		return logic.RespError(s, req, ack, enum.ErrTokenInvalid)
	}

	//获取列表
	if err := gameHandle.RecordMenu(uid, req, ack); err != nil {
		return logic.RespError(s, req, ack, err)
	}

	return s.Response(ack)
}

func (gs *GameService) RecordDetail(s *session.Session, req *pbs.RecordDetailReq) error {
	ack := &pbs.RecordDetailAck{}
	ack.Head = logic.NewAckHead(req.Head)

	uid := uint(s.UID())
	if uid == 0 {
		return logic.RespError(s, req, ack, enum.ErrTokenInvalid)
	}

	id := business.ParseSlotRecordId(req.No)
	if id == 0 {
		ack.Head.Code = pbs.Code_NotExists
		return logic.RespError(s, req, ack, enum.ErrRecordNotExist)
	}
	//获取列表
	var record *business.SlotRecord
	err := global.GVA_DB.First(&record, "id = ?", id).Error
	if err != nil {
		ack.Head.Code = pbs.Code_NotExists
		return logic.RespError(s, req, ack, enum.ErrRecordNotExist)
	}
	ack.Data = record.Ack
	return s.Response(ack)
}
