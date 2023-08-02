package local

import (
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/lonng/nano/session"
	"net"
	"slot-server/pbs"
	"slot-server/service/cluster/game"
	"slot-server/service/cluster/gate"
	"slot-server/service/cluster/master"
	"slot-server/service/logic"
)

type testEntity struct {
	Ack any
}

func (t *testEntity) Push(route string, v interface{}) error { return nil }

func (t *testEntity) RPC(route string, v interface{}) error { return nil }

func (t *testEntity) LastMid() uint64 { return 0 }

func (t *testEntity) Response(v interface{}) error {
	t.Ack = v
	return nil
}

func (t *testEntity) ResponseMid(mid uint64, v interface{}) error { return nil }

func (t *testEntity) Close() error { return nil }

func (t *testEntity) RemoteAddr() net.Addr {
	return &net.IPNet{IP: net.IPv4(127, 0, 0, 1), Mask: net.IPMask{}}
}

func (t *testEntity) GetError() error {
	v, ok := t.Ack.(logic.AckHeads)
	if !ok {
		return errors.New("it's not ack head")
	}
	head := v.GetHead()
	if head == nil {
		return nil
	}
	if head.Code != pbs.Code_Ok || head.Message != "" {
		spew.Println(t.Ack)
		return errors.New("code not ok, message:" + head.Message)
	}
	return nil
}

type LocalTest struct {
	Entity  *testEntity
	Session *session.Session
	Bind    *gate.BindService
	Game    *game.GameService
	Master  *master.MasterService
	Head    *pbs.ReqHead
}

func NewLocalTest() *LocalTest {
	entity := &testEntity{}
	return &LocalTest{
		Entity:  entity,
		Session: session.New(entity),
		Bind:    gate.Service,
		Game:    game.Service,
		Master:  master.Service,
	}
}

func (l *LocalTest) MerchantLogin() (*pbs.MerchantLoginAck, error) {
	// 商户登录
	_ = l.Bind.MerchantLogin(l.Session, &pbs.MerchantLoginReq{Head: &pbs.ReqHead{
		Uid:      1,
		Token:    "123456",
		Platform: "test",
	}})
	ack := l.Entity.Ack.(*pbs.MerchantLoginAck)
	l.Head = &pbs.ReqHead{Token: ack.Token, Platform: "test"}

	return ack, l.Entity.GetError()
}

func (l *LocalTest) Login(req *pbs.LoginReq) (*pbs.LoginAck, error) {
	// 平台登录
	_ = l.Bind.Login(l.Session, req)
	ack := l.Entity.Ack.(*pbs.LoginAck)
	l.Head = &pbs.ReqHead{Token: ack.Token, Platform: "test"}
	return ack, l.Entity.GetError()
}

func (l *LocalTest) LogoutReq() (*pbs.LogoutAck, error) {
	_ = l.Bind.Logout(l.Session, &pbs.LogoutReq{Head: l.Head})

	return l.Entity.Ack.(*pbs.LogoutAck), l.Entity.GetError()
}

func (l *LocalTest) Tracking() error {
	_ = l.Bind.Tracking(l.Session, &pbs.Tracking{Head: l.Head, Type: 1})

	return l.Entity.GetError()
}

func (l *LocalTest) IntoGame(req *pbs.IntoGame) (*pbs.IntoGameAck, error) {
	req.Head = l.Head
	_ = l.Game.IntoGame(l.Session, req)

	return l.Entity.Ack.(*pbs.IntoGameAck), l.Entity.GetError()
}

func (l *LocalTest) Recharge(req *pbs.Recharge) (*pbs.RechargeAck, error) {
	req.Head = l.Head
	_ = l.Game.Recharge(l.Session, req)

	return l.Entity.Ack.(*pbs.RechargeAck), l.Entity.GetError()
}

func (l *LocalTest) Spin(req *pbs.Spin) (*pbs.SpinAck, *pbs.MatchSpinAck, error) {
	req.Head = l.Head
	_ = l.Game.Spin(l.Session, req)

	if req.Opt.GameId == 5 || req.Opt.GameId == 6 {
		return nil, l.Entity.Ack.(*pbs.MatchSpinAck), l.Entity.GetError()
	}

	return l.Entity.Ack.(*pbs.SpinAck), nil, l.Entity.GetError()
}

func (l *LocalTest) SpinStop(req *pbs.SpinStop) (*pbs.SpinStopAck, error) {
	req.Head = l.Head
	_ = l.Game.SpinStop(l.Session, req)

	return l.Entity.Ack.(*pbs.SpinStopAck), l.Entity.GetError()
}

func (l *LocalTest) RecordMenu(req *pbs.RecordMenuReq) (*pbs.RecordMenuAck, error) {
	req.Head = l.Head
	_ = l.Game.RecordMenu(l.Session, req)

	return l.Entity.Ack.(*pbs.RecordMenuAck), l.Entity.GetError()
}
