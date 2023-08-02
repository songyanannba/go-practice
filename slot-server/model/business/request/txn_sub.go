package request

import (
	"slot-server/model/business"
	"slot-server/model/common/request"
	"time"
)

type TxnSubSearch struct{
    business.TxnSub
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    StartBet  *float64  `json:"startBet" form:"startBet"`
    EndBet  *float64  `json:"endBet" form:"endBet"`
    StartRaise  *float64  `json:"startRaise" form:"startRaise"`
    EndRaise  *float64  `json:"endRaise" form:"endRaise"`
    StartWin  *float64  `json:"startWin" form:"startWin"`
    EndWin  *float64  `json:"endWin" form:"endWin"`
    request.PageInfo
}
