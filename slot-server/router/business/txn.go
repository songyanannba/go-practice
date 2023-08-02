package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type TxnRouter struct {
}

// InitTxnRouter 初始化 Txn 路由信息
func (s *TxnRouter) InitTxnRouter(Router *gin.RouterGroup) {
	txnRouter := Router.Group("txn").Use(middleware.OperationRecord())
	txnRouterWithoutRecord := Router.Group("txn")
	var txnApi = v1.ApiGroupApp.BusinessApiGroup.TxnApi
	{
		txnRouter.POST("createTxn", txnApi.CreateTxn)             // 新建Txn
		txnRouter.DELETE("deleteTxn", txnApi.DeleteTxn)           // 删除Txn
		txnRouter.DELETE("deleteTxnByIds", txnApi.DeleteTxnByIds) // 批量删除Txn
		txnRouter.PUT("updateTxn", txnApi.UpdateTxn)              // 更新Txn
	}
	{
		txnRouterWithoutRecord.GET("findTxn", txnApi.FindTxn)       // 根据ID获取Txn
		txnRouterWithoutRecord.GET("getTxnList", txnApi.GetTxnList) // 获取Txn列表
	}
}
