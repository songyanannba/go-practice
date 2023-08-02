package router

import (
	"github.com/gin-gonic/gin"
	"slot-server/middleware"
	"slot-server/plugin/kubernetes/api"
)

type ClusterRouter struct{}

func (u *ClusterRouter) InitClusterRouter(Router *gin.RouterGroup) {
	clusterRouter := Router.Group("clusters").Use(middleware.OperationRecord())
	clusterouterWithoutRecord := Router.Group("clusters")
	clustersApi := api.ApiGroupApp.ClustersApi
	{
		clusterRouter.POST("create", clustersApi.CreateCluster)              // 创建
		clusterRouter.POST("getById", clustersApi.GetClusterById)            // 单个获取
		clusterRouter.PUT("update", clustersApi.UpdateCluster)               // 更新
		clusterRouter.DELETE("delete", clustersApi.DeleteCluster)            // 删除
		clusterRouter.DELETE("deleteByIds", clustersApi.DeleteClustersByIds) // 批量删除
	}
	{
		clusterouterWithoutRecord.POST("list", clustersApi.ClustersList) // 分页获取列表
	}
}
