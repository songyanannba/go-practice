package logic

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
	enum "slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/pbs"
	"slot-server/service/cache"
	"slot-server/utils"
	"slot-server/utils/helper"
	"strings"
)

func SignUp(ip string, req *pbs.LoginReq, ack *pbs.LoginAck) (user *business.User, err error) {
	user = &business.User{
		Username:   "",
		Password:   "123456",
		Uuid:       uuid.NewV4().String(),
		NickName:   "123",
		Phone:      req.Phone,
		Email:      "",
		Status:     enum.UserStatus1Normal,
		Ip:         ip,
		LastIp:     ip,
		MerchantId: 1,
		Currency:   helper.If(req.Currency == "", "USD", req.Currency),
	}
	m, err := cache.GetMerchant(1)
	if err != nil {
		return nil, NewErr(pbs.Code_NotExists, errors.New("the test merchant is not available"))
	}
	if !m.CheckCurrency(user.Currency) {
		return nil, NewErr(pbs.Code_SystemError, errors.New("currency does not support"))
	}
	// 附加uuid 密码hash加密 注册
	user.Password = utils.BcryptHash(user.Password)
	user.Uuid = uuid.NewV4().String()
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Create(user).Error
		if err != nil {
			return NewErr(pbs.Code_DbError, err)
		}
		user.Username = "Guest_" + fmt.Sprintf("%d", user.ID+10000)
		user.NickName = user.Username
		err = tx.Select("username", "nick_name").Updates(user).Error
		if err != nil {
			return NewErr(pbs.Code_DbError, err)
		}
		return nil
	})
	if err != nil {
		global.GVA_LOG.Error("register user failed", zap.Error(err))
		return user, NewErr(pbs.Code_DbError, err)
	}
	token, _ := GenerateToken(user.ID)
	ack.Token = token
	ack.Uid = int32(user.ID)
	ack.Username = user.Username
	ack.Currency = user.Currency
	return
}

func Login(ip string, req *pbs.LoginReq, ack *pbs.LoginAck) (user *business.User, err error) {
	verify := &business.User{
		Username: strings.TrimSpace(req.Username),
		Password: strings.TrimSpace(req.Password),
	}
	err = verify.ValidateLogin()
	if err != nil {
		err = NewErr(pbs.Code_ParameterError, err)
		return
	}
	user = &business.User{}
	err = global.GVA_DB.Where("username = ?", verify.Username).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = NewErr(pbs.Code_NotExists, errors.New("user not exists"))
			return
		}
		err = NewErr(pbs.Code_DbError, err)
		return
	}

	if ok := utils.BcryptCheck(verify.Password, user.Password); !ok {
		err = NewErr(pbs.Code_ParameterError, errors.New("wrong password"))
		return
	}
	if user.Status != enum.UserStatus1Normal {
		err = NewErr(pbs.Code_StatusError, errors.New("abnormal user status"))
		return
	}
	token, err := GenerateToken(user.ID)
	if err != nil {
		err = NewErr(pbs.Code_SystemError, err)
		return
	}
	if ip != user.LastIp {
		user.LastIp = ip
		global.GVA_DB.Select("last_ip").Updates(&user)
	}
	ack.Token = token
	ack.Uid = int32(user.ID)
	ack.Amount = user.Amount
	ack.Currency = user.Currency
	return
}

func DemoLogin(ack *pbs.LoginAck) {
	ack.Amount = 1000000000
	ack.Username = "DemoGuest"
	ack.GameList = cache.GetGameList()
	ack.Currency = "USD"
	return
}

func MerchantDemoLogin(ack *pbs.MerchantLoginAck) {
	ack.Amount = 1000000000
	ack.Username = "DemoGuest"
	ack.GameList = cache.GetGameList()
	ack.Currency = "USD"
	return
}
