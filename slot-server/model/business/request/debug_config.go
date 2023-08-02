package request

import (
	"slot-server/model/business"
	"slot-server/model/common/request"
	"time"
)

type DebugConfigSearch struct {
	business.DebugConfig
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
