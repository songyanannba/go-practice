package seamlessWallet

import (
	"errors"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/service/cache"
	"slot-server/service/logic/upper"
	"slot-server/utils"
	"strconv"
	"strings"
	"time"
)

const (
	DefaultProviderId = "HZ"
)

type ReqHeadData interface {
	GetHead() *ReqHead
}

// ReqHead 我们请求三方时的请求头
type ReqHead struct {
	ProviderId string `json:"providerId" example:"HZ"`                         // 通用参数 提供商id
	Token      string `json:"token" example:"123456"`                          // 通用参数 运营商在OpenGame中提供的token
	Sign       string `json:"sign" example:"c8687876a15b93d1251c124dbe1837a6"` // 通用参数 签名
	Timestamp  string `json:"timestamp" example:"1612345678000"`               // 通用参数 时间戳
}

func (r *ReqHead) GetHead() *ReqHead {
	return r
}

func (r ReqHead) GetSign() string {
	return r.Sign
}

// CommonRes 通用结果
type CommonRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (c CommonRes) GetRes() CommonRes {
	return c
}

type MerchantHandle struct {
	ReqHead
	Merchant *business.Merchant `json:"-"`
}

func NewMerchantReq(merchant *business.Merchant, token string) *MerchantHandle {
	return &MerchantHandle{
		ReqHead: ReqHead{
			ProviderId: DefaultProviderId,
			Token:      token,
		},
		Merchant: merchant,
	}
}

func NewMerchantReqByAgent(agent string, token string) (*MerchantHandle, error) {
	merchant, err := cache.GetMerchant(agent)
	if err != nil {
		return nil, err
	}
	return NewMerchantReq(merchant, token), nil
}

type MerchantRes[V any] struct {
	Code int    `json:"code"` // code 7 failed 0 success. 7失败 0成功
	Data V      `json:"data"` // data. 结果数据
	Msg  string `json:"msg"`  // message. 消息
}

func req[V any](m *MerchantHandle, path string, data ReqHeadData) (V, int, error) {
	var (
		resData []byte
		g       *utils.Gurl
		err     error
		res     MerchantRes[V]
		body    V
	)
	defer func() {
		go upper.WriteApiLog(m.Merchant.ID, g, err)
	}()

	head := data.GetHead()
	head.Token = m.Token
	if m.Merchant.ProviderId != "" {
		head.ProviderId = m.Merchant.ProviderId
	}
	head.Timestamp = strconv.FormatInt(time.Now().UnixMilli(), 10)
	head.Sign = makeSign(data, m.Merchant.Secret)

	g = utils.NewGurl("POST", m.Merchant.ApiUrl+"/seamless/"+path).SetData(data)
	resData, err = g.Do()
	if err != nil {
		return body, 7, err
	}

	err = global.Json.Unmarshal(resData, &res)
	if err != nil {
		return body, 7, err
	}
	if res.Code != 0 {
		err = errors.New(res.Msg)
		return body, res.Code, err
	}

	return res.Data, 0, nil
}

// CheckMerchantTokenByAgent 检查商户token并创建用户 且必须效验商户
func CheckMerchantTokenByAgent(agent string, token string) (u *business.User, m *business.Merchant, err error) {
	m, err = cache.GetMerchant(agent)
	if err != nil {
		return
	}
	u, err = CheckMerchantToken(m, token)
	if err != nil {
		return
	}
	return
}

// CheckMerchantToken 检查商户token并创建用户
func CheckMerchantToken(merchant *business.Merchant, token string) (u *business.User, err error) {
	if len(token) == 0 {
		err = enum.ErrTokenInvalid
		return
	}
	// 先从缓存中获取
	u, err = cache.GetMerchantUserByCache(merchant.Agent, token)
	if err == nil {
		return
	}

	// 请求商户端的鉴权接口
	player, _, err := NewMerchantReq(merchant, token).Authenticate()
	if err != nil {
		err = errors.New("request authenticate error: " + err.Error())
		return
	}

	// 检查商户是否支持该币种
	if !merchant.CheckCurrency(player.Currency) {
		return nil, errors.New("currency not support [" + player.Currency + "]")
	}

	// 保存用户信息
	u, err = upper.GetOrCreateUser(merchant, player.PlayerName, 0, strings.ToUpper(player.Currency))
	if err != nil {
		return
	}
	u.Token = token
	err = cache.SetMerchantUserCache(merchant.Agent, token, u)
	return
}

// PriorityCheckMerchantToken 优先效验商户token 如果存在可不获取merchant
func PriorityCheckMerchantToken(agent string, token string) (u *business.User, err error) {
	if len(token) == 0 {
		err = enum.ErrTokenInvalid
		return
	}
	// 先从缓存中获取 存在则直接退出不效验商户
	u, err = cache.GetMerchantUserByCache(agent, token)
	if err == nil {
		return
	}

	var merchant *business.Merchant
	merchant, err = cache.GetMerchant(agent)
	if err != nil {
		return
	}

	// 请求商户端的鉴权接口
	player, _, err := NewMerchantReq(merchant, token).Authenticate()
	if err != nil {
		err = errors.New("request authenticate error: " + err.Error())
		return
	}

	// 保存用户信息
	u, err = upper.GetOrCreateUser(merchant, player.PlayerName, player.PlayerBalance, player.Currency)
	if err != nil {
		return
	}
	u.Token = token
	err = cache.SetMerchantUserCache(merchant.Agent, token, u)
	return
}
