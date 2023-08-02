// 自动生成模板SlotTemplateGen
package business

import (
	"slot-server/global"
)

// SlotTemplateGen 结构体
type SlotTemplateGen struct {
	global.GVA_MODEL
	SlotId        int     `json:"slotId" form:"slotId" gorm:"column:slot_id;default:0;comment:游戏编号;size:32;"`
	Type          uint8   `json:"type" form:"type" gorm:"column:type;default:1;comment:类型;size:8;"`
	MinRatio      float64 `json:"minRatio" form:"minRatio" gorm:"column:min_ratio;default:0;type:decimal(14,6);comment:最小返奖率;size:32;"`
	MaxRatio      float64 `json:"maxRatio" form:"maxRatio" gorm:"column:max_ratio;default:0;type:decimal(14,6);comment:最大返奖率;size:32;"`
	MinScatter    float64 `json:"minScatter" form:"minScatter" gorm:"column:min_scatter;type:decimal(14,6);default:0;comment:最小消散符号个数;size:32;"`
	MaxScatter    float64 `json:"maxScatter" form:"maxScatter" gorm:"column:max_scatter;type:decimal(14,6);default:0;comment:最大消散符号个数;size:32;"`
	InitialWeight string  `json:"initialWeight" form:"initialWeight" gorm:"column:initial_weight;type:text;comment:初始权重;"`
	SpecialConfig string  `json:"specialConfig" form:"specialConfig" gorm:"column:special_config;type:text;comment:特殊配置;"`
	LargeScale    string  `json:"largeScale" form:"largeScale" gorm:"column:large_scale;type:text;comment:大范围调整;"`
	TrimDown      string  `json:"trimDown" form:"trimDown" gorm:"column:trim_down;type:text;comment:向下微调;"`
	TrimUp        string  `json:"trimUp" form:"trimUp" gorm:"column:trim_up;type:text;comment:向上微调;"`
	FinalWeight   string  `json:"finalWeight" form:"finalWeight" gorm:"column:final_weight;type:text;comment:最终权重;"`
	Template      string  `json:"template" form:"template" gorm:"column:template;comment:模版;type:text;"`
	Schedule      string  `json:"schedule" form:"schedule" gorm:"column:schedule;type:text;comment:进度;"`
	Remarks       string  `json:"remarks" form:"remarks" gorm:"column:remarks;comment:备注;size:500;"`
	State         uint8   `json:"state" form:"state" gorm:"column:state;default:0;comment:状态;size:8;"`
}

// TableName SlotTemplateGen 表名
func (SlotTemplateGen) TableName() string {
	return "b_slot_template_gen"
}
