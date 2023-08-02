package unit8

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/service/slot"
	"slot-server/service/slot/base"
	"slot-server/service/slot/component"
	"slot-server/service/slot/template/flow"
	"slot-server/service/test/public"
	"slot-server/utils"
	"slot-server/utils/helper"
	"strconv"
	"strings"
)

type Unit8 struct {
	Result Unit8Result
}

type Unit8Result struct {
	*public.BasicData
	N27         int //普通转连续消除1次的次数:::::
	N28         int //普通转连续消除1次的赢钱数:::::
	N29         int //普通转连续消除2次的次数:::::
	N30         int //普通转连续消除2次的赢钱数:::::
	N31         int //普通转连续消除3次的次数:::::
	N32         int //普通转连续消除3次的赢钱数:::::
	N33         int //普通转连续消除4次的次数:::::
	N34         int //普通转连续消除4次的赢钱数:::::
	N35         int //普通转连续消除5次的次数:::::
	N36         int //普通转连续消除5次的赢钱数:::::
	N37         int //普通转连续消除6次的次数:::::
	N38         int //普通转连续消除6次的赢钱数:::::
	N39         int //普通转连续消除7次的次数:::::
	N40         int //普通转连续消除7次的赢钱数:::::
	N41         int //普通转连续消除8次的次数:::::
	N42         int //普通转连续消除8次的赢钱数:::::
	N43         int //普通转连续消除9次的次数:::::
	N44         int //普通转连续消除9次的赢钱数:::::
	N45         int //普通转连续消除10次的次数:::::
	N46         int //普通转连续消除10次的赢钱数:::::
	N47         int //普通转连续消除10+次的次数:::::
	N48         int //普通转连续消除10+次的赢钱数:::::
	N49         int //FREESPIN初始触发的次数:::::
	N50         int //FREESPIN再次触发FREESPIN的次数:::::
	N51         int //Freespin划线赢钱0-5倍的次数:::::
	N52         int //Freespin划线赢钱0-5倍的总钱数:::::
	N53         int //Freespin划线赢钱5-10倍的次数:::::
	N54         int //Freespin划线赢钱5-10倍的总钱数:::::
	N55         int //Freespin划线赢钱10-20倍的次数:::::
	N56         int //Freespin划线赢钱10-20倍的总钱数:::::
	N57         int //Freespin划线赢钱20-30倍的次数:::::
	N58         int //Freespin划线赢钱20-30倍的总钱数:::::
	N59         int //Freespin划线赢钱30-40倍的次数:::::
	N60         int //Freespin划线赢钱30-40倍的总钱数:::::
	N61         int //Freespin划线赢钱40-50倍的次数:::::
	N62         int //Freespin划线赢钱40-50倍的总钱数:::::
	N63         int //Freespin划线赢钱50-100倍的次数:::::
	N64         int //Freespin划线赢钱50-100倍的总钱数:::::
	N65         int //Freespin划线赢钱100-200倍的次数:::::
	N66         int //Freespin划线赢钱100-200倍的总钱数:::::
	N67         int //Freespin划线赢钱200-500倍的次数:::::
	N68         int //Freespin划线赢钱200-500倍的总钱数:::::
	N69         int //Freespin划线赢钱500-1000倍的次数:::::
	N70         int //Freespin划线赢钱500-1000倍的总钱数:::::
	N71         int //Freespin划线赢钱1000以上倍的次数:::::
	N72         int //Freespin划线赢钱1000以上倍的总钱数:::::
	N73         int //Freespin连续消除1次的次数:::::
	N74         int //Freespin连续消除1次的赢钱数:::::
	N75         int //Freespin连续消除2次的次数:::::
	N76         int //Freespin连续消除2次的赢钱数:::::
	N77         int //Freespin连续消除3次的次数:::::
	N78         int //Freespin连续消除3次的赢钱数:::::
	N79         int //Freespin连续消除4次的次数:::::
	N80         int //Freespin连续消除4次的赢钱数:::::
	N81         int //Freespin连续消除5次的次数:::::
	N82         int //Freespin连续消除5次的赢钱数:::::
	N83         int //Freespin连续消除6次的次数:::::
	N84         int //Freespin连续消除6次的赢钱数:::::
	N85         int //Freespin连续消除7次的次数:::::
	N86         int //Freespin连续消除7次的赢钱数:::::
	N87         int //Freespin连续消除8次的次数:::::
	N88         int //Freespin连续消除8次的赢钱数:::::
	N89         int //Freespin连续消除9次的次数:::::
	N90         int //Freespin连续消除9次的赢钱数:::::
	N91         int //Freespin连续消除10次的次数:::::
	N92         int //Freespin连续消除10次的赢钱数:::::
	N93         int //Freespin连续消除10+次的次数:::::
	N94         int //Freespin连续消除10+次的赢钱数:::::
	N95         int //4个scatter触发Freespin的次数:::::
	N96         int //4个scatter触发Freespin的总赢钱数:::::
	MapStart    map[int]int
	MapStartWin map[int]int
}

func NewUnit() *Unit8 {
	return &Unit8{
		Result: Unit8Result{
			BasicData:   &public.BasicData{},
			MapStart:    make(map[int]int),
			MapStartWin: make(map[int]int),
		},
	}
}

func (u *Unit8) Repeat(run public.RunSlotTest, opts ...component.Option) ([]*component.Spin, error) {
	m, err := slot.Play(run.SlotId, run.Amount, opts...)
	if err != nil {
		return nil, err
	}

	spins := []*component.Spin{m.GetSpin()}
	spins = append(spins, m.GetSpins()...)
	return spins, nil
}

func (u *Unit8) Calculate(spins []*component.Spin) {
	public.AnalyticData(u.Result.BasicData, spins)

	spin := spins[0]
	u.Result.MapStart[spin.SpinInfo.GetStartValue()]++

	mSpins := make([]*component.Spin, 0)
	if len(spins) > 1 {
		mSpins = spins[1:]
	}

	sumGain := lo.SumBy(mSpins, func(s *component.Spin) int {
		return s.Gain
	}) + spin.Gain
	u.Result.MapStartWin[spin.SpinInfo.GetStartValue()] += sumGain
	if len(spin.SpinInfo.Scatter.Tags) == 4 {
		u.Result.N95++
		u.Result.N96 += sumGain - spin.Gain
	}

	//普通转消除次数和赢钱
	{
		rmCount := len(lo.Filter(spin.SpinInfo.SpinFlow, func(s flow.SpinFlow, i int) bool {
			if len(s.OmitList) > 0 {
				return true
			}
			return false
		}))
		switch rmCount {
		case 0:

		case 1:
			u.Result.N27++
			u.Result.N28 += spin.Gain
		case 2:
			u.Result.N29++
			u.Result.N30 += spin.Gain
		case 3:
			u.Result.N31++
			u.Result.N32 += spin.Gain
		case 4:
			u.Result.N33++
			u.Result.N34 += spin.Gain
		case 5:
			u.Result.N35++
			u.Result.N36 += spin.Gain
		case 6:
			u.Result.N37++
			u.Result.N38 += spin.Gain
		case 7:
			u.Result.N39++
			u.Result.N40 += spin.Gain
		case 8:
			u.Result.N41++
			u.Result.N42 += spin.Gain
		case 9:
			u.Result.N43++
			u.Result.N44 += spin.Gain
		case 10:
			u.Result.N45++
			u.Result.N46 += spin.Gain
		default:
			u.Result.N47++
			u.Result.N48 += spin.Gain

		}
	}

	//Free赢钱倍数
	{
		if len(mSpins) == 0 {
			return
		}
		if spin.FreeSpinParams.FreeNum > 0 {
			u.Result.N49++
		}
		frees := lo.Filter(mSpins, func(s *component.Spin, index int) bool {
			return s.ParentId == 0 && s.IsFree
		})

		SumFreeGain := lo.SumBy(frees, func(s *component.Spin) int {
			return s.Gain
		})

		multipleFree := float64(SumFreeGain) / float64(spin.Bet)
		if multipleFree >= 0 && multipleFree < 5 {
			u.Result.N51++
			u.Result.N52 += SumFreeGain
		} else if multipleFree >= 5 && multipleFree < 10 {
			u.Result.N53++
			u.Result.N54 += SumFreeGain
		} else if multipleFree >= 10 && multipleFree < 20 {
			u.Result.N55++
			u.Result.N56 += SumFreeGain
		} else if multipleFree >= 20 && multipleFree < 30 {
			u.Result.N57++
			u.Result.N58 += SumFreeGain
		} else if multipleFree >= 30 && multipleFree < 40 {
			u.Result.N59++
			u.Result.N60 += SumFreeGain
		} else if multipleFree >= 40 && multipleFree < 50 {
			u.Result.N61++
			u.Result.N62 += SumFreeGain
		} else if multipleFree >= 50 && multipleFree < 100 {
			u.Result.N63++
			u.Result.N64 += SumFreeGain
		} else if multipleFree >= 100 && multipleFree < 200 {
			u.Result.N65++
			u.Result.N66 += SumFreeGain
		} else if multipleFree >= 200 && multipleFree < 500 {
			u.Result.N67++
			u.Result.N68 += SumFreeGain
		} else if multipleFree >= 500 && multipleFree < 1000 {
			u.Result.N69++
			u.Result.N70 += SumFreeGain
		} else if multipleFree >= 1000 {
			u.Result.N71++
			u.Result.N72 += SumFreeGain
		}

		freeParents := lo.Filter(mSpins, func(s *component.Spin, index int) bool {
			return s.FreeSpinParams.FreeNum > 0
		})
		sumFreeSumGain := 0
		for _, parent := range freeParents {
			u.Result.N50++
			sons := lo.Filter(mSpins, func(s *component.Spin, index int) bool {
				return s.ParentId == parent.Id
			})

			freeSumGain := lo.SumBy(sons, func(s *component.Spin) int {
				return s.Gain
			})
			sumFreeSumGain += freeSumGain
			freeMultiple := float64(freeSumGain) / float64(spin.Bet)
			if freeMultiple >= 0 && freeMultiple < 5 {
				u.Result.N51++
				u.Result.N52 += freeSumGain
			} else if freeMultiple >= 5 && freeMultiple < 10 {
				u.Result.N53++
				u.Result.N54 += freeSumGain
			} else if freeMultiple >= 10 && freeMultiple < 20 {
				u.Result.N55++
				u.Result.N56 += freeSumGain
			} else if freeMultiple >= 20 && freeMultiple < 30 {
				u.Result.N57++
				u.Result.N58 += freeSumGain
			} else if freeMultiple >= 30 && freeMultiple < 40 {
				u.Result.N59++
				u.Result.N60 += freeSumGain
			} else if freeMultiple >= 40 && freeMultiple < 50 {
				u.Result.N61++
				u.Result.N62 += freeSumGain
			} else if freeMultiple >= 50 && freeMultiple < 100 {
				u.Result.N63++
				u.Result.N64 += freeSumGain
			} else if freeMultiple >= 100 && freeMultiple < 200 {
				u.Result.N65++
				u.Result.N66 += freeSumGain
			} else if freeMultiple >= 200 && freeMultiple < 500 {
				u.Result.N67++
				u.Result.N68 += freeSumGain
			} else if freeMultiple >= 500 && freeMultiple < 1000 {
				u.Result.N69++
				u.Result.N70 += freeSumGain
			} else if freeMultiple >= 1000 {
				u.Result.N71++
				u.Result.N72 += freeSumGain
			}
		}

		if sumGain-spin.Gain != sumFreeSumGain+SumFreeGain {
			global.GVA_LOG.Info("sumGain - spin.Gain!= sumFreeSumGain+ SumFreeGain", zap.Any("sumGain", sumGain), zap.Any("spin.Gain", spin.Gain), zap.Any("sumFreeSumGain", sumFreeSumGain), zap.Any("SumFreeGain", SumFreeGain))
		}

		for _, c := range mSpins {
			rmCount := len(lo.Filter(c.SpinInfo.SpinFlow, func(s flow.SpinFlow, i int) bool {
				if len(s.OmitList) > 0 {
					return true
				}
				return false
			}))
			switch rmCount {
			case 1:
				u.Result.N73++
				u.Result.N74 += c.Gain
			case 2:
				u.Result.N75++
				u.Result.N76 += c.Gain
			case 3:
				u.Result.N77++
				u.Result.N78 += c.Gain
			case 4:
				u.Result.N79++
				u.Result.N80 += c.Gain
			case 5:
				u.Result.N81++
				u.Result.N82 += c.Gain
			case 6:
				u.Result.N83++
				u.Result.N84 += c.Gain
			case 7:
				u.Result.N85++
				u.Result.N86 += c.Gain
			case 8:
				u.Result.N87++
				u.Result.N88 += c.Gain
			case 9:
				u.Result.N89++
				u.Result.N90 += c.Gain
			case 10:
				u.Result.N91++
				u.Result.N92 += c.Gain
			default:
				u.Result.N93++
				u.Result.N94 += c.Gain

			}
		}

	}
}

func (u *Unit8) GetDetail() string {
	str := `最大倍率:::: ` + fmt.Sprintf("%g", u.Result.N0) + `
总赢钱:::: ` + strconv.Itoa(u.Result.N1) + `
总押注消耗钱:::::` + strconv.Itoa(u.Result.N2) + `
总转动次数:::::` + strconv.Itoa(u.Result.N3) + `
总赢钱次数:::: ` + strconv.Itoa(u.Result.N4) + `
普通转划线赢钱0-5倍的次数:::: ` + strconv.Itoa(u.Result.N5) + `
普通转划线赢钱0-5倍的总钱数:::: ` + strconv.Itoa(u.Result.N6) + `
普通转划线赢钱5-10倍的次数:::: ` + strconv.Itoa(u.Result.N7) + `
普通转划线赢钱5-10倍的总钱数:::: ` + strconv.Itoa(u.Result.N8) + `
普通转划线赢钱10-20倍的次数:::: ` + strconv.Itoa(u.Result.N9) + `
普通转划线赢钱10-20倍的总钱数:::: ` + strconv.Itoa(u.Result.N10) + `
普通转划线赢钱20-30倍的次数:::: ` + strconv.Itoa(u.Result.N11) + `
普通转划线赢钱20-30倍的总钱数:::: ` + strconv.Itoa(u.Result.N12) + `
普通转划线赢钱30-40倍的次数:::: ` + strconv.Itoa(u.Result.N13) + `
普通转划线赢钱30-40倍的总钱数:::: ` + strconv.Itoa(u.Result.N14) + `
普通转划线赢钱40-50倍的次数:::: ` + strconv.Itoa(u.Result.N15) + `
普通转划线赢钱40-50倍的总钱数:::: ` + strconv.Itoa(u.Result.N16) + `
普通转划线赢钱50-100倍的次数:::: ` + strconv.Itoa(u.Result.N17) + `
普通转划线赢钱50-100倍的总钱数:::: ` + strconv.Itoa(u.Result.N18) + `
普通转划线赢钱100-200倍的次数:::: ` + strconv.Itoa(u.Result.N19) + `
普通转划线赢钱100-200倍的总钱数:::: ` + strconv.Itoa(u.Result.N20) + `
普通转划线赢钱200-500倍的次数:::: ` + strconv.Itoa(u.Result.N21) + `
普通转划线赢钱200-500倍的总钱数:::: ` + strconv.Itoa(u.Result.N22) + `
普通转划线赢钱500-1000倍的次数:::: ` + strconv.Itoa(u.Result.N23) + `
普通转划线赢钱500-1000倍的总钱数:::: ` + strconv.Itoa(u.Result.N24) + `
普通转划线赢钱1000以上倍的次数:::: ` + strconv.Itoa(u.Result.N25) + `
普通转划线赢钱1000以上倍的总钱数:::: ` + strconv.Itoa(u.Result.N26) + `
普通转连续消除1次的次数:::: ` + strconv.Itoa(u.Result.N27) + `
普通转连续消除1次的赢钱数:::: ` + strconv.Itoa(u.Result.N28) + `
普通转连续消除2次的次数:::: ` + strconv.Itoa(u.Result.N29) + `
普通转连续消除2次的赢钱数:::: ` + strconv.Itoa(u.Result.N30) + `
普通转连续消除3次的次数:::: ` + strconv.Itoa(u.Result.N31) + `
普通转连续消除3次的赢钱数:::: ` + strconv.Itoa(u.Result.N32) + `
普通转连续消除4次的次数:::: ` + strconv.Itoa(u.Result.N33) + `
普通转连续消除4次的赢钱数:::: ` + strconv.Itoa(u.Result.N34) + `
普通转连续消除5次的次数:::: ` + strconv.Itoa(u.Result.N35) + `
普通转连续消除5次的赢钱数:::: ` + strconv.Itoa(u.Result.N36) + `
普通转连续消除6次的次数:::: ` + strconv.Itoa(u.Result.N37) + `
普通转连续消除6次的赢钱数:::: ` + strconv.Itoa(u.Result.N38) + `
普通转连续消除7次的次数:::: ` + strconv.Itoa(u.Result.N39) + `
普通转连续消除7次的赢钱数:::: ` + strconv.Itoa(u.Result.N40) + `
普通转连续消除8次的次数:::: ` + strconv.Itoa(u.Result.N41) + `
普通转连续消除8次的赢钱数:::: ` + strconv.Itoa(u.Result.N42) + `
普通转连续消除9次的次数:::: ` + strconv.Itoa(u.Result.N43) + `
普通转连续消除9次的赢钱数:::: ` + strconv.Itoa(u.Result.N44) + `
普通转连续消除10次的次数:::: ` + strconv.Itoa(u.Result.N45) + `
普通转连续消除10次的赢钱数:::: ` + strconv.Itoa(u.Result.N46) + `
普通转连续消除10+次的次数:::: ` + strconv.Itoa(u.Result.N47) + `
普通转连续消除10+次的赢钱数:::: ` + strconv.Itoa(u.Result.N48) + `
FREESPIN初始触发的次数:::: ` + strconv.Itoa(u.Result.N49) + `
FREESPIN再次触发FREESPIN的次数:::: ` + strconv.Itoa(u.Result.N50) + `
Freespin划线赢钱0-5倍的次数:::: ` + strconv.Itoa(u.Result.N51) + `
Freespin划线赢钱0-5倍的总钱数:::: ` + strconv.Itoa(u.Result.N52) + `
Freespin划线赢钱5-10倍的次数:::: ` + strconv.Itoa(u.Result.N53) + `
Freespin划线赢钱5-10倍的总钱数:::: ` + strconv.Itoa(u.Result.N54) + `
Freespin划线赢钱10-20倍的次数:::: ` + strconv.Itoa(u.Result.N55) + `
Freespin划线赢钱10-20倍的总钱数:::: ` + strconv.Itoa(u.Result.N56) + `
Freespin划线赢钱20-30倍的次数:::: ` + strconv.Itoa(u.Result.N57) + `
Freespin划线赢钱20-30倍的总钱数:::: ` + strconv.Itoa(u.Result.N58) + `
Freespin划线赢钱30-40倍的次数:::: ` + strconv.Itoa(u.Result.N59) + `
Freespin划线赢钱30-40倍的总钱数:::: ` + strconv.Itoa(u.Result.N60) + `
Freespin划线赢钱40-50倍的次数:::: ` + strconv.Itoa(u.Result.N61) + `
Freespin划线赢钱40-50倍的总钱数:::: ` + strconv.Itoa(u.Result.N62) + `
Freespin划线赢钱50-100倍的次数:::: ` + strconv.Itoa(u.Result.N63) + `
Freespin划线赢钱50-100倍的总钱数:::: ` + strconv.Itoa(u.Result.N64) + `
Freespin划线赢钱100-200倍的次数:::: ` + strconv.Itoa(u.Result.N65) + `
Freespin划线赢钱100-200倍的总钱数:::: ` + strconv.Itoa(u.Result.N66) + `
Freespin划线赢钱200-500倍的次数:::: ` + strconv.Itoa(u.Result.N67) + `
Freespin划线赢钱200-500倍的总钱数:::: ` + strconv.Itoa(u.Result.N68) + `
Freespin划线赢钱500-1000倍的次数:::: ` + strconv.Itoa(u.Result.N69) + `
Freespin划线赢钱500-1000倍的总钱数:::: ` + strconv.Itoa(u.Result.N70) + `
Freespin划线赢钱1000以上倍的次数:::: ` + strconv.Itoa(u.Result.N71) + `
Freespin划线赢钱1000以上倍的总钱数:::: ` + strconv.Itoa(u.Result.N72) + `
Freespin连续消除1次的次数:::: ` + strconv.Itoa(u.Result.N73) + `
Freespin连续消除1次的赢钱数:::: ` + strconv.Itoa(u.Result.N74) + `
Freespin连续消除2次的次数:::: ` + strconv.Itoa(u.Result.N75) + `
Freespin连续消除2次的赢钱数:::: ` + strconv.Itoa(u.Result.N76) + `
Freespin连续消除3次的次数:::: ` + strconv.Itoa(u.Result.N77) + `
Freespin连续消除3次的赢钱数:::: ` + strconv.Itoa(u.Result.N78) + `
Freespin连续消除4次的次数:::: ` + strconv.Itoa(u.Result.N79) + `
Freespin连续消除4次的赢钱数:::: ` + strconv.Itoa(u.Result.N80) + `
Freespin连续消除5次的次数:::: ` + strconv.Itoa(u.Result.N81) + `
Freespin连续消除5次的赢钱数:::: ` + strconv.Itoa(u.Result.N82) + `
Freespin连续消除6次的次数:::: ` + strconv.Itoa(u.Result.N83) + `
Freespin连续消除6次的赢钱数:::: ` + strconv.Itoa(u.Result.N84) + `
Freespin连续消除7次的次数:::: ` + strconv.Itoa(u.Result.N85) + `
Freespin连续消除7次的赢钱数:::: ` + strconv.Itoa(u.Result.N86) + `
Freespin连续消除8次的次数:::: ` + strconv.Itoa(u.Result.N87) + `
Freespin连续消除8次的赢钱数:::: ` + strconv.Itoa(u.Result.N88) + `
Freespin连续消除9次的次数:::: ` + strconv.Itoa(u.Result.N89) + `
Freespin连续消除9次的赢钱数:::: ` + strconv.Itoa(u.Result.N90) + `
Freespin连续消除10次的次数:::: ` + strconv.Itoa(u.Result.N91) + `
Freespin连续消除10次的赢钱数:::: ` + strconv.Itoa(u.Result.N92) + `
Freespin连续消除10+次的次数:::: ` + strconv.Itoa(u.Result.N93) + `
Freespin连续消除10+次的赢钱数:::: ` + strconv.Itoa(u.Result.N94) + `
3个scatter触发Freespin的次数::::: ` + strconv.Itoa(u.Result.N95) + `
3个scatter触发Freespin的总赢钱数::::: ` + strconv.Itoa(u.Result.N96) + "\n"

	//for i := 0; i < 1000; i++ {
	//	str += fmt.Sprintf("初始 %d 的次数::::: %d\r\n", i, u.Result.MapStart[i])
	//}
	//for i := 0; i < 1000; i++ {
	//	str += fmt.Sprintf("初始 %d 的赢钱数::::: %d\r\n", i, u.Result.MapStartWin[i])
	//}
	return str
}

//func (u *Unit8) GetResult() {
//
//}

func (u *Unit8) GetBasicData() *public.BasicData {
	return u.Result.BasicData
}

func (u *Unit8) GetResult(spins []*component.Spin) (int64, int64) {
	gain := lo.SumBy(spins, func(spin *component.Spin) int64 {
		return int64(spin.Gain)
	})
	bet := int64(spins[0].Bet) + int64(spins[0].Raise) + int64(spins[0].BuyFreeCoin) + int64(spins[0].BuyReCoin)
	return gain, bet
}

func (u *Unit8) GetDataResult() *Unit8Result {
	return &u.Result
}

func TestOnce(run public.RunSlotTest, opts ...component.Option) error {
	var (
		spin         = &component.Spin{}
		slotTests    []business.SlotTests
		mainSlotTest business.SlotTests
	)
	spin, _ = component.NewSpin(run.SlotId, run.Amount, opts...)
	if run.Type == enum.SlotTestType4Result {
		strs := utils.FormatCommand(run.Result)
		if len(strs) != spin.Config.Index*spin.Config.Row {
			return errors.New("结果标签数量错误")
		}
		var arr = make([][]*base.Tag, spin.Config.Row)
		for i := 0; i < spin.Config.Row; i++ {
			for j := 0; j < spin.Config.Index; j++ {
				tagName := strings.TrimSpace(strs[i*spin.Config.Index+j])
				fillTag := spin.Config.GetTag(tagName)
				fillTag.X = i
				fillTag.Y = j
				arr[i] = append(arr[i], fillTag)
			}
		}
		opts = append(opts, component.WithSetResult(arr))
	}

	if run.IsMustFree > 0 {
		opts = append(opts, component.WithIsMustFree())
	}
	m, err := slot.Play(run.SlotId, run.Amount, opts...)
	if err != nil {
		return err
	}
	spin = m.GetSpin()
	spins := m.GetSpins()
	for _, flow := range spin.SpinInfo.SpinFlow {
		slotTest := business.SlotTests{
			Type:     uint8(run.Type),
			SlotId:   run.SlotId,
			Hold:     0,
			Amount:   run.Amount,
			Win:      helper.If(flow.Id == 0, spin.Gain, 0),
			MaxNum:   1,
			RunNum:   1,
			Detail:   flow.String(),
			Status:   enum.CommonStatusFinish,
			Bet:      spin.Bet,
			Raise:    helper.If(flow.Id == 0, int(spin.Raise)+int(spin.BuyFreeCoin)+int(spin.BuyReCoin), 0),
			GameType: helper.If(flow.Id == 0, enum.SlotSpinType1Normal, enum.SlotSpinType3Respin),
			TestId:   helper.If(flow.Id == 0, 0, int(mainSlotTest.ID)),
			Rank:     spin.Rank,
			GameData: helper.If(flow.Id == 0, "整局信息:"+spin.SpinInfo.GetInfo()+"\r\n 本次转信息: "+flow.GetInformation(), flow.GetInformation()),
			//
		}
		if flow.Id == 0 {
			mainSlotTest = slotTest
			err = global.GVA_DB.Create(&mainSlotTest).Error
		} else {
			slotTests = append(slotTests, slotTest)
		}
	}
	for _, c := range spins {
		for _, flow := range c.SpinInfo.SpinFlow {
			slotTest := business.SlotTests{
				Type:     uint8(run.Type),
				SlotId:   run.SlotId,
				Hold:     0,
				Amount:   run.Amount,
				Win:      helper.If(flow.Id == 0, c.Gain, 0),
				MaxNum:   1,
				RunNum:   1,
				Detail:   flow.String(),
				Status:   enum.CommonStatusFinish,
				Bet:      spin.Bet,
				Raise:    0,
				GameType: helper.If(flow.Id == 0, enum.SlotSpinType2Fs, enum.SlotSpinType4FsRs),
				TestId:   int(mainSlotTest.ID),
				Rank:     spin.Rank,
				GameData: helper.If(flow.Id == 0, "整局信息:"+c.SpinInfo.GetInfo()+"\r\n 本次转信息: "+flow.GetInformation(), flow.GetInformation()),
			}
			slotTests = append(slotTests, slotTest)
		}
	}

	if len(slotTests) > 0 {
		err = global.GVA_DB.Create(&slotTests).Error
	}

	return err
}

func (u *Unit8) GetReturnRatio() float64 {
	f, _ := decimal.NewFromInt(int64(u.Result.N1)).Div(decimal.NewFromInt(int64(u.Result.N2))).Float64()
	return f
}
