package cache

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/proto"
	"slot-server/global"
	"slot-server/pbs"
	"time"
)

// GetFreeSpinAckCacheKey 获取用户免费玩结果key
func GetFreeSpinAckCacheKey(userId, slotId uint, typ int) string {
	return fmt.Sprintf("{free_spin}:%d:%d:%d", typ, userId, slotId)
}

// PushFreeSpinAckCache 推送用户免费玩结果
func PushFreeSpinAckCache(userId, slotId uint, typ int, acks ...*pbs.SpinAck) error {
	var (
		arr []any
		key = GetFreeSpinAckCacheKey(userId, slotId, typ)
	)
	for _, ack := range acks {
		// 使用proto序列化
		i, err := proto.Marshal(ack)
		if err != nil {
			return err
		}
		arr = append(arr, i)
	}
	err := global.GVA_REDIS.LPush(context.Background(), key, arr...).Err()
	if err != nil {
		return err
	}
	global.GVA_REDIS.Expire(context.Background(), key, 15*time.Minute)
	return nil
}

func PopFreeSpinAckCache(userId, slotId uint, typ int, ack *pbs.SpinAck) error {
	res, err := global.GVA_REDIS.RPop(context.Background(), GetFreeSpinAckCacheKey(userId, slotId, typ)).Bytes()
	if err != nil {
		return err
	}
	err = proto.Unmarshal(res, ack)
	return err
}

func DelFreeSpinAckCache(userId uint, typ int) error {
	key := fmt.Sprintf("free_spin:%d:%d:*", typ, userId)
	keys := global.GVA_REDIS.Keys(context.Background(), key).Val()
	if len(keys) == 0 {
		return nil
	}
	return global.GVA_REDIS.Del(context.Background(), keys...).Err()
}
