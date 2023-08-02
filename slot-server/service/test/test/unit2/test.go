package unit2

import (
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"slot-server/model/business"
	"slot-server/service/slot"
	"slot-server/service/slot/component"
	"slot-server/service/slot/machine/unit2"
	"slot-server/service/test/public"
	"strconv"
)

type unit2TestResult struct {
	public.BasicData
	n27 int     //Frees pin划线赢钱0-5倍的次数:::::
	n28 int     //Frees pin划线赢钱0-5倍的总钱数:::::
	n29 int     //Frees pin划线赢钱5-10倍的次数:::::
	n30 int     //Frees pin划线赢钱5-10倍的总钱数:::::
	n31 int     //Frees pin划线赢钱10-20倍的次数:::::
	n32 int     //Frees pin划线赢钱10-20倍的总钱数:::::
	n33 int     //Frees pin划线赢钱20-30倍的次数:::::
	n34 int     //Frees pin划线赢钱20-30倍的总钱数:::::
	n35 int     //Frees pin划线赢钱30-40倍的次数:::::
	n36 int     //Frees pin划线赢钱30-40倍的总钱数:::::
	n37 int     //Frees pin划线赢钱40-50倍的次数:::::
	n38 int     //Frees pin划线赢钱40-50倍的总钱数:::::
	n39 int     //Frees pin划线赢钱50-100倍的次数:::::
	n40 int     //Frees pin划线赢钱50-100倍的总钱数:::::
	n41 int     //Frees pin划线赢钱100-200倍的次数:::::
	n42 int     //Frees pin划线赢钱100-200倍的总钱数:::::
	n43 int     //Frees pin划线赢钱200-500倍的次数:::::
	n44 int     //Frees pin划线赢钱200-500倍的总钱数:::::
	n45 int     //Frees pin划线赢钱500-1000倍的次数:::::
	n46 int     //Frees pin划线赢钱500-1000倍的总钱数:::::
	n47 int     //Frees pin划线赢钱1000以上倍的次数:::::
	n48 int     //Frees pin划线赢钱1000以上倍的总钱数:::::
	n49 int     //普通转触发Frees pin玩法次数:::::
	n50 int     //Frees pin玩法总spin转动次数:::::
	n51 int     //Frees pin内再次触发Frees pin玩法续命的次数::::
	n52 int     //Frees pin玩法内随机出1个Wild标签的次数:::::
	n53 int     //Frees pin玩法内随机出1个Wild标签的赢钱数:::::
	n54 int     //Frees pin玩法内随机出2个Wild标签的次数:::::
	n55 int     //Frees pin玩法内随机出2个Wild标签的赢钱数:::::
	n56 int     //Frees pin玩法内随机出3个Wild标签的次数:::::
	n57 int     //Frees pin玩法内随机出3个Wild标签的赢钱数:::::
	n58 int     //Frees pin玩法内随机出4个Wild标签的次数:::::
	n59 int     //Frees pin玩法内随机出4个Wild标签的赢钱数:::::
	n60 int     //Frees pin玩法内随机出5个Wild标签的次数:::::
	n61 int     //Frees pin玩法内随机出5个Wild标签的赢钱数:::::
	n62 int     //Frees pin玩法内随机出6个Wild标签的次数:::::
	n63 int     //Frees pin玩法内随机出6个Wild标签的赢钱数:::::
	n64 int     //Frees pin玩法内随机出7个Wild标签的次数:::::
	n65 int     //Frees pin玩法内随机出7个Wild标签的赢钱数:::::
	n66 int     //Frees pin玩法内随机出8个Wild标签的次数:::::
	n67 int     //Frees pin玩法内随机出8个Wild标签的赢钱数:::::
	n68 int     //Frees pin玩法内随机出9个Wild标签的次数:::::
	n69 int     //Frees pin玩法内随机出9个Wild标签的赢钱数:::::
	n70 int     //scatter 标签出现在第一列次数:::::
	n71 int     //scatter 标签出现在第二列次数:::::
	n72 int     //scatter 标签出现在第三列次数:::::
	n73 int     //scatter 标签出现在第四列次数:::::
	n74 int     //scatter 标签出现在第五列次数:::::
	n75 float64 //scatter high_1五个次数:::::
	n76 float64 //scatter high_1四个次数:::::
	n77 float64 //scatter high_1三个次数:::::
	n78 int     //Frees pin玩法内随机出0个Wild标签的次数:::::
	n79 int     //Frees pin玩法内随机出0个Wild标签的赢钱数:::::
}

var testSlotResult unit2TestResult

type A struct {
	spins []*component.Spin
}

func traversalTurntable(slotId uint, Amount int, a *A, opts ...component.Option) *component.Spin {
	m, _ := slot.Play(slotId, Amount, opts...)
	spin := m.GetSpin()
	AnalyticData(&testSlotResult, spin)
	if !spin.IsFree {
		a = &A{}
	}
	for i := 0; i < spin.FreeSpinParams.Count*8; i++ {
		opts := []component.Option{component.WithTest()}
		freeSpin := traversalTurntable(slotId, Amount, a, opts...)
		a.spins = append(a.spins, freeSpin)
	}
	if !spin.IsFree {
		freeSpinSum(&testSlotResult, a.spins, spin.Bet)
	}
	return spin
}
func freeSpinSum(result *unit2TestResult, spins []*component.Spin, bet int) {
	sumGain := lo.SumBy(spins, func(i *component.Spin) int {
		return i.Gain
	})
	if sumGain == 0 {
		return
	}
	multiple := float64(sumGain) / float64(bet)
	if multiple >= 0 && multiple < 5 {
		result.n27 += 1
		result.n28 += sumGain
	} else if multiple >= 5 && multiple < 10 {
		result.n29 += 1
		result.n30 += sumGain
	} else if multiple >= 10 && multiple < 20 {
		result.n31 += 1
		result.n32 += sumGain
	} else if multiple >= 20 && multiple < 30 {
		result.n33 += 1
		result.n34 += sumGain
	} else if multiple >= 30 && multiple < 40 {
		result.n35 += 1
		result.n36 += sumGain
	} else if multiple >= 40 && multiple < 50 {
		result.n37 += 1
		result.n38 += sumGain
	} else if multiple >= 50 && multiple < 100 {
		result.n39 += 1
		result.n40 += sumGain
	} else if multiple >= 100 && multiple < 200 {
		result.n41 += 1
		result.n42 += sumGain
	} else if multiple >= 200 && multiple < 500 {
		result.n43 += 1
		result.n44 += sumGain
	} else if multiple >= 500 && multiple < 1000 {
		result.n45 += 1
		result.n46 += sumGain
	} else if multiple >= 1000 {
		result.n47 += 1
		result.n48 += sumGain
	}
}
func TestSlot(slotTest *business.SlotTests, slotId uint, Num int, Amount int, opts ...component.Option) {
	testSlotResult = unit2TestResult{}
	defer public.EndProcessing(slotTest, &testSlotResult.BasicData)
	for i := 0; i < Num; i++ {
		traversalTurntable(slotId, Amount, &A{}, opts...)
	}
	str := public.GetInitResult(&testSlotResult.BasicData)
	slotTest.Detail = str + `
Frees pin划线赢钱0-5倍的次数:::::` + strconv.Itoa(testSlotResult.n27) + `
Frees pin划线赢钱0-5倍的总钱数:::::` + strconv.Itoa(testSlotResult.n28) + `
Frees pin划线赢钱5-10倍的次数:::::` + strconv.Itoa(testSlotResult.n29) + `
Frees pin划线赢钱5-10倍的总钱数:::::` + strconv.Itoa(testSlotResult.n30) + `
Frees pin划线赢钱10-20倍的次数:::::` + strconv.Itoa(testSlotResult.n31) + `
Frees pin划线赢钱10-20倍的总钱数:::::` + strconv.Itoa(testSlotResult.n32) + `
Frees pin划线赢钱20-30倍的次数:::::` + strconv.Itoa(testSlotResult.n33) + `
Frees pin划线赢钱20-30倍的总钱数:::::` + strconv.Itoa(testSlotResult.n34) + `
Frees pin划线赢钱30-40倍的次数:::::` + strconv.Itoa(testSlotResult.n35) + `
Frees pin划线赢钱30-40倍的总钱数:::::` + strconv.Itoa(testSlotResult.n36) + `
Frees pin划线赢钱40-50倍的次数:::::` + strconv.Itoa(testSlotResult.n37) + `
Frees pin划线赢钱40-50倍的总钱数:::::` + strconv.Itoa(testSlotResult.n38) + `
Frees pin划线赢钱50-100倍的次数:::::` + strconv.Itoa(testSlotResult.n39) + `
Frees pin划线赢钱50-100倍的总钱数:::::` + strconv.Itoa(testSlotResult.n40) + `
Frees pin划线赢钱100-200倍的次数:::::` + strconv.Itoa(testSlotResult.n41) + `
Frees pin划线赢钱100-200倍的总钱数:::::` + strconv.Itoa(testSlotResult.n42) + `
Frees pin划线赢钱200-500倍的次数:::::` + strconv.Itoa(testSlotResult.n43) + `
Frees pin划线赢钱200-500倍的总钱数:::::` + strconv.Itoa(testSlotResult.n44) + `
Frees pin划线赢钱500-1000倍的次数:::::` + strconv.Itoa(testSlotResult.n45) + `
Frees pin划线赢钱500-1000倍的总钱数:::::` + strconv.Itoa(testSlotResult.n46) + `
Frees pin划线赢钱1000以上倍的次数:::::` + strconv.Itoa(testSlotResult.n47) + `
Frees pin划线赢钱1000以上倍的总钱数:::::` + strconv.Itoa(testSlotResult.n48) + `
普通转触发Frees pin玩法次数:::::` + strconv.Itoa(testSlotResult.n49) + `
Frees pin玩法总spin转动次数:::::` + strconv.Itoa(testSlotResult.n50) + `
Frees pin内再次触发Frees pin玩法续命的次数::::` + strconv.Itoa(testSlotResult.n51) + `
Frees pin玩法内随机出0个Wild标签的次数:::::` + strconv.Itoa(testSlotResult.n78) + `
Frees pin玩法内随机出0个Wild标签的赢钱数:::::` + strconv.Itoa(testSlotResult.n79) + `
Frees pin玩法内随机出1个Wild标签的次数:::::` + strconv.Itoa(testSlotResult.n52) + `
Frees pin玩法内随机出1个Wild标签的赢钱数:::::` + strconv.Itoa(testSlotResult.n53) + `
Frees pin玩法内随机出2个Wild标签的次数:::::` + strconv.Itoa(testSlotResult.n54) + `
Frees pin玩法内随机出2个Wild标签的赢钱数:::::` + strconv.Itoa(testSlotResult.n55) + `
Frees pin玩法内随机出3个Wild标签的次数:::::` + strconv.Itoa(testSlotResult.n56) + `
Frees pin玩法内随机出3个Wild标签的赢钱数:::::` + strconv.Itoa(testSlotResult.n57) + `
Frees pin玩法内随机出4个Wild标签的次数:::::` + strconv.Itoa(testSlotResult.n58) + `
Frees pin玩法内随机出4个Wild标签的赢钱数:::::` + strconv.Itoa(testSlotResult.n59) + `
Frees pin玩法内随机出5个Wild标签的次数:::::` + strconv.Itoa(testSlotResult.n60) + `
Frees pin玩法内随机出5个Wild标签的赢钱数:::::` + strconv.Itoa(testSlotResult.n61) + `
Frees pin玩法内随机出6个Wild标签的次数:::::` + strconv.Itoa(testSlotResult.n62) + `
Frees pin玩法内随机出6个Wild标签的赢钱数:::::` + strconv.Itoa(testSlotResult.n63) + `
Frees pin玩法内随机出7个Wild标签的次数:::::` + strconv.Itoa(testSlotResult.n64) + `
Frees pin玩法内随机出7个Wild标签的赢钱数:::::` + strconv.Itoa(testSlotResult.n65) + `
Frees pin玩法内随机出8个Wild标签的次数:::::` + strconv.Itoa(testSlotResult.n66) + `
Frees pin玩法内随机出8个Wild标签的赢钱数:::::` + strconv.Itoa(testSlotResult.n67) + `
Frees pin玩法内随机出9个Wild标签的次数:::::` + strconv.Itoa(testSlotResult.n68) + `
Frees pin玩法内随机出9个Wild标签的赢钱数:::::` + strconv.Itoa(testSlotResult.n69) + `
scatter 标签出现在第一列次数::::: ` + strconv.Itoa(testSlotResult.n70) + `
scatter 标签出现在第二列次数::::: ` + strconv.Itoa(testSlotResult.n71) + `
scatter 标签出现在第三列次数::::: ` + strconv.Itoa(testSlotResult.n72) + `
scatter 标签出现在第四列次数::::: ` + strconv.Itoa(testSlotResult.n73) + `
scatter 标签出现在第五列次数::::: ` + strconv.Itoa(testSlotResult.n74) + `
high_1五个赢钱数::::: ` + decimal.NewFromFloat(testSlotResult.n75).String() + `
high_1四个赢钱数::::: ` + decimal.NewFromFloat(testSlotResult.n76).String() + `
high_1三个赢钱数::::: ` + decimal.NewFromFloat(testSlotResult.n77).String()
}

func AnalyticData(result *unit2TestResult, spin *component.Spin) {
	public.AnalyticData(&result.BasicData, []*component.Spin{spin})
	if !spin.IsFree {
		//普通转划线触发free spin次数以及赢钱
		if spin.FreeSpinParams.Count > 0 {
			result.n49 += 1
		}
		for i, v := range spin.InitDataList {
			for _, vv := range v {
				if vv.Name == "scatter" {
					switch i {
					case 0:
						result.n70 += 1
					case 1:
						result.n71 += 1
					case 2:
						result.n72 += 1
					case 3:
						result.n73 += 1
					case 4:
						result.n74 += 1
					}
				}
			}
		}
		for _, table := range spin.PayTables {
			// 判断是否为免费转标签
			if unit2.GetTagsName(table.Tags) == "high_1" {
				count := len(table.Tags)
				switch count {
				case 5:
					result.n75 += float64(spin.Bet) * table.Multiple
				case 4:
					result.n76 += float64(spin.Bet) * table.Multiple
				case 3:
					result.n77 += float64(spin.Bet) * table.Multiple
				}
			}
		}
	} else {
		result.n50 += 1
		//free spin 划线倍数次数以及赢钱
		//free spin 续命次数
		if spin.FreeSpinParams.Count > 0 {
			result.n51 += 1
		}
		//free spin wild 出现次数以及赢钱
		switch spin.FreeSpinParams.WildNum {
		case 0:
			result.n78 += 1
			result.n79 += spin.Gain
		case 1:
			result.n52 += 1
			result.n53 += spin.Gain
		case 2:
			result.n54 += 1
			result.n55 += spin.Gain
		case 3:
			result.n56 += 1
			result.n57 += spin.Gain
		case 4:
			result.n58 += 1
			result.n59 += spin.Gain
		case 5:
			result.n60 += 1
			result.n61 += spin.Gain
		case 6:
			result.n62 += 1
			result.n63 += spin.Gain
		case 7:
			result.n64 += 1
			result.n65 += spin.Gain
		case 8:
			result.n66 += 1
			result.n67 += spin.Gain
		case 9:
			result.n68 += 1
			result.n69 += spin.Gain
		default:
			break
		}

	}

}
