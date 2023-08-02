package timedtask

import (
	"github.com/samber/lo"
	"gorm.io/gorm"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/utils/helper"
	"time"
)

// MoneyUserSlotCalToday
//
//	@Description: 计算今天数据
func MoneyUserSlotCalToday() {
	defer helper.PanicRecover()
	global.GVA_LOG.Info("MoneyUserCalToday start")
	err := MoneyUserSlotCal(time.Now().Format("20060102"))
	if err != nil {
		global.GVA_LOG.Info("MoneyUserCalToday err :" + err.Error())
	}

}

// MoneyUserSlotCalYesterday
//
//	@Description: 计算昨天数据
func MoneyUserSlotCalYesterday() {
	defer helper.PanicRecover()
	global.GVA_LOG.Info("MoneyUserCalToday start")
	err := MoneyUserSlotCal(time.Now().AddDate(0, 0, -1).Format("20060102"))
	if err != nil {
		global.GVA_LOG.Info("MoneyUserCalYesterday err :" + err.Error())
	}
}
func MoneyUserSlotCal(date string) (err error) {
	defer helper.PanicRecover()
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var moneyUserSlotNews []business.MoneyUserSlot
		if err = tx.Raw(`SELECT a.user_id,a.slot_id,sum(a.bet) sum_bet,sum(a.gain) sum_gain, sum(case WHEN a.bet>0 THEN 1 ELSE 0 end) count_bet,sum(case WHEN a.gain>0 THEN 1 ELSE 0 END) gain_count, sum(
							case WHEN a.is_bk=1 THEN 1 ELSE 0 END) bk_count from b_slot_record a 
						    WHERE a.date = ?
							GROUP BY a.user_id,a.slot_id ;`, date).Scan(&moneyUserSlotNews).Error; err != nil {
			return err
		}
		var moneyUserSlots []business.MoneyUserSlot
		if err = tx.Find(&moneyUserSlots, "date = ?", date).Error; err != nil {
			return err
		}
		for _, v := range moneyUserSlotNews {
			v.Date = date
			uss := lo.Filter(moneyUserSlots, func(item business.MoneyUserSlot, index int) bool {
				if item.UserId == v.UserId && item.SlotId == v.SlotId {
					return true
				}
				return false
			})
			if len(uss) > 0 {

				v.ID = uss[0].ID
				v.CreatedAt = uss[0].CreatedAt
			}
			if err = tx.Save(&v).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return err

}
