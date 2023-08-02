package public

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/public/game"
)

type GameRouter struct{}

func (s *GameRouter) InitGameRouter(Router *gin.RouterGroup) {
	seamlessRouter := Router.Group("game")
	{
		seamlessRouter.POST("backendOperate", game.BackendOperate)
		seamlessRouter.POST("record", game.GetSlotRecord)

		seamlessRouter.GET("ping-x41nk", game.Ping)
		seamlessRouter.GET("time-x41nk", game.Time)
		seamlessRouter.POST("log-x41nk", game.Log)
	}
}
