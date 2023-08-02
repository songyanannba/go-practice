package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type CurrencyService struct {
}

// CreateCurrency 创建Currency记录
// Author [piexlmax](https://github.com/piexlmax)
func (currencyService *CurrencyService) CreateCurrency(currency business.Currency) (err error) {
	err = global.GVA_DB.Create(&currency).Error
	return err
}

// DeleteCurrency 删除Currency记录
// Author [piexlmax](https://github.com/piexlmax)
func (currencyService *CurrencyService) DeleteCurrency(currency business.Currency) (err error) {
	err = global.GVA_DB.Delete(&currency).Error
	return err
}

// DeleteCurrencyByIds 批量删除Currency记录
// Author [piexlmax](https://github.com/piexlmax)
func (currencyService *CurrencyService) DeleteCurrencyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.Currency{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCurrency 更新Currency记录
// Author [piexlmax](https://github.com/piexlmax)
func (currencyService *CurrencyService) UpdateCurrency(currency business.Currency) (err error) {
	err = global.GVA_DB.Save(&currency).Error
	return err
}

// GetCurrency 根据id获取Currency记录
// Author [piexlmax](https://github.com/piexlmax)
func (currencyService *CurrencyService) GetCurrency(id uint) (currency business.Currency, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&currency).Error
	return
}

// GetCurrencyInfoList 分页获取Currency记录
// Author [piexlmax](https://github.com/piexlmax)
func (currencyService *CurrencyService) GetCurrencyInfoList(info businessReq.CurrencySearch) (list []business.Currency, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.Currency{})
	var currencys []business.Currency
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
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
	if orderMap[info.Sort] {
		OrderStr = "`" + info.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("`id` desc")
	}

	err = db.Limit(limit).Offset(offset).Find(&currencys).Error
	return currencys, total, err
}
