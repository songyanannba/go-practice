package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type SlotFakeService struct {
}

// CreateSlotFake 创建SlotFake记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotFakeService *SlotFakeService) CreateSlotFake(slotFake business.SlotFake) (err error) {
	err = global.GVA_DB.Create(&slotFake).Error
	return err
}

// DeleteSlotFake 删除SlotFake记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotFakeService *SlotFakeService) DeleteSlotFake(slotFake business.SlotFake) (err error) {
	err = global.GVA_DB.Delete(&slotFake).Error
	return err
}

// DeleteSlotFakeByIds 批量删除SlotFake记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotFakeService *SlotFakeService) DeleteSlotFakeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.SlotFake{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSlotFake 更新SlotFake记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotFakeService *SlotFakeService) UpdateSlotFake(slotFake business.SlotFake) (err error) {
	err = global.GVA_DB.Save(&slotFake).Error
	return err
}

// GetSlotFake 根据id获取SlotFake记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotFakeService *SlotFakeService) GetSlotFake(id uint) (slotFake business.SlotFake, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&slotFake).Error
	return
}

// GetSlotFakeInfoList 分页获取SlotFake记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotFakeService *SlotFakeService) GetSlotFakeInfoList(info businessReq.SlotFakeSearch) (list []business.SlotFake, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.SlotFake{})
	var slotFakes []business.SlotFake
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.Type != 0 {
		db = db.Where("type = ?", info.Type)
	}
	if info.SlotId != 0 {
		db = db.Where("slot_id = ?", info.SlotId)
	}
	if info.Num != 0 {
		db = db.Where("num = ?", info.Num)
	}
	if info.Position != "" {
		db = db.Where("position LIKE ?", "%"+info.Position+"%")
	}
	if info.Which != 0 {
		db = db.Where("which = ?", info.Which)
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

	err = db.Limit(limit).Offset(offset).Find(&slotFakes).Error
	return slotFakes, total, err
}
