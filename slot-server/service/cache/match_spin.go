package cache

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/proto"
	"slot-server/global"
	"slot-server/pbs"
	"time"
)

// GetMatchSpinAckCacheKey 获取用户免费玩结果key
func GetMatchSpinAckCacheKey(userId, slotId uint, typ int) string {
	return fmt.Sprintf("{match_spin}:%d:%d:%d", typ, userId, slotId)
}

// PushMatchSpinAckCache  推送用户免费玩结果
func PushMatchSpinAckCache(userId, slotId uint, typ int, acks ...*pbs.MatchSpinAck) error {
	var (
		arr []any
		key = GetMatchSpinAckCacheKey(userId, slotId, typ)
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

func PopMatchSpinAckCache(userId, slotId uint, typ int, ack *pbs.MatchSpinAck) error {
	res, err := global.GVA_REDIS.RPop(context.Background(), GetMatchSpinAckCacheKey(userId, slotId, typ)).Bytes()
	if err != nil {
		return err
	}
	err = proto.Unmarshal(res, ack)
	return err
}

func DelMatchAckCache(userId uint, typ int) error {
	key := fmt.Sprintf("match_spin:%d:%d:*", typ, userId)
	keys := global.GVA_REDIS.Keys(context.Background(), key).Val()
	if len(keys) == 0 {
		return nil
	}
	return global.GVA_REDIS.Del(context.Background(), keys...).Err()
}
