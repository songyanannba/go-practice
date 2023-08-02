package unit2

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"slot-server/model/business"
	"slot-server/service/logic/gameHandle/slotHandle"
	"slot-server/service/slot"
	"slot-server/service/slot/component"
	"slot-server/service/slot/machine/unit2"
	"slot-server/service/test/public"
	"strconv"
)

type Unit2 struct {
	Result unit2Result
}

type unit2Result struct {
	*public.BasicData
	N27 int     //Frees pin划线赢钱0-5倍的次数:::::
	N28 int     //Frees pin划线赢钱0-5倍的总钱数:::::
	N29 int     //Frees pin划线赢钱5-10倍的次数:::::
	N30 int     //Frees pin划线赢钱5-10倍的总钱数:::::
	N31 int     //Frees pin划线赢钱10-20倍的次数:::::
	N32 int     //Frees pin划线赢钱10-20倍的总钱数:::::
	N33 int     //Frees pin划线赢钱20-30倍的次数:::::
	N34 int     //Frees pin划线赢钱20-30倍的总钱数:::::
	N35 int     //Frees pin划线赢钱30-40倍的次数:::::
	N36 int     //Frees pin划线赢钱30-40倍的总钱数:::::
	N37 int     //Frees pin划线赢钱40-50倍的次数:::::
	N38 int     //Frees pin划线赢钱40-50倍的总钱数:::::
	N39 int     //Frees pin划线赢钱50-100倍的次数:::::
	N40 int     //Frees pin划线赢钱50-100倍的总钱数:::::
	N41 int     //Frees pin划线赢钱100-200倍的次数:::::
	N42 int     //Frees pin划线赢钱100-200倍的总钱数:::::
	N43 int     //Frees pin划线赢钱200-500倍的次数:::::
	N44 int     //Frees pin划线赢钱200-500倍的总钱数:::::
	N45 int     //Frees pin划线赢钱500-1000倍的次数:::::
	N46 int     //Frees pin划线赢钱500-1000倍的总钱数:::::
	N47 int     //Frees pin划线赢钱1000以上倍的次数:::::
	N48 int     //Frees pin划线赢钱1000以上倍的总钱数:::::
	N49 int     //普通转触发Frees pin玩法次数:::::
	N50 int     //Frees pin玩法总spin转动次数:::::
	N51 int     //Frees pin内再次触发Frees pin玩法续命的次数::::
	N52 int     //Frees pin玩法内随机出1个Wild标签的次数:::::
	N53 int     //Frees pin玩法内随机出1个Wild标签的赢钱数:::::
	N54 int     //Frees pin玩法内随机出2个Wild标签的次数:::::
	N55 int     //Frees pin玩法内随机出2个Wild标签的赢钱数:::::
	N56 int     //Frees pin玩法内随机出3个Wild标签的次数:::::
	N57 int     //Frees pin玩法内随机出3个Wild标签的赢钱数:::::
	N58 int     //Frees pin玩法内随机出4个Wild标签的次数:::::
	N59 int     //Frees pin玩法内随机出4个Wild标签的赢钱数:::::
	N60 int     //Frees pin玩法内随机出5个Wild标签的次数:::::
	N61 int     //Frees pin玩法内随机出5个Wild标签的赢钱数:::::
	N62 int     //Frees pin玩法内随机出6个Wild标签的次数:::::
	N63 int     //Frees pin玩法内随机出6个Wild标签的赢钱数:::::
	N64 int     //Frees pin玩法内随机出7个Wild标签的次数:::::
	N65 int     //Frees pin玩法内随机出7个Wild标签的赢钱数:::::
	N66 int     //Frees pin玩法内随机出8个Wild标签的次数:::::
	N67 int     //Frees pin玩法内随机出8个Wild标签的赢钱数:::::
	N68 int     //Frees pin玩法内随机出9个Wild标签的次数:::::
	N69 int     //Frees pin玩法内随机出9个Wild标签的赢钱数:::::
	N70 int     //scatter 标签出现在第一列次数:::::
	N71 int     //scatter 标签出现在第二列次数:::::
	N72 int     //scatter 标签出现在第三列次数:::::
	N73 int     //scatter 标签出现在第四列次数:::::
	N74 int     //scatter 标签出现在第五列次数:::::
	N75 float64 //scatter high_1五个次数:::::
	N76 float64 //scatter high_1四个次数:::::
	N77 float64 //scatter high_1三个次数:::::
	N78 int     //Frees pin玩法内随机出0个Wild标签的次数:::::
	N79 int     //Frees pin玩法内随机出0个Wild标签的赢钱数:::::
}

func NewUnit(run public.RunSlotTest) *Unit2 {
	return &Unit2{Result: unit2Result{BasicData: &public.BasicData{}}}
}

func (u *Unit2) Repeat(run public.RunSlotTest, opts ...component.Option) ([]*component.Spin, error) {
	m, _ := slot.Play(run.SlotId, run.Amount, opts...)
	spin := m.GetSpin()
	mergeSpin, err := slotHandle.RunMergeSpin(0, spin, &business.SlotUserSpin{
		UserId: 0,
		SlotId: run.SlotId,
	})
	if err != nil {
		return nil, err
	}
	spins := []*component.Spin{spin}
	spins = append(spins, mergeSpin.Spins...)
	return spins, nil
}

func (u *Unit2) Calculate(spins []*component.Spin) {
	public.AnalyticData(u.Result.BasicData, spins)
	for _, spin := range spins {
		u.AnalyticData(spin)
	}
	freeSpins := lo.Filter(spins, func(item *component.Spin, index int) bool {
		return item.IsFree
	})
	for i := 0; i < len(freeSpins); i++ {
		freeSpinList := make([]*component.Spin, 0)
		for a := 0; a < 8 && a < len(freeSpins); a++ {
			freeSpinList = append(freeSpinList, freeSpins[i+a])
		}
		i += 7
		u.FreeSpinSum(freeSpinList, spins[0].Bet)
	}
}

func (u *Unit2) FreeSpinSum(spins []*component.Spin, bet int) {
	sumGain := lo.SumBy(spins, func(i *component.Spin) int {
		return i.Gain
	})
	if sumGain == 0 {
		return
	}
	multiple := float64(sumGain) / float64(bet)
	if multiple >= 0 && multiple < 5 {
		u.Result.N27 += 1
		u.Result.N28 += sumGain
	} else if multiple >= 5 && multiple < 10 {
		u.Result.N29 += 1
		u.Result.N30 += sumGain
	} else if multiple >= 10 && multiple < 20 {
		u.Result.N31 += 1
		u.Result.N32 += sumGain
	} else if multiple >= 20 && multiple < 30 {
		u.Result.N33 += 1
		u.Result.N34 += sumGain
	} else if multiple >= 30 && multiple < 40 {
		u.Result.N35 += 1
		u.Result.N36 += sumGain
	} else if multiple >= 40 && multiple < 50 {
		u.Result.N37 += 1
		u.Result.N38 += sumGain
	} else if multiple >= 50 && multiple < 100 {
		u.Result.N39 += 1
		u.Result.N40 += sumGain
	} else if multiple >= 100 && multiple < 200 {
		u.Result.N41 += 1
		u.Result.N42 += sumGain
	} else if multiple >= 200 && multiple < 500 {
		u.Result.N43 += 1
		u.Result.N44 += sumGain
	} else if multiple >= 500 && multiple < 1000 {
		u.Result.N45 += 1
		u.Result.N46 += sumGain
	} else if multiple >= 1000 {
		u.Result.N47 += 1
		u.Result.N48 += sumGain
	}
}

func (u *Unit2) AnalyticData(spin *component.Spin) {

	if !spin.IsFree {
		//普通转划线触发free spin次数以及赢钱
		if spin.FreeSpinParams.Count > 0 {
			u.Result.N49 += 1
		}
		for i, v := range spin.InitDataList {
			for _, vv := range v {
				if vv.Name == "scatter" {
					switch i {
					case 0:
						u.Result.N70 += 1
					case 1:
						u.Result.N71 += 1
					case 2:
						u.Result.N72 += 1
					case 3:
						u.Result.N73 += 1
					case 4:
						u.Result.N74 += 1
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
					u.Result.N75 += float64(spin.Bet) * table.Multiple
				case 4:
					u.Result.N76 += float64(spin.Bet) * table.Multiple
				case 3:
					u.Result.N77 += float64(spin.Bet) * table.Multiple
				}
			}
		}
	} else {
		u.Result.N50 += 1
		//free spin 划线倍数次数以及赢钱
		//free spin 续命次数
		if spin.FreeSpinParams.Count > 0 {
			u.Result.N51 += 1
		}
		//free spin wild 出现次数以及赢钱
		switch spin.FreeSpinParams.WildNum {
		case 0:
			u.Result.N78 += 1
			u.Result.N79 += spin.Gain
		case 1:
			u.Result.N52 += 1
			u.Result.N53 += spin.Gain
		case 2:
			u.Result.N54 += 1
			u.Result.N55 += spin.Gain
		case 3:
			u.Result.N56 += 1
			u.Result.N57 += spin.Gain
		case 4:
			u.Result.N58 += 1
			u.Result.N59 += spin.Gain
		case 5:
			u.Result.N60 += 1
			u.Result.N61 += spin.Gain
		case 6:
			u.Result.N62 += 1
			u.Result.N63 += spin.Gain
		case 7:
			u.Result.N64 += 1
			u.Result.N65 += spin.Gain
		case 8:
			u.Result.N66 += 1
			u.Result.N67 += spin.Gain
		case 9:
			u.Result.N68 += 1
			u.Result.N69 += spin.Gain
		default:
			break
		}

	}

}

func (u *Unit2) GetDetail() string {
	str := public.GetInitResult(u.Result.BasicData)
	str += `Freespin划线赢钱0-5倍的次数:::::` + strconv.Itoa(u.Result.N27) + `
Freespin划线赢钱0-5倍的总钱数:::::` + strconv.Itoa(u.Result.N28) + `
Freespin划线赢钱5-10倍的次数:::::` + strconv.Itoa(u.Result.N29) + `
Freespin划线赢钱5-10倍的总钱数:::::` + strconv.Itoa(u.Result.N30) + `
Freespin划线赢钱10-20倍的次数:::::` + strconv.Itoa(u.Result.N31) + `
Freespin划线赢钱10-20倍的总钱数:::::` + strconv.Itoa(u.Result.N32) + `
Freespin划线赢钱20-30倍的次数:::::` + strconv.Itoa(u.Result.N33) + `
Freespin划线赢钱20-30倍的总钱数:::::` + strconv.Itoa(u.Result.N34) + `
Freespin划线赢钱30-40倍的次数:::::` + strconv.Itoa(u.Result.N35) + `
Freespin划线赢钱30-40倍的总钱数:::::` + strconv.Itoa(u.Result.N36) + `
Freespin划线赢钱40-50倍的次数:::::` + strconv.Itoa(u.Result.N37) + `
Freespin划线赢钱40-50倍的总钱数:::::` + strconv.Itoa(u.Result.N38) + `
Freespin划线赢钱50-100倍的次数:::::` + strconv.Itoa(u.Result.N39) + `
Freespin划线赢钱50-100倍的总钱数:::::` + strconv.Itoa(u.Result.N40) + `
Freespin划线赢钱100-200倍的次数:::::` + strconv.Itoa(u.Result.N41) + `
Freespin划线赢钱100-200倍的总钱数:::::` + strconv.Itoa(u.Result.N42) + `
Freespin划线赢钱200-500倍的次数:::::` + strconv.Itoa(u.Result.N43) + `
Freespin划线赢钱200-500倍的总钱数:::::` + strconv.Itoa(u.Result.N44) + `
Freespin划线赢钱500-1000倍的次数:::::` + strconv.Itoa(u.Result.N45) + `
Freespin划线赢钱500-1000倍的总钱数:::::` + strconv.Itoa(u.Result.N46) + `
Freespin划线赢钱1000以上倍的次数:::::` + strconv.Itoa(u.Result.N47) + `
Freespin划线赢钱1000以上倍的总钱数:::::` + strconv.Itoa(u.Result.N48) + `
普通转触发Freespin玩法次数:::::` + strconv.Itoa(u.Result.N49) + `
Freespin玩法总spin转动次数:::::` + strconv.Itoa(u.Result.N50) + `
Freespin内再次触发Freespin玩法续命的次数::::` + strconv.Itoa(u.Result.N51) + `
Freespin玩法内随机出0个Wild标签的次数:::::` + strconv.Itoa(u.Result.N78) + `
Freespin玩法内随机出0个Wild标签的赢钱数:::::` + strconv.Itoa(u.Result.N79) + `
Freespin玩法内随机出1个Wild标签的次数:::::` + strconv.Itoa(u.Result.N52) + `
Freespin玩法内随机出1个Wild标签的赢钱数:::::` + strconv.Itoa(u.Result.N53) + `
Freespin玩法内随机出2个Wild标签的次数:::::` + strconv.Itoa(u.Result.N54) + `
Freespin玩法内随机出2个Wild标签的赢钱数:::::` + strconv.Itoa(u.Result.N55) + `
Freespin玩法内随机出3个Wild标签的次数:::::` + strconv.Itoa(u.Result.N56) + `
Freespin玩法内随机出3个Wild标签的赢钱数:::::` + strconv.Itoa(u.Result.N57) + `
Freespin玩法内随机出4个Wild标签的次数:::::` + strconv.Itoa(u.Result.N58) + `
Freespin玩法内随机出4个Wild标签的赢钱数:::::` + strconv.Itoa(u.Result.N59) + `
Freespin玩法内随机出5个Wild标签的次数:::::` + strconv.Itoa(u.Result.N60) + `
Freespin玩法内随机出5个Wild标签的赢钱数:::::` + strconv.Itoa(u.Result.N61) + `
Freespin玩法内随机出6个Wild标签的次数:::::` + strconv.Itoa(u.Result.N62) + `
Freespin玩法内随机出6个Wild标签的赢钱数:::::` + strconv.Itoa(u.Result.N63) + `
Freespin玩法内随机出7个Wild标签的次数:::::` + strconv.Itoa(u.Result.N64) + `
Freespin玩法内随机出7个Wild标签的赢钱数:::::` + strconv.Itoa(u.Result.N65) + `
Freespin玩法内随机出8个Wild标签的次数:::::` + strconv.Itoa(u.Result.N66) + `
Freespin玩法内随机出8个Wild标签的赢钱数:::::` + strconv.Itoa(u.Result.N67) + `
Freespin玩法内随机出9个Wild标签的次数:::::` + strconv.Itoa(u.Result.N68) + `
Freespin玩法内随机出9个Wild标签的赢钱数:::::` + strconv.Itoa(u.Result.N69) + `
scatter标签出现在第一列次数:::::` + strconv.Itoa(u.Result.N70) + `
scatter标签出现在第二列次数:::::` + strconv.Itoa(u.Result.N71) + `
scatter标签出现在第三列次数:::::` + strconv.Itoa(u.Result.N72) + `
scatter标签出现在第四列次数:::::` + strconv.Itoa(u.Result.N73) + `
scatter标签出现在第五列次数:::::` + strconv.Itoa(u.Result.N74) + `
high_1五个赢钱数:::::` + decimal.NewFromFloat(u.Result.N75).String() + `
high_1四个赢钱数:::::` + decimal.NewFromFloat(u.Result.N76).String() + `
high_1三个赢钱数:::::` + decimal.NewFromFloat(u.Result.N77).String() + "\r\n"
	str += fmt.Sprintf("最大倍率:::::%g\n", u.Result.N0)
	return str
}

func (u *Unit2) GetBasicData() *public.BasicData {
	return u.Result.BasicData
}

func (u *Unit2) GetResult(spins []*component.Spin) (int64, int64) {
	gain := lo.SumBy(spins, func(spin *component.Spin) int64 {
		return int64(spin.Gain)
	})
	bet := int64(spins[0].Bet) + int64(spins[0].Raise) + int64(spins[0].BuyFreeCoin) + int64(spins[0].BuyReCoin)
	return gain, bet
}

func (u *Unit2) GetReturnRatio() float64 {
	f, b := decimal.NewFromInt(int64(u.Result.N1)).Div(decimal.NewFromInt(int64(u.Result.N2))).Float64()
	if b {
		return f
	}
	return 0
}
