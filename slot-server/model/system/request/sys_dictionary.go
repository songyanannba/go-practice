package request

import (
	"slot-server/model/common/request"
	"slot-server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
