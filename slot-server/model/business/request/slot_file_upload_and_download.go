package request

import (
	"slot-server/model/business"
	"slot-server/model/common/request"
	"time"
)

type SlotFileUploadAndDownloadSearch struct{
    business.SlotFileUploadAndDownload
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
