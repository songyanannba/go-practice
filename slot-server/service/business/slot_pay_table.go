package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type SlotPayTableService struct {
}

// CreateSlotPayTable 创建SlotPayTable记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotPayTableService *SlotPayTableService) CreateSlotPayTable(slotPayTable business.SlotPayTable) (err error) {
	err = global.GVA_DB.Create(&slotPayTable).Error
	return err
}

// DeleteSlotPayTable 删除SlotPayTable记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotPayTableService *SlotPayTableService) DeleteSlotPayTable(slotPayTable business.SlotPayTable) (err error) {
	err = global.GVA_DB.Delete(&slotPayTable).Error
	return err
}

// DeleteSlotPayTableByIds 批量删除SlotPayTable记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotPayTableService *SlotPayTableService) DeleteSlotPayTableByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.SlotPayTable{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSlotPayTable 更新SlotPayTable记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotPayTableService *SlotPayTableService) UpdateSlotPayTable(slotPayTable business.SlotPayTable) (err error) {
	err = global.GVA_DB.Save(&slotPayTable).Error
	return err
}

// GetSlotPayTable 根据id获取SlotPayTable记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotPayTableService *SlotPayTableService) GetSlotPayTable(id uint) (slotPayTable business.SlotPayTable, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&slotPayTable).Error
	return
}

// GetSlotPayTableInfoList 分页获取SlotPayTable记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotPayTableService *SlotPayTableService) GetSlotPayTableInfoList(info businessReq.SlotPayTableSearch) (list []business.SlotPayTable, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.SlotPayTable{})
	var slotPayTables []business.SlotPayTable
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.SlotId != 0 {
		db = db.Where("slot_id = ?", info.SlotId)
	}
	if info.Type != 0 {
		db = db.Where("type = ?", info.Type)
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
	orderMap["slotId"] = true
	if orderMap[info.Sort] {
		OrderStr = "`" + info.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	err = db.Limit(limit).Offset(offset).Find(&slotPayTables).Error
	return slotPayTables, total, err
}
