package request

import (
	"slot-server/model/business"
	"slot-server/model/common/request"
	"time"
)

type SlotTestsSearch struct {
	business.SlotTests
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
