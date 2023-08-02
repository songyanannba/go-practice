package request

import (
	"slot-server/model/common/request"
	"slot-server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
