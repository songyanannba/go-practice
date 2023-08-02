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

type CurrencyApi struct {
}

var currencyService = service.ServiceGroupApp.BusinessServiceGroup.CurrencyService

// CreateCurrency 创建Currency
// @Tags Currency
// @Summary 创建Currency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Currency true "创建Currency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /currency/createCurrency [post]
func (currencyApi *CurrencyApi) CreateCurrency(c *gin.Context) {
	var currency business.Currency
	err := c.ShouldBindJSON(&currency)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := currencyService.CreateCurrency(currency); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCurrency 删除Currency
// @Tags Currency
// @Summary 删除Currency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Currency true "删除Currency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /currency/deleteCurrency [delete]
func (currencyApi *CurrencyApi) DeleteCurrency(c *gin.Context) {
	var currency business.Currency
	err := c.ShouldBindJSON(&currency)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := currencyService.DeleteCurrency(currency); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCurrencyByIds 批量删除Currency
// @Tags Currency
// @Summary 批量删除Currency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Currency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /currency/deleteCurrencyByIds [delete]
func (currencyApi *CurrencyApi) DeleteCurrencyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := currencyService.DeleteCurrencyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCurrency 更新Currency
// @Tags Currency
// @Summary 更新Currency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Currency true "更新Currency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /currency/updateCurrency [put]
func (currencyApi *CurrencyApi) UpdateCurrency(c *gin.Context) {
	var currency business.Currency
	err := c.ShouldBindJSON(&currency)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := currencyService.UpdateCurrency(currency); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCurrency 用id查询Currency
// @Tags Currency
// @Summary 用id查询Currency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.Currency true "用id查询Currency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /currency/findCurrency [get]
func (currencyApi *CurrencyApi) FindCurrency(c *gin.Context) {
	var currency business.Currency
	err := c.ShouldBindQuery(&currency)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recurrency, err := currencyService.GetCurrency(currency.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"recurrency": recurrency}, c)
	}
}

// GetCurrencyList 分页获取Currency列表
// @Tags Currency
// @Summary 分页获取Currency列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.CurrencySearch true "分页获取Currency列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /currency/getCurrencyList [get]
func (currencyApi *CurrencyApi) GetCurrencyList(c *gin.Context) {
	var pageInfo businessReq.CurrencySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := currencyService.GetCurrencyInfoList(pageInfo); err != nil {
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
