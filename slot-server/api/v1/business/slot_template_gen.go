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

type SlotTemplateGenApi struct {
}

var slotTemplateGenService = service.ServiceGroupApp.BusinessServiceGroup.SlotTemplateGenService

// CreateSlotTemplateGen 创建SlotTemplateGen
// @Tags SlotTemplateGen
// @Summary 创建SlotTemplateGen
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotTemplateGen true "创建SlotTemplateGen"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotTemplateGen/createSlotTemplateGen [post]
func (slotTemplateGenApi *SlotTemplateGenApi) CreateSlotTemplateGen(c *gin.Context) {
	var slotTemplateGen business.SlotTemplateGen
	err := c.ShouldBindJSON(&slotTemplateGen)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"SlotId": {utils.NotEmpty()},
		"Type":   {utils.NotEmpty()},
	}
	if err := utils.Verify(slotTemplateGen, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotTemplateGenService.CreateSlotTemplateGen(slotTemplateGen); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlotTemplateGen 删除SlotTemplateGen
// @Tags SlotTemplateGen
// @Summary 删除SlotTemplateGen
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotTemplateGen true "删除SlotTemplateGen"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slotTemplateGen/deleteSlotTemplateGen [delete]
func (slotTemplateGenApi *SlotTemplateGenApi) DeleteSlotTemplateGen(c *gin.Context) {
	var slotTemplateGen business.SlotTemplateGen
	err := c.ShouldBindJSON(&slotTemplateGen)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotTemplateGenService.DeleteSlotTemplateGen(slotTemplateGen); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlotTemplateGenByIds 批量删除SlotTemplateGen
// @Tags SlotTemplateGen
// @Summary 批量删除SlotTemplateGen
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SlotTemplateGen"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /slotTemplateGen/deleteSlotTemplateGenByIds [delete]
func (slotTemplateGenApi *SlotTemplateGenApi) DeleteSlotTemplateGenByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotTemplateGenService.DeleteSlotTemplateGenByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlotTemplateGen 更新SlotTemplateGen
// @Tags SlotTemplateGen
// @Summary 更新SlotTemplateGen
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotTemplateGen true "更新SlotTemplateGen"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slotTemplateGen/updateSlotTemplateGen [put]
func (slotTemplateGenApi *SlotTemplateGenApi) UpdateSlotTemplateGen(c *gin.Context) {
	var slotTemplateGen business.SlotTemplateGen
	err := c.ShouldBindJSON(&slotTemplateGen)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"SlotId": {utils.NotEmpty()},
		"Type":   {utils.NotEmpty()},
	}
	if err := utils.Verify(slotTemplateGen, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotTemplateGenService.UpdateSlotTemplateGen(slotTemplateGen); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (slotTemplateGenApi *SlotTemplateGenApi) GenerateSlotTemplateGen(c *gin.Context) {
	var slotTemplateGen business.SlotTemplateGen
	err := c.ShouldBindJSON(&slotTemplateGen)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := slotTemplateGenService.GenerateSlotTemplateGen(&slotTemplateGen); err != nil {
		global.GVA_LOG.Error("生成失败!", zap.Error(err))
		response.FailWithMessage("生成失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("生成成功", c)
	}
}

// FindSlotTemplateGen 用id查询SlotTemplateGen
// @Tags SlotTemplateGen
// @Summary 用id查询SlotTemplateGen
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.SlotTemplateGen true "用id查询SlotTemplateGen"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slotTemplateGen/findSlotTemplateGen [get]
func (slotTemplateGenApi *SlotTemplateGenApi) FindSlotTemplateGen(c *gin.Context) {
	var slotTemplateGen business.SlotTemplateGen
	err := c.ShouldBindQuery(&slotTemplateGen)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reslotTemplateGen, err := slotTemplateGenService.GetSlotTemplateGen(slotTemplateGen.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reslotTemplateGen": reslotTemplateGen}, c)
	}
}

// GetSlotTemplateGenList 分页获取SlotTemplateGen列表
// @Tags SlotTemplateGen
// @Summary 分页获取SlotTemplateGen列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.SlotTemplateGenSearch true "分页获取SlotTemplateGen列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotTemplateGen/getSlotTemplateGenList [get]
func (slotTemplateGenApi *SlotTemplateGenApi) GetSlotTemplateGenList(c *gin.Context) {
	var pageInfo businessReq.SlotTemplateGenSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := slotTemplateGenService.GetSlotTemplateGenInfoList(pageInfo); err != nil {
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
