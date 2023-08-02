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
	"slot-server/timedtask"
	"slot-server/utils"
)

type MoneyUserApi struct {
}

var userDailySumService = service.ServiceGroupApp.BusinessServiceGroup.MoneyUserService

// CreateMoneyUser 创建MoneyUser
// @Tags MoneyUser
// @Summary 创建MoneyUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MoneyUser true "创建MoneyUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userDailySum/createMoneyUser [post]
func (userDailySumApi *MoneyUserApi) CreateMoneyUser(c *gin.Context) {
	var userDailySum business.MoneyUser
	err := c.ShouldBindJSON(&userDailySum)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Date":     {utils.NotEmpty()},
		"UserId":   {utils.NotEmpty()},
		"UserName": {utils.NotEmpty()},
	}
	if err := utils.Verify(userDailySum, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userDailySumService.CreateMoneyUser(userDailySum); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMoneyUser 删除MoneyUser
// @Tags MoneyUser
// @Summary 删除MoneyUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MoneyUser true "删除MoneyUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userDailySum/deleteMoneyUser [delete]
func (userDailySumApi *MoneyUserApi) DeleteMoneyUser(c *gin.Context) {
	var userDailySum business.MoneyUser
	err := c.ShouldBindJSON(&userDailySum)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userDailySumService.DeleteMoneyUser(userDailySum); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMoneyUserByIds 批量删除MoneyUser
// @Tags MoneyUser
// @Summary 批量删除MoneyUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MoneyUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userDailySum/deleteMoneyUserByIds [delete]
func (userDailySumApi *MoneyUserApi) DeleteMoneyUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userDailySumService.DeleteMoneyUserByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMoneyUser 更新MoneyUser
// @Tags MoneyUser
// @Summary 更新MoneyUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MoneyUser true "更新MoneyUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userDailySum/updateMoneyUser [put]
func (userDailySumApi *MoneyUserApi) UpdateMoneyUser(c *gin.Context) {
	var userDailySum business.MoneyUser
	err := c.ShouldBindJSON(&userDailySum)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Date":     {utils.NotEmpty()},
		"UserId":   {utils.NotEmpty()},
		"UserName": {utils.NotEmpty()},
	}
	if err := utils.Verify(userDailySum, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userDailySumService.UpdateMoneyUser(userDailySum); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMoneyUser 用id查询MoneyUser
// @Tags MoneyUser
// @Summary 用id查询MoneyUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.MoneyUser true "用id查询MoneyUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userDailySum/findMoneyUser [get]
func (userDailySumApi *MoneyUserApi) FindMoneyUser(c *gin.Context) {
	var userDailySum business.MoneyUser
	err := c.ShouldBindQuery(&userDailySum)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reuserDailySum, err := userDailySumService.GetMoneyUser(userDailySum.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reuserDailySum": reuserDailySum}, c)
	}
}

// GetMoneyUserList 分页获取MoneyUser列表
// @Tags MoneyUser
// @Summary 分页获取MoneyUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.MoneyUserSearch true "分页获取MoneyUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userDailySum/getMoneyUserList [get]
func (userDailySumApi *MoneyUserApi) GetMoneyUserList(c *gin.Context) {
	var pageInfo businessReq.MoneyUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := userDailySumService.GetMoneyUserInfoList(pageInfo); err != nil {
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

type UserDaily struct {
	Date string `json:"date" form:"date" `
}

func (userDailySumApi *MoneyUserApi) Recalculate(c *gin.Context) {
	var UserDaily UserDaily
	err := c.ShouldBindQuery(&UserDaily)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = timedtask.MoneyUserCal(UserDaily.Date)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return

	} else {
		response.OkWithMessage("计算成功", c)
		return
	}
}
