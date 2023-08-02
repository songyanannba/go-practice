package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type JackpotService struct {
}

// CreateJackpot 创建Jackpot记录
// Author [piexlmax](https://github.com/piexlmax)
func (jackpotService *JackpotService) CreateJackpot(jackpot business.Jackpot) (err error) {
	err = global.GVA_DB.Create(&jackpot).Error
	return err
}

// DeleteJackpot 删除Jackpot记录
// Author [piexlmax](https://github.com/piexlmax)
func (jackpotService *JackpotService) DeleteJackpot(jackpot business.Jackpot) (err error) {
	err = global.GVA_DB.Delete(&jackpot).Error
	return err
}

// DeleteJackpotByIds 批量删除Jackpot记录
// Author [piexlmax](https://github.com/piexlmax)
func (jackpotService *JackpotService) DeleteJackpotByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.Jackpot{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateJackpot 更新Jackpot记录
// Author [piexlmax](https://github.com/piexlmax)
func (jackpotService *JackpotService) UpdateJackpot(jackpot business.Jackpot) (err error) {
	err = global.GVA_DB.Save(&jackpot).Error
	return err
}

// GetJackpot 根据id获取Jackpot记录
// Author [piexlmax](https://github.com/piexlmax)
func (jackpotService *JackpotService) GetJackpot(id uint) (jackpot business.Jackpot, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&jackpot).Error
	return
}

// GetJackpotInfoList 分页获取Jackpot记录
// Author [piexlmax](https://github.com/piexlmax)
func (jackpotService *JackpotService) GetJackpotInfoList(info businessReq.JackpotSearch) (list []business.Jackpot, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.Jackpot{})
	var jackpots []business.Jackpot
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.SlotId != 0 {
		ruleNos := ""
		global.GVA_DB.Model(business.Slot{}).Where("id = ?", info.SlotId).Pluck("jackpot_rule", &ruleNos)
		db = db.Where("id in (" + ruleNos + ")")
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

	err = db.Limit(limit).Offset(offset).Find(&jackpots).Error
	return jackpots, total, err
}
