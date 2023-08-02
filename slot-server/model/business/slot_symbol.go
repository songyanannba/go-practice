// 自动生成模板SlotSymbol
package business

import (
	"slot-server/global"
	"strings"
)

// SlotSymbol 结构体
type SlotSymbol struct {
	global.GVA_MODEL
	SlotId      int    `json:"slotId" form:"slotId" gorm:"column:slot_id;default:0;comment:机器id;size:32;"`
	Name        string `json:"name" form:"name" gorm:"column:name;comment:标签名;size:50;"`
	Size        string `json:"size" form:"size" gorm:"column:size;comment:标签尺寸;size:50;"`
	Hierarchy   int    `json:"hierarchy" form:"hierarchy" gorm:"column:hierarchy;default:0;comment:标签层级;size:32;"`
	IsWild      uint8  `json:"isWild" form:"isWild" gorm:"column:is_wild;default:1;comment:标签层级;size:8;"`
	Sub         string `json:"sub" form:"sub" gorm:"column:sub;comment:替代标签;size:1000;"`
	Multiple    int    `json:"multiple" form:"multiple" gorm:"column:multiple;default:0;comment:翻倍;size:32;"`
	IsSingleWin uint8  `json:"isSingleWin" form:"isSingleWin" gorm:"column:is_single_win;default:1;comment:单出是否赢钱;size:8;"`
	Status      uint8  `json:"status" form:"status" gorm:"column:status;default:1;comment:状态;size:8;"`
}

// TableName SlotSymbol 表名
func (SlotSymbol) TableName() string {
	return "b_slot_symbol"
}

func (s *SlotSymbol) GetCondition() string {
	return ""
}

func (s SlotSymbol) ParseInclude() []string {
	subs := strings.Trim(s.Sub, " ")
	if subs == "" {
		return nil
	}
	return strings.Split(subs, "&")
}
