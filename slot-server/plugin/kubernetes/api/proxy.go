package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"slot-server/global"
	"slot-server/model/common/response"
	kubernetesReq "slot-server/plugin/kubernetes/model/kubernetes/request"
	"slot-server/plugin/kubernetes/utils"
	sutils "slot-server/utils"
)

type ProxyApi struct{}

// @Tags ProxyApi
// @Summary 分页获取ProxyApi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router  /kubernetes/proxy/:cluster_id/*path [Any]
func (pr *ProxyApi) K8sAPIProxy(c *gin.Context) {
	// 解析参数
	var proxy kubernetesReq.ProxyRequest
	_ = c.ShouldBindQuery(&proxy)
	var urlParam kubernetesReq.ProxyParamRequest
	_ = c.BindUri(&urlParam)

	//参数校验
	if err := sutils.Verify(urlParam, utils.ProxyVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	ret, err := proxyService.Option(c, proxy, urlParam)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
	} else {
		response.OkWithDetailed(ret, "获取成功", c)
	}

	return
}
