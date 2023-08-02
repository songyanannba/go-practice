package seamlessWallet

// PlayerAmount 通用金额结果
type PlayerAmount struct {
	PlayerName    string  `json:"playerName" example:"1"`         // 玩家名称
	Currency      string  `json:"currency" example:"USD"`         // 货币
	PlayerBalance float64 `json:"playerBalance" example:"800.85"` // 玩家余额
}

// AuthReq 鉴权请求
type AuthReq struct {
	ReqHead
}

func (m *MerchantHandle) Authenticate() (PlayerAmount, int, error) {
	return req[PlayerAmount](m, "authenticate", &AuthReq{})
}

// BalanceReq 查询玩家余额请求
type BalanceReq struct {
	ReqHead
	PlayerName string `json:"playerName" example:"test"` // 玩家名称
}

func (m *MerchantHandle) Balance(playerName string) (PlayerAmount, int, error) {
	return req[PlayerAmount](m, "balance", &BalanceReq{PlayerName: playerName})
}

// BetReq 下注请求
type BetReq struct {
	ReqHead
	PlayerName string `json:"playerName" example:"test"` // 玩家在运营商系统中的唯一标识
	GameId     int    `json:"gameId" example:"1"`        // 游戏id

	Currency      string  `json:"currency" example:"USD"`              // 货币类型
	BetAmount     float64 `json:"betAmount" example:"10.00"`           // 下注总金额
	TransactionId string  `json:"transactionId" example:"tx1"`         // 供应商交易id
	UpdatedTime   int64   `json:"updatedTime" example:"1234567890123"` // 更新时间戳
}

type BetAck struct {
	AgentTxid     string  `json:"agentTxid" example:"tx2"`        // 运营商交易id
	Currency      string  `json:"currency" example:"USD"`         // 货币类型
	PlayerBalance float64 `json:"playerBalance" example:"800.85"` // 玩家余额(扣款后)
}

func (m *MerchantHandle) Bet(betReq BetReq) (BetAck, int, error) {
	return req[BetAck](m, "bet", &betReq)
}

// ResultReq 下注结果请求
type ResultReq struct {
	ReqHead
	PlayerName string `json:"playerName" example:"test"` // 玩家在运营商系统中的唯一标识
	GameId     int    `json:"gameId" example:"1"`        // 机台id

	Currency      string  `json:"currency" example:"USD"`              // 货币类型
	BetAmount     float64 `json:"betAmount" example:"10.00"`           // 下注总金额
	WinAmount     float64 `json:"winAmount" example:"1.00"`            // 赢得金额
	TransactionId string  `json:"transactionId" example:"tx1"`         // 供应商交易id
	AgentTxid     string  `json:"agentTxid" example:"tx2"`             // 运营商交易id
	UpdatedTime   int64   `json:"updatedTime" example:"1234567890123"` // 更新时间戳
}

// ResultAck 下注结果响应
type ResultAck struct {
	AgentTxid     string  `json:"agentTxid" example:"tx3"`        // 运营商交易id
	Currency      string  `json:"currency" example:"USD"`         // 货币类型
	PlayerBalance float64 `json:"playerBalance" example:"800.85"` // 玩家余额(扣款后)
}

func (m *MerchantHandle) Result(resultReq ResultReq) (ResultAck, int, error) {
	return req[ResultAck](m, "result", &resultReq)
}

// RefundReq 退款请求
type RefundReq struct {
	ReqHead
	PlayerName string `json:"playerName" example:"test"` // 玩家在运营商系统中的唯一标识

	GameId        int     `json:"gameId" example:"1"`                  // 游戏id
	TransactionId string  `json:"transactionId" example:"tx1"`         // 供应商交易id
	AgentTxid     string  `json:"agentTxid" example:"tx3"`             // 运营商交易id
	UpdatedTime   int64   `json:"updatedTime" example:"1234567890123"` // 更新时间戳
	BetAmount     float64 `json:"betAmount" example:"10.00"`           // 下注总金额
}

// RefundAck 退款响应
type RefundAck struct {
	//AgentTxid string `json:"agentTxid" example:"tx3"` // 运营商内部的流水id
	Currency      string  `json:"currency" example:"USD"`         // 货币类型
	PlayerBalance float64 `json:"playerBalance" example:"800.85"` // 玩家余额(退款后)
}

func (m *MerchantHandle) Refund(refundReq RefundReq) (RefundAck, int, error) {
	return req[RefundAck](m, "refund", &refundReq)
}

func BonusWin() {

}

func Adjustment() {

}
