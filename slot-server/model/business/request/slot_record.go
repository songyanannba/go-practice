package request

import (
	"slot-server/model/business"
	"slot-server/model/common/request"
	"time"
)

type SlotRecordSearch struct {
	business.SlotRecord
	DateChoice     *string    `json:"dateChoice" form:"dateChoice"`
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartGain      *int       `json:"startGain" form:"startGain"`
	EndGain        *int       `json:"endGain" form:"endGain"`
	TxnNo          string     `json:"txnNo" form:"txnNo"`
	request.PageInfo
}

type SlotRecordPublicSearch struct {
	No string `json:"no" form:"no"`
}
