package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type MoneySlotService struct {
}

// CreateMoneySlot 创建MoneySlot记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneySlotService *MoneySlotService) CreateMoneySlot(moneySlot business.MoneySlot) (err error) {
	err = global.GVA_DB.Create(&moneySlot).Error
	return err
}

// DeleteMoneySlot 删除MoneySlot记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneySlotService *MoneySlotService) DeleteMoneySlot(moneySlot business.MoneySlot) (err error) {
	err = global.GVA_DB.Delete(&moneySlot).Error
	return err
}

// DeleteMoneySlotByIds 批量删除MoneySlot记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneySlotService *MoneySlotService) DeleteMoneySlotByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.MoneySlot{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMoneySlot 更新MoneySlot记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneySlotService *MoneySlotService) UpdateMoneySlot(moneySlot business.MoneySlot) (err error) {
	err = global.GVA_DB.Save(&moneySlot).Error
	return err
}

// GetMoneySlot 根据id获取MoneySlot记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneySlotService *MoneySlotService) GetMoneySlot(id uint) (moneySlot business.MoneySlot, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&moneySlot).Error
	return
}

// GetMoneySlotInfoList 分页获取MoneySlot记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneySlotService *MoneySlotService) GetMoneySlotInfoList(info businessReq.MoneySlotSearch) (list []business.MoneySlot, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.MoneySlot{})
	var moneySlots []business.MoneySlot
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.Date != "" {
		db = db.Where("date = ?", info.Date)
	}
	if info.SlotId != 0 {
		db = db.Where("slot_id = ?", info.SlotId)
	}
	if info.StartRecentSpins != nil && info.EndRecentSpins != nil {
		db = db.Where("recent_spins BETWEEN ? AND ? ", info.StartRecentSpins, info.EndRecentSpins)
	}
	if info.StartAvgSpins != nil && info.EndAvgSpins != nil {
		db = db.Where("avg_spins BETWEEN ? AND ? ", info.StartAvgSpins, info.EndAvgSpins)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["date"] = true
	orderMap["slotId"] = true
	orderMap["coinReduce"] = true
	orderMap["coinIncrease"] = true
	orderMap["rtpRatio"] = true
	orderMap["recentPlayers"] = true
	orderMap["recentSpins"] = true
	orderMap["avgSpins"] = true
	orderMap["bkrpPeoples"] = true
	orderMap["bkrpTimes"] = true
	orderMap["bkrpAddTimes"] = true
	orderMap["bkrpAddAmount"] = true
	orderMap["topUpPeoples"] = true
	orderMap["topUpTimes"] = true
	orderMap["topUpAmount"] = true
	if orderMap[info.Sort] {
		OrderStr = "`" + info.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("`id` desc")
	}

	err = db.Limit(limit).Offset(offset).Find(&moneySlots).Error
	return moneySlots, total, err
}
