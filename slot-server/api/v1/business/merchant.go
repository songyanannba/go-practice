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

type MerchantApi struct {
}

var merchantService = service.ServiceGroupApp.BusinessServiceGroup.MerchantService

// CreateMerchant 创建Merchant
// @Tags Merchant
// @Summary 创建Merchant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Merchant true "创建Merchant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /merchant/createMerchant [post]
func (merchantApi *MerchantApi) CreateMerchant(c *gin.Context) {
	var merchant business.Merchant
	err := c.ShouldBindJSON(&merchant)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"Currency": {utils.NotEmpty()},
		"Type":     {utils.NotEmpty()},
	}
	if err := utils.Verify(merchant, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := merchantService.CreateMerchant(merchant); err != nil {
		global.GVA_LOG.Error("fail!", zap.Error(err))
		response.FailWithMessage("fail: "+err.Error(), c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteMerchant 删除Merchant
// @Tags Merchant
// @Summary 删除Merchant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Merchant true "删除Merchant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /merchant/deleteMerchant [delete]
func (merchantApi *MerchantApi) DeleteMerchant(c *gin.Context) {
	var merchant business.Merchant
	err := c.ShouldBindJSON(&merchant)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := merchantService.DeleteMerchant(merchant); err != nil {
		global.GVA_LOG.Error("fail!", zap.Error(err))
		response.FailWithMessage("fail: "+err.Error(), c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteMerchantByIds 批量删除Merchant
// @Tags Merchant
// @Summary 批量删除Merchant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Merchant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /merchant/deleteMerchantByIds [delete]
func (merchantApi *MerchantApi) DeleteMerchantByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := merchantService.DeleteMerchantByIds(IDS); err != nil {
		global.GVA_LOG.Error("fail!", zap.Error(err))
		response.FailWithMessage("fail: "+err.Error(), c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// UpdateMerchant 更新Merchant
// @Tags Merchant
// @Summary 更新Merchant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Merchant true "更新Merchant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /merchant/updateMerchant [put]
func (merchantApi *MerchantApi) UpdateMerchant(c *gin.Context) {
	var merchant business.Merchant
	err := c.ShouldBindJSON(&merchant)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"Currency": {utils.NotEmpty()},
		"Type":     {utils.NotEmpty()},
	}
	if err := utils.Verify(merchant, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := merchantService.UpdateMerchant(merchant); err != nil {
		global.GVA_LOG.Error("fail!", zap.Error(err))
		response.FailWithMessage("fail: "+err.Error(), c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// FindMerchant 用id查询Merchant
// @Tags Merchant
// @Summary 用id查询Merchant
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.Merchant true "用id查询Merchant"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /merchant/findMerchant [get]
func (merchantApi *MerchantApi) FindMerchant(c *gin.Context) {
	var merchant business.Merchant
	err := c.ShouldBindQuery(&merchant)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if remerchant, err := merchantService.GetMerchant(merchant.ID); err != nil {
		global.GVA_LOG.Error("fail!", zap.Error(err))
		response.FailWithMessage("fail: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"remerchant": remerchant}, c)
	}
}

// GetMerchantList 分页获取Merchant列表
// @Tags Merchant
// @Summary 分页获取Merchant列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.MerchantSearch true "分页获取Merchant列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /merchant/getMerchantList [get]
func (merchantApi *MerchantApi) GetMerchantList(c *gin.Context) {
	var pageInfo businessReq.MerchantSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := merchantService.GetMerchantInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("fail!", zap.Error(err))
		response.FailWithMessage("fail: "+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "success", c)
	}
}
