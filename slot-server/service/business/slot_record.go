package business

import (
	"google.golang.org/protobuf/proto"
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
	"slot-server/pbs"
	"slot-server/service/logic"
)

type SlotRecordService struct {
}

// CreateSlotRecord 创建SlotRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (RecordService *SlotRecordService) CreateSlotRecord(Record business.SlotRecord) (err error) {
	err = global.GVA_DB.Create(&Record).Error
	return err
}

// DeleteSlotRecord 删除SlotRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (RecordService *SlotRecordService) DeleteSlotRecord(Record business.SlotRecord) (err error) {
	err = global.GVA_DB.Delete(&Record).Error
	return err
}

// DeleteSlotRecordByIds 批量删除SlotRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (RecordService *SlotRecordService) DeleteSlotRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.SlotRecord{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSlotRecord 更新SlotRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (RecordService *SlotRecordService) UpdateSlotRecord(Record business.SlotRecord) (err error) {
	err = global.GVA_DB.Save(&Record).Error
	return err
}

// GetSlotRecord 根据id获取SlotRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (RecordService *SlotRecordService) GetSlotRecord(id uint) (Record business.SlotRecord, err error) {
	err = global.GVA_READ_DB.Where("id = ?", id).First(&Record).Error
	if Record.SlotId < 5 {
		ack := pbs.SpinAck{}
		proto.Unmarshal(Record.Ack, &ack)
		Record.AckStr = logic.DumpSpinAck(&ack)
		Record.Ack = nil
	}
	return
}

// GetSlotRecordInfoList 分页获取SlotRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (RecordService *SlotRecordService) GetSlotRecordInfoList(info businessReq.SlotRecordSearch) (list []business.SlotRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_READ_DB.Model(&business.SlotRecord{})
	var records []business.SlotRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	//if info.ResultData != "" {
	//	db = db.Where("result_data LIKE ?", "%"+info.ResultData+"%")
	//}
	if info.MerchantId != 0 {
		db = db.Where("merchant_id = ?", info.MerchantId)
	}
	if info.JackpotId != 0 {
		db = db.Where("jackpot_id = ?", info.JackpotId)
	}
	if info.StartGain != nil && info.EndGain != nil {
		db = db.Where("gain BETWEEN ? AND ? ", info.StartGain, info.EndGain)
	}
	if info.DateChoice != nil {
		db = db.Where("date = ? ", info.DateChoice)
	}
	if info.UserId != 0 {
		db = db.Where("user_id = ? ", info.UserId)
	}
	if info.SlotId != 0 {
		db = db.Where("slot_id = ? ", info.SlotId)
	}
	if info.ID != 0 {
		db = db.Where("id = ? ", info.ID)
	}
	if info.TxnId != 0 {
		db = db.Where("txn_id = ? ", info.TxnId)
	}
	if info.TxnNo != "" {
		var txnId int
		global.GVA_READ_DB.Model(&business.Txn{}).Where("txn_id = ?", info.TxnNo).Pluck("id", &txnId)
		db = db.Where("txn_id = ? ", txnId)
	}
	//if info.PlayType != 0 {
	//	db = db.Where("play_type = ? ", info.PlayType)
	//}
	if info.Status != 0 {
		db = db.Where("status = ? ", info.Status)
	}
	if info.No != "" {
		id := business.ParseSlotRecordId(info.No)
		if id == 0 {
			return
		}
		db = db.Where(id)
	}
	//if info.MoneyLogId != 0 {
	//	db = db.Where("money_log_id = ? ", info.MoneyLogId)
	//}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
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

	err = db.Limit(limit).Offset(offset).Omit("ack").Find(&records).Error
	for i, record := range records {
		records[i].No = business.FmtSlotRecordNo(record.ID)
	}
	return records, total, err
}
