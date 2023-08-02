package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type SlotUserSpinService struct {
}

// CreateSlotUserSpin 创建SlotUserSpin记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotUserSpinService *SlotUserSpinService) CreateSlotUserSpin(slotUserSpin business.SlotUserSpin) (err error) {
	err = global.GVA_DB.Create(&slotUserSpin).Error
	return err
}

// DeleteSlotUserSpin 删除SlotUserSpin记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotUserSpinService *SlotUserSpinService) DeleteSlotUserSpin(slotUserSpin business.SlotUserSpin) (err error) {
	err = global.GVA_DB.Delete(&slotUserSpin).Error
	return err
}

// DeleteSlotUserSpinByIds 批量删除SlotUserSpin记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotUserSpinService *SlotUserSpinService) DeleteSlotUserSpinByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.SlotUserSpin{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSlotUserSpin 更新SlotUserSpin记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotUserSpinService *SlotUserSpinService) UpdateSlotUserSpin(slotUserSpin business.SlotUserSpin) (err error) {
	err = global.GVA_DB.Save(&slotUserSpin).Error
	return err
}

// GetSlotUserSpin 根据id获取SlotUserSpin记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotUserSpinService *SlotUserSpinService) GetSlotUserSpin(id uint) (slotUserSpin business.SlotUserSpin, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&slotUserSpin).Error
	return
}

// GetSlotUserSpinInfoList 分页获取SlotUserSpin记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotUserSpinService *SlotUserSpinService) GetSlotUserSpinInfoList(info businessReq.SlotUserSpinSearch) (list []business.SlotUserSpin, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.SlotUserSpin{})
	var slotUserSpins []business.SlotUserSpin
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.SlotId != 0 {
		db = db.Where("slot_id = ?", info.SlotId)
	}
	if info.UserId != 0 {
		db = db.Where("user_id = ?", info.UserId)
	}
	if info.StartFree != nil && info.EndFree != nil {
		db = db.Where("free BETWEEN ? AND ? ", info.StartFree, info.EndFree)
	}
	if info.StartPlayNum != nil && info.EndPlayNum != nil {
		db = db.Where("play_num BETWEEN ? AND ? ", info.StartPlayNum, info.EndPlayNum)
	}
	if info.StartFreeNum != nil && info.EndFreeNum != nil {
		db = db.Where("free_num BETWEEN ? AND ? ", info.StartFreeNum, info.EndFreeNum)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["userId"] = true
	orderMap["slotId"] = true
	orderMap["free"] = true
	orderMap["playNum"] = true
	orderMap["freeNum"] = true
	if orderMap[info.Sort] {
		OrderStr = "`" + info.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("`id` desc")
	}

	err = db.Limit(limit).Offset(offset).Find(&slotUserSpins).Error
	return slotUserSpins, total, err
}
