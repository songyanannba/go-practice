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
)

type SlotGenTplApi struct {
}

var slotGenTplService = service.ServiceGroupApp.BusinessServiceGroup.SlotGenTplService

// CreateSlotGenTpl 创建SlotGenTpl
// @Tags SlotGenTpl
// @Summary 创建SlotGenTpl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotGenTpl true "创建SlotGenTpl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotGenTpl/createSlotGenTpl [post]
func (slotGenTplApi *SlotGenTplApi) CreateSlotGenTpl(c *gin.Context) {
	var slotGenTpl business.SlotGenTpl
	err := c.ShouldBindJSON(&slotGenTpl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotGenTplService.CreateSlotGenTpl(slotGenTpl); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlotGenTpl 删除SlotGenTpl
// @Tags SlotGenTpl
// @Summary 删除SlotGenTpl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotGenTpl true "删除SlotGenTpl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slotGenTpl/deleteSlotGenTpl [delete]
func (slotGenTplApi *SlotGenTplApi) DeleteSlotGenTpl(c *gin.Context) {
	var slotGenTpl business.SlotGenTpl
	err := c.ShouldBindJSON(&slotGenTpl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotGenTplService.DeleteSlotGenTpl(slotGenTpl); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlotGenTplByIds 批量删除SlotGenTpl
// @Tags SlotGenTpl
// @Summary 批量删除SlotGenTpl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SlotGenTpl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /slotGenTpl/deleteSlotGenTplByIds [delete]
func (slotGenTplApi *SlotGenTplApi) DeleteSlotGenTplByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotGenTplService.DeleteSlotGenTplByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlotGenTpl 更新SlotGenTpl
// @Tags SlotGenTpl
// @Summary 更新SlotGenTpl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotGenTpl true "更新SlotGenTpl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slotGenTpl/updateSlotGenTpl [put]
func (slotGenTplApi *SlotGenTplApi) UpdateSlotGenTpl(c *gin.Context) {
	var slotGenTpl business.SlotGenTpl
	err := c.ShouldBindJSON(&slotGenTpl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotGenTplService.UpdateSlotGenTpl(slotGenTpl); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSlotGenTpl 用id查询SlotGenTpl
// @Tags SlotGenTpl
// @Summary 用id查询SlotGenTpl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.SlotGenTpl true "用id查询SlotGenTpl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slotGenTpl/findSlotGenTpl [get]
func (slotGenTplApi *SlotGenTplApi) FindSlotGenTpl(c *gin.Context) {
	var slotGenTpl business.SlotGenTpl
	err := c.ShouldBindQuery(&slotGenTpl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reslotGenTpl, err := slotGenTplService.GetSlotGenTpl(slotGenTpl.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reslotGenTpl": reslotGenTpl}, c)
	}
}

// GetSlotGenTplList 分页获取SlotGenTpl列表
// @Tags SlotGenTpl
// @Summary 分页获取SlotGenTpl列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.SlotGenTplSearch true "分页获取SlotGenTpl列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotGenTpl/getSlotGenTplList [get]
func (slotGenTplApi *SlotGenTplApi) GetSlotGenTplList(c *gin.Context) {
	var pageInfo businessReq.SlotGenTplSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := slotGenTplService.GetSlotGenTplInfoList(pageInfo); err != nil {
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
