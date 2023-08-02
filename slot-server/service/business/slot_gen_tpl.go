package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
	"slot-server/service/business/slotTpl"
)

type SlotGenTplService struct {
}

// CreateSlotGenTpl 创建SlotGenTpl记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotGenTplService *SlotGenTplService) CreateSlotGenTpl(slotGenTpl business.SlotGenTpl) (err error) {
	err = slotTpl.Run(&slotGenTpl)
	if err != nil {
		return
	}

	err = global.GVA_DB.Create(&slotGenTpl).Error
	return err
}

// DeleteSlotGenTpl 删除SlotGenTpl记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotGenTplService *SlotGenTplService) DeleteSlotGenTpl(slotGenTpl business.SlotGenTpl) (err error) {
	err = global.GVA_DB.Delete(&slotGenTpl).Error
	return err
}

// DeleteSlotGenTplByIds 批量删除SlotGenTpl记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotGenTplService *SlotGenTplService) DeleteSlotGenTplByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.SlotGenTpl{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSlotGenTpl 更新SlotGenTpl记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotGenTplService *SlotGenTplService) UpdateSlotGenTpl(slotGenTpl business.SlotGenTpl) (err error) {
	err = global.GVA_DB.Save(&slotGenTpl).Error
	return err
}

// GetSlotGenTpl 根据id获取SlotGenTpl记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotGenTplService *SlotGenTplService) GetSlotGenTpl(id uint) (slotGenTpl business.SlotGenTpl, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&slotGenTpl).Error
	return
}

// GetSlotGenTplInfoList 分页获取SlotGenTpl记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotGenTplService *SlotGenTplService) GetSlotGenTplInfoList(info businessReq.SlotGenTplSearch) (list []business.SlotGenTpl, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.SlotGenTpl{})
	var slotGenTpls []business.SlotGenTpl
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.Type != 0 {
		db = db.Where("type = ?", info.Type)
	}
	if info.SlotId != 0 {
		db = db.Where("slot_id = ?", info.SlotId)
	}
	if info.Size != "" {
		db = db.Where("size LIKE ?", "%"+info.Size+"%")
	}
	if info.Num != 0 {
		db = db.Where("num = ?", info.Num)
	}
	if info.Params != "" {
		db = db.Where("params LIKE ?", "%"+info.Params+"%")
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

	err = db.Limit(limit).Offset(offset).Find(&slotGenTpls).Error
	return slotGenTpls, total, err
}
