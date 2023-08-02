// 自动生成模板Jackpot
package business

import (
	"slot-server/global"
	"strings"
)

// Jackpot 结构体
type Jackpot struct {
	global.GVA_MODEL
	//Start   int     `json:"start" form:"start" gorm:"column:start;default:0;comment:启始值;size:32;"`
	//Inc     float64 `json:"inc" form:"inc" gorm:"column:inc;type:decimal(14,2);default:0;comment:增长速度;"`
	End     float64 `json:"end" form:"end" gorm:"column:end;default:0;comment:最终值;size:32;"` // 原jackpot有增长逻辑 现改为固定值
	Combine string  `json:"combine" form:"combine" gorm:"column:combine;comment:组合;size:255;"`
}

// TableName Jackpot 表名
func (Jackpot) TableName() string {
	return "b_jackpot"
}

func (j Jackpot) ParseCombine() []string {
	return strings.Split(j.Combine, "&")
}

func GetJackpotListBySlot(slot *Slot) (list []*Jackpot, err error) {
	if slot.JackpotRule == "" {
		return nil, nil
	}
	err = global.GVA_DB.Find(&list, "id in ("+slot.JackpotRule+")").Error
	return
}
