// Package business 自动生成模板UserDailySum
package business

import (
	"slot-server/global"
	"time"
)

// MoneyUser 结构体
type MoneyUser struct {
	global.GVA_MODEL
	Date         string    `json:"date" form:"date" gorm:"column:date;comment:时间;size:10;"`
	UserId       int       `json:"userId" form:"userId" gorm:"column:user_id;default:0;comment:玩家ID;size:32;"`
	UserName     string    `json:"userName" form:"userName" gorm:"column:user_name;comment:玩家名称;size:30;"`
	RegTime      time.Time `json:"regTime" form:"regTime" gorm:"column:reg_time;comment:注册时间;"`
	Amount       int       `json:"amount" form:"amount" gorm:"column:amount;default:0;comment:用户金额;size:32;"`
	SpinDay      int       `json:"spinDay" form:"spinDay" gorm:"column:spin_day;default:0;comment:当天Spin次数;size:32;"`
	BetCommon    int       `json:"betCommon" form:"betCommon" gorm:"column:bet_common;default:0;comment:常用押注金额;size:32;"`
	BetAmount    int       `json:"betAmount" form:"betAmount" gorm:"column:bet_amount;default:0;comment:押注消耗;size:32;"`
	GainAmount   int       `json:"gainAmount" form:"gainAmount" gorm:"column:gain_amount;default:0;comment:赢钱金额;size:32;"`
	LastStand    int       `json:"lastStand" form:"lastStand" gorm:"column:last_stand;default:0;comment:最后游玩机台;size:32;"`
	LastPlayTime time.Time `json:"lastPlayTime" form:"lastPlayTime" gorm:"column:last_play_time;comment:最后游玩时间;"`
}

// TableName MoneyUser 表名
func (MoneyUser) TableName() string {
	return "b_money_user"
}
