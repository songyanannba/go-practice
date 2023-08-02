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

type TxnApi struct {
}

var txnService = service.ServiceGroupApp.BusinessServiceGroup.TxnService

// CreateTxn 创建Txn
// @Tags Txn
// @Summary 创建Txn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Txn true "创建Txn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /txn/createTxn [post]
func (txnApi *TxnApi) CreateTxn(c *gin.Context) {
	var txn business.Txn
	err := c.ShouldBindJSON(&txn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := txnService.CreateTxn(txn); err != nil {
		global.GVA_LOG.Error("fail!", zap.Error(err))
		response.FailWithMessage("fail: "+err.Error(), c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteTxn 删除Txn
// @Tags Txn
// @Summary 删除Txn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Txn true "删除Txn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /txn/deleteTxn [delete]
func (txnApi *TxnApi) DeleteTxn(c *gin.Context) {
	var txn business.Txn
	err := c.ShouldBindJSON(&txn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := txnService.DeleteTxn(txn); err != nil {
		global.GVA_LOG.Error("fail!", zap.Error(err))
		response.FailWithMessage("fail: "+err.Error(), c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteTxnByIds 批量删除Txn
// @Tags Txn
// @Summary 批量删除Txn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Txn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /txn/deleteTxnByIds [delete]
func (txnApi *TxnApi) DeleteTxnByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := txnService.DeleteTxnByIds(IDS); err != nil {
		global.GVA_LOG.Error("fail!", zap.Error(err))
		response.FailWithMessage("fail: "+err.Error(), c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// UpdateTxn 更新Txn
// @Tags Txn
// @Summary 更新Txn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.Txn true "更新Txn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /txn/updateTxn [put]
func (txnApi *TxnApi) UpdateTxn(c *gin.Context) {
	var txn business.Txn
	err := c.ShouldBindJSON(&txn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := txnService.UpdateTxn(txn); err != nil {
		global.GVA_LOG.Error("fail!", zap.Error(err))
		response.FailWithMessage("fail: "+err.Error(), c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// FindTxn 用id查询Txn
// @Tags Txn
// @Summary 用id查询Txn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.Txn true "用id查询Txn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /txn/findTxn [get]
func (txnApi *TxnApi) FindTxn(c *gin.Context) {
	var txn business.Txn
	err := c.ShouldBindQuery(&txn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retxn, err := txnService.GetTxn(txn.ID); err != nil {
		global.GVA_LOG.Error("fail!", zap.Error(err))
		response.FailWithMessage("fail: "+err.Error(), c)
	} else {
		response.OkWithData(gin.H{"retxn": retxn}, c)
	}
}

// GetTxnList 分页获取Txn列表
// @Tags Txn
// @Summary 分页获取Txn列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.TxnSearch true "分页获取Txn列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /txn/getTxnList [get]
func (txnApi *TxnApi) GetTxnList(c *gin.Context) {
	var pageInfo businessReq.TxnSearch
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
	if list, total, err := txnService.GetTxnInfoList(pageInfo); err != nil {
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
