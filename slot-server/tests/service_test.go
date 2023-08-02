package tests

import (
	"errors"
	"fmt"
	"github.com/lonng/nano/benchmark/testdata"
	"go.uber.org/zap"
	"slot-server/core"
	"slot-server/global"
	"slot-server/pbs"
	"slot-server/service/cluster/backend"
	"slot-server/service/logic"
	"slot-server/utils/helper"
	"sync"
	"testing"
	"time"
)

func TestServiceReq(t *testing.T) {
	core.BaseInit()
	var err error
	connector := backend.CreateClusterConn("pre", "34.36.25.220", "/gate", "ws")
	// ping
	err = logic.Req(connector, "MasterService.Test", &testdata.Ping{Content: "a"}, &testdata.Pong{})
	helper.FatalError(err)

	// 登录
	loginReq := &pbs.LoginReq{Username: "Guest_10001", Password: "123456", Pkg: "1"}
	loginAck := &pbs.LoginAck{}
	err = logic.Req(connector, "BindService.Login", loginReq, loginAck)
	helper.FatalError(err)

	// 商户登录
	//loginReq := &pbs.MerchantLoginReq{Head: &pbs.ReqHead{Token: "123456", Platform: "test"}}
	//loginAck := &pbs.MerchantLoginAck{}
	//err = logic.Req(connector, "BindService.MerchantLogin", loginReq, loginAck)
	//helper.FatalError(err)
	//return
	// 获取token
	head := &pbs.ReqHead{Token: loginAck.Token, Platform: "test"}

	// 登出
	//logoutReq := &pbs.LogoutReq{Head: head}
	//logoutAck := &pbs.LogoutAck{}
	//err = logic.Req(connector, "BindService.Logout", logoutReq, logoutAck)
	//helper.FatalError(err)

	// 余额
	amountReq := &pbs.Amount{Head: head}
	amountAck := &pbs.AmountAck{}
	err = logic.Req(connector, "GameService.Amount", amountReq, amountAck)
	helper.FatalError(err)
	t.Log(amountAck)

	// 埋点通知 无回复
	//tracking := &pbs.Tracking{Type: pbs.TrackingType_login}
	//err = logic.Notify(connector, "BindService.Tracking", tracking)
	//helper.FatalError(err)

	// 进入游戏
	//intoGame := &pbs.IntoGame{Head: head, GameId: 1}
	//intoGameAck := &pbs.IntoGameAck{}
	//err = logic.Req(connector, "GameService.IntoGame", intoGame, intoGameAck)
	//helper.FatalError(err)

	// 充值
	//recharge := &pbs.Recharge{Head: head, Amount: 1000, Type: 1}
	//rechargeAck := &pbs.RechargeAck{}
	//err = logic.Req(connector, "GameService.Recharge", recharge, rechargeAck)
	//helper.FatalError(err)

	// spin
	//spin := &pbs.Spin{Head: head, Opt: &pbs.SpinOpt{
	//	GameId:  2,
	//	Bet:     100,
	//	Raise:   false,
	//	BuyFree: false,
	//}}
	//spinAck := pbs.SpinAck{}
	//err = logic.Req(connector, "GameService.Spin", spin, &spinAck)
	//helper.FatalError(err)

	// spin结束 该逻辑目前只在jackpot中使用
	//spinStop := &pbs.SpinStop{Head: head, GameId: 1}
	//spinStopAck := &pbs.SpinStopAck{}
	//err = logic.Req(connector, "GameService.SpinStop", spinStop, spinStopAck)
	//helper.FatalError(err)

	global.GVA_LOG.Info("out")
	time.Sleep(time.Second * 15)
	connector.Close()
}

func TestServicePressure(t *testing.T) {
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG.Logger)
	wg := sync.WaitGroup{}
	userCount := 2
	for i := 1; i <= userCount; i++ {
		wg.Add(1)
		go func(i int) {
			err := userPlay(i)
			t.Log(i, err)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func userPlay(id int) error {
	connector := backend.CreateClusterConn("name", "api-pre.bigwin.money", "/gate", "wss")
	loginAck := &pbs.LoginAck{}
	err := logic.Req(connector, "BindService.Login", &pbs.LoginReq{
		Head:     &pbs.ReqHead{},
		Pkg:      "1",
		Username: "Guest_" + helper.Itoa(id+10000),
		Password: "123456",
	}, loginAck)
	if err != nil {
		return err
	}

	head := &pbs.ReqHead{Token: loginAck.Token}
	ch, errCh, _ := helper.Parallel(1, 5, func() (int, error) {
		err = spinOnce(connector, head)
		if err != nil {
			global.GVA_LOG.Error("spinOnce", zap.Error(err))
		}
		return 1, err
	})
	count := 0
	for {
		select {
		case err = <-errCh:
			global.GVA_LOG.Error("err", zap.Error(err))
			return errors.New("count " + helper.Itoa(count) + " err: " + err.Error())
		case _, beforeClosed := <-ch:
			if !beforeClosed {
				goto end
			}
			count++
			fmt.Println(id, " - ", count)
		}
	}
end:
	connector.Close()
	return err
}

func spinOnce(connector *backend.Connector, head *pbs.ReqHead) error {
	ack := pbs.SpinAck{}
	slotId := 3
	err := logic.Req(connector, "GameService.Spin", &pbs.Spin{Head: head, Opt: &pbs.SpinOpt{
		GameId:  int32(slotId),
		Bet:     100,
		Raise:   false,
		BuyFree: false,
	}}, &ack)
	return err
}
