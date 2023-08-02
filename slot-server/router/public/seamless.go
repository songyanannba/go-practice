package public

import (
	"github.com/gin-gonic/gin"
	"slot-server/api/public/seamless"
)

type SeamlessRouter struct{}

func (s *SeamlessRouter) InitSeamlessRouter(Router *gin.RouterGroup) {
	seamlessRouter := Router.Group("seamless")
	{
		seamlessRouter.POST("openGame", seamless.OpenGame)
		seamlessRouter.POST("gameList", seamless.GameList)
		seamlessRouter.POST("closeGame", seamless.CloseGame)
		seamlessRouter.POST("historyList", seamless.HistoryList)
	}
}

func (s *SeamlessRouter) InitTestSeamlessRouter(Router *gin.RouterGroup) {
	seamlessRouter := Router.Group("seamless")
	{
		seamlessRouter.POST("authenticate", seamless.Authenticate)
		seamlessRouter.POST("balance", seamless.Balance)
		seamlessRouter.POST("bet", seamless.Bet)
		seamlessRouter.POST("result", seamless.Result)
		seamlessRouter.POST("refund", seamless.Refund)
	}
}
