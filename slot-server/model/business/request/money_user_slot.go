package request

import (
	"slot-server/model/business"
	"slot-server/model/common/request"
	"time"
)

type MoneyUserSlotSearch struct {
	business.MoneyUserSlot
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
