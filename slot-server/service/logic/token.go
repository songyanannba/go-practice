package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	enum "slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/pbs"
	"slot-server/service/cache"
	"slot-server/service/logic/upper/seamlessWallet"
	"slot-server/utils"
	"slot-server/utils/env"
	"strconv"
	"time"
)

const (
	tokenSecret = "young people998"          //加密密钥
	tokenIssuer = "dhn998"                   //加密签名
	tokenExpire = 24 * 60 * 60 * time.Second //token老化时间
)

type Claims struct {
	Uid         uint `json:"uid"`
	UUID        string
	Username    string
	NickName    string
	AuthorityId uint
	jwt.StandardClaims
}

func (c Claims) Valid() error {
	return nil
}

func GenerateToken(userId uint) (string, error) {
	now := time.Now()
	exp := now.AddDate(0, 0, 2).Unix()
	claims := Claims{
		Uid: userId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: now.Unix(),
			ExpiresAt: exp,
			Issuer:    tokenIssuer,
			IssuedAt:  now.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenStr, _ := token.SignedString([]byte(tokenSecret))
	// 顶号通知，旧token踢下线
	//singleton.HallOfflineUserReq(userId)
	err := global.GVA_REDIS.Set(context.Background(), AuthUserCacheKey(claims.Uid), tokenStr, tokenExpire).Err()
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ParseToken(tokenStr string) (claims *Claims, err error) {
	keyFn := func(token *jwt.Token) (interface{}, error) {
		// 验证加密方式是否正确
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			global.GVA_LOG.Error("invalid signing method")
			return nil, errors.New("invalid signing method")
		}
		return []byte(tokenSecret), nil
	}

	// 解析
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, keyFn)
	if err != nil {
		//global.GVA_LOG.Warn("parse token warning: ", err)
		return nil, errors.New("parse token error")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		global.GVA_LOG.Warn("invalid token claims: ", zap.Any("claims", token.Claims))
		return nil, errors.New("invalid token claims")
	}
	// 验证是否过期
	if err := claims.StandardClaims.Valid(); err != nil {
		global.GVA_LOG.Info("token has expired: " + err.Error())
		return nil, err
	}
	// 验证签发人
	if !claims.VerifyIssuer(tokenIssuer, true) {
		global.GVA_LOG.Warn("invalid issuer: " + claims.Issuer)
		return nil, errors.New("invalid token issuer")
	}

	return claims, nil
}

func RequireLogin(token string) (uint, error) {
	if token == "" {
		return 0, enum.ErrTokenInvalid
	}

	if env.Mode != env.Prod && token == "123456" {
		return 1, nil
	}

	claims, err := ParseToken(token)
	if err != nil {
		return 0, enum.ErrTokenInvalid
	}

	cmd := global.GVA_REDIS.Get(context.Background(), AuthUserCacheKey(claims.Uid))
	if cmd.Err() != nil {
		return 0, enum.ErrTokenInvalid
	}

	cacheToken := cmd.Val()
	if token != cacheToken {
		return 0, enum.ErrTokenInvalid
	}
	//SetUserSession(claims.Uid, ses)
	return claims.Uid, nil
}

func RequireMerchantToken(head *pbs.ReqHead, needMerchant bool) (u *business.User, m *business.Merchant, err error) {
	u = &business.User{}
	m = &business.Merchant{}
	if head == nil || len(head.Token) == 0 {
		err = enum.ErrTokenInvalid
		return
	}

	if env.Mode != env.Prod && head.Token == "123456" {
		u.ID = 1
		u.Token = head.Token
		u.Username = "Guest_10001"
		u.GetCurrency()
		if needMerchant {
			m, err = cache.GetMerchant(1)
		}
		return
	}

	if head.Platform == "" {
		// 平台为空则走自身平台的token验证
		u.ID, err = RequireLogin(head.Token)
		if err != nil {
			return
		}
		u.Token = head.Token
		u.Username = "Guest_" + strconv.Itoa(int(10000+u.ID))
		u.GetCurrency()
		if needMerchant {
			m, err = cache.GetMerchant(1)
		}
		return
	}

	if needMerchant {
		// 必须获取商户
		u, m, err = seamlessWallet.CheckMerchantTokenByAgent(head.Platform, head.Token)
	} else {
		// 可以不获取商户
		u, err = seamlessWallet.PriorityCheckMerchantToken(head.Platform, head.Token)
	}
	if err != nil {
		err = enum.ErrTokenInvalid
		return
	}
	return
}

func CleanUser(userId uint) error {
	if err := global.GVA_REDIS.Del(context.Background(), AuthUserCacheKey(userId)).Err(); err != nil {
		global.GVA_LOG.Error("clean user error: " + err.Error())
		return err
	}
	return nil
}

func AuthUserCacheKey(userId uint) string {
	// branch:string:user:auth:userId:123
	return fmt.Sprintf("%s:user:auth:userId:%d", utils.PlaceString, userId)
}
