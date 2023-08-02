package seamless

import (
	"github.com/gin-gonic/gin"
	"slot-server/model/common/response"
	"slot-server/service/logic"
	"slot-server/service/logic/upper/seamlessWallet"
)

// Authenticate [测试用] 通过安全令牌对玩家进行身份验证
// @Tags Seamless-Operator
// @Summary Authenticate
// @Description 打开游戏时，供应商将收到由娱乐场运营商生成的 URL 安全令牌。使用此令牌将要求娱乐场运营商进行玩家身份验证并获取玩家余额。
// @Param RequestParams body seamlessWallet.AuthReq true "请求参数"
// @Success 0 {object} response.Response{data=seamlessWallet.PlayerAmount}
// @Router /seamless/authenticate [post]
func Authenticate(c *gin.Context) {
	var params seamlessWallet.AuthReq
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId, err := logic.RequireLogin(params.Token)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := seamlessWallet.TestAuthenticate(userId, &params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(res, c)
}

// Balance 查询玩家余额
// @Tags Seamless-Operator
// @Summary Balance
// @Description 通过此方法，供应商系统能知道玩家的当前余额并显示在游戏中
// @Param RequestParams body seamlessWallet.BalanceReq true "请求参数"
// @Success 0 {object} response.Response{data=seamlessWallet.PlayerAmount}
// @Router /seamless/balance [post]
func Balance(c *gin.Context) {
	var params seamlessWallet.BalanceReq
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId, err := logic.RequireLogin(params.Token)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := seamlessWallet.TestBalance(userId, &params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(res, c)
}

// Bet 下注
// @Tags Seamless-Operator
// @Summary Bet
// @Description 此方法将会返回玩家的下注金额，运营商需确保玩家余额足够支付下注并扣除,如失败或异常请返回code 7
// @Description 余额不足时导致的错误可返回code 10，并需在data中携带playerBalance字段(此时玩家的真实余额)，客户端将依据此字段更新余额显示
// @Param RequestParams body seamlessWallet.BetReq true "请求参数"
// @Success 0 {object} response.Response{data=seamlessWallet.BetAck} "playerBalance为扣除下注后的余额"
// @Router /seamless/bet [post]
func Bet(c *gin.Context) {
	var params seamlessWallet.BetReq
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId, err := logic.RequireLogin(params.Token)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := seamlessWallet.TestBet(userId, &params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(res, c)
}

// Result 下注结果
// @Tags Seamless-Operator
// @Summary Result
// @Description 此接口在失败时会定时重试,token字段为空，运营商需根据交易id保证在玩家赢得奖金时支付奖金的幂等性
// @Param RequestParams body seamlessWallet.ResultReq true "请求参数"
// @Success 0 {object} response.Response{data=seamlessWallet.ResultAck}
// @Router /seamless/result [post]
func Result(c *gin.Context) {
	var params seamlessWallet.ResultReq
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := seamlessWallet.TestResult(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(res, c)
}

// Refund 退款通知
// @Tags Seamless-Operator
// @Summary Refund
// @Description 当收到退款请求时，运营商必须将钱退还给玩家 调用是幂等的，例如再次为现有投注发送退款只会创建一笔交易。此接口在失败时会定时重试,token字段为空。
// @Param RequestParams body seamlessWallet.RefundReq true "请求参数"
// @Success 0 {object} response.Response{data=seamlessWallet.RefundAck}
// @Router /seamless/refund [post]
func Refund(c *gin.Context) {
	var params seamlessWallet.RefundReq
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId, err := logic.RequireLogin(params.Token)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := seamlessWallet.TestRefund(userId, &params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(res, c)
}
