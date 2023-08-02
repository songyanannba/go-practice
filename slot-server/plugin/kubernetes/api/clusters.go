package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"slot-server/global"
	"slot-server/model/common/request"
	"slot-server/model/common/response"
	"slot-server/plugin/kubernetes/model"
	kubernetesReq "slot-server/plugin/kubernetes/model/kubernetes/request"
	kubernetesRes "slot-server/plugin/kubernetes/model/kubernetes/response"
	"slot-server/utils"
)

type ClustersApi struct{}

// @Tags ClustersApi
// @Summary 分页获取ClustersApi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取API列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取API列表,返回包括列表,总数,页码,每页数量"
// @Router  /kubernetes/clusters/list [post]
func (cl *ClustersApi) ClustersList(c *gin.Context) {

	var pageInfo kubernetesReq.SearchClusterParams
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customerList, total, err := clusterService.GetClustersInfoList(pageInfo.Cluster, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     customerList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags ClustersApi
// @Summary 根据id获取Clusters
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {object} response.Response{data=systemRes.SysAPIResponse} "根据id获取api,返回包括api详情"
// @Router  /kubernetes/clusters/getById [post]
func (cl *ClustersApi) GetClusterById(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cluster, err := clusterService.GetGlusterById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(kubernetesRes.ClusterResponse{Cluster: cluster}, "获取成功", c)
	}
}

// @Tags ClustersApi
// @Summary 创建基础Cluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body kubernetes.Cluster true "创建集群信息"
// @Success 200 {object} response.Response{msg=string} "创建集群信息"
// @Router  /kubernetes/clusters/create [post]
func (cl *ClustersApi) CreateCluster(c *gin.Context) {
	var cluster model.Cluster
	_ = c.ShouldBindJSON(&cluster)
	if err := clusterService.CreateCluster(cluster); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ClustersApi
// @Summary 修改基础ClustersApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body  kubernetes.Cluster true "更新"
// @Success 200 {object} response.Response{msg=string} "更新"
// @Router  /kubernetes/clusters/update [put]
func (cl *ClustersApi) UpdateCluster(c *gin.Context) {
	var cluster model.Cluster
	_ = c.ShouldBindJSON(&cluster)
	if err := clusterService.UpdateCluster(cluster); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags ClustersApi
// @Summary 删除Clusters
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body kubernetes.Cluster true "ID"
// @Success 200 {object} response.Response{msg=string} "删除Clusters"
// @Router  /kubernetes/clusters/delete [delete]
func (cl *ClustersApi) DeleteCluster(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := clusterService.DeleteCluster(idInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ClustersApi
// @Summary 删除选中集群
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID"
// @Success 200 {object} response.Response{msg=string} "删除选中集群"
// @Router  /kubernetes/clusters/deleteByIds [delete]
func (cl *ClustersApi) DeleteClustersByIds(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err := clusterService.DeleteClustersByIds(ids); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
