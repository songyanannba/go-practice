package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type MoneyLogService struct {
}

// CreateMoneyLog 创建MoneyLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyLogService *MoneyLogService) CreateMoneyLog(moneyLog business.MoneyLog) (err error) {
	err = global.GVA_DB.Create(&moneyLog).Error
	return err
}

// DeleteMoneyLog 删除MoneyLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyLogService *MoneyLogService) DeleteMoneyLog(moneyLog business.MoneyLog) (err error) {
	err = global.GVA_DB.Delete(&moneyLog).Error
	return err
}

// DeleteMoneyLogByIds 批量删除MoneyLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyLogService *MoneyLogService) DeleteMoneyLogByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.MoneyLog{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMoneyLog 更新MoneyLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyLogService *MoneyLogService) UpdateMoneyLog(moneyLog business.MoneyLog) (err error) {
	err = global.GVA_DB.Save(&moneyLog).Error
	return err
}

// GetMoneyLog 根据id获取MoneyLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyLogService *MoneyLogService) GetMoneyLog(id uint) (moneyLog business.MoneyLog, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&moneyLog).Error
	return
}

// GetMoneyLogInfoList 分页获取MoneyLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (moneyLogService *MoneyLogService) GetMoneyLogInfoList(info businessReq.MoneyLogSearch) (list []business.MoneyLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.MoneyLog{})
	var moneyLogs []business.MoneyLog
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
	if info.Action != 0 {
		db = db.Where("action = ?", info.Action)
	}
	if info.ActionType != 0 {
		db = db.Where("action_type = ?", info.ActionType)
	}
	if info.GameId != 0 {
		db = db.Where("game_id = ?", info.GameId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["date"] = true
	orderMap["user_id"] = true
	orderMap["action"] = true
	orderMap["coin_initial"] = true
	orderMap["coin_change"] = true
	orderMap["game_id"] = true
	if orderMap[info.Sort] {
		OrderStr = "`" + info.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("`id` desc")
	}

	err = db.Limit(limit).Offset(offset).Find(&moneyLogs).Error
	return moneyLogs, total, err
}
