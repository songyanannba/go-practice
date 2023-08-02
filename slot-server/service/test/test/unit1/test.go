package unit1

import (
	"slot-server/model/business"
	"slot-server/service/slot"
	"slot-server/service/slot/component"
	"slot-server/service/test/public"
	"strconv"
)

type unit1TestResult struct {
	public.BasicData
	n27 int //jackpot_1赢钱次数:::::
	n28 int //jackpot_1赢钱总数:::::
	n29 int //jackpot_2赢钱次数:::::
	n30 int //jackpot_2赢钱总数:::::
	n31 int //jackpot_3赢钱次数:::::
	n32 int //jackpot_3赢钱总数:::::
	n33 int //jackpot_4赢钱次数:::::
	n34 int //jackpot_4赢钱总数:::::
}

var statistics = make([]map[string]int, 0)

func TestSlot(slotTest *business.SlotTests, slotId uint, Num int, Amount int, opts ...component.Option) {
	var testSlotResult unit1TestResult
	defer public.EndProcessing(slotTest, &testSlotResult.BasicData)
	statistics = make([]map[string]int, 0)
	for i := 0; i < Num; i++ {
		m, _ := slot.Play(slotId, Amount, opts...)
		spin := m.GetSpin()
		AnalyticData(&testSlotResult, spin)
	}

	str := public.GetInitResult(&testSlotResult.BasicData)
	slotTest.Detail = str + `
jackpot_1赢钱次数:::::` + strconv.Itoa(testSlotResult.n27) + `
jackpot_1赢钱总数:::::` + strconv.Itoa(testSlotResult.n28) + `
jackpot_2赢钱次数:::::` + strconv.Itoa(testSlotResult.n29) + `
jackpot_2赢钱总数:::::` + strconv.Itoa(testSlotResult.n30) + `
jackpot_3赢钱次数:::::` + strconv.Itoa(testSlotResult.n31) + `
jackpot_3赢钱总数:::::` + strconv.Itoa(testSlotResult.n32) + `
jackpot_4赢钱次数:::::` + strconv.Itoa(testSlotResult.n33) + `
jackpot_4赢钱总数:::::` + strconv.Itoa(testSlotResult.n34) + `
`
	//拼接统计数据
	for i, v := range statistics {
		slotTest.Detail += "第" + strconv.Itoa(i+1) + "列标签出现次数统计:\r"
		for vv, ii := range v {
			slotTest.Detail += vv + ":::::" + strconv.Itoa(ii) + "\r"
		}
	}

}

func AnalyticData(testSlotResult *unit1TestResult, spin *component.Spin) {
	public.AnalyticData(&testSlotResult.BasicData, []*component.Spin{spin})
	if spin.Jackpot != nil {
		if spin.Jackpot.Id == 1 {
			testSlotResult.n27 += 1
			testSlotResult.n28 += spin.Gain
		} else if spin.Jackpot.Id == 2 {
			testSlotResult.n29 += 1
			testSlotResult.n30 += spin.Gain
		} else if spin.Jackpot.Id == 3 {
			testSlotResult.n31 += 1
			testSlotResult.n32 += spin.Gain
		} else if spin.Jackpot.Id == 4 {
			testSlotResult.n33 += 1
			testSlotResult.n34 += spin.Gain
		}
	}
	//初始化标签统计
	for i := 0; len(statistics) < spin.Config.Index; i++ {
		var tags = make(map[string]int)
		for _, v := range spin.Config.GetAllTag() {
			tags[v.Name] = 0
		}
		statistics = append(statistics, tags)
	}
	//统计标签
	if spin.Config.SlotId == 1 {
		for i, v := range spin.InitDataList[1] {
			tags := statistics[i]
			tags[v.Name] += 1
		}
	} else {
		for i, v := range spin.InitDataList {
			tags := statistics[i]
			for _, vv := range v {
				tags[vv.Name] += 1
			}
		}
	}

}
