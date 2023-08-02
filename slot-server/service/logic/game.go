package logic

import (
	"errors"
	"github.com/lonng/nano/session"
	"google.golang.org/protobuf/reflect/protoreflect"
	"gorm.io/gorm"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/pbs"
	"slot-server/service/cache"
	"slot-server/service/logic/gameHandle"
	"slot-server/service/slot/component"
	"time"
)

func IntoGameFunc(userId uint, currency string, req *pbs.IntoGame, ack *pbs.IntoGameAck) error {
	//if !req.Head.Demo {
	//	// 设置用户进入时间
	//	_ = cache.SetUserIntoGameTime(userId, time.Now().UnixMilli())
	//}
	//// 删除免费转回复列表
	//_ = cache.DelFreeSpinAckCache(userId, helper.If(req.Head.Demo, enum.SpinAckType2Demo, enum.SpinAckType1Normal))

	c, err := component.GetSlotConfig(uint(req.GameId), false)
	if err != nil {
		return err
	}

	if c.Status == enum.No {
		ack.Head.Code = pbs.Code_StatusError
		err = enum.ErrNoServer
		return err
	}

	ack.Bet = c.BetMap.Get(currency).Bets
	return nil
}

func SpinFunc(u *business.User, m *business.Merchant, req *pbs.Spin, ack protoreflect.ProtoMessage, handler gameHandle.Handler, s *session.Session) (h *gameHandle.Handle, err error) {
	h, err = gameHandle.Start(u, m, req, ack, handler, s)
	return
}

func SpinStopFunc(userId uint, merchant *business.Merchant, req *pbs.SpinStop, ack *pbs.SpinStopAck) (err error) {
	if req.TxnId == 0 {
		return nil
	}
	if !cache.SetUserSpinLock(userId) {
		ack.Head.Code = pbs.Code_AmountInsufficientError
		return enum.ErrBusy
	}

	defer cache.DelUserSpinLock(userId)
	var txn *business.Txn
	txn, err = gameHandle.StopHandle(req, merchant)
	if err != nil {
		ack.Head.Code = pbs.Code_SystemError
		ack.Head.Message = err.Error()
		return err
	}
	if txn != nil {
		ack.Amount = txn.RealBal
		_ = cache.DeleteRecordListCache(userId, req.GameId, time.Now().Format("2006-01-02 15"))
	}
	return
}

func UserTrackingFunc(uid uint, typ pbs.TrackingType) (err error) {
	if uid == 0 {
		return nil
	}
	datetime := time.Now().Format("2006-01-02")
	var tracking business.Tracking
	err = global.GVA_DB.Where("user_id = ? and date = ? and type = ?", uid, datetime, uint8(typ)).First(&tracking).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tracking = business.Tracking{
				Date:   datetime,
				Type:   uint8(typ),
				UserId: uid,
				Num:    0,
			}
		} else {
			return
		}
	}
	tracking.Num++
	err = global.GVA_DB.Save(&tracking).Error
	return
}
