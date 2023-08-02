package request

import (
	"slot-server/model/business"
	"slot-server/model/common/request"
	"time"
)

type JackpotSearch struct {
	business.Jackpot
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
	SlotId int `json:"slotId" form:"slotId"`
}
