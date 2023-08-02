package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"slot-server/global"
	"slot-server/model/common/response"
	kubernetesReq "slot-server/plugin/kubernetes/model/kubernetes/request"
	"slot-server/plugin/kubernetes/utils"
	sutils "slot-server/utils"
)

type WsApi struct{}

// 终端token验证
func (w *WsApi) TerminalJWTAuth(terminal kubernetesReq.TerminalRequest) (err error) {
	j := sutils.NewJWT()
	claims, err := j.ParseToken(terminal.XToken)
	if err != nil {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}

	if claims.AuthorityId == 0 {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析角色信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return err
}

// @Tags WsApi
// @Summary  终端
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body kubernetesReq.TerminalRequest true "根据id获取api"
// @Success 200 {object} response.Response{data=systemRes.SysAPIResponse} "根据id获取api,返回包括api详情"
// @Router  /kubernetes/pods/terminal [get]
func (w *WsApi) Terminal(c *gin.Context) {
	var terminal kubernetesReq.TerminalRequest
	_ = c.ShouldBindQuery(&terminal)
	if err := w.TerminalJWTAuth(terminal); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	if err := sutils.Verify(terminal, utils.TerminalVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	client, err := utils.NewKubeClient(terminal.ClusterId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("连接k8s client失败!，%v", err.Error()), c)
	}

	//校验pod
	kubeshell, err := utils.NewKubeShell(c.Writer, c.Request, nil)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	cmd := []string{
		"/bin/sh", "-c", fmt.Sprintf("clear;(bash || sh); export LINES=%d ; export COLUMNS=%d;", terminal.Rows, terminal.Cols),
	}
	if err := client.Pod.Exec(cmd, kubeshell, terminal.Namespace, terminal.PodName, terminal.Name); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
}

// @Tags WsApi
// @Summary  终端日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body kubernetesReq.TerminalRequest true "根据id获取api"
// @Success 200 {object} response.Response{data=systemRes.SysAPIResponse} "根据id获取api,返回包括api详情"
// @Router  /kubernetes/pods/logs [get]
func (w *WsApi) ContainerLog(c *gin.Context) {
	var terminal kubernetesReq.TerminalRequest
	_ = c.ShouldBindQuery(&terminal)
	if err := w.TerminalJWTAuth(terminal); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	if err := sutils.Verify(terminal, utils.TerminalVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	kubeLogger, err := utils.NewKubeLogger(c.Writer, c.Request, nil)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("连接终端websocket升级失败!，%v", err.Error()), c)
	}

	client, err := utils.NewKubeClient(terminal.ClusterId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("连接k8s client失败!，%v", err.Error()), c)
	}

	client.Pod.ContainerLog(kubeLogger, terminal.Name, terminal.PodName, terminal.Namespace)
}
