package request

import (
	"slot-server/model/common/request"
	"slot-server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
