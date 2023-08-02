package public

import (
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/service/slot"
	"slot-server/service/slot/component"
	"strconv"
)

func TestDeath(slotTest *business.SlotTests, slotId uint, Num int, Amount int, Hold int, opts ...component.Option) {
	defer EndProcessing(slotTest, &BasicData{N1: 0, N3: 0})
	str := ""
	for a := 0; a < Num; a++ {
		count := 0
		HoldCop := Hold
		AmountCop := Amount
		for i := 0; i < 100000 && HoldCop >= AmountCop; i++ {
			HoldCop -= AmountCop
			count++
			m, err := slot.Play(slotId, AmountCop, opts...)
			if err != nil {
				global.GVA_LOG.Error("死亡测试错误:" + err.Error())
				return
			}
			spin := m.GetSpin()
			HoldCop += spin.Gain
		}
		str += "编号:" + strconv.Itoa(a+1) + " 余额:" + strconv.Itoa(HoldCop) + " 死亡次数:" + strconv.Itoa(count) + "\r\n"
	}
	slotTest.Detail = str
}
