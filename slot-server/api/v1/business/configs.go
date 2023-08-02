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

type ConfigsApi struct {
}

var configsService = service.ServiceGroupApp.BusinessServiceGroup.ConfigsService

// CreateConfigs 创建Configs
// @Tags Configs
// @Summary 创建Configs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Configs true "创建Configs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /configs/createConfigs [post]
func (configsApi *ConfigsApi) CreateConfigs(c *gin.Context) {
	var configs business.Configs
	err := c.ShouldBindJSON(&configs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":  {utils.NotEmpty()},
		"Value": {utils.NotEmpty()},
	}
	if err := utils.Verify(configs, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := configsService.CreateConfigs(configs); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteConfigs 删除Configs
// @Tags Configs
// @Summary 删除Configs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Configs true "删除Configs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /configs/deleteConfigs [delete]
func (configsApi *ConfigsApi) DeleteConfigs(c *gin.Context) {
	var configs business.Configs
	err := c.ShouldBindJSON(&configs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := configsService.DeleteConfigs(configs); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteConfigsByIds 批量删除Configs
// @Tags Configs
// @Summary 批量删除Configs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Configs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /configs/deleteConfigsByIds [delete]
func (configsApi *ConfigsApi) DeleteConfigsByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := configsService.DeleteConfigsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateConfigs 更新Configs
// @Tags Configs
// @Summary 更新Configs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Configs true "更新Configs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /configs/updateConfigs [put]
func (configsApi *ConfigsApi) UpdateConfigs(c *gin.Context) {
	var configs business.Configs
	err := c.ShouldBindJSON(&configs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":  {utils.NotEmpty()},
		"Value": {utils.NotEmpty()},
	}
	if err := utils.Verify(configs, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := configsService.UpdateConfigs(configs); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindConfigs 用id查询Configs
// @Tags Configs
// @Summary 用id查询Configs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.Configs true "用id查询Configs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /configs/findConfigs [get]
func (configsApi *ConfigsApi) FindConfigs(c *gin.Context) {
	var configs business.Configs
	err := c.ShouldBindQuery(&configs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reconfigs, err := configsService.GetConfigs(configs.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reconfigs": reconfigs}, c)
	}
}

// GetConfigsList 分页获取Configs列表
// @Tags Configs
// @Summary 分页获取Configs列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.ConfigsSearch true "分页获取Configs列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /configs/getConfigsList [get]
func (configsApi *ConfigsApi) GetConfigsList(c *gin.Context) {
	var pageInfo businessReq.ConfigsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := configsService.GetConfigsInfoList(pageInfo); err != nil {
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
