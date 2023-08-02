package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type MoneyUserSlotService struct {
}

// CreateMoneyUserSlot 创建MoneyUserSlot记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyUserSlotService *MoneyUserSlotService) CreateMoneyUserSlot(moneyUserSlot business.MoneyUserSlot) (err error) {
	err = global.GVA_DB.Create(&moneyUserSlot).Error
	return err
}

// DeleteMoneyUserSlot 删除MoneyUserSlot记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyUserSlotService *MoneyUserSlotService) DeleteMoneyUserSlot(moneyUserSlot business.MoneyUserSlot) (err error) {
	err = global.GVA_DB.Delete(&moneyUserSlot).Error
	return err
}

// DeleteMoneyUserSlotByIds 批量删除MoneyUserSlot记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyUserSlotService *MoneyUserSlotService) DeleteMoneyUserSlotByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.MoneyUserSlot{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMoneyUserSlot 更新MoneyUserSlot记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyUserSlotService *MoneyUserSlotService) UpdateMoneyUserSlot(moneyUserSlot business.MoneyUserSlot) (err error) {
	err = global.GVA_DB.Save(&moneyUserSlot).Error
	return err
}

// GetMoneyUserSlot 根据id获取MoneyUserSlot记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyUserSlotService *MoneyUserSlotService) GetMoneyUserSlot(id uint) (moneyUserSlot business.MoneyUserSlot, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&moneyUserSlot).Error
	return
}

// GetMoneyUserSlotInfoList 分页获取MoneyUserSlot记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyUserSlotService *MoneyUserSlotService) GetMoneyUserSlotInfoList(info businessReq.MoneyUserSlotSearch) (list []business.MoneyUserSlot, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.MoneyUserSlot{})
	var moneyUserSlots []business.MoneyUserSlot
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.Date != "" {
		db = db.Where("date = ?", info.Date)
	}
	if info.UserId != 0 {
		db = db.Where("user_id = ?", info.UserId)
	}
	if info.SlotId != 0 {
		db = db.Where("slot_id = ?", info.SlotId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["date"] = true
	orderMap["userId"] = true
	orderMap["slotId"] = true
	orderMap["sumBet"] = true
	orderMap["sumGain"] = true
	orderMap["countBet"] = true
	orderMap["countGain"] = true
	orderMap["countBk"] = true
	if orderMap[info.Sort] {
		OrderStr = "`" + info.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("`id` desc")
	}

	err = db.Limit(limit).Offset(offset).Find(&moneyUserSlots).Error
	return moneyUserSlots, total, err
}
