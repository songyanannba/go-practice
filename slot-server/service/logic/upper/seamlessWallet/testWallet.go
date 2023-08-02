package seamlessWallet

import (
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/service/cache"
	"slot-server/utils/helper"
	"strconv"
)

const (
	TestMerchant = "test"
	TestToken    = "123456"
)

func testCheck(params ReqHeadData) (*business.Merchant, error) {
	merchant, err := cache.GetMerchant(TestMerchant)
	if err != nil {
		return nil, err
	}
	return merchant, CheckSign(params, params.GetHead().Sign, merchant.Secret)
}

// TestAuthenticate 模拟鉴权  OpenGame 时返回的token -> Auth
func TestAuthenticate(userId uint, params *AuthReq) (*PlayerAmount, error) {
	_, err := testCheck(params)
	if err != nil {
		return nil, err
	}

	user, err := business.Get[*business.User](userId)
	if err != nil {
		return nil, err
	}

	return &PlayerAmount{
		PlayerName:    user.Username,
		Currency:      user.Currency,
		PlayerBalance: helper.Div100(user.Amount),
	}, nil
}

func TestBalance(userId uint, params *BalanceReq) (*PlayerAmount, error) {
	_, err := testCheck(params)
	if err != nil {
		return nil, err
	}

	user, err := business.Get[*business.User](userId)
	if err != nil {
		return nil, err
	}

	return &PlayerAmount{
		PlayerName:    user.Username,
		Currency:      user.Currency,
		PlayerBalance: helper.Div100(user.Amount),
	}, nil
}

func TestBet(userId uint, params *BetReq) (*BetAck, error) {
	_, err := testCheck(params)
	if err != nil {
		return nil, err
	}

	user, err := business.Get[*business.User](userId)
	if err != nil {
		return nil, err
	}

	bet := helper.Mul100(params.BetAmount)

	log, err := cache.ChangeMoney(user.ID, -bet, nil,
		business.MoneyLogWithTxnId(uint(params.GameId), params.TransactionId, enum.MoneyType1Spin),
	)
	if err != nil {
		return nil, err
	}

	return &BetAck{
		Currency:      user.Currency,
		PlayerBalance: helper.Div100(log.CoinResult),
		AgentTxid:     strconv.Itoa(int(log.ID)),
	}, nil
}

func TestResult(params *ResultReq) (*ResultAck, error) {
	merchant, err := testCheck(params)
	if err != nil {
		return nil, err
	}

	var user business.User
	err = global.GVA_DB.First(&user, "username = ? and merchant_id = ?", params.PlayerName, merchant.ID).Error
	if err != nil {
		return nil, err
	}
	win := helper.Mul100(params.WinAmount)

	log, err := cache.ChangeMoney(user.ID, win, nil,
		business.MoneyLogWithTxnId(uint(params.GameId), params.TransactionId, enum.MoneyType1Spin),
	)
	if err != nil {
		return nil, err
	}
	if params.WinAmount == 0 {
		log.CoinResult, _ = cache.GetUserAmount(user.ID)
	}

	return &ResultAck{
		Currency:      user.Currency,
		PlayerBalance: helper.Div100(log.CoinResult),
		AgentTxid:     strconv.Itoa(int(log.ID)),
	}, nil
}

func TestRefund(userId uint, params *RefundReq) (*RefundAck, error) {
	_, err := testCheck(params)
	if err != nil {
		return nil, err
	}
	var moneyLog business.MoneyLog
	err = global.GVA_DB.First(&moneyLog, "txn_id = ?", params.TransactionId).Error
	if err != nil {
		return nil, err
	}

	log, err := cache.ChangeMoney(moneyLog.UserId, -moneyLog.CoinChange, nil,
		business.MoneyLogWithTxnId(moneyLog.GameId, params.TransactionId, enum.MoneyType5Refund),
	)
	if err != nil {
		return nil, err
	}

	var user business.User
	err = global.GVA_DB.First(&user, "id = ?", moneyLog.UserId).Error
	if err != nil {
		return nil, err
	}

	return &RefundAck{
		Currency:      user.Currency,
		PlayerBalance: helper.Div100(log.CoinResult),
	}, nil
}
