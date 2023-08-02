package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type SlotEventService struct {
}

// CreateSlotEvent 创建SlotEvent记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotEventService *SlotEventService) CreateSlotEvent(slotEvent business.SlotEvent) (err error) {
	err = global.GVA_DB.Create(&slotEvent).Error
	return err
}

// DeleteSlotEvent 删除SlotEvent记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotEventService *SlotEventService) DeleteSlotEvent(slotEvent business.SlotEvent) (err error) {
	err = global.GVA_DB.Delete(&slotEvent).Error
	return err
}

// DeleteSlotEventByIds 批量删除SlotEvent记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotEventService *SlotEventService) DeleteSlotEventByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.SlotEvent{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSlotEvent 更新SlotEvent记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotEventService *SlotEventService) UpdateSlotEvent(slotEvent business.SlotEvent) (err error) {
	err = global.GVA_DB.Save(&slotEvent).Error
	return err
}

// GetSlotEvent 根据id获取SlotEvent记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotEventService *SlotEventService) GetSlotEvent(id uint) (slotEvent business.SlotEvent, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&slotEvent).Error
	return
}

// GetSlotEventInfoList 分页获取SlotEvent记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotEventService *SlotEventService) GetSlotEventInfoList(info businessReq.SlotEventSearch) (list []business.SlotEvent, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.SlotEvent{})
	var slotEvents []business.SlotEvent
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.SlotId != 0 {
		db = db.Where("slot_id = ?", info.SlotId)
	}
	if info.Demo != 0 {
		db = db.Where("demo = ?", info.Demo)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["slotId"] = true
	if orderMap[info.Sort] {
		OrderStr = "`" + info.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	err = db.Limit(limit).Offset(offset).Find(&slotEvents).Error
	return slotEvents, total, err
}
