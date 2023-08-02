package request

type ProxyParamRequest struct {
	ClusterId int    `uri:"cluster_id" binding:"required"`
	Path      string `uri:"path" binding:"required"`
}

type ProxyResult struct {
	Items    interface{} `json:"items"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type ProxyRequest struct {
	Search        string `form:"search" binding:"required"`
	Namespace     string `form:"namespace" binding:"required"`
	Keywords      string `form:"keywords" binding:"required"`
	FieldSelector string `form:"fieldSelector" binding:"required"`
	LabelSelector string `form:"labelSelector" binding:"required"`
	Page          int    `form:"page" binding:"required"`
	PageSize      int    `form:"pageSize" binding:"required"`
}
