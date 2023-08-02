package business

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
	"slot-server/model/common/response"
	"slot-server/service"
	"slot-server/utils"
)

type TrackingApi struct {
}

var trackingService = service.ServiceGroupApp.BusinessServiceGroup.TrackingService

// CreateTracking 创建Tracking
// @Tags Tracking
// @Summary 创建Tracking
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Tracking true "创建Tracking"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tracking/createTracking [post]
func (trackingApi *TrackingApi) CreateTracking(c *gin.Context) {
	var tracking business.Tracking
	err := c.ShouldBindJSON(&tracking)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Date": {utils.NotEmpty()},
		"Type": {utils.NotEmpty()},
	}
	if err := utils.Verify(tracking, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := trackingService.CreateTracking(tracking); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTracking 删除Tracking
// @Tags Tracking
// @Summary 删除Tracking
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Tracking true "删除Tracking"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tracking/deleteTracking [delete]
func (trackingApi *TrackingApi) DeleteTracking(c *gin.Context) {
	var tracking business.Tracking
	err := c.ShouldBindJSON(&tracking)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := trackingService.DeleteTracking(tracking); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTrackingByIds 批量删除Tracking
// @Tags Tracking
// @Summary 批量删除Tracking
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Tracking"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tracking/deleteTrackingByIds [delete]
func (trackingApi *TrackingApi) DeleteTrackingByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := trackingService.DeleteTrackingByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTracking 更新Tracking
// @Tags Tracking
// @Summary 更新Tracking
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Tracking true "更新Tracking"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tracking/updateTracking [put]
func (trackingApi *TrackingApi) UpdateTracking(c *gin.Context) {
	var tracking business.Tracking
	err := c.ShouldBindJSON(&tracking)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Date": {utils.NotEmpty()},
		"Type": {utils.NotEmpty()},
	}
	if err := utils.Verify(tracking, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := trackingService.UpdateTracking(tracking); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTracking 用id查询Tracking
// @Tags Tracking
// @Summary 用id查询Tracking
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.Tracking true "用id查询Tracking"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tracking/findTracking [get]
func (trackingApi *TrackingApi) FindTracking(c *gin.Context) {
	var tracking business.Tracking
	err := c.ShouldBindQuery(&tracking)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retracking, err := trackingService.GetTracking(tracking.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"retracking": retracking}, c)
	}
}

// GetTrackingList 分页获取Tracking列表
// @Tags Tracking
// @Summary 分页获取Tracking列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.TrackingSearch true "分页获取Tracking列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tracking/getTrackingList [get]
func (trackingApi *TrackingApi) GetTrackingList(c *gin.Context) {
	var pageInfo businessReq.TrackingSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := trackingService.GetTrackingInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
