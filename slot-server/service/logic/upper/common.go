package upper

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/utils"
	"slot-server/utils/helper"
)

func WriteApiLog(merchantId uint, gurl *utils.Gurl, err error) {
	var (
		remark = ""
		status = enum.Yes
	)
	if err != nil {
		remark = err.Error()
		status = enum.No
	}
	log := business.ApiLog{
		MerchantId: merchantId,
		Type:       1,
		Url:        gurl.Url,
		Method:     gurl.Method,
		Request:    helper.LeaveOutStr(string(gurl.Payload), 1000, " ..."),
		Response:   helper.LeaveOutStr(string(gurl.ResBody), 1000, " ..."),
		Status:     uint8(status),
		Remark:     helper.LeaveOutStr(remark, 150, " ..."),
		Consume:    gurl.Consuming,
	}
	err = global.GVA_DB.Create(&log).Error
	if err != nil {
		global.GVA_LOG.Error("写api日志失败", zap.Error(err), zap.Any("log", log))
	}
	return
}

func GetOrCreateUser(m *business.Merchant, playerName string, balance float64, currency string) (*business.User, error) {
	var user = &business.User{}
	err := global.GVA_DB.First(user, "merchant_id = ? and username = ?", m.ID, playerName).Error
	intBal := helper.Mul100(balance)
	if err == nil {
		user.Amount = intBal
		user.Online = enum.Yes
		user.Currency = currency
		global.GVA_DB.Select("amount", "online").Updates(user)
		return user, nil
	}
	if err != gorm.ErrRecordNotFound {
		return nil, err
	}
	user.Amount = intBal
	user.Username = playerName
	user.MerchantId = m.ID
	user.Currency = currency
	err = global.GVA_DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
