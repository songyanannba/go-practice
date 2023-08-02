package request

import (
	"slot-server/model/business"
	"slot-server/model/common/request"
	"time"
)

type TxnSearch struct {
	business.Txn
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
