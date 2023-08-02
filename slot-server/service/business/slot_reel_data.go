package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type SlotReelDataService struct {
}

// CreateSlotReelData 创建SlotReelData记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotReelDataService *SlotReelDataService) CreateSlotReelData(slotReelData business.SlotReelData) (err error) {
	err = global.GVA_DB.Create(&slotReelData).Error
	return err
}

// DeleteSlotReelData 删除SlotReelData记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotReelDataService *SlotReelDataService) DeleteSlotReelData(slotReelData business.SlotReelData) (err error) {
	err = global.GVA_DB.Delete(&slotReelData).Error
	return err
}

// DeleteSlotReelDataByIds 批量删除SlotReelData记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotReelDataService *SlotReelDataService) DeleteSlotReelDataByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.SlotReelData{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSlotReelData 更新SlotReelData记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotReelDataService *SlotReelDataService) UpdateSlotReelData(slotReelData business.SlotReelData) (err error) {
	err = global.GVA_DB.Save(&slotReelData).Error
	return err
}

// GetSlotReelData 根据id获取SlotReelData记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotReelDataService *SlotReelDataService) GetSlotReelData(id uint) (slotReelData business.SlotReelData, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&slotReelData).Error
	return
}

// GetSlotReelDataInfoList 分页获取SlotReelData记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotReelDataService *SlotReelDataService) GetSlotReelDataInfoList(info businessReq.SlotReelDataSearch) (list []business.SlotReelData, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.SlotReelData{})
	var slotReelDatas []business.SlotReelData
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.SlotId != 0 {
		db = db.Where("slot_id = ?", info.SlotId)
	}
	if info.Group != 0 {
		db = db.Where("`group` = ?", info.Group)
	}
	if info.Which != 0 {
		db = db.Where("which = ?", info.Which)
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
	orderMap["group"] = true
	orderMap["which"] = true
	if orderMap[info.Sort] {
		OrderStr = "`" + info.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("`id` desc")
	}

	err = db.Limit(limit).Offset(offset).Find(&slotReelDatas).Error
	return slotReelDatas, total, err
}
