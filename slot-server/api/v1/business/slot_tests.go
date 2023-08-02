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
	"slot-server/service/test/public"
)

type SlotTestsApi struct {
}

var slotTestsService = service.ServiceGroupApp.BusinessServiceGroup.SlotTestsService

// CreateSlotTests 创建SlotTests
// @Tags SlotTests
// @Summary 创建SlotTests
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotTests true "创建SlotTests"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotTests/createSlotTests [post]
func (slotTestsApi *SlotTestsApi) CreateSlotTests(c *gin.Context) {
	var run public.RunSlotTest
	err := c.ShouldBindJSON(&run)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotTestsService.CreateSlotTests(run); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlotTests 删除SlotTests
// @Tags SlotTests
// @Summary 删除SlotTests
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotTests true "删除SlotTests"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slotTests/deleteSlotTests [delete]
func (slotTestsApi *SlotTestsApi) DeleteSlotTests(c *gin.Context) {
	var slotTests business.SlotTests
	err := c.ShouldBindJSON(&slotTests)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotTestsService.DeleteSlotTests(slotTests); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlotTestsByIds 批量删除SlotTests
// @Tags SlotTests
// @Summary 批量删除SlotTests
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SlotTests"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /slotTests/deleteSlotTestsByIds [delete]
func (slotTestsApi *SlotTestsApi) DeleteSlotTestsByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotTestsService.DeleteSlotTestsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlotTests 更新SlotTests
// @Tags SlotTests
// @Summary 更新SlotTests
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotTests true "更新SlotTests"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slotTests/updateSlotTests [put]
func (slotTestsApi *SlotTestsApi) UpdateSlotTests(c *gin.Context) {
	var slotTests business.SlotTests
	err := c.ShouldBindJSON(&slotTests)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotTestsService.UpdateSlotTests(slotTests); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSlotTests 用id查询SlotTests
// @Tags SlotTests
// @Summary 用id查询SlotTests
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.SlotTests true "用id查询SlotTests"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slotTests/findSlotTests [get]
func (slotTestsApi *SlotTestsApi) FindSlotTests(c *gin.Context) {
	var slotTests business.SlotTests
	err := c.ShouldBindQuery(&slotTests)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reslotTests, err := slotTestsService.GetSlotTests(slotTests.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reslotTests": reslotTests}, c)
	}
}

// GetSlotTestsList 分页获取SlotTests列表
// @Tags SlotTests
// @Summary 分页获取SlotTests列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.SlotTestsSearch true "分页获取SlotTests列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotTests/getSlotTestsList [get]
func (slotTestsApi *SlotTestsApi) GetSlotTestsList(c *gin.Context) {
	var pageInfo businessReq.SlotTestsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := slotTestsService.GetSlotTestsInfoList(pageInfo); err != nil {
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

// Truncate 清空表
// @Tags SlotTests
// @Summary 清空表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.SlotTestsSearch true "清空表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"清空表成功"}"
// @Router /slotTests/getSlotTestsList [get]
func (slotTestsApi *SlotTestsApi) Truncate(c *gin.Context) {
	global.GVA_DB.Exec("truncate table b_slot_test;")
	response.OkWithDetailed(response.PageResult{}, "清空成功", c)
}
