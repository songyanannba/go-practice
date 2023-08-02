package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type TxnService struct {
}

// CreateTxn 创建Txn记录
// Author [piexlmax](https://github.com/piexlmax)
func (txnService *TxnService) CreateTxn(txn business.Txn) (err error) {
	err = global.GVA_DB.Create(&txn).Error
	return err
}

// DeleteTxn 删除Txn记录
// Author [piexlmax](https://github.com/piexlmax)
func (txnService *TxnService) DeleteTxn(txn business.Txn) (err error) {
	err = global.GVA_DB.Delete(&txn).Error
	return err
}

// DeleteTxnByIds 批量删除Txn记录
// Author [piexlmax](https://github.com/piexlmax)
func (txnService *TxnService) DeleteTxnByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.Txn{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTxn 更新Txn记录
// Author [piexlmax](https://github.com/piexlmax)
func (txnService *TxnService) UpdateTxn(txn business.Txn) (err error) {
	err = global.GVA_DB.Save(&txn).Error
	return err
}

// GetTxn 根据id获取Txn记录
// Author [piexlmax](https://github.com/piexlmax)
func (txnService *TxnService) GetTxn(id uint) (txn business.Txn, err error) {
	err = global.GVA_READ_DB.Where("id = ?", id).First(&txn).Error
	return
}

// GetTxnInfoList 分页获取Txn记录
// Author [piexlmax](https://github.com/piexlmax)
func (txnService *TxnService) GetTxnInfoList(info businessReq.TxnSearch) (list []business.Txn, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_READ_DB.Model(&business.Txn{})
	var txns []business.Txn
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.ID != 0 {
		db = db.Where("id = ?", info.ID)
	}
	if info.MerchantId != 0 {
		db = db.Where("merchant_id = ?", info.MerchantId)
	}
	if info.UserId != 0 {
		db = db.Where("user_id = ?", info.UserId)
	}
	if info.GameId != 0 {
		db = db.Where("game_id = ?", info.GameId)
	}
	if info.TxnId != "" {
		db = db.Where("txn_id = ?", info.TxnId)
	}
	if info.PlatformTxnId != "" {
		db = db.Where("platform_txn_id = ?", info.PlatformTxnId)
	}
	if info.Currency != "" {
		db = db.Where("currency = ?", info.Currency)
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
	orderMap["merchantId"] = true
	orderMap["gameId"] = true
	orderMap["amount"] = true
	orderMap["beforeBal"] = true
	orderMap["afterBal"] = true
	if orderMap[info.Sort] {
		OrderStr = "`" + info.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("`id` desc")
	}

	err = db.Limit(limit).Offset(offset).Find(&txns).Error
	return txns, total, err
}
