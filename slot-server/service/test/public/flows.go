package public

import (
	"github.com/samber/lo"
	"slot-server/service/slot/component"
	"strconv"
)

// AnalyticData 基础数据分析
func AnalyticData(result *BasicData, spins []*component.Spin) {

	sumGain := lo.SumBy(spins, func(spin *component.Spin) int {
		return spin.Gain
	})

	result.N1 += sumGain
	if sumGain > 0 {
		result.N4++
	}
	spin := spins[0]
	result.N2 += spin.Bet + int(spin.Options.Raise) + int(spin.Options.BuyFreeCoin) + int(spin.Options.BuyReCoin)
	result.N3 += len(spins)
	multiple := float64(sumGain) / float64(spin.Bet)
	if multiple >= 0 && multiple < 5 {
		result.N5 += 1
		result.N6 += sumGain
	} else if multiple >= 5 && multiple < 10 {
		result.N7 += 1
		result.N8 += sumGain
	} else if multiple >= 10 && multiple < 20 {
		result.N9 += 1
		result.N10 += sumGain
	} else if multiple >= 20 && multiple < 30 {
		result.N11 += 1
		result.N12 += sumGain
	} else if multiple >= 30 && multiple < 40 {
		result.N13 += 1
		result.N14 += sumGain
	} else if multiple >= 40 && multiple < 50 {
		result.N15 += 1
		result.N16 += sumGain
	} else if multiple >= 50 && multiple < 100 {
		result.N17 += 1
		result.N18 += sumGain
	} else if multiple >= 100 && multiple < 200 {
		result.N19 += 1
		result.N20 += sumGain
	} else if multiple >= 200 && multiple < 500 {
		result.N21 += 1
		result.N22 += sumGain
	} else if multiple >= 500 && multiple < 1000 {
		result.N23 += 1
		result.N24 += sumGain
	} else if multiple >= 1000 {
		result.N25 += 1
		result.N26 += sumGain
	}

	if multiple > result.N0 {
		result.N0 = multiple
	}
}

// GetInitResult 获取初始化结果
func GetInitResult(result *BasicData) (str string) {
	str = `总赢钱::::` + strconv.Itoa(result.N1) + `
总押注消耗钱:::::` + strconv.Itoa(result.N2) + `
总转动次数:::::` + strconv.Itoa(result.N3) + `
总赢钱次数:::::` + strconv.Itoa(result.N4) + `
普通转划线赢钱0-5倍的次数:::::` + strconv.Itoa(result.N5) + `
普通转划线赢钱0-5倍的总钱数:::::` + strconv.Itoa(result.N6) + `
普通转划线赢钱5-10倍的次数:::::` + strconv.Itoa(result.N7) + `
普通转划线赢钱5-10倍的总钱数:::::` + strconv.Itoa(result.N8) + `
普通转划线赢钱10-20倍的次数:::::` + strconv.Itoa(result.N9) + `
普通转划线赢钱10-20倍的总钱数:::::` + strconv.Itoa(result.N10) + `
普通转划线赢钱20-30倍的次数:::::` + strconv.Itoa(result.N11) + `
普通转划线赢钱20-30倍的总钱数:::::` + strconv.Itoa(result.N12) + `
普通转划线赢钱30-40倍的次数:::::` + strconv.Itoa(result.N13) + `
普通转划线赢钱30-40倍的总钱数:::::` + strconv.Itoa(result.N14) + `
普通转划线赢钱40-50倍的次数:::::` + strconv.Itoa(result.N15) + `
普通转划线赢钱40-50倍的总钱数:::::` + strconv.Itoa(result.N16) + `
普通转划线赢钱50-100倍的次数:::::` + strconv.Itoa(result.N17) + `
普通转划线赢钱50-100倍的总钱数:::::` + strconv.Itoa(result.N18) + `
普通转划线赢钱100-200倍的次数:::::` + strconv.Itoa(result.N19) + `
普通转划线赢钱100-200倍的总钱数:::::` + strconv.Itoa(result.N20) + `
普通转划线赢钱200-500倍的次数:::::` + strconv.Itoa(result.N21) + `
普通转划线赢钱200-500倍的总钱数:::::` + strconv.Itoa(result.N22) + `
普通转划线赢钱500-1000倍的次数:::::` + strconv.Itoa(result.N23) + `
普通转划线赢钱500-1000倍的总钱数:::::` + strconv.Itoa(result.N24) + `
普通转划线赢钱1000以上倍的次数:::::` + strconv.Itoa(result.N25) + `
普通转划线赢钱1000以上倍的总钱数:::::` + strconv.Itoa(result.N26) + "\n"
	return str
}
