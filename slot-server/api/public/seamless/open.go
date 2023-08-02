package seamless

import (
	"github.com/gin-gonic/gin"
	"slot-server/model/common/response"
	"slot-server/service/logic/upper/seamlessWallet"
)

// OpenGame
// @Tags Seamless-Provider
// @Summary	OpenGame
// @Description 通过这种方法，运营商可以接收到所请求游戏的有效启动 URL。
// @Param RequestParams body seamlessWallet.OpenGameParams true "请求参数"
// @Success	0 {object} response.Response{data=string}	"URL将在data中返回"
// @Router /seamless/openGame [post]
func OpenGame(c *gin.Context) {
	var (
		params = &seamlessWallet.OpenGameParams{}
		h      = seamlessWallet.NewOpenHandle(params, "/seamless/openGame")
		err    error
	)
	//defer func() {
	//	h.WriteLog(err)
	//}()
	err = c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = h.OpenGame()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(h.Res, c)
}

// GameList
// @Tags Seamless-Provider
// @Summary	GameList
// @Description 游戏列表
// @Param RequestParams body seamlessWallet.GameListParams true "请求参数"
// @Success 0 {object} response.Response{data=[]seamlessWallet.GameListRes}
// @Router /seamless/gameList [post]
func GameList(c *gin.Context) {
	var (
		params = &seamlessWallet.GameListParams{}
		h      = seamlessWallet.NewOpenHandle(params, "/seamless/gameList")
		err    error
	)
	//defer func() {
	//	h.WriteLog(err)
	//}()
	err = c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = h.GameList()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(h.Res, c)
}

// CloseGame
// @Tags Seamless-Provider
// @Summary CloseGame
// @Description 关闭游戏
// @Param RequestParams body seamlessWallet.CloseGameParams true "请求参数"
// @Success 0 {object} response.Response
// @Router /seamless/closeGame [post]
func CloseGame(c *gin.Context) {
	var (
		params = &seamlessWallet.CloseGameParams{}
		h      = seamlessWallet.NewOpenHandle(params, "/seamless/closeGame")
		err    error
	)
	defer func() {
		h.WriteLog(err)
	}()
	err = c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = h.CloseGame()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// HistoryList 历史记录列表
// @Tags Seamless-Provider
// @Summary HistoryList
// @Description 历史记录列表 playerName gameId transactionId date 为可选字段. status 1 玩家已开始游戏回合但尚未结束 2 玩家已完成游戏回合 3 游戏回合在数据库中被标记为已完成；但是Result请求没有得到正确回复 9 游戏回合被游戏回合结束流程⾃动关闭 10退款处于异步队列中并正被发送给运营商
// @Param RequestParams body seamlessWallet.HistoryListParams true "请求参数"
// @Success 0 {object} response.Response{data=[]seamlessWallet.History}
// @Router /seamless/historyList [post]
func HistoryList(c *gin.Context) {
	var (
		params = &seamlessWallet.HistoryListParams{}
		h      = seamlessWallet.NewOpenHandle(params, "/seamless/historyList")
		err    error
	)
	//defer func() {
	//	h.WriteLog(err)
	//}()
	err = c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = h.HistoryList(params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(h.Res, c)
}
