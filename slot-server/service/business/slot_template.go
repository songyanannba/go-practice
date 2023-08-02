package business

import (
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
)

type SlotTemplateService struct {
}

// CreateSlotTemplate 创建SlotTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotTemplateService *SlotTemplateService) CreateSlotTemplate(slotTemplate business.SlotTemplate) (err error) {
	err = global.GVA_DB.Create(&slotTemplate).Error
	return err
}

// DeleteSlotTemplate 删除SlotTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotTemplateService *SlotTemplateService) DeleteSlotTemplate(slotTemplate business.SlotTemplate) (err error) {
	err = global.GVA_DB.Delete(&slotTemplate).Error
	return err
}

// DeleteSlotTemplateByIds 批量删除SlotTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotTemplateService *SlotTemplateService) DeleteSlotTemplateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.SlotTemplate{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSlotTemplate 更新SlotTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotTemplateService *SlotTemplateService) UpdateSlotTemplate(slotTemplate business.SlotTemplate) (err error) {
	err = global.GVA_DB.Save(&slotTemplate).Error
	return err
}

// GetSlotTemplate 根据id获取SlotTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotTemplateService *SlotTemplateService) GetSlotTemplate(id uint) (slotTemplate business.SlotTemplate, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&slotTemplate).Error
	return
}

// GetSlotTemplateInfoList 分页获取SlotTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotTemplateService *SlotTemplateService) GetSlotTemplateInfoList(info businessReq.SlotTemplateSearch) (list []business.SlotTemplate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.SlotTemplate{})
	var slotTemplates []business.SlotTemplate
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
	if info.Column != 0 {
		db = db.Where("column = ?", info.Column)
	}
	if info.GenId != 0 {
		db = db.Where("gen_id = ?", info.GenId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["slotId"] = true
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

	err = db.Limit(limit).Offset(offset).Find(&slotTemplates).Error
	return slotTemplates, total, err
}
