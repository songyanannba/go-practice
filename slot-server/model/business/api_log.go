// 自动生成模板ApiLog
package business

import (
	"slot-server/global"
)

// ApiLog 结构体
type ApiLog struct {
	global.GVA_MODEL
	MerchantId uint   `json:"merchantId" form:"merchantId" gorm:"column:merchant_id;default:0;comment:Merchant ID;size:32;"`
	Type       uint8  `json:"type" form:"type" gorm:"column:type;default:0;comment:Type;size:8;"`
	Agent      string `json:"agent" form:"agent" gorm:"column:agent;comment:agent;size:30;"`
	Url        string `json:"url" form:"url" gorm:"column:url;comment:Url;size:255;"`
	Method     string `json:"method" form:"method" gorm:"column:method;comment:Method;size:10;"`
	Request    string `json:"request" form:"request" gorm:"type:text;column:request;comment:Request;"`
	Response   string `json:"response" form:"response" gorm:"type:text;column:response;comment:Response;"`
	Status     uint8  `json:"status" form:"status" gorm:"column:status;default:1;comment:Status;size:8;"`
	Consume    int64  `json:"consume" form:"consume" gorm:"column:consume;default:0;comment:Consume;size:32;"` // 消耗时间，单位微秒
	Remark     string `json:"remark" form:"remark" gorm:"column:remark;type:text;comment:Remark;"`
}

// TableName ApiLog 表名
func (ApiLog) TableName() string {
	return "b_api_log"
}
