package cache

import (
	"context"
	"gorm.io/gorm"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/utils/env"
	"strconv"
	"time"
)

// getUserSpinCacheKey 获取Dome用户游玩统计key
func getUserSpinCacheKey(userId int64, slotId uint) string {
	return "{slot_user_spin}:" + env.IPLast + ":" + strconv.Itoa(int(userId)) + ":" + strconv.Itoa(int(slotId))
}

// getUserSpinLockKey 获取用户游玩锁key
func getUserSpinLockKey(userId uint) string {
	return "{slot_user_lock}:" + strconv.Itoa(int(userId))
}

// SetUserSpinLock 设置用户游玩锁
func SetUserSpinLock(userId uint) bool {
	key := getUserSpinLockKey(userId)
	ok, err := global.GVA_REDIS.SetNX(context.Background(), key, 1, 10*time.Second).Result()
	if err != nil {
		return false
	}
	return ok
}

// DelUserSpinLock 删除用户游玩锁
func DelUserSpinLock(userId uint) {
	global.GVA_REDIS.Del(context.Background(), getUserSpinLockKey(userId))
}

// GetSlotUserSpinInfo 获取用户游玩统计
func GetSlotUserSpinInfo(sessionId int64, slotId uint) (res *business.SlotUserSpin, err error) {
	res, err = GetUserSpinInfoByCache(sessionId, slotId)
	if err != nil {
		res.SlotId = slotId
		res.UserId = sessionId
		if err == nil {
			_ = SetUserSpinInfoCache(res)
		}
		return
	}
	return
}

// GetUserSpinInfoByCache 获取Dome用户游玩统计缓存
func GetUserSpinInfoByCache(userId int64, slotId uint) (res *business.SlotUserSpin, err error) {
	key := getUserSpinCacheKey(userId, slotId)
	res = &business.SlotUserSpin{}
	err = global.GVA_REDIS.HGetAll(context.Background(), key).Scan(res)
	return
}

// SetUserSpinInfoCache 设置Dome用户游玩统计缓存
func SetUserSpinInfoCache(s *business.SlotUserSpin) (err error) {
	key := getUserSpinCacheKey(s.UserId, s.SlotId)
	data := map[string]interface{}{
		"freeNum": s.FreeNum,
		"playNum": s.PlayNum,
	}
	err = global.GVA_REDIS.HSet(context.Background(), key, data).Err()
	return
}

// UserSpinInfoInc 用户游玩统计
func UserSpinInfoInc(userId int64, slotId uint, free, freeNum, playNum int64, tx *gorm.DB) (err error) {
	m := map[string]int64{
		"free":    free,
		"freeNum": freeNum,
		"playNum": playNum,
	}
	//err = business.AddUserSpinNum(userId, slotId, int(free), int(freeNum), int(playNum), tx)
	//if err != nil {
	//	return
	//}
	key := getUserSpinCacheKey(userId, slotId)
	err = HMapIncr(key, m)
	if err == nil {
		global.GVA_REDIS.Expire(context.Background(), key, 15*time.Minute)
	}
	return
}
