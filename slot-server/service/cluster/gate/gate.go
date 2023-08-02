package gate

import (
	"errors"
	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
	"go.uber.org/zap"
	"log"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/pbs"
	"slot-server/service/cache"
	"slot-server/service/logic"
	"slot-server/service/logic/upper"
	"slot-server/service/logic/upper/seamlessWallet"
	"slot-server/utils"
	"slot-server/utils/helper"
)

var (
	Component = &component.Components{}
	Service   = newService()
)

func init() {
	Component.Register(Service)
}

type BindService struct {
	component.Base
	group *nano.Group
}

func OnSessionClosed(s *session.Session) {
	Service.userDisconnected(s)
}

func newService() *BindService {
	return &BindService{
		group: nano.NewGroup("users"),
	}
}

func (bs *BindService) Init() {}

func (bs *BindService) AfterInit() {}

func (bs *BindService) BeforeShutdown() {}

func (bs *BindService) Shutdown() {
	global.GVA_DB.Where("online = 1").
		Model(&business.User{}).Update("online", enum.No)
}

func (bs *BindService) Kick(uid int64) bool {
	session, err := bs.group.Member(uid)
	if err != nil {
		return false
	}
	bs.group.Leave(session)
	session.Close()
	return true
}

func (bs *BindService) userDisconnected(s *session.Session) {
	uid := s.UID()
	if uid == 0 {
		return
	}
	bs.group.Leave(s)
	global.GVA_DB.Model(&business.User{}).Where("id = ?", uid).Update("online", enum.No)
	log.Println("User session disconnected", uid)

}

func (bs *BindService) Login(s *session.Session, req *pbs.LoginReq) error {
	ack := &pbs.LoginAck{}
	ack.Head = logic.NewAckHead(nil)
	ip := utils.RemoteIpString(s)

	if req.Head == nil {
		return logic.RespError(s, req, ack, errors.New("head is nil"))
	}

	if req.Head.Demo {
		logic.DemoLogin(ack)
		return s.Response(ack)
	}

	var (
		err     error
		user    *business.User
		isLogin = req.Username != ""
	)
	if isLogin {
		user, err = logic.Login(ip, req, ack)
	} else {
		user, err = logic.SignUp(ip, req, ack)
	}
	if err != nil {
		return logic.RespError(s, req, ack, err)
	}

	if s.Bind(int64(user.ID)) == nil {
		// 如果已经在线，踢下线
		if isLogin && !bs.Kick(int64(user.ID)) {
			// 通知其他节点踢下线
			_ = logic.PublishCluster(enum.ClusterOperateType1KickAccount, helper.Itoa(user.ID))
		}
		_ = bs.group.Add(s)
		global.GVA_DB.Model(&business.User{}).Where("id = ?", user.ID).Update("online", enum.Yes)
	}

	ack.GameList = cache.GetGameList()
	err = logic.UserTrackingFunc(user.ID, pbs.TrackingType_login_success)
	if err != nil {
		return logic.RespError(s, nil, ack, err)
	}
	return s.Response(ack)
}

func (bs *BindService) Logout(s *session.Session, req *pbs.LogoutReq) error {
	ack := &pbs.LogoutAck{}
	ack.Head = logic.NewAckHead(req.Head)
	u, _, err := logic.RequireMerchantToken(req.Head, false)
	if err != nil {
		return logic.RespError(s, req, ack, err)
	}

	logic.CleanUser(u.ID)
	return s.Response(ack)
}

// Tracking 数据埋点
func (bs *BindService) Tracking(s *session.Session, req *pbs.Tracking) error {
	uid := uint(0)
	if req.Head != nil && req.Head.Token != "" && req.Head.Platform != "" {
		u, _, err := logic.RequireMerchantToken(req.Head, false)
		if err != nil {
			return logic.RespError(s, req, nil, err)
		}
		uid = u.ID
	}

	err := logic.UserTrackingFunc(uid, req.Type)
	if err != nil {
		return logic.RespError(s, nil, &pbs.LoginAck{Token: req.Head.Token}, err)
	}

	return nil
}

func (bs *BindService) MerchantLogin(s *session.Session, req *pbs.MerchantLoginReq) error {
	ack := &pbs.MerchantLoginAck{}
	ack.Head = logic.NewAckHead(nil)
	global.GVA_LOG.Info("Login", zap.Any("req", req))

	if req.Head == nil {
		ack.Head.Code = pbs.Code_TokenInvalid
		return logic.RespError(s, req, ack, enum.ErrTokenInvalid)
	}

	if req.Head.Demo {
		logic.MerchantDemoLogin(ack)
		return s.Response(ack)
	}

	merchant, err := cache.GetMerchant(req.Head.Platform)
	if err != nil {
		ack.Head.Code = pbs.Code_NotExists
		return logic.RespError(s, req, ack, errors.New("platform not exists"))
	}

	// 请求商户端的鉴权接口
	h := seamlessWallet.NewMerchantReq(merchant, req.Head.Token)
	player, _, err := h.Authenticate()
	if err != nil {
		ack.Head.Code = pbs.Code_TokenInvalid
		return logic.RespError(s, req, ack, errors.New("request authenticate error: "+err.Error()))
	}

	// 保存用户信息
	u, err := upper.GetOrCreateUser(merchant, player.PlayerName, player.PlayerBalance, player.Currency)
	if err != nil {
		ack.Head.Code = pbs.Code_SystemError
		return logic.RespError(s, req, ack, enum.ErrSysError)
	}
	u.Token = req.Head.Token
	err = cache.SetMerchantUserCache(merchant.Agent, u.Token, u)
	if err != nil {
		ack.Head.Code = pbs.Code_SystemError
		return logic.RespError(s, req, ack, enum.ErrSysError)
	}
	err = s.Bind(int64(u.ID))
	if err == nil {
		// 如果已经在线，踢下线
		if !bs.Kick(int64(u.ID)) {
			// 通知其他节点踢下线
			_ = logic.PublishCluster(enum.ClusterOperateType1KickAccount, helper.Itoa(u.ID))
		}
		_ = bs.group.Add(s)
		global.GVA_DB.Model(&business.User{}).Where("id = ?", u.ID).Update("online", enum.Yes)
	}

	ack.GameList = cache.GetGameList()
	err = logic.UserTrackingFunc(u.ID, pbs.TrackingType_login_success)
	if err != nil {
		return logic.RespError(s, nil, ack, err)
	}
	ack.Token = u.Token
	ack.Username = u.Username
	ack.Amount = u.Amount
	ack.Head.Uid = int32(u.ID)
	ack.Currency = u.Currency
	return s.Response(ack)
}

// BackendOperate gate后台操作
func (bs *BindService) BackendOperate(s *session.Session, req *pbs.BackendOperate) error {
	ack := &pbs.BackendOperateAck{}
	ack.Head = logic.NewAckHead(req.Head)
	_, err := logic.ParseBackendToken(req.Head)
	if err != nil {
		ack.Head.Code = pbs.Code_TokenInvalid
		return logic.RespError(s, req, ack, err)
	}
	err = logic.HandleBackendOperate(req)
	if err != nil {
		ack.Head.Code = pbs.Code_SystemError
		return logic.RespError(s, req, ack, err)
	}
	err = s.RPC("GameService.BackendOperate", req)
	if err != nil {
		ack.Head.Code = pbs.Code_Unknown
		return logic.RespError(s, req, ack, err)
	}

	return s.Response(ack)
}
