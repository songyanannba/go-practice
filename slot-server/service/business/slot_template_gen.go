package business

import (
	"go.uber.org/zap"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
	"slot-server/service/template"
)

type SlotTemplateGenService struct {
}

// CreateSlotTemplateGen 创建SlotTemplateGen记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotTemplateGenService *SlotTemplateGenService) CreateSlotTemplateGen(slotTemplateGen business.SlotTemplateGen) (err error) {
	err = global.GVA_DB.Create(&slotTemplateGen).Error
	return err
}

// DeleteSlotTemplateGen 删除SlotTemplateGen记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotTemplateGenService *SlotTemplateGenService) DeleteSlotTemplateGen(slotTemplateGen business.SlotTemplateGen) (err error) {
	err = global.GVA_DB.Delete(&slotTemplateGen).Error
	return err
}

// DeleteSlotTemplateGenByIds 批量删除SlotTemplateGen记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotTemplateGenService *SlotTemplateGenService) DeleteSlotTemplateGenByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.SlotTemplateGen{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSlotTemplateGen 更新SlotTemplateGen记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotTemplateGenService *SlotTemplateGenService) UpdateSlotTemplateGen(slotTemplateGen business.SlotTemplateGen) (err error) {
	err = global.GVA_DB.Save(&slotTemplateGen).Error
	return err
}

func (slotTemplateGenService *SlotTemplateGenService) GenerateSlotTemplateGen(slotTemplateGen *business.SlotTemplateGen) (err error) {
	go func() {
		err := template.CreateTemplate(slotTemplateGen)
		if err != nil {
			global.GVA_LOG.Error("生成模板失败!", zap.Error(err))
			slotTemplateGen.State = enum.CommonStatusError
			slotTemplateGen.Remarks = err.Error()
			global.GVA_DB.Save(slotTemplateGen)
		}
	}()
	return nil
}

// GetSlotTemplateGen 根据id获取SlotTemplateGen记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotTemplateGenService *SlotTemplateGenService) GetSlotTemplateGen(id uint) (slotTemplateGen business.SlotTemplateGen, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&slotTemplateGen).Error
	return
}

// GetSlotTemplateGenInfoList 分页获取SlotTemplateGen记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotTemplateGenService *SlotTemplateGenService) GetSlotTemplateGenInfoList(info businessReq.SlotTemplateGenSearch) (list []business.SlotTemplateGen, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.SlotTemplateGen{})
	var slotTemplateGens []business.SlotTemplateGen
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
	if info.Remarks != "" {
		db = db.Where("remarks LIKE ?", "%"+info.Remarks+"%")
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

	err = db.Limit(limit).Offset(offset).Find(&slotTemplateGens).Error
	return slotTemplateGens, total, err
}
