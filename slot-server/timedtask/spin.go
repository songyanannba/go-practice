package timedtask

import (
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/service/cache"
	"slot-server/utils/helper"
	"time"
)

func FinishSpin() {
	defer helper.PanicRecover()
	global.GVA_LOG.Info("FinishSpin start")
	date := time.Now().Add(-6 * time.Minute).Format("2006-01-02")
	t := time.Now().Add(-1 * time.Minute)
	var records []*business.SlotRecord
	global.GVA_DB.Find(&records, "status = ? and date = ? and created_at > ?", enum.CommonStatusProcessing, date, t)
	for _, record := range records {
		global.GVA_DB.Transaction(func(tx *gorm.DB) error {
			record.Gain = int(decimal.NewFromInt(int64(record.Gain)).
				Sub(decimal.NewFromInt(int64(record.Bet))).
				Sub(decimal.NewFromInt(record.Raise)).
				IntPart())
			record.Status = enum.CommonStatusFinish
			err := tx.Select("gain", "status").Updates(&record).Error
			if err != nil {
				global.GVA_LOG.Error("FinishSpin error", zap.Error(err))
				return err
			}
			_, err = cache.ChangeMoney(record.UserId, int64(record.Gain), tx, business.MoneyLogWithSlot(record.SlotId))
			if err != nil {
				global.GVA_LOG.Error("FinishSpin changeMoney error", zap.Error(err))
				return err
			}
			return nil
		})
	}
}
