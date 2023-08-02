package timedtask

import (
	"fmt"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/utils/helper"
	"strconv"
	"time"
)

// MoneySlotCalToday
//
//	@Description: 计算今天数据
func MoneySlotCalToday() {
	defer helper.PanicRecover()
	global.GVA_LOG.Info("MoneySlotCalToday start")
	err := MoneySlotCal(time.Now().Format("20060102"))
	if err != nil {
		return
	}
}

// MoneySlotCalYesterday
//
//	@Description: 计算昨天数据
func MoneySlotCalYesterday() {
	defer helper.PanicRecover()
	global.GVA_LOG.Info("MoneySlotCalYesterday start")
	err := MoneySlotCal(time.Now().AddDate(0, 0, -1).Format("20060102"))
	if err != nil {
		return
	}
}

// MoneySlotCal
//
//	@Description: 计算数据
//	@param summaryType 1:今天 2:昨天
//	@return err error
func MoneySlotCal(date string) (err error) {
	defer helper.PanicRecover()
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var slotRecords []*business.SlotRecord
		if err = tx.Order("slot_id").Find(&slotRecords, "date = ?", date).Error; err != nil {
			return err
		}
		var MoneySlots []business.MoneySlot
		if err = tx.Find(&MoneySlots, "date = ?", date).Error; err != nil {
			return err
		}
		var players []int
		var bkrpPeoples []int
		var slotDaySum business.MoneySlot
		for _, slotRecord := range slotRecords {
			if slotDaySum.SlotId != slotRecord.SlotId {
				if slotDaySum != (business.MoneySlot{}) {
					slotDaySum.RtpRatio, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(slotDaySum.CoinIncrease)/float64(slotDaySum.CoinReduce)), 64)
					slotDaySum.RecentPlayers = len(lo.Uniq(players))
					slotDaySum.AvgSpins, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(slotDaySum.RecentSpins/slotDaySum.RecentPlayers)), 64)
					slotDaySum.BkrpPeoples = len(lo.Uniq(bkrpPeoples))
					uss := lo.Filter(MoneySlots, func(item business.MoneySlot, index int) bool {
						if item.SlotId == slotDaySum.SlotId {
							return true
						}
						return false
					})
					if len(uss) > 0 {
						slotDaySum.ID = uss[0].ID
						slotDaySum.CreatedAt = uss[0].CreatedAt
					}
					if err := tx.Save(&slotDaySum).Error; err != nil {
						return err
					}
				}
				players = []int{}
				bkrpPeoples = []int{}
				slotDaySum = business.MoneySlot{
					Date:   date,
					SlotId: slotRecord.SlotId,
				}
			}
			slotDaySum.CoinReduce += slotRecord.Bet
			slotDaySum.CoinIncrease += slotRecord.Gain
			players = append(players, int(slotRecord.UserId))
			slotDaySum.RecentSpins += 1
			if slotRecord.IsBk == 1 {
				bkrpPeoples = append(bkrpPeoples, int(slotRecord.UserId))
				slotDaySum.BkrpTimes += 1
			}
		}
		if slotDaySum != (business.MoneySlot{}) {
			slotDaySum.RtpRatio, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(slotDaySum.CoinIncrease)/float64(slotDaySum.CoinReduce)), 64)
			slotDaySum.RecentPlayers = len(lo.Uniq(players))
			slotDaySum.AvgSpins, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(slotDaySum.RecentSpins/slotDaySum.RecentPlayers)), 64)
			slotDaySum.BkrpPeoples = len(lo.Uniq(bkrpPeoples))
			uss := lo.Filter(MoneySlots, func(item business.MoneySlot, index int) bool {
				if item.SlotId == slotDaySum.SlotId {
					return true
				}
				return false
			})
			if len(uss) > 0 {
				slotDaySum.ID = uss[0].ID
				slotDaySum.CreatedAt = uss[0].CreatedAt
			}
			if err := tx.Save(&slotDaySum).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
