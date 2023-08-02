package cache

import (
	"fmt"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/utils"
	"time"
)

// GetMerchant 根据商户id或agent获取商户信息
func GetMerchant(anyKey any) (m *business.Merchant, err error) {
	key := GetMerchantAnyKey(anyKey)
	err = utils.GetCache(key, &m)
	if err == nil {
		return m, nil
	}

	var where []any
	switch k := anyKey.(type) {
	case string:
		where = []any{"agent = ?", k}
	case uint, int:
		where = []any{"id = ?", k}
	default:
		return nil, fmt.Errorf("merchant: %v not found", anyKey)
	}

	err = global.GVA_DB.First(&m, where...).Error
	if err != nil {
		return nil, fmt.Errorf("merchant: %v not found", anyKey)
	}

	SetMerchantCache(m)
	return
}

func GetMerchantMap(ids ...uint) map[uint]*business.Merchant {
	m := map[uint]*business.Merchant{}
	for _, id := range ids {
		if _, ok := m[id]; ok {
			continue
		}
		m[id], _ = GetMerchant(id)
	}
	return m
}

func GetMerchantAnyKey(key any) string {
	switch k := key.(type) {
	case string:
		return fmt.Sprintf("{merchant}:agent_%s", k)
	case uint, int:
		return fmt.Sprintf("{merchant}:id_%d", k)
	}
	return ""
}

func SetMerchantCache(m *business.Merchant) {
	utils.SetCache(GetMerchantAnyKey(m.ID), m, 24*time.Hour)
	utils.SetCache(GetMerchantAnyKey(m.Agent), m, 24*time.Hour)
}

func DelMerchantCache(m *business.Merchant) error {
	return utils.DelCache(GetMerchantAnyKey(m.ID), GetMerchantAnyKey(m.Agent))
}

func DelMerchantCacheById(id uint) error {
	var merchant business.Merchant
	err := global.GVA_DB.First(&merchant, id).Error
	if err != nil {
		return err
	}
	return DelMerchantCache(&merchant)
}

//func DelMerchantCache(key any) {
//	FuzzyDel(GetMerchantAnyKey(key))
//}

//func ClearMerchantCache() {
//	FuzzyDel("merchant:*")
//}
