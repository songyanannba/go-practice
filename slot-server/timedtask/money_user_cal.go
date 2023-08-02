package timedtask

import (
	"github.com/samber/lo"
	"gorm.io/gorm"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/utils/helper"
	"time"
)

// MoneyUserCalToday
//
//	@Description: 计算今天数据
func MoneyUserCalToday() {
	defer helper.PanicRecover()
	global.GVA_LOG.Info("MoneyUserCalToday start")
	err := MoneyUserCal(time.Now().Format("20060102"))
	if err != nil {
		global.GVA_LOG.Info("MoneyUserCalToday err :" + err.Error())
	}

}

// MoneyUserCalYesterday
//
//	@Description: 计算昨天数据
func MoneyUserCalYesterday() {
	defer helper.PanicRecover()
	global.GVA_LOG.Info("MoneyUserCalToday start")
	err := MoneyUserCal(time.Now().AddDate(0, 0, -1).Format("20060102"))
	if err != nil {
		global.GVA_LOG.Info("MoneyUserCalYesterday err :" + err.Error())
	}
}

//SELECT a.id as userId,a.username userName,a.amount ,
//b.slot_id slotId,b.created_at playTime,a.created_at
//as regTime,b.bet,b.gain FROM b_user a
//LEFT JOIN b_slot_record b ON a.id = b.user_id

type UserInfo struct {
	UserId   int       `gorm:"column:userId"`   //用户编号
	UserName string    `gorm:"column:userName"` //用户名称
	Amount   int       `gorm:"column:amount"`   //用户金额
	SlotId   int       `gorm:"column:slotId"`   //机台编号
	PlayTime time.Time `gorm:"column:playTime"` //游戏时间
	RegTime  time.Time `gorm:"column:regTime"`  //注册时间
	Bet      int       `gorm:"column:bet"`      //押注消耗
	Gain     int       `gorm:"column:gain"`     //赢钱金额
}

// MoneyUserCal
//
//	@Description: 计算用户数据
//	@param Type 1:今天 2:昨天
//	@return err error

func MoneyUserCal(date string) (err error) {
	defer helper.PanicRecover()
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var userInfos []UserInfo
		err = tx.Raw(`SELECT a.id as userId,
		a.username as userName,a.amount ,b.slot_id slotId,b.created_at as
		playTime,a.created_at as regTime,b.bet,b.gain FROM b_user a 
		LEFT JOIN b_slot_record b ON a.id = b.user_id where b.date = ? order by b.created_at`, date).Scan(&userInfos).Error
		if err != nil {
			return err
		}
		var UserDailySums []business.MoneyUser
		err = tx.Find(&UserDailySums, "date = ?", date).Error
		if err != nil {
			return err
		}
		UserDailySum := business.MoneyUser{}
		var bets = make(map[int]int)
		for _, userInfo := range userInfos {
			if UserDailySum.UserId != userInfo.UserId {
				if UserDailySum != (business.MoneyUser{}) {
					UserDailySum.BetCommon = findMaxValue(bets)
					uss := lo.Filter(UserDailySums, func(item business.MoneyUser, index int) bool {
						if item.UserId == UserDailySum.UserId {
							return true
						}
						return false
					})
					if len(uss) > 0 {
						UserDailySum.ID = uss[0].ID
						UserDailySum.CreatedAt = uss[0].CreatedAt
					}
					err = tx.Save(&UserDailySum).Error
					if err != nil {
						return err
					}
				}
				bets = make(map[int]int)
				UserDailySum = business.MoneyUser{
					Date:     date,
					UserId:   userInfo.UserId,
					UserName: userInfo.UserName,
					Amount:   userInfo.Amount,
					RegTime:  userInfo.RegTime,
				}
			}
			bets[userInfo.Bet]++
			UserDailySum.SpinDay++
			UserDailySum.BetAmount += userInfo.Bet
			UserDailySum.GainAmount += userInfo.Gain
			UserDailySum.LastStand = userInfo.SlotId
			UserDailySum.LastPlayTime = userInfo.PlayTime
		}
		if UserDailySum != (business.MoneyUser{}) {
			UserDailySum.BetCommon = findMaxValue(bets)
			uss := lo.Filter(UserDailySums, func(item business.MoneyUser, index int) bool {
				if item.UserId == UserDailySum.UserId {
					return true
				}
				return false
			})
			if len(uss) > 0 {
				UserDailySum.ID = uss[0].ID
				UserDailySum.CreatedAt = uss[0].CreatedAt
			}
			if err = tx.Save(&UserDailySum).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return err

}
func findMaxValue(m map[int]int) int {
	var maxKey int
	var maxValue int
	for key, value := range m {
		if value > maxValue {
			maxValue = value
			maxKey = key
		}
	}
	return maxKey
}
