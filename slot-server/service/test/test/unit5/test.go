package unit5

import (
	"fmt"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/service/slot"
	"slot-server/service/slot/component"
	"slot-server/service/test/public"
	"slot-server/utils/helper"
	"strconv"
)

type unit5TestResult struct {
	public.BasicData
	N27 int //普通转连续消除1次的次数:::::
	N28 int //普通转连续消除1次的赢钱数:::::
	N29 int //普通转连续消除2次的次数:::::
	N30 int //普通转连续消除2次的赢钱数:::::
	N31 int //普通转连续消除3次的次数:::::
	N32 int //普通转连续消除3次的赢钱数:::::
	N33 int //普通转连续消除4次的次数:::::
	N34 int //普通转连续消除4次的赢钱数:::::
	N35 int //普通转连续消除5次的次数:::::
	N36 int //普通转连续消除5次的赢钱数:::::
	N37 int //普通转连续消除6次的次数:::::
	N38 int //普通转连续消除6次的赢钱数:::::
	N39 int //普通转连续消除7次的次数:::::
	N40 int //普通转连续消除7次的赢钱数:::::
	N41 int //普通转连续消除8次的次数:::::
	N42 int //普通转连续消除8次的赢钱数:::::
	N43 int //普通转连续消除9次的次数:::::
	N44 int //普通转连续消除9次的赢钱数:::::
	N45 int //普通转连续消除10次的次数:::::
	N46 int //普通转连续消除10次的赢钱数:::::
	N47 int //普通转连续消除10+次的次数:::::
	N48 int //普通转连续消除10+次的赢钱数:::::
	N49 int //FREESPIN初始触发的次数:::::
	N50 int //FREESPIN再次触发FREESPIN的次数:::::
	N51 int //Freespin划线赢钱0-5倍的次数:::::
	N52 int //Freespin划线赢钱0-5倍的总钱数:::::
	N53 int //Freespin划线赢钱5-10倍的次数:::::
	N54 int //Freespin划线赢钱5-10倍的总钱数:::::
	N55 int //Freespin划线赢钱10-20倍的次数:::::
	N56 int //Freespin划线赢钱10-20倍的总钱数:::::
	N57 int //Freespin划线赢钱20-30倍的次数:::::
	N58 int //Freespin划线赢钱20-30倍的总钱数:::::
	N59 int //Freespin划线赢钱30-40倍的次数:::::
	N60 int //Freespin划线赢钱30-40倍的总钱数:::::
	N61 int //Freespin划线赢钱40-50倍的次数:::::
	N62 int //Freespin划线赢钱40-50倍的总钱数:::::
	N63 int //Freespin划线赢钱50-100倍的次数:::::
	N64 int //Freespin划线赢钱50-100倍的总钱数:::::
	N65 int //Freespin划线赢钱100-200倍的次数:::::
	N66 int //Freespin划线赢钱100-200倍的总钱数:::::
	N67 int //Freespin划线赢钱200-500倍的次数:::::
	N68 int //Freespin划线赢钱200-500倍的总钱数:::::
	N69 int //Freespin划线赢钱500-1000倍的次数:::::
	N70 int //Freespin划线赢钱500-1000倍的总钱数:::::
	N71 int //Freespin划线赢钱1000以上倍的次数:::::
	N72 int //Freespin划线赢钱1000以上倍的总钱数:::::
	N73 int //Freespin连续消除1次的次数:::::
	N74 int //Freespin连续消除1次的赢钱数:::::
	N75 int //Freespin连续消除2次的次数:::::
	N76 int //Freespin连续消除2次的赢钱数:::::
	N77 int //Freespin连续消除3次的次数:::::
	N78 int //Freespin连续消除3次的赢钱数:::::
	N79 int //Freespin连续消除4次的次数:::::
	N80 int //Freespin连续消除4次的赢钱数:::::
	N81 int //Freespin连续消除5次的次数:::::
	N82 int //Freespin连续消除5次的赢钱数:::::
	N83 int //Freespin连续消除6次的次数:::::
	N84 int //Freespin连续消除6次的赢钱数:::::
	N85 int //Freespin连续消除7次的次数:::::
	N86 int //Freespin连续消除7次的赢钱数:::::
	N87 int //Freespin连续消除8次的次数:::::
	N88 int //Freespin连续消除8次的赢钱数:::::
	N89 int //Freespin连续消除9次的次数:::::
	N90 int //Freespin连续消除9次的赢钱数:::::
	N91 int //Freespin连续消除10次的次数:::::
	N92 int //Freespin连续消除10次的赢钱数:::::
	N93 int //Freespin连续消除10+次的次数:::::
	N94 int //Freespin连续消除10+次的赢钱数:::::
	N95 int //3个scatter触发Freespin的次数:::::
	N96 int //3个scatter触发Freespin的总赢钱数:::::

}

func TestSlot(slotTest *business.SlotTests, run public.RunSlotTest, opts ...component.Option) error {

	normalResult := unit5TestResult{}
	ch, errCh, _ := helper.Parallel[slot.Machine](run.Num, 1000, func() (slot.Machine, error) {
		m, err := slot.Play(run.SlotId, run.Amount, opts...)
		return m, err
	})
	defer public.EndProcessing(slotTest, &normalResult.BasicData)
	count := 0
	a := 0
	for {
		a++
		select {
		case err := <-errCh:
			slotTest.Detail = err.Error()
			return err
		case v, beforeClosed := <-ch:
			count++
			if run.Num >= 10000 && count%(run.Num/10) == 0 && count < run.Num-1 {
				go func() {
					slotTest.Detail = fmt.Sprintf("进度:%d/%d", count, run.Num)
					global.GVA_DB.Save(&slotTest)
				}()
			}
			if !beforeClosed {
				goto end
			}
			normalResult.Record(v)
		}
	}
end:
	slotTest.Detail = `最高倍率:::: ` + fmt.Sprintf("%g", normalResult.N0) + `
总赢钱:::: ` + strconv.Itoa(normalResult.N1) + `
总押注消耗钱:::::` + strconv.Itoa(normalResult.N2) + `
总转动次数:::::` + strconv.Itoa(normalResult.N3) + `
总赢钱次数:::: ` + strconv.Itoa(normalResult.N4) + `
普通转划线赢钱0-5倍的次数:::: ` + strconv.Itoa(normalResult.N5) + `
普通转划线赢钱0-5倍的总钱数:::: ` + strconv.Itoa(normalResult.N6) + `
普通转划线赢钱5-10倍的次数:::: ` + strconv.Itoa(normalResult.N7) + `
普通转划线赢钱5-10倍的总钱数:::: ` + strconv.Itoa(normalResult.N8) + `
普通转划线赢钱10-20倍的次数:::: ` + strconv.Itoa(normalResult.N9) + `
普通转划线赢钱10-20倍的总钱数:::: ` + strconv.Itoa(normalResult.N10) + `
普通转划线赢钱20-30倍的次数:::: ` + strconv.Itoa(normalResult.N11) + `
普通转划线赢钱20-30倍的总钱数:::: ` + strconv.Itoa(normalResult.N12) + `
普通转划线赢钱30-40倍的次数:::: ` + strconv.Itoa(normalResult.N13) + `
普通转划线赢钱30-40倍的总钱数:::: ` + strconv.Itoa(normalResult.N14) + `
普通转划线赢钱40-50倍的次数:::: ` + strconv.Itoa(normalResult.N15) + `
普通转划线赢钱40-50倍的总钱数:::: ` + strconv.Itoa(normalResult.N16) + `
普通转划线赢钱50-100倍的次数:::: ` + strconv.Itoa(normalResult.N17) + `
普通转划线赢钱50-100倍的总钱数:::: ` + strconv.Itoa(normalResult.N18) + `
普通转划线赢钱100-200倍的次数:::: ` + strconv.Itoa(normalResult.N19) + `
普通转划线赢钱100-200倍的总钱数:::: ` + strconv.Itoa(normalResult.N20) + `
普通转划线赢钱200-500倍的次数:::: ` + strconv.Itoa(normalResult.N21) + `
普通转划线赢钱200-500倍的总钱数:::: ` + strconv.Itoa(normalResult.N22) + `
普通转划线赢钱500-1000倍的次数:::: ` + strconv.Itoa(normalResult.N23) + `
普通转划线赢钱500-1000倍的总钱数:::: ` + strconv.Itoa(normalResult.N24) + `
普通转划线赢钱1000以上倍的次数:::: ` + strconv.Itoa(normalResult.N25) + `
普通转划线赢钱1000以上倍的总钱数:::: ` + strconv.Itoa(normalResult.N26) + `
普通转连续消除1次的次数:::: ` + strconv.Itoa(normalResult.N27) + `
普通转连续消除1次的赢钱数:::: ` + strconv.Itoa(normalResult.N28) + `
普通转连续消除2次的次数:::: ` + strconv.Itoa(normalResult.N29) + `
普通转连续消除2次的赢钱数:::: ` + strconv.Itoa(normalResult.N30) + `
普通转连续消除3次的次数:::: ` + strconv.Itoa(normalResult.N31) + `
普通转连续消除3次的赢钱数:::: ` + strconv.Itoa(normalResult.N32) + `
普通转连续消除4次的次数:::: ` + strconv.Itoa(normalResult.N33) + `
普通转连续消除4次的赢钱数:::: ` + strconv.Itoa(normalResult.N34) + `
普通转连续消除5次的次数:::: ` + strconv.Itoa(normalResult.N35) + `
普通转连续消除5次的赢钱数:::: ` + strconv.Itoa(normalResult.N36) + `
普通转连续消除6次的次数:::: ` + strconv.Itoa(normalResult.N37) + `
普通转连续消除6次的赢钱数:::: ` + strconv.Itoa(normalResult.N38) + `
普通转连续消除7次的次数:::: ` + strconv.Itoa(normalResult.N39) + `
普通转连续消除7次的赢钱数:::: ` + strconv.Itoa(normalResult.N40) + `
普通转连续消除8次的次数:::: ` + strconv.Itoa(normalResult.N41) + `
普通转连续消除8次的赢钱数:::: ` + strconv.Itoa(normalResult.N42) + `
普通转连续消除9次的次数:::: ` + strconv.Itoa(normalResult.N43) + `
普通转连续消除9次的赢钱数:::: ` + strconv.Itoa(normalResult.N44) + `
普通转连续消除10次的次数:::: ` + strconv.Itoa(normalResult.N45) + `
普通转连续消除10次的赢钱数:::: ` + strconv.Itoa(normalResult.N46) + `
普通转连续消除10+次的次数:::: ` + strconv.Itoa(normalResult.N47) + `
普通转连续消除10+次的赢钱数:::: ` + strconv.Itoa(normalResult.N48) + `
FREESPIN初始触发的次数:::: ` + strconv.Itoa(normalResult.N49) + `
FREESPIN再次触发FREESPIN的次数:::: ` + strconv.Itoa(normalResult.N50) + `
Freespin划线赢钱0-5倍的次数:::: ` + strconv.Itoa(normalResult.N51) + `
Freespin划线赢钱0-5倍的总钱数:::: ` + strconv.Itoa(normalResult.N52) + `
Freespin划线赢钱5-10倍的次数:::: ` + strconv.Itoa(normalResult.N53) + `
Freespin划线赢钱5-10倍的总钱数:::: ` + strconv.Itoa(normalResult.N54) + `
Freespin划线赢钱10-20倍的次数:::: ` + strconv.Itoa(normalResult.N55) + `
Freespin划线赢钱10-20倍的总钱数:::: ` + strconv.Itoa(normalResult.N56) + `
Freespin划线赢钱20-30倍的次数:::: ` + strconv.Itoa(normalResult.N57) + `
Freespin划线赢钱20-30倍的总钱数:::: ` + strconv.Itoa(normalResult.N58) + `
Freespin划线赢钱30-40倍的次数:::: ` + strconv.Itoa(normalResult.N59) + `
Freespin划线赢钱30-40倍的总钱数:::: ` + strconv.Itoa(normalResult.N60) + `
Freespin划线赢钱40-50倍的次数:::: ` + strconv.Itoa(normalResult.N61) + `
Freespin划线赢钱40-50倍的总钱数:::: ` + strconv.Itoa(normalResult.N62) + `
Freespin划线赢钱50-100倍的次数:::: ` + strconv.Itoa(normalResult.N63) + `
Freespin划线赢钱50-100倍的总钱数:::: ` + strconv.Itoa(normalResult.N64) + `
Freespin划线赢钱100-200倍的次数:::: ` + strconv.Itoa(normalResult.N65) + `
Freespin划线赢钱100-200倍的总钱数:::: ` + strconv.Itoa(normalResult.N66) + `
Freespin划线赢钱200-500倍的次数:::: ` + strconv.Itoa(normalResult.N67) + `
Freespin划线赢钱200-500倍的总钱数:::: ` + strconv.Itoa(normalResult.N68) + `
Freespin划线赢钱500-1000倍的次数:::: ` + strconv.Itoa(normalResult.N69) + `
Freespin划线赢钱500-1000倍的总钱数:::: ` + strconv.Itoa(normalResult.N70) + `
Freespin划线赢钱1000以上倍的次数:::: ` + strconv.Itoa(normalResult.N71) + `
Freespin划线赢钱1000以上倍的总钱数:::: ` + strconv.Itoa(normalResult.N72) + `
Freespin连续消除1次的次数:::: ` + strconv.Itoa(normalResult.N73) + `
Freespin连续消除1次的赢钱数:::: ` + strconv.Itoa(normalResult.N74) + `
Freespin连续消除2次的次数:::: ` + strconv.Itoa(normalResult.N75) + `
Freespin连续消除2次的赢钱数:::: ` + strconv.Itoa(normalResult.N76) + `
Freespin连续消除3次的次数:::: ` + strconv.Itoa(normalResult.N77) + `
Freespin连续消除3次的赢钱数:::: ` + strconv.Itoa(normalResult.N78) + `
Freespin连续消除4次的次数:::: ` + strconv.Itoa(normalResult.N79) + `
Freespin连续消除4次的赢钱数:::: ` + strconv.Itoa(normalResult.N80) + `
Freespin连续消除5次的次数:::: ` + strconv.Itoa(normalResult.N81) + `
Freespin连续消除5次的赢钱数:::: ` + strconv.Itoa(normalResult.N82) + `
Freespin连续消除6次的次数:::: ` + strconv.Itoa(normalResult.N83) + `
Freespin连续消除6次的赢钱数:::: ` + strconv.Itoa(normalResult.N84) + `
Freespin连续消除7次的次数:::: ` + strconv.Itoa(normalResult.N85) + `
Freespin连续消除7次的赢钱数:::: ` + strconv.Itoa(normalResult.N86) + `
Freespin连续消除8次的次数:::: ` + strconv.Itoa(normalResult.N87) + `
Freespin连续消除8次的赢钱数:::: ` + strconv.Itoa(normalResult.N88) + `
Freespin连续消除9次的次数:::: ` + strconv.Itoa(normalResult.N89) + `
Freespin连续消除9次的赢钱数:::: ` + strconv.Itoa(normalResult.N90) + `
Freespin连续消除10次的次数:::: ` + strconv.Itoa(normalResult.N91) + `
Freespin连续消除10次的赢钱数:::: ` + strconv.Itoa(normalResult.N92) + `
Freespin连续消除10+次的次数:::: ` + strconv.Itoa(normalResult.N93) + `
Freespin连续消除10+次的赢钱数:::: ` + strconv.Itoa(normalResult.N94) + `
3个scatter触发Freespin的次数::::: ` + strconv.Itoa(normalResult.N95) + `
3个scatter触发Freespin的总赢钱数::::: ` + strconv.Itoa(normalResult.N96)
	return nil
}

func (r *unit5TestResult) Record(machine slot.Machine) {
	spins := machine.GetSpins()
	spin := machine.GetSpin()

	sumGain := lo.SumBy(spins, func(s *component.Spin) int {
		return s.Gain
	}) + spin.Gain
	r.N1 += sumGain
	r.N2 += spin.Bet
	r.N3++
	r.N3 += len(spins)
	if sumGain > 0 {
		r.N4++
	}
	Multiple := float64(sumGain) / float64(spin.Bet)
	if Multiple > r.N0 {
		r.N0 = Multiple
	}
	if len(spin.Table.QueryTags("scatter")) == 3 {
		r.N95++
		r.N96 += sumGain - spin.Gain
	}

	//普通转赢钱倍数
	{
		if Multiple >= 0 && Multiple < 5 {
			r.N5++
			r.N6 += sumGain
		} else if Multiple >= 5 && Multiple < 10 {
			r.N7++
			r.N8 += sumGain
		} else if Multiple >= 10 && Multiple < 20 {
			r.N9++
			r.N10 += sumGain
		} else if Multiple >= 20 && Multiple < 30 {
			r.N11++
			r.N12 += sumGain
		} else if Multiple >= 30 && Multiple < 40 {
			r.N13++
			r.N14 += sumGain
		} else if Multiple >= 40 && Multiple < 50 {
			r.N15++
			r.N16 += sumGain
		} else if Multiple >= 50 && Multiple < 100 {
			r.N17++
			r.N18 += sumGain
		} else if Multiple >= 100 && Multiple < 200 {
			r.N19++
			r.N20 += sumGain
		} else if Multiple >= 200 && Multiple < 500 {
			r.N21++
			r.N22 += sumGain
		} else if Multiple >= 500 && Multiple < 1000 {
			r.N23++
			r.N24 += sumGain
		} else if Multiple >= 1000 {
			r.N25++
			r.N26 += sumGain
		}
	}

	//普通转消除次数和赢钱
	{
		switch len(spin.Table.AlterFlows) {
		case 2:
			r.N27++
			r.N28 += spin.Gain
		case 3:
			r.N29++
			r.N30 += spin.Gain
		case 4:
			r.N31++
			r.N32 += spin.Gain
		case 5:
			r.N33++
			r.N34 += spin.Gain
		case 6:
			r.N35++
			r.N36 += spin.Gain
		case 7:
			r.N37++
			r.N38 += spin.Gain
		case 8:
			r.N39++
			r.N40 += spin.Gain
		case 9:
			r.N41++
			r.N42 += spin.Gain
		case 10:
			r.N43++
			r.N44 += spin.Gain
		case 11:
			r.N45++
			r.N46 += spin.Gain
		default:
			r.N47++
			r.N48 += spin.Gain

		}
	}

	//Free赢钱倍数
	{
		if len(spins) == 0 {
			return
		}
		if spin.FreeSpinParams.FreeNum > 0 {
			r.N49++
		}
		frees := lo.Filter(spins, func(s *component.Spin, index int) bool {
			return s.ParentId == 0 && s.IsFree
		})

		SumFreeGain := lo.SumBy(frees, func(s *component.Spin) int {
			return s.Gain
		})

		multipleFree := float64(SumFreeGain) / float64(spin.Bet)
		if multipleFree >= 0 && multipleFree < 5 {
			r.N51++
			r.N52 += SumFreeGain
		} else if multipleFree >= 5 && multipleFree < 10 {
			r.N53++
			r.N54 += SumFreeGain
		} else if multipleFree >= 10 && multipleFree < 20 {
			r.N55++
			r.N56 += SumFreeGain
		} else if multipleFree >= 20 && multipleFree < 30 {
			r.N57++
			r.N58 += SumFreeGain
		} else if multipleFree >= 30 && multipleFree < 40 {
			r.N59++
			r.N60 += SumFreeGain
		} else if multipleFree >= 40 && multipleFree < 50 {
			r.N61++
			r.N62 += SumFreeGain
		} else if multipleFree >= 50 && multipleFree < 100 {
			r.N63++
			r.N64 += SumFreeGain
		} else if multipleFree >= 100 && multipleFree < 200 {
			r.N65++
			r.N66 += SumFreeGain
		} else if multipleFree >= 200 && multipleFree < 500 {
			r.N67++
			r.N68 += SumFreeGain
		} else if multipleFree >= 500 && multipleFree < 1000 {
			r.N69++
			r.N70 += SumFreeGain
		} else if multipleFree >= 1000 {
			r.N71++
			r.N72 += SumFreeGain
		}

		freeParents := lo.Filter(spins, func(s *component.Spin, index int) bool {
			return s.FreeSpinParams.FreeNum > 0
		})
		sumFreeSumGain := 0
		for _, parent := range freeParents {
			r.N50++
			sons := lo.Filter(spins, func(s *component.Spin, index int) bool {
				return s.ParentId == parent.Id
			})

			freeSumGain := lo.SumBy(sons, func(s *component.Spin) int {
				return s.Gain
			})
			sumFreeSumGain += freeSumGain
			freeMultiple := float64(freeSumGain) / float64(spin.Bet)
			if freeMultiple >= 0 && freeMultiple < 5 {
				r.N51++
				r.N52 += freeSumGain
			} else if freeMultiple >= 5 && freeMultiple < 10 {
				r.N53++
				r.N54 += freeSumGain
			} else if freeMultiple >= 10 && freeMultiple < 20 {
				r.N55++
				r.N56 += freeSumGain
			} else if freeMultiple >= 20 && freeMultiple < 30 {
				r.N57++
				r.N58 += freeSumGain
			} else if freeMultiple >= 30 && freeMultiple < 40 {
				r.N59++
				r.N60 += freeSumGain
			} else if freeMultiple >= 40 && freeMultiple < 50 {
				r.N61++
				r.N62 += freeSumGain
			} else if freeMultiple >= 50 && freeMultiple < 100 {
				r.N63++
				r.N64 += freeSumGain
			} else if freeMultiple >= 100 && freeMultiple < 200 {
				r.N65++
				r.N66 += freeSumGain
			} else if freeMultiple >= 200 && freeMultiple < 500 {
				r.N67++
				r.N68 += freeSumGain
			} else if freeMultiple >= 500 && freeMultiple < 1000 {
				r.N69++
				r.N70 += freeSumGain
			} else if freeMultiple >= 1000 {
				r.N71++
				r.N72 += freeSumGain
			}
		}

		if sumGain-spin.Gain != sumFreeSumGain+SumFreeGain {
			global.GVA_LOG.Info("sumGain - spin.Gain!= sumFreeSumGain+ SumFreeGain", zap.Any("sumGain", sumGain), zap.Any("spin.Gain", spin.Gain), zap.Any("sumFreeSumGain", sumFreeSumGain), zap.Any("SumFreeGain", SumFreeGain))
		}

		for _, c := range spins {
			switch len(c.Table.AlterFlows) {
			case 2:
				r.N73++
				r.N74 += c.Gain
			case 3:
				r.N75++
				r.N76 += c.Gain
			case 4:
				r.N77++
				r.N78 += c.Gain
			case 5:
				r.N79++
				r.N80 += c.Gain
			case 6:
				r.N81++
				r.N82 += c.Gain
			case 7:
				r.N83++
				r.N84 += c.Gain
			case 8:
				r.N85++
				r.N86 += c.Gain
			case 9:
				r.N87++
				r.N88 += c.Gain
			case 10:
				r.N89++
				r.N90 += c.Gain
			case 11:
				r.N91++
				r.N92 += c.Gain
			default:
				r.N93++
				r.N94 += c.Gain

			}
		}

	}

}
