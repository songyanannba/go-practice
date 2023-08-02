// 自动生成模板Slot
package business

import (
	"slot-server/global"
)

// Slot 结构体
type Slot struct {
	global.GVA_MODEL
	Name        string  `json:"name" form:"name" gorm:"column:name;comment:名称;size:100;"`
	NamePkg     string  `json:"namePkg" form:"namePkg" gorm:"column:name_pkg;comment:分包名称;size:100;"`
	Icon        string  `json:"icon" form:"icon" gorm:"column:icon;comment:使用的标签图集;size:255;"`
	IconShadow  string  `json:"iconShadow" form:"iconShadow" gorm:"column:icon_shadow;comment:使用的模糊标签图集;size:255;"`
	PaylineNo   int     `json:"paylineNo" form:"paylineNo" gorm:"column:payline_no;comment:线数编号;default:0;size:32;"`
	BigWin      string  `json:"bigWin" form:"bigWin" gorm:"column:big_win;comment:赢钱最大区间;size:255;"`
	BetNum      string  `json:"betNum" form:"betNum" gorm:"type:text;column:bet_num;comment:机器押注;"`
	Raise       float64 `json:"raise" form:"raise" gorm:"column:raise;comment:加注;type:decimal(14,2);default:0;"`
	BuyFreeSpin float64 `json:"buyFreeSpin" form:"buyFreeSpin" gorm:"column:buy_free_spin;comment:购买免费旋转;type:decimal(14,2);default:0;"`
	BuyReSpin   float64 `json:"buyReSpin" form:"buyReSpin" gorm:"column:buy_re_spin;comment:购买重旋转;type:decimal(14,2);default:0;"`
	JackpotRule string  `json:"jackpotRule" form:"jackpotRule" gorm:"column:jackpot_rule;comment:奖池规则;size:255;"`
	Status      uint8   `json:"status" form:"status" gorm:"column:status;default:1;comment:状态;size:8;"`
	Url         string  `json:"url" form:"url" gorm:"column:url;comment:游戏地址;size:255;"`
	TopMul      int     `json:"topMul" form:"topMul" gorm:"column:top_mul;comment:最高倍数;default:0;size:32;"`
	ClientConf  string  `json:"clientConf" form:"clientConf" gorm:"column:client_conf;comment:客户端配置;type:text;"`
}

// TableName Slot 表名
func (Slot) TableName() string {
	return "b_slot"
}
