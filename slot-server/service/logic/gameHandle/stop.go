package gameHandle

import (
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/pbs"
	"slot-server/service/cache"
	"slot-server/utils/helper"
	"time"
)

// StopHandle 结束游戏操作
func StopHandle(req *pbs.SpinStop, merchant *business.Merchant) (*business.Txn, error) {
	var (
		txn business.Txn
		err error
	)

	err = global.GVA_DB.Last(&txn, "id = ?", req.TxnId).Error
	if err != nil {
		return nil, err
	}
	if txn.Status == enum.TxnStatus4Completed {
		return nil, nil
	}
	// 通知第三方结果
	err = NotifyResult(merchant, &txn)
	if err != nil {
		global.GVA_LOG.Error("通知第三方结果失败", zap.Error(err))
		return nil, err
	}
	return &txn, nil
}

// NotifyTxns 通知第三方结果
func NotifyTxns(txns ...*business.Txn) {
	merchantMap := map[uint]*business.Merchant{}
	merchantIds := helper.DistinctByFunc(txns, func(txn *business.Txn) uint {
		return txn.MerchantId
	})
	var err error
	for _, id := range merchantIds {
		merchantMap[id], err = cache.GetMerchant(id)
		if err != nil {
			global.GVA_LOG.Error("查询商户", zap.Uint("商户ID", id), zap.Error(err))
		}
	}

	for _, txn := range txns {
		merchant, ok := merchantMap[txn.MerchantId]
		if !ok {
			continue
		}
		err = NotifyResult(merchant, txn)
		if err != nil {
			global.GVA_LOG.Error("通知第三方结果失败", zap.Error(err))
			continue
		}
	}
}

func SumJackpot(userId, slotId uint) (*business.SlotRecord, error) {
	intoTime, _ := cache.GetUserIntoGameTime(userId)
	if intoTime == 0 {
		return nil, nil
	}
	var (
		record  business.SlotRecord
		jackpot business.Jackpot
	)
	err := global.GVA_DB.Omit("result_data").
		Last(&record, "user_id = ? and slot_id = ?", userId, slotId).Error
	if err != nil {
		return nil, err
	}

	err = global.GVA_DB.First(&jackpot, "id = ?", record.JackpotId).Error
	if err != nil {
		return nil, err
	}

	// 走过的时间 (毫秒)
	totalMilli := time.Now().UnixMilli() - intoTime
	// 走完一圈所需的时间 (毫秒)
	onceMilli := decimal.NewFromInt(int64(jackpot.End)).
		//Sub(decimal.NewFromInt(int64(jackpot.Start))).
		//Div(decimal.NewFromFloat(jackpot.Inc)).
		Mul(decimal.NewFromInt(1000)).
		IntPart()
	realMilli := totalMilli % onceMilli
	mul, _ := decimal.NewFromInt(realMilli).
		//Mul(decimal.NewFromFloat(jackpot.Inc)).
		Div(decimal.NewFromInt(1000)).
		//Add(decimal.NewFromInt(int64(jackpot.Start))).
		Float64()
	record.JackpotMul = mul
	record.Gain = int(decimal.NewFromInt(int64(record.Bet)).Mul(decimal.NewFromFloat(mul)).IntPart())
	record.Status = enum.CommonStatusFinish
	err = global.GVA_DB.Select("jackpot_mul", "gain", "status").Updates(&record).Error
	if err != nil {
		return nil, err
	}

	_ = cache.SetUserIntoGameTime(userId, time.Now().UnixMilli())
	return &record, nil
}
