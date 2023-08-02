package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type ConfigsService struct {
}

// CreateConfigs 创建Configs记录
// Author [piexlmax](https://github.com/piexlmax)
func (configsService *ConfigsService) CreateConfigs(configs business.Configs) (err error) {
	err = global.GVA_DB.Create(&configs).Error
	if err != nil {
		business.DelXConfigCacheById(configs.ID)
	}
	return err
}

// DeleteConfigs 删除Configs记录
// Author [piexlmax](https://github.com/piexlmax)
func (configsService *ConfigsService) DeleteConfigs(configs business.Configs) (err error) {
	err = global.GVA_DB.Delete(&configs).Error
	if err == nil {
		business.DelXConfigCacheById(configs.ID)
	}
	return err
}

// DeleteConfigsByIds 批量删除Configs记录
// Author [piexlmax](https://github.com/piexlmax)
func (configsService *ConfigsService) DeleteConfigsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.Configs{}, "id in ?", ids.Ids).Error
	for _, id := range ids.Ids {
		business.DelXConfigCacheById(uint(id))
	}
	return err
}

// UpdateConfigs 更新Configs记录
// Author [piexlmax](https://github.com/piexlmax)
func (configsService *ConfigsService) UpdateConfigs(configs business.Configs) (err error) {
	err = global.GVA_DB.Save(&configs).Error
	if err == nil {
		business.DelXConfigCacheById(configs.ID)
	}
	return err
}

// GetConfigs 根据id获取Configs记录
// Author [piexlmax](https://github.com/piexlmax)
func (configsService *ConfigsService) GetConfigs(id uint) (configs business.Configs, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&configs).Error
	return
}

// GetConfigsInfoList 分页获取Configs记录
// Author [piexlmax](https://github.com/piexlmax)
func (configsService *ConfigsService) GetConfigsInfoList(info businessReq.ConfigsSearch) (list []business.Configs, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.Configs{})
	var configss []business.Configs
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	if info.Value != "" {
		db = db.Where("value LIKE ?", "%"+info.Value+"%")
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
	}

	err = db.Limit(limit).Offset(offset).Find(&configss).Error
	return configss, total, err
}
