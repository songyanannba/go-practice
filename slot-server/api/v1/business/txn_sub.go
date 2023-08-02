package business

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
	"slot-server/model/common/response"
	"slot-server/model/system"
	"slot-server/service"
	"slot-server/utils"
)

type TxnSubApi struct {
}

var txnSubService = service.ServiceGroupApp.BusinessServiceGroup.TxnSubService

// CreateTxnSub 创建TxnSub
// @Tags TxnSub
// @Summary 创建TxnSub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.TxnSub true "创建TxnSub"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /txnSub/createTxnSub [post]
func (txnSubApi *TxnSubApi) CreateTxnSub(c *gin.Context) {
	var txnSub business.TxnSub
	err := c.ShouldBindJSON(&txnSub)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := txnSubService.CreateTxnSub(txnSub); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTxnSub 删除TxnSub
// @Tags TxnSub
// @Summary 删除TxnSub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.TxnSub true "删除TxnSub"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /txnSub/deleteTxnSub [delete]
func (txnSubApi *TxnSubApi) DeleteTxnSub(c *gin.Context) {
	var txnSub business.TxnSub
	err := c.ShouldBindJSON(&txnSub)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := txnSubService.DeleteTxnSub(txnSub); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTxnSubByIds 批量删除TxnSub
// @Tags TxnSub
// @Summary 批量删除TxnSub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TxnSub"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /txnSub/deleteTxnSubByIds [delete]
func (txnSubApi *TxnSubApi) DeleteTxnSubByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := txnSubService.DeleteTxnSubByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTxnSub 更新TxnSub
// @Tags TxnSub
// @Summary 更新TxnSub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.TxnSub true "更新TxnSub"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /txnSub/updateTxnSub [put]
func (txnSubApi *TxnSubApi) UpdateTxnSub(c *gin.Context) {
	var txnSub business.TxnSub
	err := c.ShouldBindJSON(&txnSub)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := txnSubService.UpdateTxnSub(txnSub); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTxnSub 用id查询TxnSub
// @Tags TxnSub
// @Summary 用id查询TxnSub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.TxnSub true "用id查询TxnSub"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /txnSub/findTxnSub [get]
func (txnSubApi *TxnSubApi) FindTxnSub(c *gin.Context) {
	var txnSub business.TxnSub
	err := c.ShouldBindQuery(&txnSub)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retxnSub, err := txnSubService.GetTxnSub(txnSub.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"retxnSub": retxnSub}, c)
	}
}

// GetTxnSubList 分页获取TxnSub列表
// @Tags TxnSub
// @Summary 分页获取TxnSub列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.TxnSubSearch true "分页获取TxnSub列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /txnSub/getTxnSubList [get]
func (txnSubApi *TxnSubApi) GetTxnSubList(c *gin.Context) {
	var pageInfo businessReq.TxnSubSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	info := utils.GetUserInfo(c)
	if info.AuthorityId == 10 {
		err = global.GVA_DB.Model(&system.SysUser{}).
			Where("id = ?", info.ID).
			Pluck("merchant_id", &pageInfo.MerchantId).Error
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}
	if list, total, err := txnSubService.GetTxnSubInfoList(pageInfo); err != nil {
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
