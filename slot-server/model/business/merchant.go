// 自动生成模板Merchant
package business

import (
	"errors"
	"slot-server/global"
	"slot-server/utils"
	"slot-server/utils/helper"
	"strings"
)

// Merchant 结构体
type Merchant struct {
	global.GVA_MODEL
	Agent      string `json:"agent" form:"agent" gorm:"column:agent;comment:agent;size:30;"`
	Name       string `json:"name" form:"name" gorm:"column:name;comment:name;size:255;"`
	Currency   string `json:"currency" form:"currency" gorm:"column:currency;comment:currency;size:30;"`
	Type       uint8  `json:"type" form:"type" gorm:"column:type;default:1;comment:type;size:8;"`
	ApiUrl     string `json:"apiUrl" form:"apiUrl" gorm:"column:api_url;comment:api_url;size:255;"`
	Appkey     string `json:"appkey" form:"appkey" gorm:"column:appkey;comment:appKey;size:50;"`
	Secret     string `json:"secret" form:"secret" gorm:"column:secret;comment:secret;size:50;"`
	ProviderId string `json:"providerId" form:"providerId" gorm:"column:provider_id;comment:provider_id;size:50;"`
	Status     uint8  `json:"status" form:"status" gorm:"column:status;default:1;comment:status;size:8;"`
}

// TableName Merchant 表名
func (Merchant) TableName() string {
	return "b_merchant"
}

func MerchantIsExist(merchant *Merchant) bool {
	var count int64
	q := global.GVA_DB.Table(merchant.TableName()).Select("id").
		Where("(name = ? || agent = ? || appkey = ?)",
			merchant.Name, merchant.Agent, merchant.Appkey)
	if merchant.ID > 0 {
		q.Where("id != ?", merchant.ID)
	}
	q.Count(&count)
	if count > 0 {
		return true
	}
	return false
}

func (m *Merchant) FormatCurrency() error {
	m.Currency = utils.FormatCommandStr(m.Currency)
	if m.Currency == "" {
		return errors.New("currency is empty")
	}
	m.Currency = strings.ToUpper(m.Currency)
	return nil
}

// CheckCurrency 检查商户是否支持该币种
func (m Merchant) CheckCurrency(currency string) bool {
	supportCurrency := strings.Split(m.Currency, " ")
	return helper.CaseInsensitiveInArr(currency, supportCurrency)
}
