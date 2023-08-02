package test

import (
	"github.com/shopspring/decimal"
	"slot-server/model/business"
	"slot-server/service/slot/component"
	"slot-server/service/test/public"
	unitTest5 "slot-server/service/test/test/unit5"
	unitTest6 "slot-server/service/test/test/unit6"
	unitTest7 "slot-server/service/test/test/unit7"
	unitTest8 "slot-server/service/test/test/unit8"
)

// Death
//
//	@Description: 死循环测试
//	@param slotTest 测试结果保存位置
//	@param slotId 机台id
//	@param Num 测试次数
//	@param Amount 测试金额
//	@param Hold 测试保持次数
//	@param opts 测试选项
func Death(slotTest *business.SlotTests, run public.RunSlotTest, opts ...component.Option) {
	//public.TestDeath(slotTest, run.SlotId, run.Num, run.Amount, run.Hold, opts...)
	err := DeathTest(slotTest, run, opts...)
	if err != nil {
		return
	}
}

// Once
//
//	@Description: 单次测试
//	@param slotId 机台id
//	@param amount 测试金额
//	@param opts 测试选项
//	@return err 错误信息
func Once(run public.RunSlotTest, opts ...component.Option) (err error) {
	switch run.SlotId {
	case 5:
		return unitTest5.TestOnce(run, opts...)
	case 6:
		return unitTest6.TestOnce(run, opts...)
	case 7:
		return unitTest7.TestOnce(run, opts...)
	case 8:
		return unitTest8.TestOnce(run, opts...)
	default:
		return public.TestOnce(run, opts...)
	}

}

// Appoint
//
//	@Description: 指定次数测试
//	@param slotTest 测试结果保存位置
//	@param slotId 机台id
//	@param Num 测试次数
//	@param Amount 测试金额
//	@param opts 测试选项
func Appoint(slotTest *business.SlotTests, slotId uint, Num int, Amount int, run public.RunSlotTest, opts ...component.Option) {
	if run.Raise == 1 {
		c, _ := component.GetSlotConfig(slotId, false)
		raise := decimal.NewFromInt(int64(Amount)).Mul(decimal.NewFromFloat(c.Raise)).IntPart()
		opts = append(opts, component.WithRaise(raise))
	}
	if run.IsMustFree > 0 {
		opts = append(opts, component.WithIsMustFree())
	}
	if run.IsMustRes > 0 {
		opts = append(opts, component.WithIsMustRes())
	}

	err := RepeatTest(slotTest, run, opts...)
	if err != nil {
		return
	}
}

// User 真实用户测试
func User(run public.RunSlotTest, opts ...component.Option) (err error) {
	return public.TestUser(run, opts...)
}
