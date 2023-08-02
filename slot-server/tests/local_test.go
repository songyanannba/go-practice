package tests

import (
	"github.com/stretchr/testify/assert"
	"slot-server/core"
	"slot-server/pbs"
	"slot-server/service/test/local"
	"testing"
)

func TestLocal(t *testing.T) {
	core.BaseInit()
	lt := local.NewLocalTest()

	// 平台登录 demo 用户只能使用spin功能
	//loginAck, err := lt.Login(&pbs.LoginReq{Username: "demo"})
	//if !assert.Equal(t, err, nil) {
	//	return
	//}
	//lt.Head = &pbs.ReqHead{Token: loginAck.Token, Platform: "test", Demo: true}

	// 商户登录
	mLoginAck, err := lt.MerchantLogin()
	if !assert.Equal(t, err, nil) {
		return
	}
	lt.Head = &pbs.ReqHead{Token: mLoginAck.Token, Platform: "test"}

	//// 登出
	//_, err = lt.LogoutReq()
	//if !assert.Equal(t, err, nil) {
	//	return
	//}
	//
	//// 埋点通知 无回复
	//err = lt.Tracking()
	//if !assert.Equal(t, err, nil) {
	//	return
	//}
	//
	//// 进入游戏
	intoGameAck, err := lt.IntoGame(&pbs.IntoGame{GameId: 1, PlayType: 1})
	if !assert.Equal(t, err, nil) {
		return
	}
	t.Log(intoGameAck)

	//// 历史记录
	//recordMenuAck, err := lt.RecordMenu(&pbs.RecordMenuReq{})
	//if !assert.Equal(t, err, nil) {
	//	return
	//}
	//t.Log(recordMenuAck)
	//
	//// 充值
	//_, err = lt.Recharge(&pbs.Recharge{Amount: 1000, Type: 1})
	//if !assert.Equal(t, err, nil) {
	//	return
	//}

	// spin
	//spinAck, matchSpinAck, err := lt.Spin(&pbs.Spin{Opt: &pbs.SpinOpt{GameId: 3, Bet: 100, Raise: true}})
	//if !assert.Equal(t, err, nil) {
	//	return
	//}
	//t.Log(spinAck)
	//t.Log(matchSpinAck)
	// spinStop
	//_, err = lt.SpinStop(&pbs.SpinStop{GameId: 2})
	//if !assert.Equal(t, err, nil) {
	//	return
	//}
}
