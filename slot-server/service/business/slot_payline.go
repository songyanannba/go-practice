package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type SlotPaylineService struct {
}

// CreateSlotPayline 创建SlotPayline记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotPaylineService *SlotPaylineService) CreateSlotPayline(slotPayline business.SlotPayline) (err error) {
	err = global.GVA_DB.Create(&slotPayline).Error
	return err
}

// DeleteSlotPayline 删除SlotPayline记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotPaylineService *SlotPaylineService) DeleteSlotPayline(slotPayline business.SlotPayline) (err error) {
	err = global.GVA_DB.Delete(&slotPayline).Error
	return err
}

// DeleteSlotPaylineByIds 批量删除SlotPayline记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotPaylineService *SlotPaylineService) DeleteSlotPaylineByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.SlotPayline{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSlotPayline 更新SlotPayline记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotPaylineService *SlotPaylineService) UpdateSlotPayline(slotPayline business.SlotPayline) (err error) {
	err = global.GVA_DB.Save(&slotPayline).Error
	return err
}

// GetSlotPayline 根据id获取SlotPayline记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotPaylineService *SlotPaylineService) GetSlotPayline(id uint) (slotPayline business.SlotPayline, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&slotPayline).Error
	return
}

// GetSlotPaylineInfoList 分页获取SlotPayline记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotPaylineService *SlotPaylineService) GetSlotPaylineInfoList(info businessReq.SlotPaylineSearch) (list []business.SlotPayline, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.SlotPayline{})
	var slotPaylines []business.SlotPayline
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.ID != 0 {
		db = db.Where("no = ?", info.No)
	}
	if info.Num != 0 {
		db = db.Where("num = ?", info.Num)
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["no"] = true
	orderMap["sorted"] = true
	if orderMap[info.PageInfo.Sort] {
		OrderStr = "`" + info.PageInfo.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	err = db.Limit(limit).Offset(offset).Find(&slotPaylines).Error
	return slotPaylines, total, err
}
