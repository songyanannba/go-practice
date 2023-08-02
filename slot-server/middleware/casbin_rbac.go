package middleware

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"slot-server/global"
	"slot-server/model/common/response"
	"slot-server/service"
	"slot-server/utils"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.GVA_CONFIG.System.Env != "develop" {
			waitUse, _ := utils.GetClaims(c)
			//获取请求的PATH
			path := c.Request.URL.Path
			obj := strings.TrimPrefix(path, global.GVA_CONFIG.System.RouterPrefix)
			// 获取请求方法
			act := c.Request.Method
			// 获取用户的角色
			if waitUse.AuthorityId == 888 {
				// 如果为超级管理员则直接放行
				c.Next()
				return
			}
			sub := strconv.Itoa(int(waitUse.AuthorityId))
			e := casbinService.Casbin() // 判断策略中是否存在
			success, _ := e.Enforce(sub, obj, act)
			if !success {
				response.FailWithDetailed(gin.H{}, "权限不足", c)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}