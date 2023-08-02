package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"slot-server/global"
	"slot-server/model/common/response"
	"slot-server/plugin/kubernetes/model"
	kubernetesRes "slot-server/plugin/kubernetes/model/kubernetes/response"
)

type MetricsApi struct{}

// @Tags MetricsApi
// @Summary  MetricsGet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body kubernetesReq.ResourceParamRequest true "分页获取API列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取API列表,返回包括列表,总数,页码,每页数量"
// @Router  /kubernetes/metrics/get [post]
func (m *MetricsApi) MetricsGet(c *gin.Context) {
	var metricsQuery model.MetricsQuery
	_ = c.ShouldBindJSON(&metricsQuery)
	if queryResp, err := metricService.GetMetrics(metricsQuery); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
	} else {
		response.OkWithDetailed(kubernetesRes.MetricsResponse{
			Metrics: queryResp,
		}, "获取成功", c)
	}
}
