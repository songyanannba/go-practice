package timedtask

import (
	"go.uber.org/zap"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/service/cache"
	"slot-server/service/logic/gameHandle"
	"slot-server/utils/helper"
	"time"
)

func FinishTxn() {
	var (
		txns      []*business.Txn
		err       error
		now       = time.Now()
		layout    = "2006-01-02 15:04:05"
		startTime = now.Add(-1 * time.Hour).Format(layout)
		endTime   = now.Add(-5 * time.Minute).Format(layout)
	)
	global.GVA_DB.
		Where("created_at between ? and ?", startTime, endTime).
		Where("status in ?", []int{enum.TxnStatus2CompleteInProcess, enum.TxnStatus3CancelInProcess}).
		Find(&txns)
	merchantMap := cache.GetMerchantMap(helper.ArrMap(txns, func(txn *business.Txn) uint {
		return txn.MerchantId
	})...)
	for _, txn := range txns {
		merchant := merchantMap[txn.MerchantId]
		if merchant == nil {
			global.GVA_LOG.Error("完成订单时查询商户失败", zap.Uint("商户ID", txn.MerchantId))
			continue
		}
		if txn.Status == enum.TxnStatus2CompleteInProcess {
			err = gameHandle.NotifyResult(merchant, txn)
		} else {
			err = gameHandle.NotifyRefund(merchant, txn)
		}
		if err != nil {
			global.GVA_LOG.Error("完成订单时通知第三方失败", zap.Uint("商户ID", txn.MerchantId), zap.Error(err))
		}
	}
}
