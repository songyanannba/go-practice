package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"slot-server/global"
	"time"
)

func Pipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return global.GVA_REDIS.TxPipelined(context.Background(), fn)
}

func HMapIncr(key string, m map[string]int64) (err error) {
	for k, v := range m {
		if v == 0 {
			delete(m, k)
		}
	}
	if len(m) == 0 {
		return nil
	}
	_, err = Pipelined(func(pipe redis.Pipeliner) error {
		for k, v := range m {
			err = pipe.HIncrBy(context.Background(), key, k, v).Err()
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		global.GVA_REDIS.Del(context.Background(), key)
	}
	return
}

// SleepRetry 睡眠一定时间后重试
func SleepRetry[T any](fn func(T) bool, key T, t time.Duration, max int) bool {
	for i := 0; i < max; i++ {
		if fn(key) {
			return true
		}
		time.Sleep(t)
	}
	return false
}

// FuzzyDel 模糊删除 不建议使用
func FuzzyDel(key string) {
	if key == "" {
		return
	}
	keys := global.GVA_REDIS.Keys(context.Background(), key).Val()
	if len(keys) > 0 {
		global.GVA_REDIS.Del(context.Background(), keys...)
	}
}
