package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type SlotReelService struct {
}

// CreateSlotReel 创建SlotReel记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotReelService *SlotReelService) CreateSlotReel(slotReel business.SlotReel) (err error) {
	err = global.GVA_DB.Create(&slotReel).Error
	return err
}

// DeleteSlotReel 删除SlotReel记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotReelService *SlotReelService) DeleteSlotReel(slotReel business.SlotReel) (err error) {
	err = global.GVA_DB.Delete(&slotReel).Error
	return err
}

// DeleteSlotReelByIds 批量删除SlotReel记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotReelService *SlotReelService) DeleteSlotReelByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.SlotReel{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSlotReel 更新SlotReel记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotReelService *SlotReelService) UpdateSlotReel(slotReel business.SlotReel) (err error) {
	err = global.GVA_DB.Save(&slotReel).Error
	return err
}

// GetSlotReel 根据id获取SlotReel记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotReelService *SlotReelService) GetSlotReel(id uint) (slotReel business.SlotReel, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&slotReel).Error
	return
}

// GetSlotReelInfoList 分页获取SlotReel记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotReelService *SlotReelService) GetSlotReelInfoList(info businessReq.SlotReelSearch) (list []business.SlotReel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.SlotReel{})
	var slotReels []business.SlotReel
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
	if orderMap[info.Sort] {
		OrderStr = "`" + info.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	err = db.Limit(limit).Offset(offset).Find(&slotReels).Error
	return slotReels, total, err
}
