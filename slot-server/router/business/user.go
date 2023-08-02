package business

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/v1"
	"slot-server/middleware"
)

type UserRouter struct {
}

// InitUserRouter 初始化 User 路由信息
func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("bUser").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("bUser")
	var userApi = v1.ApiGroupApp.BusinessApiGroup.UserApi
	{
		userRouter.POST("createUser", userApi.CreateUser)             // 新建User
		userRouter.DELETE("deleteUser", userApi.DeleteUser)           // 删除User
		userRouter.DELETE("deleteUserByIds", userApi.DeleteUserByIds) // 批量删除User
		userRouter.PUT("updateUser", userApi.UpdateUser)              // 更新User
		userRouter.PUT("changePassword", userApi.ChangePassword)      // 修改密码
		userRouter.PUT("changeAmount", userApi.ChangeAmount)          // 修改密码
	}
	{
		userRouterWithoutRecord.GET("findUser", userApi.FindUser)       // 根据ID获取User
		userRouterWithoutRecord.GET("getUserList", userApi.GetUserList) // 获取User列表
	}
}
