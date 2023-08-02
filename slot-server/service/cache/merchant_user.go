package cache

import (
	"fmt"
	"slot-server/model/business"
	"slot-server/utils"
	"time"
)

func GetMerchantUserKey(agent string, token string) string {
	return fmt.Sprintf("{merchant_user}:%s:%s", agent, token)
}

func GetMerchantUserByCache(agent string, token string) (u *business.User, err error) {
	err = utils.GetCache(GetMerchantUserKey(agent, token), &u)
	return
}

func SetMerchantUserCache(agent string, token string, u *business.User) error {
	return utils.SetCache(GetMerchantUserKey(agent, token), u, 24*time.Hour)
}

//func ClearMerchantUserCache() {
//	FuzzyDel("merchant_user:*")
//}
