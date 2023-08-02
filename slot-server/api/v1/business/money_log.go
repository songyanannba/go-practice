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

type MoneyLogApi struct {
}

var moneyLogService = service.ServiceGroupApp.BusinessServiceGroup.MoneyLogService

// CreateMoneyLog 创建MoneyLog
// @Tags MoneyLog
// @Summary 创建MoneyLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MoneyLog true "创建MoneyLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moneyLog/createMoneyLog [post]
func (moneyLogApi *MoneyLogApi) CreateMoneyLog(c *gin.Context) {
	var moneyLog business.MoneyLog
	err := c.ShouldBindJSON(&moneyLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Date":   {utils.NotEmpty()},
		"UserId": {utils.NotEmpty()},
		"Action": {utils.NotEmpty()},
	}
	if err := utils.Verify(moneyLog, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := moneyLogService.CreateMoneyLog(moneyLog); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMoneyLog 删除MoneyLog
// @Tags MoneyLog
// @Summary 删除MoneyLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MoneyLog true "删除MoneyLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /moneyLog/deleteMoneyLog [delete]
func (moneyLogApi *MoneyLogApi) DeleteMoneyLog(c *gin.Context) {
	var moneyLog business.MoneyLog
	err := c.ShouldBindJSON(&moneyLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := moneyLogService.DeleteMoneyLog(moneyLog); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMoneyLogByIds 批量删除MoneyLog
// @Tags MoneyLog
// @Summary 批量删除MoneyLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MoneyLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /moneyLog/deleteMoneyLogByIds [delete]
func (moneyLogApi *MoneyLogApi) DeleteMoneyLogByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := moneyLogService.DeleteMoneyLogByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMoneyLog 更新MoneyLog
// @Tags MoneyLog
// @Summary 更新MoneyLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MoneyLog true "更新MoneyLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /moneyLog/updateMoneyLog [put]
func (moneyLogApi *MoneyLogApi) UpdateMoneyLog(c *gin.Context) {
	var moneyLog business.MoneyLog
	err := c.ShouldBindJSON(&moneyLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Date":   {utils.NotEmpty()},
		"UserId": {utils.NotEmpty()},
		"Action": {utils.NotEmpty()},
	}
	if err := utils.Verify(moneyLog, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := moneyLogService.UpdateMoneyLog(moneyLog); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMoneyLog 用id查询MoneyLog
// @Tags MoneyLog
// @Summary 用id查询MoneyLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.MoneyLog true "用id查询MoneyLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /moneyLog/findMoneyLog [get]
func (moneyLogApi *MoneyLogApi) FindMoneyLog(c *gin.Context) {
	var moneyLog business.MoneyLog
	err := c.ShouldBindQuery(&moneyLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if remoneyLog, err := moneyLogService.GetMoneyLog(moneyLog.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"remoneyLog": remoneyLog}, c)
	}
}

// GetMoneyLogList 分页获取MoneyLog列表
// @Tags MoneyLog
// @Summary 分页获取MoneyLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.MoneyLogSearch true "分页获取MoneyLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moneyLog/getMoneyLogList [get]
func (moneyLogApi *MoneyLogApi) GetMoneyLogList(c *gin.Context) {
	var pageInfo businessReq.MoneyLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := moneyLogService.GetMoneyLogInfoList(pageInfo); err != nil {
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
