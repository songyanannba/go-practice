package request

import (
	"slot-server/model/business"
	"slot-server/model/common/request"
	"time"
)

type SlotUserSpinSearch struct {
	business.SlotUserSpin
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartFree      *int       `json:"startFree" form:"startFree"`
	EndFree        *int       `json:"endFree" form:"endFree"`
	StartPlayNum   *int       `json:"startPlayNum" form:"startPlayNum"`
	EndPlayNum     *int       `json:"endPlayNum" form:"endPlayNum"`
	StartFreeNum   *int       `json:"startFreeNum" form:"startFreeNum"`
	EndFreeNum     *int       `json:"endFreeNum" form:"endFreeNum"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
