package request

import (
	"slot-server/model/business"
	"slot-server/model/common/request"
	"time"
)

type UserSearch struct {
	business.User
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
