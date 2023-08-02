package initialize

import (
	"context"
	"fmt"
	"slot-server/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.GVA_CONFIG.Redis
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{redisCfg.Addr},
		Password: redisCfg.Password, // no password set
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("redis connect ping failed, err:%s", err.Error()))
	} else {
		global.GVA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.GVA_REDIS = client
	}
}
