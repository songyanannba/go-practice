package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type ApiLogService struct {
}

// CreateApiLog 创建ApiLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiLogService *ApiLogService) CreateApiLog(apiLog business.ApiLog) (err error) {
	err = global.GVA_DB.Create(&apiLog).Error
	return err
}

// DeleteApiLog 删除ApiLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiLogService *ApiLogService) DeleteApiLog(apiLog business.ApiLog) (err error) {
	err = global.GVA_DB.Delete(&apiLog).Error
	return err
}

// DeleteApiLogByIds 批量删除ApiLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiLogService *ApiLogService) DeleteApiLogByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.ApiLog{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateApiLog 更新ApiLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiLogService *ApiLogService) UpdateApiLog(apiLog business.ApiLog) (err error) {
	err = global.GVA_DB.Save(&apiLog).Error
	return err
}

// GetApiLog 根据id获取ApiLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiLogService *ApiLogService) GetApiLog(id uint) (apiLog business.ApiLog, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&apiLog).Error
	return
}

// GetApiLogInfoList 分页获取ApiLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiLogService *ApiLogService) GetApiLogInfoList(info businessReq.ApiLogSearch) (list []business.ApiLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.ApiLog{})
	var apiLogs []business.ApiLog
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.MerchantId != 0 {
		db = db.Where("merchant_id = ?", info.MerchantId)
	}
	if info.Agent != "" {
		db = db.Where("agent LIKE ?", "%"+info.Agent+"%")
	}
	if info.Url != "" {
		db = db.Where("url LIKE ?", "%"+info.Url+"%")
	}
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Request != "" {
		db = db.Where("request LIKE ?", "%"+info.Request+"%")
	}
	if info.Response != "" {
		db = db.Where("response LIKE ?", "%"+info.Response+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	if info.Type != 0 {
		db = db.Where("type = ?", info.Type)
	}
	if info.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+info.Remark+"%")
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

	err = db.Limit(limit).Offset(offset).Find(&apiLogs).Error

	return apiLogs, total, err
}
