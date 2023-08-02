package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type TxnSubService struct {
}

// CreateTxnSub 创建TxnSub记录
// Author [piexlmax](https://github.com/piexlmax)
func (txnSubService *TxnSubService) CreateTxnSub(txnSub business.TxnSub) (err error) {
	err = global.GVA_DB.Create(&txnSub).Error
	return err
}

// DeleteTxnSub 删除TxnSub记录
// Author [piexlmax](https://github.com/piexlmax)
func (txnSubService *TxnSubService) DeleteTxnSub(txnSub business.TxnSub) (err error) {
	err = global.GVA_DB.Delete(&txnSub).Error
	return err
}

// DeleteTxnSubByIds 批量删除TxnSub记录
// Author [piexlmax](https://github.com/piexlmax)
func (txnSubService *TxnSubService) DeleteTxnSubByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.TxnSub{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTxnSub 更新TxnSub记录
// Author [piexlmax](https://github.com/piexlmax)
func (txnSubService *TxnSubService) UpdateTxnSub(txnSub business.TxnSub) (err error) {
	err = global.GVA_DB.Save(&txnSub).Error
	return err
}

// GetTxnSub 根据id获取TxnSub记录
// Author [piexlmax](https://github.com/piexlmax)
func (txnSubService *TxnSubService) GetTxnSub(id uint) (txnSub business.TxnSub, err error) {
	err = global.GVA_READ_DB.Where("id = ?", id).First(&txnSub).Error
	return
}

// GetTxnSubInfoList 分页获取TxnSub记录
// Author [piexlmax](https://github.com/piexlmax)
func (txnSubService *TxnSubService) GetTxnSubInfoList(info businessReq.TxnSubSearch) (list []business.TxnSub, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_READ_DB.Model(&business.TxnSub{})
	var txnSubs []business.TxnSub
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.MerchantId != 0 {
		db = db.Where("merchant_id = ?", info.MerchantId)
	}
	if info.Pid != 0 {
		db = db.Where("pid = ?", info.Pid)
	}
	if info.RecordId != 0 {
		db = db.Where("record_id = ?", info.RecordId)
	}
	if info.Type != 0 {
		db = db.Where("type = ?", info.Type)
	}
	if info.StartBet != nil && info.EndBet != nil {
		db = db.Where("bet BETWEEN ? AND ? ", info.StartBet, info.EndBet)
	}
	if info.StartRaise != nil && info.EndRaise != nil {
		db = db.Where("raise BETWEEN ? AND ? ", info.StartRaise, info.EndRaise)
	}
	if info.StartWin != nil && info.EndWin != nil {
		db = db.Where("win BETWEEN ? AND ? ", info.StartWin, info.EndWin)
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

	err = db.Limit(limit).Offset(offset).Find(&txnSubs).Error
	return txnSubs, total, err
}
