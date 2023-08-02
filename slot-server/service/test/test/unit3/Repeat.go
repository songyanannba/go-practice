package unit3

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"math"
	"slot-server/model/business"
	"slot-server/service/logic/gameHandle/slotHandle"
	"slot-server/service/slot"
	"slot-server/service/slot/component"
	"slot-server/service/test/public"
	"strconv"
)

type Unit3 struct {
	Result unit3Result
}

type unit3Result struct {
	*public.BasicData
	N27 int     //普通转进度条为0时的赢钱次数
	N28 int     //普通转进度条为0时的赢钱数
	N29 int     //普通转进度条为1时的赢钱次数
	N30 int     //普通转进度条为1时的赢钱数
	N31 int     //普通转进度条为2时的赢钱次数
	N32 int     //普通转进度条为2时的赢钱数
	N33 int     //普通转进度条为3时的赢钱次数
	N34 int     //普通转进度条为3时的赢钱数
	N35 int     //普通转进度条为4时的赢钱次数
	N36 int     //普通转进度条为4时的赢钱数
	N37 int     //普通转进度条为5时的赢钱次数
	N38 int     //普通转进度条为5时的赢钱数
	N39 int     //普通转进度条为6时的赢钱次数
	N40 int     //普通转进度条为6时的赢钱数
	N41 int     //Freespin划线赢钱1-5倍的次数:::::
	N42 int     //Freespin划线赢钱1-5倍的总钱数:::::
	N43 int     //Freespin划线赢钱5-10倍的次数:::::
	N44 int     //Freespin划线赢钱5-10倍的总钱数:::::
	N45 int     //Freespin划线赢钱10-20倍的次数:::::
	N46 int     //Freespin划线赢钱10-20倍的总钱数:::::
	N47 int     //Freespin划线赢钱20-30倍的次数:::::
	N48 int     //Freespin划线赢钱20-30倍的总钱数:::::
	N49 int     //Freespin划线赢钱30-40倍的次数:::::
	N50 int     //Freespin划线赢钱30-40倍的总钱数:::::
	N51 int     //Freespin划线赢钱40-50倍的次数:::::
	N52 int     //Freespin划线赢钱40-50倍的总钱数:::::
	N53 int     //Freespin划线赢钱50-100倍的次数:::::
	N54 int     //Freespin划线赢钱50-100倍的总钱数:::::
	N55 int     //Freespin划线赢钱100-200倍的次数:::::
	N56 int     //Freespin划线赢钱100-200倍的总钱数:::::
	N57 int     //Freespin划线赢钱200-500倍的次数:::::
	N58 int     //Freespin划线赢钱200-500倍的总钱数:::::
	N59 int     //Freespin划线赢钱500-1000倍的次数:::::
	N60 int     //Freespin划线赢钱500-1000倍的总钱数:::::
	N61 int     //Freespin划线赢钱1000以上倍的次数:::::
	N62 int     //Freespin划线赢钱1000以上倍的总钱数:::::
	N63 int     //Freespin进度条为0时的赢钱次数
	N64 int     //Freespin进度条为0时的赢钱数
	N65 int     //Freespin进度条为1时的赢钱次数
	N66 int     //Freespin进度条为1时的赢钱数
	N67 int     //Freespin进度条为2时的赢钱次数
	N68 int     //Freespin进度条为2时的赢钱数
	N69 int     //Freespin进度条为3时的赢钱次数
	N70 int     //Freespin进度条为3时的赢钱数
	N71 int     //Freespin进度条为4时的赢钱次数
	N72 int     //Freespin进度条为4时的赢钱数
	N73 int     //Freespin进度条为5时的赢钱次数
	N74 int     //Freespin进度条为5时的赢钱数
	N75 int     //Freespin进度条为6时的赢钱次数
	N76 int     //Freespin进度条为6时的赢钱数
	N77 int     //普通转触发Freespin玩法次数:::::
	N78 int     //Freespin玩法总spin转动次数:::::
	N79 int     //总次数
	N80 float64 //平方和
	N81 bool    //是否为加注
}

func NewUnit(run public.RunSlotTest) *Unit3 {
	return &Unit3{Result: unit3Result{BasicData: &public.BasicData{}}}
}

func (u *Unit3) Repeat(run public.RunSlotTest, opts ...component.Option) ([]*component.Spin, error) {
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

func (u *Unit3) Calculate(spins []*component.Spin) {
	public.AnalyticData(u.Result.BasicData, spins)

	spin := spins[0]
	mSpins := make([]*component.Spin, 0)
	if len(spins) > 1 {
		mSpins = spins[1:]
	}
	noFreespin := lo.Filter(mSpins, func(item *component.Spin, index int) bool {
		return !item.IsFree
	})
	sumSpin := lo.SumBy(noFreespin, func(i *component.Spin) int {
		return i.Gain
	})
	sumZSpin := lo.SumBy(mSpins, func(i *component.Spin) int {
		return i.Gain
	})
	sumSpin += spin.Gain
	u.Result.N79 += len(mSpins)
	npf := math.Pow(float64(sumZSpin+spin.Gain), 2)

	u.Result.N80 += npf

	if sumSpin > 0 {
		u.Result.N4++
		ss := append(mSpins, spin)
		spins := lo.Filter(ss, func(spin *component.Spin, index int) bool {
			return (!spin.IsFree) && spin.Gain > 0
		})

		for _, v := range spins {
			if v.Gain == 0 {
				continue
			}
			switch v.Rank {
			case 0:
				u.Result.N27++
				u.Result.N28 += v.Gain
			case 1:
				u.Result.N29++
				u.Result.N30 += v.Gain
			case 2:
				u.Result.N31++
				u.Result.N32 += v.Gain
			case 3:
				u.Result.N33++
				u.Result.N34 += v.Gain
			case 4:
				u.Result.N35++
				u.Result.N36 += v.Gain
			case 5:
				u.Result.N37++
				u.Result.N38 += v.Gain
			case 6:
				u.Result.N39++
				u.Result.N40 += v.Gain
			default:
				fmt.Print("rank error")
			}
		}
		freeSpins := lo.Filter(ss, func(spin *component.Spin, index int) bool {
			return spin.IsFree
		})
		for _, v := range freeSpins {
			freeMultiple := float64(v.Gain) / float64(v.Bet)
			if freeMultiple >= 0 && freeMultiple < 5 {
				u.Result.N41++
				u.Result.N42 += v.Gain
			} else if freeMultiple >= 5 && freeMultiple < 10 {
				u.Result.N43++
				u.Result.N44 += v.Gain
			} else if freeMultiple >= 10 && freeMultiple < 20 {
				u.Result.N45++
				u.Result.N46 += v.Gain
			} else if freeMultiple >= 20 && freeMultiple < 30 {
				u.Result.N47++
				u.Result.N48 += v.Gain
			} else if freeMultiple >= 30 && freeMultiple < 40 {
				u.Result.N49++
				u.Result.N50 += v.Gain
			} else if freeMultiple >= 40 && freeMultiple < 50 {
				u.Result.N51++
				u.Result.N52 += v.Gain
			} else if freeMultiple >= 50 && freeMultiple < 100 {
				u.Result.N53++
				u.Result.N54 += v.Gain
			} else if freeMultiple >= 100 && freeMultiple < 200 {
				u.Result.N55++
				u.Result.N56 += v.Gain
			} else if freeMultiple >= 200 && freeMultiple < 500 {
				u.Result.N57++
				u.Result.N58 += v.Gain
			} else if freeMultiple >= 500 && freeMultiple < 1000 {
				u.Result.N59++
				u.Result.N60 += v.Gain
			} else if freeMultiple >= 1000 {
				u.Result.N61++
				u.Result.N62 += v.Gain
			}
			if v.Gain > 0 {
				switch v.Rank {
				case 0:
					u.Result.N63++
					u.Result.N64 += v.Gain
				case 1:
					u.Result.N65++
					u.Result.N66 += v.Gain
				case 2:
					u.Result.N67++
					u.Result.N68 += v.Gain
				case 3:
					u.Result.N69++
					u.Result.N70 += v.Gain
				case 4:
					u.Result.N71++
					u.Result.N72 += v.Gain
				case 5:
					u.Result.N73++
					u.Result.N74 += v.Gain
				case 6:
					u.Result.N75++
					u.Result.N76 += v.Gain
				}
			}

		}
		u.Result.N77 += len(lo.Filter(spins, func(spin *component.Spin, index int) bool {
			return spin.FreeSpinParams.Count > 0
		}))
		u.Result.N78 += len(lo.Filter(ss, func(spin *component.Spin, index int) bool {
			return spin.IsFree
		}))
	}
}

func (u *Unit3) GetDetail() string {
	str := ""
	if u.Result.N81 {
		str += "加注玩法:"
	}
	str +=
		`总次数::::: ` + strconv.Itoa(u.Result.N79) + `
总赢钱:::: ` + strconv.Itoa(u.Result.N1) + `
总押注消耗钱::::: ` + strconv.Itoa(u.Result.N2) + `
总转动次数::::: ` + strconv.Itoa(u.Result.N3) + `
总赢钱次数::::: ` + strconv.Itoa(u.Result.N4) + `
普通转划线赢钱0-5倍的次数::::: ` + strconv.Itoa(u.Result.N5) + `
普通转划线赢钱0-5倍的总钱数::::: ` + strconv.Itoa(u.Result.N6) + `
普通转划线赢钱5-10倍的次数::::: ` + strconv.Itoa(u.Result.N7) + `
普通转划线赢钱5-10倍的总钱数:::::  ` + strconv.Itoa(u.Result.N8) + `
普通转划线赢钱10-20倍的次数::::: ` + strconv.Itoa(u.Result.N9) + `
普通转划线赢钱10-20倍的总钱数::::: ` + strconv.Itoa(u.Result.N10) + `
普通转划线赢钱20-30倍的次数::::: ` + strconv.Itoa(u.Result.N11) + `
普通转划线赢钱20-30倍的总钱数::::: ` + strconv.Itoa(u.Result.N12) + `
普通转划线赢钱30-40倍的次数::::: ` + strconv.Itoa(u.Result.N13) + `
普通转划线赢钱30-40倍的总钱数::::: ` + strconv.Itoa(u.Result.N14) + `
普通转划线赢钱40-50倍的次数::::: ` + strconv.Itoa(u.Result.N15) + `
普通转划线赢钱40-50倍的总钱数::::: ` + strconv.Itoa(u.Result.N16) + `
普通转划线赢钱50-100倍的次数::::: ` + strconv.Itoa(u.Result.N17) + `
普通转划线赢钱50-100倍的总钱数::::: ` + strconv.Itoa(u.Result.N18) + `
普通转划线赢钱100-200倍的次数::::: ` + strconv.Itoa(u.Result.N19) + `
普通转划线赢钱100-200倍的总钱数::::: ` + strconv.Itoa(u.Result.N20) + `
普通转划线赢钱200-500倍的次数::::: ` + strconv.Itoa(u.Result.N21) + `
普通转划线赢钱200-500倍的总钱数::::: ` + strconv.Itoa(u.Result.N22) + `
普通转划线赢钱500-1000倍的次数::::: ` + strconv.Itoa(u.Result.N23) + `
普通转划线赢钱500-1000倍的总钱数::::: ` + strconv.Itoa(u.Result.N24) + `
普通转划线赢钱1000以上倍的次数::::: ` + strconv.Itoa(u.Result.N25) + `
普通转划线赢钱1000以上倍的总钱数::::: ` + strconv.Itoa(u.Result.N26) + `
普通转进度条为0时的赢钱次数::::: ` + strconv.Itoa(u.Result.N27) + `
普通转进度条为0时的赢钱数::::: ` + strconv.Itoa(u.Result.N28) + `
普通转进度条为1时的赢钱次数::::: ` + strconv.Itoa(u.Result.N29) + `
普通转进度条为1时的赢钱数::::: ` + strconv.Itoa(u.Result.N30) + `
普通转进度条为2时的赢钱次数::::: ` + strconv.Itoa(u.Result.N31) + `
普通转进度条为2时的赢钱数::::: ` + strconv.Itoa(u.Result.N32) + `
普通转进度条为3时的赢钱次数::::: ` + strconv.Itoa(u.Result.N33) + `
普通转进度条为3时的赢钱数::::: ` + strconv.Itoa(u.Result.N34) + `
普通转进度条为4时的赢钱次数::::: ` + strconv.Itoa(u.Result.N35) + `
普通转进度条为4时的赢钱数::::: ` + strconv.Itoa(u.Result.N36) + `
普通转进度条为5时的赢钱次数::::: ` + strconv.Itoa(u.Result.N37) + `
普通转进度条为5时的赢钱数::::: ` + strconv.Itoa(u.Result.N38) + `
普通转进度条为6时的赢钱次数::::: ` + strconv.Itoa(u.Result.N39) + `
普通转进度条为6时的赢钱数::::: ` + strconv.Itoa(u.Result.N40) + `
Freespin划线赢钱0-5倍的次数::::: ` + strconv.Itoa(u.Result.N41) + `
Freespin划线赢钱0-5倍的总钱数::::: ` + strconv.Itoa(u.Result.N42) + `
Freespin划线赢钱5-10倍的次数:::::  ` + strconv.Itoa(u.Result.N43) + `
Freespin划线赢钱5-10倍的总钱数::::: ` + strconv.Itoa(u.Result.N44) + `
Freespin划线赢钱10-20倍的次数::::: ` + strconv.Itoa(u.Result.N45) + `
Freespin划线赢钱10-20倍的总钱数::::: ` + strconv.Itoa(u.Result.N46) + `
Freespin划线赢钱20-30倍的次数::::: ` + strconv.Itoa(u.Result.N47) + `
Freespin划线赢钱20-30倍的总钱数::::: ` + strconv.Itoa(u.Result.N48) + `
Freespin划线赢钱30-40倍的次数::::: ` + strconv.Itoa(u.Result.N49) + `
Freespin划线赢钱30-40倍的总钱数::::: ` + strconv.Itoa(u.Result.N50) + `
Freespin划线赢钱40-50倍的次数::::: ` + strconv.Itoa(u.Result.N51) + `
Freespin划线赢钱40-50倍的总钱数::::: ` + strconv.Itoa(u.Result.N52) + `
Freespin划线赢钱50-100倍的次数::::: ` + strconv.Itoa(u.Result.N53) + `
Freespin划线赢钱50-100倍的总钱数::::: ` + strconv.Itoa(u.Result.N54) + `
Freespin划线赢钱100-200倍的次数::::: ` + strconv.Itoa(u.Result.N55) + `
Freespin划线赢钱100-200倍的总钱数::::: ` + strconv.Itoa(u.Result.N56) + ` 
Freespin划线赢钱200-500倍的次数::::: ` + strconv.Itoa(u.Result.N57) + `
Freespin划线赢钱200-500倍的总钱数::::: ` + strconv.Itoa(u.Result.N58) + `
Freespin划线赢钱500-1000倍的次数::::: ` + strconv.Itoa(u.Result.N59) + `
Freespin划线赢钱500-1000倍的总钱数::::: ` + strconv.Itoa(u.Result.N60) + `
Freespin划线赢钱1000以上倍的次数::::: ` + strconv.Itoa(u.Result.N61) + `
Freespin划线赢钱1000以上倍的总钱数::::: ` + strconv.Itoa(u.Result.N62) + `
Freespin进度条为0时的赢钱次数::::: ` + strconv.Itoa(u.Result.N63) + `
Freespin进度条为0时的赢钱数::::: ` + strconv.Itoa(u.Result.N64) + `
Freespin进度条为1时的赢钱次数::::: ` + strconv.Itoa(u.Result.N65) + `
Freespin进度条为1时的赢钱数::::: ` + strconv.Itoa(u.Result.N66) + `
Freespin进度条为2时的赢钱次数::::: ` + strconv.Itoa(u.Result.N67) + `
Freespin进度条为2时的赢钱数::::: ` + strconv.Itoa(u.Result.N68) + `
Freespin进度条为3时的赢钱次数::::: ` + strconv.Itoa(u.Result.N69) + `
Freespin进度条为3时的赢钱数::::: ` + strconv.Itoa(u.Result.N70) + `
Freespin进度条为4时的赢钱次数::::: ` + strconv.Itoa(u.Result.N71) + `
Freespin进度条为4时的赢钱数::::: ` + strconv.Itoa(u.Result.N72) + `
Freespin进度条为5时的赢钱次数::::: ` + strconv.Itoa(u.Result.N73) + `
Freespin进度条为5时的赢钱数::::: ` + strconv.Itoa(u.Result.N74) + `
Freespin进度条为6时的赢钱次数::::: ` + strconv.Itoa(u.Result.N75) + `
Freespin进度条为6时的赢钱数::::: ` + strconv.Itoa(u.Result.N76) + `
普通转触发Freespin玩法次数::::: ` + strconv.Itoa(u.Result.N77) + `
Freespin玩法总spin转动次数:::::` + strconv.Itoa(u.Result.N78) + `
平方和:::::` + fmt.Sprintf("%v", u.Result.N80) + "\r\n"
	str += fmt.Sprintf("最大倍率:::::%g\n", u.Result.N0)
	return str
}

func (u *Unit3) GetBasicData() *public.BasicData {
	return u.Result.BasicData
}

func (u *Unit3) GetResult(spins []*component.Spin) (int64, int64) {
	gain := lo.SumBy(spins, func(spin *component.Spin) int64 {
		return int64(spin.Gain)
	})
	bet := int64(spins[0].Bet) + int64(spins[0].Raise) + int64(spins[0].BuyFreeCoin) + int64(spins[0].BuyReCoin)
	return gain, bet
}

func (u *Unit3) GetReturnRatio() float64 {
	f, b := decimal.NewFromInt(int64(u.Result.N1)).Div(decimal.NewFromInt(int64(u.Result.N2))).Float64()
	if b {
		return f
	}
	return 0
}
