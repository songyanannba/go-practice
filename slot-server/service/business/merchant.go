package business

import (
	"errors"
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
	"slot-server/service/cache"
)

type MerchantService struct {
}

// CreateMerchant 创建Merchant记录
// Author [piexlmax](https://github.com/piexlmax)
func (merchantService *MerchantService) CreateMerchant(merchant business.Merchant) (err error) {
	if business.MerchantIsExist(&merchant) {
		return errors.New("merchant params repeat")
	}
	err = merchant.FormatCurrency()
	if err != nil {
		return err
	}
	err = global.GVA_DB.Create(&merchant).Error
	return err
}

// DeleteMerchant 删除Merchant记录
// Author [piexlmax](https://github.com/piexlmax)
func (merchantService *MerchantService) DeleteMerchant(merchant business.Merchant) (err error) {
	err = cache.DelMerchantCacheById(merchant.ID)
	if err != nil {
		return err
	}
	err = global.GVA_DB.Delete(&merchant).Error
	return err
}

// DeleteMerchantByIds 批量删除Merchant记录
// Author [piexlmax](https://github.com/piexlmax)
func (merchantService *MerchantService) DeleteMerchantByIds(ids request.IdsReq) (err error) {
	for _, id := range ids.Ids {
		err = cache.DelMerchantCacheById(uint(id))
		if err != nil {
			return err
		}
	}
	err = global.GVA_DB.Delete(&[]business.Merchant{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMerchant 更新Merchant记录
// Author [piexlmax](https://github.com/piexlmax)
func (merchantService *MerchantService) UpdateMerchant(merchant business.Merchant) (err error) {
	if business.MerchantIsExist(&merchant) {
		return errors.New("merchant params repeat")
	}
	err = cache.DelMerchantCacheById(merchant.ID)
	if err != nil {
		return err
	}
	err = merchant.FormatCurrency()
	if err != nil {
		return err
	}
	err = global.GVA_DB.Save(&merchant).Error
	return err
}

// GetMerchant 根据id获取Merchant记录
// Author [piexlmax](https://github.com/piexlmax)
func (merchantService *MerchantService) GetMerchant(id uint) (merchant business.Merchant, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&merchant).Error
	return
}

// GetMerchantInfoList 分页获取Merchant记录
// Author [piexlmax](https://github.com/piexlmax)
func (merchantService *MerchantService) GetMerchantInfoList(info businessReq.MerchantSearch) (list []business.Merchant, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.Merchant{})
	var merchants []business.Merchant
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Currency != "" {
		db = db.Where("currency = ?", info.Currency)
	}
	if info.Type != 0 {
		db = db.Where("type = ?", info.Type)
	}
	if info.ApiUrl != "" {
		db = db.Where("api_url LIKE ?", "%"+info.ApiUrl+"%")
	}
	if info.Appkey != "" {
		db = db.Where("appkey LIKE ?", "%"+info.Appkey+"%")
	}
	if info.Secret != "" {
		db = db.Where("secret LIKE ?", "%"+info.Secret+"%")
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
	orderMap["type"] = true
	if orderMap[info.Sort] {
		OrderStr = "`" + info.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("`id` desc")
	}

	err = db.Limit(limit).Offset(offset).Find(&merchants).Error
	return merchants, total, err
}
