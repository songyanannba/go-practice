package tests

import (
	"fmt"
	"go.uber.org/zap"
	"slot-server/core"
	"slot-server/global"
	"slot-server/service/slot"
	"testing"
)

func TestRandSlot6(t *testing.T) {
	core.BaseInit()
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		_, err := slot.Play(6, 100) // 6
		if err != nil {
			global.GVA_LOG.Error("slot.Play err", zap.Error(err))
		}
		//fmt.Println(m)
	}
}
