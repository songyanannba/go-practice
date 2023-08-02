package cache

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/utils"
	"strconv"
	"time"
)

func GetUserAmountCacheKey(id uint) string {
	return utils.GetCacheKey(utils.PlaceString, "{user_amount}", strconv.Itoa(int(id)))
}

func GetUserIntoGameTimeCacheKey(id uint) string {
	return utils.GetCacheKey(utils.PlaceString, "{user_into_time}", strconv.Itoa(int(id)))
}

// GetUserAmount 获取用户余额并缓存
func GetUserAmount(id uint) (int64, error) {
	amount, err := global.GVA_REDIS.Get(context.Background(), GetUserAmountCacheKey(id)).Int64()
	if err == nil {
		return amount, nil
	}
	err = global.GVA_DB.Model(&business.User{}).Where("id = ?", id).Pluck("amount", &amount).Error
	if err != nil {
		return 0, err
	}
	global.GVA_REDIS.Set(context.Background(), GetUserAmountCacheKey(id), amount, 15*time.Minute)
	return amount, nil
}

func GetUserMoneyLockKey(id uint) string {
	return utils.GetCacheKey(utils.PlaceLock, "user_money", strconv.Itoa(int(id)))
}

func SetUserMoneyLock(id uint) bool {
	ok, err := global.GVA_REDIS.SetNX(context.Background(), GetUserMoneyLockKey(id), 1, time.Second*10).Result()
	if err != nil {
		return false
	}
	return ok
}

func DelUserMoneyLock(id uint) error {
	return global.GVA_REDIS.Del(context.Background(), GetUserMoneyLockKey(id)).Err()
}

// ChangeMoney 添加金币流水
func ChangeMoney(userId uint, change int64, tx *gorm.DB, opts ...business.MoneyLogOption) (*business.MoneyLog, error) {
	if change == 0 {
		return &business.MoneyLog{}, nil
	}
	ok := SleepRetry(SetUserMoneyLock, userId, 200*time.Millisecond, 3)
	if !ok {
		global.GVA_LOG.Error("快速流水变动", zap.Uint("user_id", userId), zap.Int64("change", change))
		return nil, enum.ErrBusy
	}
	defer DelUserMoneyLock(userId)
	if tx == nil {
		tx = global.GVA_DB
	}
	resVal, err := changeAmount(userId, change, tx)
	if err != nil {
		return nil, err
	}
	// 默认为赠送操作
	moneyLog := business.MoneyLog{
		Date:        time.Now().Format("20060102"),
		UserId:      userId,
		Action:      enum.MoneyAction3System,
		ActionType:  enum.MoneyType3Give,
		CoinInitial: resVal - change,
		CoinChange:  change,
		CoinResult:  resVal,
	}
	for _, opt := range opts {
		opt(&moneyLog)
	}
	err = tx.Create(&moneyLog).Error
	return &moneyLog, err
}

// changeAmount 变更金额并返回余额
func changeAmount(id uint, change int64, tx *gorm.DB) (val int64, err error) {
	// 先把余额加入缓存
	val, err = GetUserAmount(id)
	if err != nil || change == 0 {
		return
	}
	val, err = global.GVA_REDIS.IncrBy(context.Background(), GetUserAmountCacheKey(id), change).Result()
	if err != nil {
		return
	}
	if tx == nil {
		tx = global.GVA_DB
	}
	err = tx.Model(&business.User{}).Where("id = ?", id).Update("amount", gorm.Expr("amount + ?", change)).Error
	if err != nil {
		global.GVA_REDIS.Del(context.Background(), GetUserAmountCacheKey(id))
		return
	}
	return
}

// GetUserIntoGameTime 获取用户进入游戏时间
func GetUserIntoGameTime(id uint) (int64, error) {
	t, err := global.GVA_REDIS.Get(context.Background(), GetUserIntoGameTimeCacheKey(id)).Int64()
	return t, err
}

// SetUserIntoGameTime 设置用户进入游戏时间
func SetUserIntoGameTime(id uint, t int64) error {
	err := global.GVA_REDIS.Set(context.Background(), GetUserIntoGameTimeCacheKey(id), t, 3*24*time.Hour).Err()
	return err
}
