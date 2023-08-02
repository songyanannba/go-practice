package tests

import (
	"context"
	"slot-server/core"
	"slot-server/global"
	"slot-server/utils/helper"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	core.BaseInit()
	//cache.ChangeMoney(1, 100, nil)
	for i := 0; i < 5; i++ {
		s := "test" + helper.Itoa(i)
		err := global.GVA_REDIS.Set(context.Background(), s, s, 60*time.Second).Err()
		t.Log(err)
	}
	for i := 0; i < 5; i++ {
		s := "test" + helper.Itoa(i)
		a, err := global.GVA_REDIS.Get(context.Background(), s).Result()
		t.Log(a, err)
	}
}
