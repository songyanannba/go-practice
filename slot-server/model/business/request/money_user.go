package request

import (
	"slot-server/model/business"
	"slot-server/model/common/request"
	"time"
)

type MoneyUserSearch struct {
	business.MoneyUser
	StartCreatedAt  *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt    *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartBetAmount  *int       `json:"startBetAmount" form:"startBetAmount"`
	EndBetAmount    *int       `json:"endBetAmount" form:"endBetAmount"`
	StartGainAmount *int       `json:"startGainAmount" form:"startGainAmount"`
	EndGainAmount   *int       `json:"endGainAmount" form:"endGainAmount"`
	request.PageInfo
}
