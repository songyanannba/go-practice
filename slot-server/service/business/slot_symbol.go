package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type SlotSymbolService struct {
}

// CreateSlotSymbol 创建SlotSymbol记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotSymbolService *SlotSymbolService) CreateSlotSymbol(slotSymbol business.SlotSymbol) (err error) {
	err = global.GVA_DB.Create(&slotSymbol).Error
	return err
}

// DeleteSlotSymbol 删除SlotSymbol记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotSymbolService *SlotSymbolService) DeleteSlotSymbol(slotSymbol business.SlotSymbol) (err error) {
	err = global.GVA_DB.Delete(&slotSymbol).Error
	return err
}

// DeleteSlotSymbolByIds 批量删除SlotSymbol记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotSymbolService *SlotSymbolService) DeleteSlotSymbolByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.SlotSymbol{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSlotSymbol 更新SlotSymbol记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotSymbolService *SlotSymbolService) UpdateSlotSymbol(slotSymbol business.SlotSymbol) (err error) {
	err = global.GVA_DB.Save(&slotSymbol).Error
	return err
}

// GetSlotSymbol 根据id获取SlotSymbol记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotSymbolService *SlotSymbolService) GetSlotSymbol(id uint) (slotSymbol business.SlotSymbol, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&slotSymbol).Error
	return
}

// GetSlotSymbolInfoList 分页获取SlotSymbol记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotSymbolService *SlotSymbolService) GetSlotSymbolInfoList(info businessReq.SlotSymbolSearch) (list []business.SlotSymbol, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.SlotSymbol{})
	var slotSymbols []business.SlotSymbol
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.SlotId != 0 {
		db = db.Where("slot_id = ?", info.SlotId)
	}
	if info.IsSingleWin != 0 {
		db = db.Where("is_single_win = ?", info.IsSingleWin)
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
	orderMap["multiple"] = true
	if orderMap[info.Sort] {
		OrderStr = "`" + info.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	err = db.Limit(limit).Offset(offset).Find(&slotSymbols).Error
	return slotSymbols, total, err
}
