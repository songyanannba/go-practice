package business

import (
	"slot-server/api/v1"
	"slot-server/middleware"
	"github.com/gin-gonic/gin"
)

type TxnSubRouter struct {
}

// InitTxnSubRouter 初始化 TxnSub 路由信息
func (s *TxnSubRouter) InitTxnSubRouter(Router *gin.RouterGroup) {
	txnSubRouter := Router.Group("txnSub").Use(middleware.OperationRecord())
	txnSubRouterWithoutRecord := Router.Group("txnSub")
	var txnSubApi = v1.ApiGroupApp.BusinessApiGroup.TxnSubApi
	{
		txnSubRouter.POST("createTxnSub", txnSubApi.CreateTxnSub)   // 新建TxnSub
		txnSubRouter.DELETE("deleteTxnSub", txnSubApi.DeleteTxnSub) // 删除TxnSub
		txnSubRouter.DELETE("deleteTxnSubByIds", txnSubApi.DeleteTxnSubByIds) // 批量删除TxnSub
		txnSubRouter.PUT("updateTxnSub", txnSubApi.UpdateTxnSub)    // 更新TxnSub
	}
	{
		txnSubRouterWithoutRecord.GET("findTxnSub", txnSubApi.FindTxnSub)        // 根据ID获取TxnSub
		txnSubRouterWithoutRecord.GET("getTxnSubList", txnSubApi.GetTxnSubList)  // 获取TxnSub列表
	}
}
