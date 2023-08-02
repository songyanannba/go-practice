package seamlessWallet

import (
	"errors"
	"github.com/fatih/structs"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/service/cache"
	"slot-server/utils"
	"slot-server/utils/helper"
)

type AcceptHeadData interface {
	GetSign() string
	GetAcceptHead() AcceptHead
	CheckParams() error
}

// AcceptHead 三方请求我们时的请求头
type AcceptHead struct {
	Agent     string `json:"agent" example:"test"`                            // 通用参数 运营商id
	Sign      string `json:"sign" example:"c623f86755262355f3b948ba96ae705e"` // 通用参数 签名
	Timestamp string `json:"timestamp" example:"1612345678000"`               // 通用参数 时间戳
}

func (h AcceptHead) GetSign() string {
	return h.Sign
}

func (h AcceptHead) GetAcceptHead() AcceptHead {
	return h
}

func (h AcceptHead) CheckHead() error {
	if h.Agent == "" {
		return errors.New("agent is empty")
	}
	if h.Sign == "" {
		return errors.New("sign is empty")
	}
	if h.Timestamp == "" {
		return errors.New("timestamp is empty")
	}
	return nil
}

func (h AcceptHead) CheckSign(params any) (*business.Merchant, error) {
	merchant, err := cache.GetMerchant(h.Agent)
	if err != nil {
		return nil, err
	}
	if err = CheckSign(params, h.Sign, merchant.Secret); err != nil {
		return merchant, err
	}
	return merchant, nil
}

func makeSign(data any, secret string) string {
	var (
		s = structs.New(data)
		m = make(map[string]any)
	)
	for _, f := range s.Fields() {
		name := f.Tag("json")
		if name == "" {
			name = f.Name()
		}
		if name == "AcceptHead" || name == "ReqHead" {
			for _, v := range f.Fields() {
				name = v.Tag("json")
				if name == "" {
					name = v.Name()
				}
				m[name] = v.Value()
			}
			continue
		}
		if f.IsExported() {
			m[name] = f.Value()
		}
	}
	return utils.Md5Sign(m, secret)
}

func CheckSign(data any, oSign, secret string) error {
	sign := makeSign(data, secret)
	if oSign != sign {
		global.GVA_LOG.Warn("original sign: " + oSign + " ,need sign: " + sign)
		return errors.New("sign authentication failed")
	}
	return nil
}

func CheckLanguage(language string) error {
	if !helper.InArr(language, []string{"", "en"}) {
		return errors.New("language " + language + " is not supported")
	}
	return nil
}
