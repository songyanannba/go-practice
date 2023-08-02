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
	"slot-server/utils/helper"
)

type DebugConfigApi struct {
}

var debugConfigService = service.ServiceGroupApp.BusinessServiceGroup.DebugConfigService

// CreateDebugConfig 创建DebugConfig
// @Tags DebugConfig
// @Summary 创建DebugConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.DebugConfig true "创建DebugConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /debugConfig/createDebugConfig [post]
func (debugConfigApi *DebugConfigApi) CreateDebugConfig(c *gin.Context) {
	var debugConfig business.DebugConfig
	err := c.ShouldBindJSON(&debugConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"SlotId":    {utils.NotEmpty()},
		"PalyType":  {utils.NotEmpty()},
		"DebugType": {utils.NotEmpty()},
		"Start":     {utils.NotEmpty()},
	}
	if err := utils.Verify(debugConfig, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := debugConfigService.CreateDebugConfig(debugConfig); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDebugConfig 删除DebugConfig
// @Tags DebugConfig
// @Summary 删除DebugConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.DebugConfig true "删除DebugConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /debugConfig/deleteDebugConfig [delete]
func (debugConfigApi *DebugConfigApi) DeleteDebugConfig(c *gin.Context) {
	var debugConfig business.DebugConfig
	err := c.ShouldBindJSON(&debugConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := debugConfigService.DeleteDebugConfig(debugConfig); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDebugConfigByIds 批量删除DebugConfig
// @Tags DebugConfig
// @Summary 批量删除DebugConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DebugConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /debugConfig/deleteDebugConfigByIds [delete]
func (debugConfigApi *DebugConfigApi) DeleteDebugConfigByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := debugConfigService.DeleteDebugConfigByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDebugConfig 更新DebugConfig
// @Tags DebugConfig
// @Summary 更新DebugConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.DebugConfig true "更新DebugConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /debugConfig/updateDebugConfig [put]
func (debugConfigApi *DebugConfigApi) UpdateDebugConfig(c *gin.Context) {
	var debugConfig business.DebugConfig
	err := c.ShouldBindJSON(&debugConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"SlotId":    {utils.NotEmpty()},
		"PalyType":  {utils.NotEmpty()},
		"DebugType": {utils.NotEmpty()},
		"Start":     {utils.NotEmpty()},
	}
	if err := utils.Verify(debugConfig, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := debugConfigService.UpdateDebugConfig(debugConfig); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDebugConfig 用id查询DebugConfig
// @Tags DebugConfig
// @Summary 用id查询DebugConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.DebugConfig true "用id查询DebugConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /debugConfig/findDebugConfig [get]
func (debugConfigApi *DebugConfigApi) FindDebugConfig(c *gin.Context) {
	var debugConfig business.DebugConfig
	err := c.ShouldBindQuery(&debugConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redebugConfig, err := debugConfigService.GetDebugConfig(debugConfig.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"redebugConfig": redebugConfig}, c)
	}
}

// GetDebugConfigList 分页获取DebugConfig列表
// @Tags DebugConfig
// @Summary 分页获取DebugConfig列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.DebugConfigSearch true "分页获取DebugConfig列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /debugConfig/getDebugConfigList [get]
func (debugConfigApi *DebugConfigApi) GetDebugConfigList(c *gin.Context) {
	var pageInfo businessReq.DebugConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := debugConfigService.GetDebugConfigInfoList(pageInfo); err != nil {
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

type GetTags struct {
	SlotId int `json:"slot_id" form:"slot_id"`
}
type SlotConfig struct {
	SlotId      int                   `json:"slot_id" form:"slot_id"`
	Row         int                   `json:"row" form:"row"`
	Col         int                   `json:"col" form:"col"`
	SlotSymbols []business.SlotSymbol `json:"slot_symbols" form:"slot_symbols"`
}
type SlotSize struct {
	Size string `json:"size" form:"size"`
}

func (debugConfigApi *DebugConfigApi) GetSlotTags(c *gin.Context) {
	var getTags GetTags
	err := c.ShouldBindQuery(&getTags)
	slotConfig := SlotConfig{
		SlotId:      getTags.SlotId,
		SlotSymbols: []business.SlotSymbol{},
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	slotSize := SlotSize{}
	err = global.GVA_DB.Raw("SELECT b.size FROM b_slot a LEFT JOIN "+
		"b_slot_payline b ON a.payline_no=b.NO WHERE a.id=? LIMIT 1", getTags.SlotId).Scan(&slotSize).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	arr := helper.SplitInt[int](slotSize.Size, "*")
	slotConfig.Row = helper.SliceVal(arr, 0)
	slotConfig.Col = helper.SliceVal(arr, 1)

	var SlotSymbols []business.SlotSymbol
	if err = global.GVA_DB.Where("slot_id = ?", getTags.SlotId).Find(&SlotSymbols).Error; err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	slotConfig.SlotSymbols = SlotSymbols
	response.OkWithDetailed(slotConfig, "获取成功", c)
}
