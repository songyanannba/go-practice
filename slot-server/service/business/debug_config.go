package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type DebugConfigService struct {
}

// CreateDebugConfig 创建DebugConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (debugConfigService *DebugConfigService) CreateDebugConfig(debugConfig business.DebugConfig) (err error) {
	err = global.GVA_DB.Create(&debugConfig).Error
	return err
}

// DeleteDebugConfig 删除DebugConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (debugConfigService *DebugConfigService) DeleteDebugConfig(debugConfig business.DebugConfig) (err error) {
	err = global.GVA_DB.Delete(&debugConfig).Error
	return err
}

// DeleteDebugConfigByIds 批量删除DebugConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (debugConfigService *DebugConfigService) DeleteDebugConfigByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.DebugConfig{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDebugConfig 更新DebugConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (debugConfigService *DebugConfigService) UpdateDebugConfig(debugConfig business.DebugConfig) (err error) {
	err = global.GVA_DB.Save(&debugConfig).Error
	return err
}

// GetDebugConfig 根据id获取DebugConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (debugConfigService *DebugConfigService) GetDebugConfig(id uint) (debugConfig business.DebugConfig, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&debugConfig).Error
	return
}

// GetDebugConfigInfoList 分页获取DebugConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (debugConfigService *DebugConfigService) GetDebugConfigInfoList(info businessReq.DebugConfigSearch) (list []business.DebugConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.DebugConfig{})
	var debugConfigs []business.DebugConfig
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.SlotId != 0 {
		db = db.Where("slot_id = ?", info.SlotId)
	}
	if info.PalyType != 0 {
		db = db.Where("paly_type = ?", info.PalyType)
	}
	if info.DebugType != 0 {
		db = db.Where("debug_type = ?", info.DebugType)
	}
	if info.Start != 0 {
		db = db.Where("start = ?", info.Start)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	if orderMap[info.Sort] {
		OrderStr = "`" + info.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("`id` desc")
	}

	err = db.Limit(limit).Offset(offset).Find(&debugConfigs).Error
	return debugConfigs, total, err
}
