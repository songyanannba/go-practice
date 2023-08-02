package unit4

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"slot-server/model/business"
	"slot-server/service/logic/gameHandle/slotHandle"
	"slot-server/service/slot"
	"slot-server/service/slot/base"
	"slot-server/service/slot/component"
	"slot-server/service/test/public"
	"slot-server/utils/helper"
	"sort"
	"strconv"
)

type Unit4 struct {
	Result unit4Result
}

type unit4Result struct {
	*public.BasicData
	N27               int //	普通转触发的RESPIN触发次数:::::
	N28               int //	普通转触发的RESPIN划线赢钱赢钱0-5倍的次数:::::
	N29               int //	普通转触发的RESPIN划线赢钱赢钱0-5倍的总钱数:::::
	N30               int //	普通转触发的RESPIN划线赢钱赢钱5-10倍的次数:::::
	N31               int //	普通转触发的RESPIN划线赢钱赢钱5-10倍的总钱数:::::
	N32               int //	普通转触发的RESPIN划线赢钱赢钱10-20倍的次数:::::
	N33               int //	普通转触发的RESPIN划线赢钱赢钱10-20倍的总钱数:::::
	N34               int //	普通转触发的RESPIN划线赢钱赢钱20-30倍的次数:::::
	N35               int //	普通转触发的RESPIN划线赢钱赢钱20-30倍的总钱数:::::
	N36               int //	普通转触发的RESPIN划线赢钱赢钱30-40倍的次数:::::
	N37               int //	普通转触发的RESPIN划线赢钱赢钱30-40倍的总钱数:::::
	N38               int //	普通转触发的RESPIN划线赢钱赢钱40-50倍的次数:::::
	N39               int //	普通转触发的RESPIN划线赢钱赢钱40-50倍的总钱数:::::
	N40               int //	普通转触发的RESPIN划线赢钱赢钱50-100倍的次数:::::
	N41               int //	普通转触发的RESPIN划线赢钱赢钱50-100倍的总钱数:::::
	N42               int //	普通转触发的RESPIN划线赢钱赢钱100-200倍的次数:::::
	N43               int //	普通转触发的RESPIN划线赢钱赢钱100-200倍的总钱数:::::
	N44               int //	普通转触发的RESPIN划线赢钱赢钱200-500倍的次数:::::
	N45               int //	普通转触发的RESPIN划线赢钱赢钱200-500倍的总钱数:::::
	N46               int //	普通转触发的RESPIN划线赢钱赢钱500-1000倍的次数:::::
	N47               int //	普通转触发的RESPIN划线赢钱赢钱500-1000倍的总钱数:::::
	N48               int //	普通转触发的RESPIN划线赢钱赢钱1000以上倍的次数:::::
	N49               int //	普通转触发的RESPIN划线赢钱赢钱1000以上倍的总钱数:::::
	N50               int //	FREESPIN初始触发的次数:::::
	N51               int //	FREESPIN再次触发FREESPIN的次数:::::
	N52               int //	Freespin划线赢钱0-5倍的次数:::::
	N53               int //	Freespin划线赢钱0-5倍的总钱数:::::
	N54               int //	Freespin划线赢钱5-10倍的次数:::::
	N55               int //	Freespin划线赢钱5-10倍的总钱数:::::
	N56               int //	Freespin划线赢钱10-20倍的次数:::::
	N57               int //	Freespin划线赢钱10-20倍的总钱数:::::
	N58               int //	Freespin划线赢钱20-30倍的次数:::::
	N59               int //	Freespin划线赢钱20-30倍的总钱数:::::
	N60               int //	Freespin划线赢钱30-40倍的次数:::::
	N61               int //	Freespin划线赢钱30-40倍的总钱数:::::
	N62               int //	Freespin划线赢钱40-50倍的次数:::::
	N63               int //	Freespin划线赢钱40-50倍的总钱数:::::
	N64               int //	Freespin划线赢钱50-100倍的次数:::::
	N65               int //	Freespin划线赢钱50-100倍的总钱数:::::
	N66               int //	Freespin划线赢钱100-200倍的次数:::::
	N67               int //	Freespin划线赢钱100-200倍的总钱数:::::
	N68               int //	Freespin划线赢钱200-500倍的次数:::::
	N69               int //	Freespin划线赢钱200-500倍的总钱数:::::
	N70               int //	Freespin划线赢钱500-1000倍的次数:::::
	N71               int //	Freespin划线赢钱500-1000倍的总钱数:::::
	N72               int //	Freespin划线赢钱1000以上倍的次数:::::
	N73               int //	Freespin划线赢钱1000以上倍的总钱数:::::
	N74               int //	FREESPIN触发RESPIN的次数:::::
	N75               int //	FREESPIN触发的RESPIN划线赢钱0-5倍的次数:::::
	N76               int //	FREESPIN触发的RESPIN划线赢钱0-5倍的总钱数:::::
	N77               int //	FREESPIN触发的RESPIN划线赢钱5-10倍的次数:::::
	N78               int //	FREESPIN触发的RESPIN划线赢钱5-10倍的总钱数:::::
	N79               int //	FREESPIN触发的RESPIN划线赢钱10-20倍的次数:::::
	N80               int //	FREESPIN触发的RESPIN划线赢钱10-20倍的总钱数:::::
	N81               int //	FREESPIN触发的RESPIN划线赢钱20-30倍的次数:::::
	N82               int //	FREESPIN触发的RESPIN划线赢钱20-30倍的总钱数:::::
	N83               int //	FREESPIN触发的RESPIN划线赢钱30-40倍的次数:::::
	N84               int //	FREESPIN触发的RESPIN划线赢钱30-40倍的总钱数:::::
	N85               int //	FREESPIN触发的RESPIN划线赢钱40-50倍的次数:::::
	N86               int //	FREESPIN触发的RESPIN划线赢钱40-50倍的总钱数:::::
	N87               int //	FREESPIN触发的RESPIN划线赢钱50-100倍的次数:::::
	N88               int //	FREESPIN触发的RESPIN划线赢钱50-100倍的总钱数:::::
	N89               int //	FREESPIN触发的RESPIN划线赢钱100-200倍的次数:::::
	N90               int //	FREESPIN触发的RESPIN划线赢钱100-200倍的总钱数:::::
	N91               int //	FREESPIN触发的RESPIN划线赢钱200-500倍的次数:::::
	N92               int //	FREESPIN触发的RESPIN划线赢钱200-500倍的总钱数:::::
	N93               int //	FREESPIN触发的RESPIN划线赢钱500-1000倍的次数:::::
	N94               int //	FREESPIN触发的RESPIN划线赢钱500-1000倍的总钱数:::::
	N95               int //	FREESPIN触发的RESPIN划线赢钱1000以上倍的次数:::::
	N96               int //	FREESPIN触发的RESPIN划线赢钱1000以上倍的总钱数:::::
	N97               int //	单独普通转赢钱金额:::::
	N98               int //	初始进入Link_coin标签个数6的次数:::::
	N99               int //	初始进入Link_coin标签个数6的赢钱:::::
	N100              int //初始进入Link_coin标签个数7的次数:::::
	N101              int //初始进入Link_coin标签个数7的赢钱:::::
	N102              int //初始进入Link_coin标签个数8的次数:::::
	N103              int //初始进入Link_coin标签个数8的赢钱:::::
	N104              int //初始进入Link_coin标签个数9的次数:::::
	N105              int //初始进入Link_coin标签个数9的赢钱:::::
	N106              int //初始进入Link_coin标签个数10的次数:::::
	N107              int //初始进入Link_coin标签个数10的赢钱:::::
	N108              int //初始进入Link_coin标签个数11的次数:::::
	N109              int //初始进入Link_coin标签个数11的赢钱:::::
	N110              int //初始进入Link_coin标签个数12的次数:::::
	N111              int //初始进入Link_coin标签个数12的赢钱:::::
	N112              int //初始进入Link_coin标签个数13的次数:::::
	N113              int //初始进入Link_coin标签个数13的赢钱:::::
	N114              int //初始进入Link_coin标签个数14的次数:::::
	N115              int //初始进入Link_coin标签个数14的赢钱:::::
	N116              int //初始进入Link_coin标签个数15的次数:::::
	N117              int //初始进入Link_coin标签个数15的赢钱:::::
	coinMultiplierMap map[float64]int
	coin6SumCoin      map[int]int
}

func NewUnit(run public.RunSlotTest) *Unit4 {
	return &Unit4{Result: unit4Result{BasicData: &public.BasicData{},
		coinMultiplierMap: make(map[float64]int),
		coin6SumCoin:      make(map[int]int),
	}}
}

func (u *Unit4) Repeat(run public.RunSlotTest, opts ...component.Option) ([]*component.Spin, error) {
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

func (u *Unit4) Calculate(mSpins []*component.Spin) {
	public.AnalyticData(u.Result.BasicData, mSpins)
	spin := mSpins[0]
	u.Result.N97 += spin.Gain
	if spin.FreeSpinParams.ReNum > 0 {
		u.Result.N27++
	}
	//普通转触发ReSpin次数统计
	{
		firstReSpins := lo.Filter(mSpins, func(item *component.Spin, index int) bool {
			return item.ParentId == 1 && item.Gain > 0 && item.IsReSpin
		})
		for _, re := range firstReSpins {
			Multiple := float64(re.Gain) / float64(spin.Bet)
			if Multiple >= 0 && Multiple < 5 {
				u.Result.N28++
				u.Result.N29 += re.Gain
			} else if Multiple >= 5 && Multiple < 10 {
				u.Result.N30++
				u.Result.N31 += re.Gain
			} else if Multiple >= 10 && Multiple < 20 {
				u.Result.N32++
				u.Result.N33 += re.Gain
			} else if Multiple >= 20 && Multiple < 30 {
				u.Result.N34++
				u.Result.N35 += re.Gain
			} else if Multiple >= 30 && Multiple < 40 {
				u.Result.N36++
				u.Result.N37 += re.Gain
			} else if Multiple >= 40 && Multiple < 50 {
				u.Result.N38++
				u.Result.N39 += re.Gain
			} else if Multiple >= 50 && Multiple < 100 {
				u.Result.N40++
				u.Result.N41 += re.Gain
			} else if Multiple >= 100 && Multiple < 200 {
				u.Result.N42++
				u.Result.N43 += re.Gain
			} else if Multiple >= 200 && Multiple < 500 {
				u.Result.N44++
				u.Result.N45 += re.Gain
			} else if Multiple >= 500 && Multiple < 1000 {
				u.Result.N46++
				u.Result.N47 += re.Gain
			} else if Multiple >= 1000 {
				u.Result.N48++
				u.Result.N49 += re.Gain
			}
		}
	}

	if spin.FreeSpinParams.FreeNum > 0 {
		u.Result.N50++
	}

	FreeFrees := lo.Filter(mSpins, func(item *component.Spin, index int) bool {
		return item.IsFree && item.FreeSpinParams.FreeNum > 0
	})
	u.Result.N51 += len(FreeFrees)

	triggerFrees := lo.Filter(mSpins, func(item *component.Spin, index int) bool {
		return item.FreeSpinParams.FreeNum > 0
	})

	for _, tFree := range triggerFrees {
		frees := lo.Filter(mSpins, func(item *component.Spin, index int) bool {
			return item.ParentId == tFree.Id && item.IsFree
		})

		fSum := lo.SumBy(frees, func(item *component.Spin) int {
			return item.Gain
		})
		Multiple := float64(fSum) / float64(spin.Bet)
		if Multiple >= 0 && Multiple < 5 {
			u.Result.N52++
			u.Result.N53 += fSum
		} else if Multiple >= 5 && Multiple < 10 {
			u.Result.N54++
			u.Result.N55 += fSum

		} else if Multiple >= 10 && Multiple < 20 {
			u.Result.N56++
			u.Result.N57 += fSum
		} else if Multiple >= 20 && Multiple < 30 {
			u.Result.N58++
			u.Result.N59 += fSum
		} else if Multiple >= 30 && Multiple < 40 {
			u.Result.N60++
			u.Result.N61 += fSum
		} else if Multiple >= 40 && Multiple < 50 {
			u.Result.N62++
			u.Result.N63 += fSum
		} else if Multiple >= 50 && Multiple < 100 {
			u.Result.N64++
			u.Result.N65 += fSum
		} else if Multiple >= 100 && Multiple < 200 {
			u.Result.N66++
			u.Result.N67 += fSum
		} else if Multiple >= 200 && Multiple < 500 {
			u.Result.N68++
			u.Result.N69 += fSum
		} else if Multiple >= 500 && Multiple < 1000 {
			u.Result.N70++
			u.Result.N71 += fSum
		} else if Multiple >= 1000 {
			u.Result.N72++
			u.Result.N73 += fSum
		}
	}

	freeRes := lo.Filter(mSpins, func(item *component.Spin, index int) bool {
		return item.IsFree && item.FreeSpinParams.ReNum > 0
	})

	u.Result.N74 += len(freeRes)

	for _, freeRe := range freeRes {

		FreeSpinRes := lo.Filter(mSpins, func(item *component.Spin, index int) bool {
			return item.ParentId == freeRe.Id && item.IsReSpin && item.Gain > 0
		})

		sumFreeSpinRes := lo.SumBy(FreeSpinRes, func(item *component.Spin) int {
			return item.Gain
		})
		{
			Multiple := float64(sumFreeSpinRes) / float64(spin.Bet)
			if Multiple >= 0 && Multiple < 5 {
				u.Result.N75++
				u.Result.N76 += sumFreeSpinRes
			} else if Multiple >= 5 && Multiple < 10 {
				u.Result.N77++
				u.Result.N78 += sumFreeSpinRes
			} else if Multiple >= 10 && Multiple < 20 {
				u.Result.N79++
				u.Result.N80 += sumFreeSpinRes
			} else if Multiple >= 20 && Multiple < 30 {
				u.Result.N81++
				u.Result.N82 += sumFreeSpinRes
			} else if Multiple >= 30 && Multiple < 40 {
				u.Result.N83++
				u.Result.N84 += sumFreeSpinRes
			} else if Multiple >= 40 && Multiple < 50 {
				u.Result.N85++
				u.Result.N86 += sumFreeSpinRes
			} else if Multiple >= 50 && Multiple < 100 {
				u.Result.N87++
				u.Result.N88 += sumFreeSpinRes
			} else if Multiple >= 100 && Multiple < 200 {
				u.Result.N89++
				u.Result.N90 += sumFreeSpinRes
			} else if Multiple >= 200 && Multiple < 500 {
				u.Result.N91++
				u.Result.N92 += sumFreeSpinRes
			} else if Multiple >= 500 && Multiple < 1000 {
				u.Result.N93++
				u.Result.N94 += sumFreeSpinRes
			} else if Multiple >= 1000 {
				u.Result.N95++
				u.Result.N96 += sumFreeSpinRes
			}
		}
	}
	u.Result.N97 += spin.Gain

	likeList := base.GetSpecialTags(spin.InitDataList, "link_coin", "link_collect")

	firstRespins := lo.Filter(mSpins, func(item *component.Spin, index int) bool {
		return item.ParentId == 1 && item.Gain > 0 && item.IsReSpin
	})
	firstRespinSum := lo.SumBy(firstRespins, func(item *component.Spin) int {
		return item.Gain
	})
	switch len(likeList) {
	case 6:
		u.Result.N98++
		u.Result.N99 += firstRespinSum
		coin6SumNum := base.GetSpecialTags(helper.SliceVal(firstRespins, 0).InitDataList, "link_coin", "link_collect")
		u.Result.coin6SumCoin[len(coin6SumNum)]++
	case 7:
		u.Result.N100++
		u.Result.N101 += firstRespinSum
	case 8:
		u.Result.N102++
		u.Result.N103 += firstRespinSum
	case 9:
		u.Result.N104++
		u.Result.N105 += firstRespinSum
	case 10:
		u.Result.N106++
		u.Result.N107 += firstRespinSum
	case 11:
		u.Result.N108++
		u.Result.N109 += firstRespinSum
	case 12:
		u.Result.N110++
		u.Result.N111 += firstRespinSum
	case 13:
		u.Result.N112++
		u.Result.N113 += firstRespinSum
	case 14:
		u.Result.N114++
		u.Result.N115 += firstRespinSum
	case 15:
		u.Result.N116++
		u.Result.N117 += firstRespinSum + spin.Gain
	}
}

func (u *Unit4) GetDetail() string {
	str :=
		`总赢钱:::: ` + strconv.Itoa(u.Result.N1) + `
总押注消耗钱::::: ` + strconv.Itoa(u.Result.N2) + `
总转动次数::::: ` + strconv.Itoa(u.Result.N3) + `
总赢钱次数::::: ` + strconv.Itoa(u.Result.N4) + `
普通转划线赢钱0-5倍的次数::::: ` + strconv.Itoa(u.Result.N5) + `
普通转划线赢钱0-5倍的总钱数::::: ` + strconv.Itoa(u.Result.N6) + `
普通转划线赢钱5-10倍的次数::::: ` + strconv.Itoa(u.Result.N7) + `
普通转划线赢钱5-10倍的总钱数::::: ` + strconv.Itoa(u.Result.N8) + `
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
普通转触发的RESPIN触发次数::::: ` + strconv.Itoa(u.Result.N27) + `
普通转触发的RESPIN划线赢钱赢钱0-5倍的次数::::: ` + strconv.Itoa(u.Result.N28) + `
普通转触发的RESPIN划线赢钱赢钱0-5倍的总钱数::::: ` + strconv.Itoa(u.Result.N29) + `
普通转触发的RESPIN划线赢钱赢钱5-10倍的次数::::: ` + strconv.Itoa(u.Result.N30) + `
普通转触发的RESPIN划线赢钱赢钱5-10倍的总钱数::::: ` + strconv.Itoa(u.Result.N31) + `
普通转触发的RESPIN划线赢钱赢钱10-20倍的次数::::: ` + strconv.Itoa(u.Result.N32) + `
普通转触发的RESPIN划线赢钱赢钱10-20倍的总钱数::::: ` + strconv.Itoa(u.Result.N33) + `
普通转触发的RESPIN划线赢钱赢钱20-30倍的次数::::: ` + strconv.Itoa(u.Result.N34) + `
普通转触发的RESPIN划线赢钱赢钱20-30倍的总钱数::::: ` + strconv.Itoa(u.Result.N35) + `
普通转触发的RESPIN划线赢钱赢钱30-40倍的次数::::: ` + strconv.Itoa(u.Result.N36) + `
普通转触发的RESPIN划线赢钱赢钱30-40倍的总钱数::::: ` + strconv.Itoa(u.Result.N37) + `
普通转触发的RESPIN划线赢钱赢钱40-50倍的次数::::: ` + strconv.Itoa(u.Result.N38) + `
普通转触发的RESPIN划线赢钱赢钱40-50倍的总钱数::::: ` + strconv.Itoa(u.Result.N39) + `
普通转触发的RESPIN划线赢钱赢钱50-100倍的次数::::: ` + strconv.Itoa(u.Result.N40) + `
普通转触发的RESPIN划线赢钱赢钱50-100倍的总钱数::::: ` + strconv.Itoa(u.Result.N41) + `
普通转触发的RESPIN划线赢钱赢钱100-200倍的次数::::: ` + strconv.Itoa(u.Result.N42) + `
普通转触发的RESPIN划线赢钱赢钱100-200倍的总钱数::::: ` + strconv.Itoa(u.Result.N43) + `
普通转触发的RESPIN划线赢钱赢钱200-500倍的次数::::: ` + strconv.Itoa(u.Result.N44) + `
普通转触发的RESPIN划线赢钱赢钱200-500倍的总钱数::::: ` + strconv.Itoa(u.Result.N45) + `
普通转触发的RESPIN划线赢钱赢钱500-1000倍的次数::::: ` + strconv.Itoa(u.Result.N46) + `
普通转触发的RESPIN划线赢钱赢钱500-1000倍的总钱数::::: ` + strconv.Itoa(u.Result.N47) + `
普通转触发的RESPIN划线赢钱赢钱1000以上倍的次数::::: ` + strconv.Itoa(u.Result.N48) + `
普通转触发的RESPIN划线赢钱赢钱1000以上倍的总钱数::::: ` + strconv.Itoa(u.Result.N49) + `
FREESPIN初始触发的次数::::: ` + strconv.Itoa(u.Result.N50) + `
FREESPIN再次触发FREESPIN的次数::::: ` + strconv.Itoa(u.Result.N51) + `
Freespin划线赢钱0-5倍的次数::::: ` + strconv.Itoa(u.Result.N52) + `
Freespin划线赢钱0-5倍的总钱数::::: ` + strconv.Itoa(u.Result.N53) + `
Freespin划线赢钱5-10倍的次数::::: ` + strconv.Itoa(u.Result.N54) + `
Freespin划线赢钱5-10倍的总钱数::::: ` + strconv.Itoa(u.Result.N55) + `
Freespin划线赢钱10-20倍的次数::::: ` + strconv.Itoa(u.Result.N56) + `
Freespin划线赢钱10-20倍的总钱数::::: ` + strconv.Itoa(u.Result.N57) + `
Freespin划线赢钱20-30倍的次数::::: ` + strconv.Itoa(u.Result.N58) + `
Freespin划线赢钱20-30倍的总钱数::::: ` + strconv.Itoa(u.Result.N59) + `
Freespin划线赢钱30-40倍的次数::::: ` + strconv.Itoa(u.Result.N60) + `
Freespin划线赢钱30-40倍的总钱数::::: ` + strconv.Itoa(u.Result.N61) + `
Freespin划线赢钱40-50倍的次数::::: ` + strconv.Itoa(u.Result.N62) + `
Freespin划线赢钱40-50倍的总钱数::::: ` + strconv.Itoa(u.Result.N63) + `
Freespin划线赢钱50-100倍的次数::::: ` + strconv.Itoa(u.Result.N64) + `
Freespin划线赢钱50-100倍的总钱数::::: ` + strconv.Itoa(u.Result.N65) + `
Freespin划线赢钱100-200倍的次数::::: ` + strconv.Itoa(u.Result.N66) + `
Freespin划线赢钱100-200倍的总钱数::::: ` + strconv.Itoa(u.Result.N67) + `
Freespin划线赢钱200-500倍的次数::::: ` + strconv.Itoa(u.Result.N68) + `
Freespin划线赢钱200-500倍的总钱数::::: ` + strconv.Itoa(u.Result.N69) + `
Freespin划线赢钱500-1000倍的次数::::: ` + strconv.Itoa(u.Result.N70) + `
Freespin划线赢钱500-1000倍的总钱数::::: ` + strconv.Itoa(u.Result.N71) + `
Freespin划线赢钱1000以上倍的次数::::: ` + strconv.Itoa(u.Result.N72) + `
Freespin划线赢钱1000以上倍的总钱数::::: ` + strconv.Itoa(u.Result.N73) + `
FREESPIN触发RESPIN的次数::::: ` + strconv.Itoa(u.Result.N74) + `
FREESPIN触发的RESPIN划线赢钱0-5倍的次数::::: ` + strconv.Itoa(u.Result.N75) + `
FREESPIN触发的RESPIN划线赢钱0-5倍的总钱数::::: ` + strconv.Itoa(u.Result.N76) + `
FREESPIN触发的RESPIN划线赢钱5-10倍的次数::::: ` + strconv.Itoa(u.Result.N77) + `
FREESPIN触发的RESPIN划线赢钱5-10倍的总钱数::::: ` + strconv.Itoa(u.Result.N78) + `
FREESPIN触发的RESPIN划线赢钱10-20倍的次数::::: ` + strconv.Itoa(u.Result.N79) + `
FREESPIN触发的RESPIN划线赢钱10-20倍的总钱数::::: ` + strconv.Itoa(u.Result.N80) + `
FREESPIN触发的RESPIN划线赢钱20-30倍的次数::::: ` + strconv.Itoa(u.Result.N81) + `
FREESPIN触发的RESPIN划线赢钱20-30倍的总钱数::::: ` + strconv.Itoa(u.Result.N82) + `
FREESPIN触发的RESPIN划线赢钱30-40倍的次数::::: ` + strconv.Itoa(u.Result.N83) + `
FREESPIN触发的RESPIN划线赢钱30-40倍的总钱数::::: ` + strconv.Itoa(u.Result.N84) + `
FREESPIN触发的RESPIN划线赢钱40-50倍的次数::::: ` + strconv.Itoa(u.Result.N85) + `
FREESPIN触发的RESPIN划线赢钱40-50倍的总钱数::::: ` + strconv.Itoa(u.Result.N86) + `
FREESPIN触发的RESPIN划线赢钱50-100倍的次数::::: ` + strconv.Itoa(u.Result.N87) + `
FREESPIN触发的RESPIN划线赢钱50-100倍的总钱数::::: ` + strconv.Itoa(u.Result.N88) + `
FREESPIN触发的RESPIN划线赢钱100-200倍的次数::::: ` + strconv.Itoa(u.Result.N89) + `
FREESPIN触发的RESPIN划线赢钱100-200倍的总钱数::::: ` + strconv.Itoa(u.Result.N90) + `
FREESPIN触发的RESPIN划线赢钱200-500倍的次数::::: ` + strconv.Itoa(u.Result.N91) + `
FREESPIN触发的RESPIN划线赢钱200-500倍的总钱数::::: ` + strconv.Itoa(u.Result.N92) + `
FREESPIN触发的RESPIN划线赢钱500-1000倍的次数::::: ` + strconv.Itoa(u.Result.N93) + `
FREESPIN触发的RESPIN划线赢钱500-1000倍的总钱数::::: ` + strconv.Itoa(u.Result.N94) + `
FREESPIN触发的RESPIN划线赢钱1000以上倍的次数::::: ` + strconv.Itoa(u.Result.N95) + `
FREESPIN触发的RESPIN划线赢钱1000以上倍的总钱数::::: ` + strconv.Itoa(u.Result.N96) + `
单独普通转赢钱金额::::: ` + strconv.Itoa(u.Result.N97) + `
初始进入Link_coin标签个数6的次数::::: ` + strconv.Itoa(u.Result.N98) + `
初始进入Link_coin标签个数6的总钱数::::: ` + strconv.Itoa(u.Result.N99) + `
初始进入Link_coin标签个数7的次数::::: ` + strconv.Itoa(u.Result.N100) + `
初始进入Link_coin标签个数7的总钱数::::: ` + strconv.Itoa(u.Result.N101) + `
初始进入Link_coin标签个数8的次数::::: ` + strconv.Itoa(u.Result.N102) + `
初始进入Link_coin标签个数8的总钱数::::: ` + strconv.Itoa(u.Result.N103) + `
初始进入Link_coin标签个数9的次数::::: ` + strconv.Itoa(u.Result.N104) + `
初始进入Link_coin标签个数9的总钱数::::: ` + strconv.Itoa(u.Result.N105) + `
初始进入Link_coin标签个数10的次数::::: ` + strconv.Itoa(u.Result.N106) + `
初始进入Link_coin标签个数10的总钱数::::: ` + strconv.Itoa(u.Result.N107) + `
初始进入Link_coin标签个数11的次数::::: ` + strconv.Itoa(u.Result.N108) + `
初始进入Link_coin标签个数11的总钱数::::: ` + strconv.Itoa(u.Result.N109) + `
初始进入Link_coin标签个数12的次数::::: ` + strconv.Itoa(u.Result.N110) + `
初始进入Link_coin标签个数12的总钱数::::: ` + strconv.Itoa(u.Result.N111) + `
初始进入Link_coin标签个数13的次数::::: ` + strconv.Itoa(u.Result.N112) + `
初始进入Link_coin标签个数13的总钱数::::: ` + strconv.Itoa(u.Result.N113) + `
初始进入Link_coin标签个数14的次数::::: ` + strconv.Itoa(u.Result.N114) + `
初始进入Link_coin标签个数14的总钱数::::: ` + strconv.Itoa(u.Result.N115) + `
初始进入Link_coin标签个数15的次数::::: ` + strconv.Itoa(u.Result.N116) + `
初始进入Link_coin标签个数15的总钱数::::: ` + strconv.Itoa(u.Result.N117) + "\r\n"

	var coinMultipliers []float64
	for f, _ := range u.Result.coinMultiplierMap {
		coinMultipliers = append(coinMultipliers, f)
	}
	sort.Slice(coinMultipliers, func(i, j int) bool {
		return coinMultipliers[i] < coinMultipliers[j]
	})
	for _, multiplier := range coinMultipliers {
		str += fmt.Sprintf("Link_coin标签倍率 %g 的次数::::: %d\r\n", multiplier, u.Result.coinMultiplierMap[multiplier])
	}

	var coin6SumCoin []int
	for f, _ := range u.Result.coin6SumCoin {
		coin6SumCoin = append(coin6SumCoin, f)
	}
	sort.Slice(coin6SumCoin, func(i, j int) bool {
		return coin6SumCoin[i] < coin6SumCoin[j]
	})
	for _, multiplier := range coin6SumCoin {
		str += fmt.Sprintf("Link_coin初始标签数量6总标签数%d的次数::::: %d\r\n", multiplier, u.Result.coin6SumCoin[multiplier])
	}

	str += fmt.Sprintf("最大倍率:::::%g\n", u.Result.N0)
	return str
}

func (u *Unit4) GetBasicData() *public.BasicData {
	return u.Result.BasicData
}

func (u *Unit4) GetResult(spins []*component.Spin) (int64, int64) {
	gain := lo.SumBy(spins, func(spin *component.Spin) int64 {
		return int64(spin.Gain)
	})
	bet := int64(spins[0].Bet) + int64(spins[0].Raise) + int64(spins[0].BuyFreeCoin) + int64(spins[0].BuyReCoin)
	return gain, bet
}

func (u *Unit4) GetReturnRatio() float64 {
	f, b := decimal.NewFromInt(int64(u.Result.N1)).Div(decimal.NewFromInt(int64(u.Result.N2))).Float64()
	if b {
		return f
	}
	return 0
}
