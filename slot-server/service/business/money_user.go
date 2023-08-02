package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type MoneyUserService struct {
}

// CreateMoneyUser 创建MoneyUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDailySumService *MoneyUserService) CreateMoneyUser(userDailySum business.MoneyUser) (err error) {
	err = global.GVA_DB.Create(&userDailySum).Error
	return err
}

// DeleteMoneyUser 删除MoneyUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDailySumService *MoneyUserService) DeleteMoneyUser(userDailySum business.MoneyUser) (err error) {
	err = global.GVA_DB.Delete(&userDailySum).Error
	return err
}

// DeleteMoneyUserByIds 批量删除MoneyUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDailySumService *MoneyUserService) DeleteMoneyUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.MoneyUser{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMoneyUser 更新MoneyUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDailySumService *MoneyUserService) UpdateMoneyUser(userDailySum business.MoneyUser) (err error) {
	err = global.GVA_DB.Save(&userDailySum).Error
	return err
}

// GetMoneyUser 根据id获取MoneyUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDailySumService *MoneyUserService) GetMoneyUser(id uint) (userDailySum business.MoneyUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userDailySum).Error
	return
}

// GetMoneyUserInfoList 分页获取MoneyUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDailySumService *MoneyUserService) GetMoneyUserInfoList(info businessReq.MoneyUserSearch) (list []business.MoneyUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.MoneyUser{})
	var userDailySums []business.MoneyUser
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
	if info.UserName != "" {
		db = db.Where("user_name LIKE ?", "%"+info.UserName+"%")
	}
	if info.StartBetAmount != nil && info.EndBetAmount != nil {
		db = db.Where("bet_amount BETWEEN ? AND ? ", info.StartBetAmount, info.EndBetAmount)
	}
	if info.StartGainAmount != nil && info.EndGainAmount != nil {
		db = db.Where("gain_amount BETWEEN ? AND ? ", info.StartGainAmount, info.EndGainAmount)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["betCommon"] = true
	orderMap["betAmount"] = true
	orderMap["gainAmount"] = true
	orderMap["lastStand"] = true
	orderMap["date"] = true
	if orderMap[info.Sort] {
		OrderStr = "`" + info.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("`id` desc")
	}

	err = db.Limit(limit).Offset(offset).Find(&userDailySums).Error
	return userDailySums, total, err
}
