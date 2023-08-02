package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type TrackingService struct {
}

// CreateTracking 创建Tracking记录
// Author [piexlmax](https://github.com/piexlmax)
func (trackingService *TrackingService) CreateTracking(tracking business.Tracking) (err error) {
	err = global.GVA_DB.Create(&tracking).Error
	return err
}

// DeleteTracking 删除Tracking记录
// Author [piexlmax](https://github.com/piexlmax)
func (trackingService *TrackingService) DeleteTracking(tracking business.Tracking) (err error) {
	err = global.GVA_DB.Delete(&tracking).Error
	return err
}

// DeleteTrackingByIds 批量删除Tracking记录
// Author [piexlmax](https://github.com/piexlmax)
func (trackingService *TrackingService) DeleteTrackingByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.Tracking{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTracking 更新Tracking记录
// Author [piexlmax](https://github.com/piexlmax)
func (trackingService *TrackingService) UpdateTracking(tracking business.Tracking) (err error) {
	err = global.GVA_DB.Save(&tracking).Error
	return err
}

// GetTracking 根据id获取Tracking记录
// Author [piexlmax](https://github.com/piexlmax)
func (trackingService *TrackingService) GetTracking(id uint) (tracking business.Tracking, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&tracking).Error
	return
}

// GetTrackingInfoList 分页获取Tracking记录
// Author [piexlmax](https://github.com/piexlmax)
func (trackingService *TrackingService) GetTrackingInfoList(info businessReq.TrackingSearch) (list []business.Tracking, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.Tracking{})
	var trackings []business.Tracking
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.Date != "" {
		db = db.Where("date = ?", info.Date)
	}
	if info.Type != 0 {
		db = db.Where("type = ?", info.Type)
	}
	if info.UserId != 0 {
		db = db.Where("user_id = ?", info.UserId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["date"] = true
	orderMap["type"] = true
	orderMap["userId"] = true
	orderMap["num"] = true
	if orderMap[info.Sort] {
		OrderStr = "`" + info.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("`id` desc")
	}

	err = db.Limit(limit).Offset(offset).Find(&trackings).Error
	return trackings, total, err
}
