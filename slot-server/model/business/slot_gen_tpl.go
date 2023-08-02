// 自动生成模板SlotGenTpl
package business

import (
	"slot-server/global"
)

// SlotGenTpl 结构体
type SlotGenTpl struct {
	global.GVA_MODEL
	Type   uint8  `json:"type" form:"type" gorm:"column:type;default:1;comment:类型;size:8;"`
	SlotId uint   `json:"slotId" form:"slotId" gorm:"column:slot_id;default:0;comment:slot_id;size:32;"`
	Size   string `json:"size" form:"size" gorm:"column:size;comment:尺寸;size:10;"`
	Num    int    `json:"num" form:"num" gorm:"column:num;default:0;comment:数量;size:32;"`
	Params string `json:"params" form:"params" gorm:"column:params;comment:参数;"`
	Result string `json:"result" form:"result" gorm:"column:result;comment:结果;"`
	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`
	Status uint8  `json:"status" form:"status" gorm:"column:status;default:1;comment:状态;size:8;"`
}

// TableName SlotGenTpl 表名
func (s *SlotGenTpl) TableName() string {
	return "b_slot_gen_tpl"
}
