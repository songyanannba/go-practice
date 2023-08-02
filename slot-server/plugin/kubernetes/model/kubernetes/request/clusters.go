package request

import (
	"slot-server/model/common/request"
	"slot-server/plugin/kubernetes/model"
)

type SearchClusterParams struct {
	model.Cluster
	request.PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}
