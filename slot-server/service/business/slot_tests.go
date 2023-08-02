package business

import (
	"errors"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
	"slot-server/service/slot/component"
	"slot-server/service/test"
	"slot-server/service/test/public"
	"time"
)

type SlotTestsService struct {
}

const slotTestLock = "slot_test_lock"

// CreateSlotTests 创建SlotTests记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotTestsService *SlotTestsService) CreateSlotTests(run public.RunSlotTest) (err error) {
	opts := []component.Option{component.WithTest(), component.WithNeedSpecify(true)}
	opts = append(opts, component.SetDebugConfig(0))
	for _, opt := range run.Opts {
		switch opt {
		case 1:
			opts = append(opts)
		case 2:
			opts = append(opts, component.WithDemo())
		}
	}

	switch run.Type {
	case enum.SlotTestType3Once, enum.SlotTestType4Result:
		return test.Once(run, opts...)
	case enum.SlotTestType6User:
		return test.User(run, opts...)
	}
	_, err = component.GetSlotConfig(run.SlotId, false)
	if err != nil {
		return errors.New("构建slot配置失败: " + err.Error())
	}
	_, ok := global.BlackCache.Get(slotTestLock)
	if ok {
		return errors.New("测试进行中，稍后再试")
	}
	slotTest := business.SlotTests{
		Type:   uint8(run.Type),
		SlotId: run.SlotId,
		Hold:   run.Hold,
		Amount: run.Amount,
		Win:    0,
		MaxNum: run.Num,
		RunNum: 0,
		Detail: "",
		Status: enum.CommonStatusBegin,
	}
	err = global.GVA_DB.Create(&slotTest).Error
	if err != nil {
		return err
	}
	global.BlackCache.Set(slotTestLock, 1, time.Minute*30)
	switch run.Type {
	case enum.SlotTestType2Die:
		go test.Death(&slotTest, run, opts...)
	case enum.SlotTestType1Time:
		go test.Appoint(&slotTest, run.SlotId, run.Num, run.Amount, run, opts...)
	default:
		return errors.New("测试类型错误")
	}
	return err
}

// DeleteSlotTests 删除SlotTests记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotTestsService *SlotTestsService) DeleteSlotTests(slotTests business.SlotTests) (err error) {
	err = global.GVA_DB.Delete(&slotTests).Error
	return err
}

// DeleteSlotTestsByIds 批量删除SlotTests记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotTestsService *SlotTestsService) DeleteSlotTestsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.SlotTests{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSlotTests 更新SlotTests记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotTestsService *SlotTestsService) UpdateSlotTests(slotTests business.SlotTests) (err error) {
	err = global.GVA_DB.Save(&slotTests).Error
	return err
}

// GetSlotTests 根据id获取SlotTests记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotTestsService *SlotTestsService) GetSlotTests(id uint) (slotTests business.SlotTests, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&slotTests).Error
	return
}

// GetSlotTestsInfoList 分页获取SlotTests记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotTestsService *SlotTestsService) GetSlotTestsInfoList(info businessReq.SlotTestsSearch) (list []business.SlotTests, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.SlotTests{})
	var slotTestss []business.SlotTests
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
	if info.Amount != 0 {
		db = db.Where("amount > ?", info.Amount)
	}
	if info.Win != 0 {
		db = db.Where("win > ?", info.Win)
	}
	if info.TestId != 0 {
		db = db.Where("id = ? or test_id = ?", info.TestId, info.TestId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["amount"] = true
	orderMap["win"] = true
	orderMap["maxNum"] = true
	orderMap["runNum"] = true
	orderMap["hold"] = true
	orderMap["type"] = true
	if orderMap[info.Sort] {
		OrderStr = "`" + info.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("id desc")
	}

	err = db.Limit(limit).Offset(offset).Find(&slotTestss).Error
	return slotTestss, total, err
}
