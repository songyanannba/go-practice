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

type JackpotApi struct {
}

var jackpotService = service.ServiceGroupApp.BusinessServiceGroup.JackpotService

// CreateJackpot 创建Jackpot
// @Tags Jackpot
// @Summary 创建Jackpot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Jackpot true "创建Jackpot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jackpot/createJackpot [post]
func (jackpotApi *JackpotApi) CreateJackpot(c *gin.Context) {
	var jackpot business.Jackpot
	err := c.ShouldBindJSON(&jackpot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"slotId": {utils.NotEmpty()},
	}
	if err := utils.Verify(jackpot, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := jackpotService.CreateJackpot(jackpot); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteJackpot 删除Jackpot
// @Tags Jackpot
// @Summary 删除Jackpot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Jackpot true "删除Jackpot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jackpot/deleteJackpot [delete]
func (jackpotApi *JackpotApi) DeleteJackpot(c *gin.Context) {
	var jackpot business.Jackpot
	err := c.ShouldBindJSON(&jackpot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := jackpotService.DeleteJackpot(jackpot); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteJackpotByIds 批量删除Jackpot
// @Tags Jackpot
// @Summary 批量删除Jackpot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Jackpot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /jackpot/deleteJackpotByIds [delete]
func (jackpotApi *JackpotApi) DeleteJackpotByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := jackpotService.DeleteJackpotByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateJackpot 更新Jackpot
// @Tags Jackpot
// @Summary 更新Jackpot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Jackpot true "更新Jackpot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /jackpot/updateJackpot [put]
func (jackpotApi *JackpotApi) UpdateJackpot(c *gin.Context) {
	var jackpot business.Jackpot
	err := c.ShouldBindJSON(&jackpot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"slotId": {utils.NotEmpty()},
	}
	if err := utils.Verify(jackpot, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := jackpotService.UpdateJackpot(jackpot); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindJackpot 用id查询Jackpot
// @Tags Jackpot
// @Summary 用id查询Jackpot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.Jackpot true "用id查询Jackpot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /jackpot/findJackpot [get]
func (jackpotApi *JackpotApi) FindJackpot(c *gin.Context) {
	var jackpot business.Jackpot
	err := c.ShouldBindQuery(&jackpot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rejackpot, err := jackpotService.GetJackpot(jackpot.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"rejackpot": rejackpot}, c)
	}
}

// GetJackpotList 分页获取Jackpot列表
// @Tags Jackpot
// @Summary 分页获取Jackpot列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.JackpotSearch true "分页获取Jackpot列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jackpot/getJackpotList [get]
func (jackpotApi *JackpotApi) GetJackpotList(c *gin.Context) {
	var pageInfo businessReq.JackpotSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := jackpotService.GetJackpotInfoList(pageInfo); err != nil {
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

type SaveJackpot struct {
	Jackpots []business.Jackpot `json:"jackpots" form:"jackpots" `
}

func (jackpotApi *JackpotApi) SaveJackpotList(c *gin.Context) {
	var SaveJackpot SaveJackpot
	err := c.ShouldBindJSON(&SaveJackpot)
	if err = global.GVA_DB.Exec("truncate table b_jackpot").Error; err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	if err = global.GVA_DB.CreateInBatches(SaveJackpot.Jackpots, 1000).Error; err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithMessage("保存成功", c)
}
