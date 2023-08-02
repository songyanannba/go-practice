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

type MoneyUserSlotApi struct {
}

var moneyUserSlotService = service.ServiceGroupApp.BusinessServiceGroup.MoneyUserSlotService

// CreateMoneyUserSlot 创建MoneyUserSlot
// @Tags MoneyUserSlot
// @Summary 创建MoneyUserSlot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MoneyUserSlot true "创建MoneyUserSlot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moneyUserSlot/createMoneyUserSlot [post]
func (moneyUserSlotApi *MoneyUserSlotApi) CreateMoneyUserSlot(c *gin.Context) {
	var moneyUserSlot business.MoneyUserSlot
	err := c.ShouldBindJSON(&moneyUserSlot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Date":   {utils.NotEmpty()},
		"UserId": {utils.NotEmpty()},
		"SlotId": {utils.NotEmpty()},
	}
	if err := utils.Verify(moneyUserSlot, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := moneyUserSlotService.CreateMoneyUserSlot(moneyUserSlot); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMoneyUserSlot 删除MoneyUserSlot
// @Tags MoneyUserSlot
// @Summary 删除MoneyUserSlot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MoneyUserSlot true "删除MoneyUserSlot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /moneyUserSlot/deleteMoneyUserSlot [delete]
func (moneyUserSlotApi *MoneyUserSlotApi) DeleteMoneyUserSlot(c *gin.Context) {
	var moneyUserSlot business.MoneyUserSlot
	err := c.ShouldBindJSON(&moneyUserSlot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := moneyUserSlotService.DeleteMoneyUserSlot(moneyUserSlot); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMoneyUserSlotByIds 批量删除MoneyUserSlot
// @Tags MoneyUserSlot
// @Summary 批量删除MoneyUserSlot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MoneyUserSlot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /moneyUserSlot/deleteMoneyUserSlotByIds [delete]
func (moneyUserSlotApi *MoneyUserSlotApi) DeleteMoneyUserSlotByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := moneyUserSlotService.DeleteMoneyUserSlotByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMoneyUserSlot 更新MoneyUserSlot
// @Tags MoneyUserSlot
// @Summary 更新MoneyUserSlot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MoneyUserSlot true "更新MoneyUserSlot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /moneyUserSlot/updateMoneyUserSlot [put]
func (moneyUserSlotApi *MoneyUserSlotApi) UpdateMoneyUserSlot(c *gin.Context) {
	var moneyUserSlot business.MoneyUserSlot
	err := c.ShouldBindJSON(&moneyUserSlot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Date":   {utils.NotEmpty()},
		"UserId": {utils.NotEmpty()},
		"SlotId": {utils.NotEmpty()},
	}
	if err := utils.Verify(moneyUserSlot, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := moneyUserSlotService.UpdateMoneyUserSlot(moneyUserSlot); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMoneyUserSlot 用id查询MoneyUserSlot
// @Tags MoneyUserSlot
// @Summary 用id查询MoneyUserSlot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.MoneyUserSlot true "用id查询MoneyUserSlot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /moneyUserSlot/findMoneyUserSlot [get]
func (moneyUserSlotApi *MoneyUserSlotApi) FindMoneyUserSlot(c *gin.Context) {
	var moneyUserSlot business.MoneyUserSlot
	err := c.ShouldBindQuery(&moneyUserSlot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if remoneyUserSlot, err := moneyUserSlotService.GetMoneyUserSlot(moneyUserSlot.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"remoneyUserSlot": remoneyUserSlot}, c)
	}
}

// GetMoneyUserSlotList 分页获取MoneyUserSlot列表
// @Tags MoneyUserSlot
// @Summary 分页获取MoneyUserSlot列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.MoneyUserSlotSearch true "分页获取MoneyUserSlot列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moneyUserSlot/getMoneyUserSlotList [get]
func (moneyUserSlotApi *MoneyUserSlotApi) GetMoneyUserSlotList(c *gin.Context) {
	var pageInfo businessReq.MoneyUserSlotSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := moneyUserSlotService.GetMoneyUserSlotInfoList(pageInfo); err != nil {
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
func (moneyUserSlotApi *MoneyUserSlotApi) Recalculate(c *gin.Context) {
	var UserDaily UserDaily
	err := c.ShouldBindQuery(&UserDaily)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = timedtask.MoneyUserSlotCal(UserDaily.Date)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return

	} else {
		response.OkWithMessage("计算成功", c)
		return
	}
}
