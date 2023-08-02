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

type MoneySlotApi struct {
}

var moneySlotService = service.ServiceGroupApp.BusinessServiceGroup.MoneySlotService

// CreateMoneySlot 创建MoneySlot
// @Tags MoneySlot
// @Summary 创建MoneySlot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MoneySlot true "创建MoneySlot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moneySlot/createMoneySlot [post]
func (moneySlotApi *MoneySlotApi) CreateMoneySlot(c *gin.Context) {
	var moneySlot business.MoneySlot
	err := c.ShouldBindJSON(&moneySlot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Date":   {utils.NotEmpty()},
		"SlotId": {utils.NotEmpty()},
	}
	if err := utils.Verify(moneySlot, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := moneySlotService.CreateMoneySlot(moneySlot); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMoneySlot 删除MoneySlot
// @Tags MoneySlot
// @Summary 删除MoneySlot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MoneySlot true "删除MoneySlot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /moneySlot/deleteMoneySlot [delete]
func (moneySlotApi *MoneySlotApi) DeleteMoneySlot(c *gin.Context) {
	var moneySlot business.MoneySlot
	err := c.ShouldBindJSON(&moneySlot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := moneySlotService.DeleteMoneySlot(moneySlot); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMoneySlotByIds 批量删除MoneySlot
// @Tags MoneySlot
// @Summary 批量删除MoneySlot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MoneySlot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /moneySlot/deleteMoneySlotByIds [delete]
func (moneySlotApi *MoneySlotApi) DeleteMoneySlotByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := moneySlotService.DeleteMoneySlotByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMoneySlot 更新MoneySlot
// @Tags MoneySlot
// @Summary 更新MoneySlot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.MoneySlot true "更新MoneySlot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /moneySlot/updateMoneySlot [put]
func (moneySlotApi *MoneySlotApi) UpdateMoneySlot(c *gin.Context) {
	var moneySlot business.MoneySlot
	err := c.ShouldBindJSON(&moneySlot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Date":   {utils.NotEmpty()},
		"SlotId": {utils.NotEmpty()},
	}
	if err := utils.Verify(moneySlot, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := moneySlotService.UpdateMoneySlot(moneySlot); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMoneySlot 用id查询MoneySlot
// @Tags MoneySlot
// @Summary 用id查询MoneySlot
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.MoneySlot true "用id查询MoneySlot"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /moneySlot/findMoneySlot [get]
func (moneySlotApi *MoneySlotApi) FindMoneySlot(c *gin.Context) {
	var moneySlot business.MoneySlot
	err := c.ShouldBindQuery(&moneySlot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if remoneySlot, err := moneySlotService.GetMoneySlot(moneySlot.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"remoneySlot": remoneySlot}, c)
	}
}

// GetMoneySlotList 分页获取MoneySlot列表
// @Tags MoneySlot
// @Summary 分页获取MoneySlot列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.MoneySlotSearch true "分页获取MoneySlot列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moneySlot/getMoneySlotList [get]
func (moneySlotApi *MoneySlotApi) GetMoneySlotList(c *gin.Context) {
	var pageInfo businessReq.MoneySlotSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := moneySlotService.GetMoneySlotInfoList(pageInfo); err != nil {
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
func (moneySlotApi *MoneySlotApi) Recalculate(c *gin.Context) {
	var UserDaily UserDaily
	err := c.ShouldBindQuery(&UserDaily)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = timedtask.MoneySlotCal(UserDaily.Date)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return

	} else {
		response.OkWithMessage("计算成功", c)
		return
	}
}
