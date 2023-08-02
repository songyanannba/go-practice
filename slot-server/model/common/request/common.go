package request

import "time"

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page        int         `json:"page" form:"page"`         // 页码
	PageSize    int         `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword     string      `json:"keyword" form:"keyword"`   //关键字
	Sort        string      `json:"sort" form:"sort"`         //排序字段
	Order       string      `json:"order" form:"order"`       //排序
	BetweenTime []time.Time `json:"betweenTime[]" form:"betweenTime[]" time_format:"20060102150405"`
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

type ModelIdReqUint struct {
	ModelId int `json:"modelId" form:"modelId"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}
