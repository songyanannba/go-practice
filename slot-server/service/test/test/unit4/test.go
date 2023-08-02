package unit4

import (
	"fmt"
	"github.com/samber/lo"
	"slot-server/global"
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

type unit4TestResult struct {
	public.BasicData
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

func TestSlot(slotTest *business.SlotTests, run public.RunSlotTest, opts ...component.Option) error {
	normalResult := unit4TestResult{}
	normalResult.coinMultiplierMap = map[float64]int{}
	normalResult.coin6SumCoin = map[int]int{}
	UserSpin := &business.SlotUserSpin{
		UserId: 0,
		SlotId: run.SlotId,
	}

	defer public.EndProcessing(slotTest, &normalResult.BasicData)
	ch, errCh, _ := helper.Parallel[*slotHandle.MergeSpin](run.Num, 1000, func() (*slotHandle.MergeSpin, error) {

		m, _ := slot.Play(run.SlotId, run.Amount, opts...)
		spin := m.GetSpin()
		mergeSpin, err := slotHandle.RunMergeSpin(0, spin, UserSpin)
		if err != nil {
			slotTest.Detail = err.Error()
			return nil, err
		}
		return mergeSpin, nil
	})
	count := 0
	for {
		select {
		case err := <-errCh:
			slotTest.Detail = err.Error()
			return err
		case v, beforeClosed := <-ch:
			count++
			if run.Num > 10000 && count%(run.Num/100) == 0 && count < run.Num-1 {
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
	slotTest.Bet = run.Amount
	if true {
		slotTest.Detail =
			`总赢钱:::: ` + strconv.Itoa(normalResult.N1) + `
总押注消耗钱::::: ` + strconv.Itoa(normalResult.N2) + `
总转动次数::::: ` + strconv.Itoa(normalResult.N3) + `
总赢钱次数::::: ` + strconv.Itoa(normalResult.N4) + `
普通转划线赢钱0-5倍的次数::::: ` + strconv.Itoa(normalResult.N5) + `
普通转划线赢钱0-5倍的总钱数::::: ` + strconv.Itoa(normalResult.N6) + `
普通转划线赢钱5-10倍的次数::::: ` + strconv.Itoa(normalResult.N7) + `
普通转划线赢钱5-10倍的总钱数::::: ` + strconv.Itoa(normalResult.N8) + `
普通转划线赢钱10-20倍的次数::::: ` + strconv.Itoa(normalResult.N9) + `
普通转划线赢钱10-20倍的总钱数::::: ` + strconv.Itoa(normalResult.N10) + `
普通转划线赢钱20-30倍的次数::::: ` + strconv.Itoa(normalResult.N11) + `
普通转划线赢钱20-30倍的总钱数::::: ` + strconv.Itoa(normalResult.N12) + `
普通转划线赢钱30-40倍的次数::::: ` + strconv.Itoa(normalResult.N13) + `
普通转划线赢钱30-40倍的总钱数::::: ` + strconv.Itoa(normalResult.N14) + `
普通转划线赢钱40-50倍的次数::::: ` + strconv.Itoa(normalResult.N15) + `
普通转划线赢钱40-50倍的总钱数::::: ` + strconv.Itoa(normalResult.N16) + `
普通转划线赢钱50-100倍的次数::::: ` + strconv.Itoa(normalResult.N17) + `
普通转划线赢钱50-100倍的总钱数::::: ` + strconv.Itoa(normalResult.N18) + `
普通转划线赢钱100-200倍的次数::::: ` + strconv.Itoa(normalResult.N19) + `
普通转划线赢钱100-200倍的总钱数::::: ` + strconv.Itoa(normalResult.N20) + `
普通转划线赢钱200-500倍的次数::::: ` + strconv.Itoa(normalResult.N21) + `
普通转划线赢钱200-500倍的总钱数::::: ` + strconv.Itoa(normalResult.N22) + `
普通转划线赢钱500-1000倍的次数::::: ` + strconv.Itoa(normalResult.N23) + `
普通转划线赢钱500-1000倍的总钱数::::: ` + strconv.Itoa(normalResult.N24) + `
普通转划线赢钱1000以上倍的次数::::: ` + strconv.Itoa(normalResult.N25) + `
普通转划线赢钱1000以上倍的总钱数::::: ` + strconv.Itoa(normalResult.N26) + `
普通转触发的RESPIN触发次数::::: ` + strconv.Itoa(normalResult.N27) + `
普通转触发的RESPIN划线赢钱赢钱0-5倍的次数::::: ` + strconv.Itoa(normalResult.N28) + `
普通转触发的RESPIN划线赢钱赢钱0-5倍的总钱数::::: ` + strconv.Itoa(normalResult.N29) + `
普通转触发的RESPIN划线赢钱赢钱5-10倍的次数::::: ` + strconv.Itoa(normalResult.N30) + `
普通转触发的RESPIN划线赢钱赢钱5-10倍的总钱数::::: ` + strconv.Itoa(normalResult.N31) + `
普通转触发的RESPIN划线赢钱赢钱10-20倍的次数::::: ` + strconv.Itoa(normalResult.N32) + `
普通转触发的RESPIN划线赢钱赢钱10-20倍的总钱数::::: ` + strconv.Itoa(normalResult.N33) + `
普通转触发的RESPIN划线赢钱赢钱20-30倍的次数::::: ` + strconv.Itoa(normalResult.N34) + `
普通转触发的RESPIN划线赢钱赢钱20-30倍的总钱数::::: ` + strconv.Itoa(normalResult.N35) + `
普通转触发的RESPIN划线赢钱赢钱30-40倍的次数::::: ` + strconv.Itoa(normalResult.N36) + `
普通转触发的RESPIN划线赢钱赢钱30-40倍的总钱数::::: ` + strconv.Itoa(normalResult.N37) + `
普通转触发的RESPIN划线赢钱赢钱40-50倍的次数::::: ` + strconv.Itoa(normalResult.N38) + `
普通转触发的RESPIN划线赢钱赢钱40-50倍的总钱数::::: ` + strconv.Itoa(normalResult.N39) + `
普通转触发的RESPIN划线赢钱赢钱50-100倍的次数::::: ` + strconv.Itoa(normalResult.N40) + `
普通转触发的RESPIN划线赢钱赢钱50-100倍的总钱数::::: ` + strconv.Itoa(normalResult.N41) + `
普通转触发的RESPIN划线赢钱赢钱100-200倍的次数::::: ` + strconv.Itoa(normalResult.N42) + `
普通转触发的RESPIN划线赢钱赢钱100-200倍的总钱数::::: ` + strconv.Itoa(normalResult.N43) + `
普通转触发的RESPIN划线赢钱赢钱200-500倍的次数::::: ` + strconv.Itoa(normalResult.N44) + `
普通转触发的RESPIN划线赢钱赢钱200-500倍的总钱数::::: ` + strconv.Itoa(normalResult.N45) + `
普通转触发的RESPIN划线赢钱赢钱500-1000倍的次数::::: ` + strconv.Itoa(normalResult.N46) + `
普通转触发的RESPIN划线赢钱赢钱500-1000倍的总钱数::::: ` + strconv.Itoa(normalResult.N47) + `
普通转触发的RESPIN划线赢钱赢钱1000以上倍的次数::::: ` + strconv.Itoa(normalResult.N48) + `
普通转触发的RESPIN划线赢钱赢钱1000以上倍的总钱数::::: ` + strconv.Itoa(normalResult.N49) + `
FREESPIN初始触发的次数::::: ` + strconv.Itoa(normalResult.N50) + `
FREESPIN再次触发FREESPIN的次数::::: ` + strconv.Itoa(normalResult.N51) + `
Freespin划线赢钱0-5倍的次数::::: ` + strconv.Itoa(normalResult.N52) + `
Freespin划线赢钱0-5倍的总钱数::::: ` + strconv.Itoa(normalResult.N53) + `
Freespin划线赢钱5-10倍的次数::::: ` + strconv.Itoa(normalResult.N54) + `
Freespin划线赢钱5-10倍的总钱数::::: ` + strconv.Itoa(normalResult.N55) + `
Freespin划线赢钱10-20倍的次数::::: ` + strconv.Itoa(normalResult.N56) + `
Freespin划线赢钱10-20倍的总钱数::::: ` + strconv.Itoa(normalResult.N57) + `
Freespin划线赢钱20-30倍的次数::::: ` + strconv.Itoa(normalResult.N58) + `
Freespin划线赢钱20-30倍的总钱数::::: ` + strconv.Itoa(normalResult.N59) + `
Freespin划线赢钱30-40倍的次数::::: ` + strconv.Itoa(normalResult.N60) + `
Freespin划线赢钱30-40倍的总钱数::::: ` + strconv.Itoa(normalResult.N61) + `
Freespin划线赢钱40-50倍的次数::::: ` + strconv.Itoa(normalResult.N62) + `
Freespin划线赢钱40-50倍的总钱数::::: ` + strconv.Itoa(normalResult.N63) + `
Freespin划线赢钱50-100倍的次数::::: ` + strconv.Itoa(normalResult.N64) + `
Freespin划线赢钱50-100倍的总钱数::::: ` + strconv.Itoa(normalResult.N65) + `
Freespin划线赢钱100-200倍的次数::::: ` + strconv.Itoa(normalResult.N66) + `
Freespin划线赢钱100-200倍的总钱数::::: ` + strconv.Itoa(normalResult.N67) + `
Freespin划线赢钱200-500倍的次数::::: ` + strconv.Itoa(normalResult.N68) + `
Freespin划线赢钱200-500倍的总钱数::::: ` + strconv.Itoa(normalResult.N69) + `
Freespin划线赢钱500-1000倍的次数::::: ` + strconv.Itoa(normalResult.N70) + `
Freespin划线赢钱500-1000倍的总钱数::::: ` + strconv.Itoa(normalResult.N71) + `
Freespin划线赢钱1000以上倍的次数::::: ` + strconv.Itoa(normalResult.N72) + `
Freespin划线赢钱1000以上倍的总钱数::::: ` + strconv.Itoa(normalResult.N73) + `
FREESPIN触发RESPIN的次数::::: ` + strconv.Itoa(normalResult.N74) + `
FREESPIN触发的RESPIN划线赢钱0-5倍的次数::::: ` + strconv.Itoa(normalResult.N75) + `
FREESPIN触发的RESPIN划线赢钱0-5倍的总钱数::::: ` + strconv.Itoa(normalResult.N76) + `
FREESPIN触发的RESPIN划线赢钱5-10倍的次数::::: ` + strconv.Itoa(normalResult.N77) + `
FREESPIN触发的RESPIN划线赢钱5-10倍的总钱数::::: ` + strconv.Itoa(normalResult.N78) + `
FREESPIN触发的RESPIN划线赢钱10-20倍的次数::::: ` + strconv.Itoa(normalResult.N79) + `
FREESPIN触发的RESPIN划线赢钱10-20倍的总钱数::::: ` + strconv.Itoa(normalResult.N80) + `
FREESPIN触发的RESPIN划线赢钱20-30倍的次数::::: ` + strconv.Itoa(normalResult.N81) + `
FREESPIN触发的RESPIN划线赢钱20-30倍的总钱数::::: ` + strconv.Itoa(normalResult.N82) + `
FREESPIN触发的RESPIN划线赢钱30-40倍的次数::::: ` + strconv.Itoa(normalResult.N83) + `
FREESPIN触发的RESPIN划线赢钱30-40倍的总钱数::::: ` + strconv.Itoa(normalResult.N84) + `
FREESPIN触发的RESPIN划线赢钱40-50倍的次数::::: ` + strconv.Itoa(normalResult.N85) + `
FREESPIN触发的RESPIN划线赢钱40-50倍的总钱数::::: ` + strconv.Itoa(normalResult.N86) + `
FREESPIN触发的RESPIN划线赢钱50-100倍的次数::::: ` + strconv.Itoa(normalResult.N87) + `
FREESPIN触发的RESPIN划线赢钱50-100倍的总钱数::::: ` + strconv.Itoa(normalResult.N88) + `
FREESPIN触发的RESPIN划线赢钱100-200倍的次数::::: ` + strconv.Itoa(normalResult.N89) + `
FREESPIN触发的RESPIN划线赢钱100-200倍的总钱数::::: ` + strconv.Itoa(normalResult.N90) + `
FREESPIN触发的RESPIN划线赢钱200-500倍的次数::::: ` + strconv.Itoa(normalResult.N91) + `
FREESPIN触发的RESPIN划线赢钱200-500倍的总钱数::::: ` + strconv.Itoa(normalResult.N92) + `
FREESPIN触发的RESPIN划线赢钱500-1000倍的次数::::: ` + strconv.Itoa(normalResult.N93) + `
FREESPIN触发的RESPIN划线赢钱500-1000倍的总钱数::::: ` + strconv.Itoa(normalResult.N94) + `
FREESPIN触发的RESPIN划线赢钱1000以上倍的次数::::: ` + strconv.Itoa(normalResult.N95) + `
FREESPIN触发的RESPIN划线赢钱1000以上倍的总钱数::::: ` + strconv.Itoa(normalResult.N96) + `
单独普通转赢钱金额::::: ` + strconv.Itoa(normalResult.N97) + `
初始进入Link_coin标签个数6的次数::::: ` + strconv.Itoa(normalResult.N98) + `
初始进入Link_coin标签个数6的总钱数::::: ` + strconv.Itoa(normalResult.N99) + `
初始进入Link_coin标签个数7的次数::::: ` + strconv.Itoa(normalResult.N100) + `
初始进入Link_coin标签个数7的总钱数::::: ` + strconv.Itoa(normalResult.N101) + `
初始进入Link_coin标签个数8的次数::::: ` + strconv.Itoa(normalResult.N102) + `
初始进入Link_coin标签个数8的总钱数::::: ` + strconv.Itoa(normalResult.N103) + `
初始进入Link_coin标签个数9的次数::::: ` + strconv.Itoa(normalResult.N104) + `
初始进入Link_coin标签个数9的总钱数::::: ` + strconv.Itoa(normalResult.N105) + `
初始进入Link_coin标签个数10的次数::::: ` + strconv.Itoa(normalResult.N106) + `
初始进入Link_coin标签个数10的总钱数::::: ` + strconv.Itoa(normalResult.N107) + `
初始进入Link_coin标签个数11的次数::::: ` + strconv.Itoa(normalResult.N108) + `
初始进入Link_coin标签个数11的总钱数::::: ` + strconv.Itoa(normalResult.N109) + `
初始进入Link_coin标签个数12的次数::::: ` + strconv.Itoa(normalResult.N110) + `
初始进入Link_coin标签个数12的总钱数::::: ` + strconv.Itoa(normalResult.N111) + `
初始进入Link_coin标签个数13的次数::::: ` + strconv.Itoa(normalResult.N112) + `
初始进入Link_coin标签个数13的总钱数::::: ` + strconv.Itoa(normalResult.N113) + `
初始进入Link_coin标签个数14的次数::::: ` + strconv.Itoa(normalResult.N114) + `
初始进入Link_coin标签个数14的总钱数::::: ` + strconv.Itoa(normalResult.N115) + `
初始进入Link_coin标签个数15的次数::::: ` + strconv.Itoa(normalResult.N116) + `
初始进入Link_coin标签个数15的总钱数::::: ` + strconv.Itoa(normalResult.N117) + "\r\n"
		var coinMultipliers []float64
		for f, _ := range normalResult.coinMultiplierMap {
			coinMultipliers = append(coinMultipliers, f)
		}
		sort.Slice(coinMultipliers, func(i, j int) bool {
			return coinMultipliers[i] < coinMultipliers[j]
		})
		for _, multiplier := range coinMultipliers {
			slotTest.Detail += fmt.Sprintf("Link_coin标签倍率 %g 的次数::::: %d\r\n", multiplier, normalResult.coinMultiplierMap[multiplier])
		}

		var coin6SumCoin []int
		for f, _ := range normalResult.coin6SumCoin {
			coin6SumCoin = append(coin6SumCoin, f)
		}
		sort.Slice(coin6SumCoin, func(i, j int) bool {
			return coin6SumCoin[i] < coin6SumCoin[j]
		})
		for _, multiplier := range coin6SumCoin {
			slotTest.Detail += fmt.Sprintf("Link_coin初始标签数量6总标签数%d的次数::::: %d\r\n", multiplier, normalResult.coin6SumCoin[multiplier])
		}
	}
	return nil
}

func (r *unit4TestResult) Record(m *slotHandle.MergeSpin) {

	sumGain := lo.SumBy(m.Spins, func(item *component.Spin) int {
		return item.Gain
	}) + m.First.Gain
	r.N97 += m.First.Gain
	r.N1 += sumGain
	r.N2 += m.First.Bet
	r.N3 += len(m.Spins) + 1
	if m.First.FreeSpinParams.ReNum > 0 {
		r.N27++
	}
	if sumGain > 0 {
		r.N4++
	}
	Multiple := float64(sumGain) / float64(m.First.Bet)
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
	likeList := base.GetSpecialTags(m.First.InitDataList, "link_coin", "link_collect")
	firstRespins := lo.Filter(m.Spins, func(item *component.Spin, index int) bool {
		return item.TriggerMode == 0 && item.Gain > 0 && item.IsReSpin
	})
	firstRespinSum := lo.SumBy(firstRespins, func(item *component.Spin) int {
		return item.Gain
	})
	switch len(likeList) {
	case 6:
		r.N98++
		r.N99 += firstRespinSum
		coin6SumNum := base.GetSpecialTags(firstRespins[0].InitDataList, "link_coin", "link_collect")
		r.coin6SumCoin[len(coin6SumNum)]++
	case 7:
		r.N100++
		r.N101 += firstRespinSum
	case 8:
		r.N102++
		r.N103 += firstRespinSum
	case 9:
		r.N104++
		r.N105 += firstRespinSum
	case 10:
		r.N106++
		r.N107 += firstRespinSum
	case 11:
		r.N108++
		r.N109 += firstRespinSum
	case 12:
		r.N110++
		r.N111 += firstRespinSum
	case 13:
		r.N112++
		r.N113 += firstRespinSum
	case 14:
		r.N114++
		r.N115 += firstRespinSum
	case 15:
		r.N116++
		r.N117 += firstRespinSum + m.First.Gain
	}
	OrdinaryRes := lo.Filter(m.Spins, func(item *component.Spin, index int) bool {
		return item.TriggerMode == 0 && item.Gain > 0 && item.IsReSpin
	})
	if m.First.Jackpot != nil {
		r.N48++
		r.N49 += m.First.Gain
	}
	for _, re := range OrdinaryRes {
		Multiple = float64(re.Gain) / float64(m.First.Bet)
		if Multiple >= 0 && Multiple < 5 {
			r.N28++
			r.N29 += re.Gain
		} else if Multiple >= 5 && Multiple < 10 {
			r.N30++
			r.N31 += re.Gain
		} else if Multiple >= 10 && Multiple < 20 {
			r.N32++
			r.N33 += re.Gain
		} else if Multiple >= 20 && Multiple < 30 {
			r.N34++
			r.N35 += re.Gain
		} else if Multiple >= 30 && Multiple < 40 {
			r.N36++
			r.N37 += re.Gain
		} else if Multiple >= 40 && Multiple < 50 {
			r.N38++
			r.N39 += re.Gain
		} else if Multiple >= 50 && Multiple < 100 {
			r.N40++
			r.N41 += re.Gain
		} else if Multiple >= 100 && Multiple < 200 {
			r.N42++
			r.N43 += re.Gain
		} else if Multiple >= 200 && Multiple < 500 {
			r.N44++
			r.N45 += re.Gain
		} else if Multiple >= 500 && Multiple < 1000 {
			r.N46++
			r.N47 += re.Gain
		} else if Multiple >= 1000 {
			r.N48++
			r.N49 += re.Gain
		}
	}

	if m.First.FreeSpinParams.FreeNum > 0 {
		r.N50++
	}
	FreeFrees := lo.Filter(m.Spins, func(item *component.Spin, index int) bool {
		return item.IsFree && item.FreeSpinParams.FreeNum > 0
	})
	r.N51 += len(FreeFrees)
	frees := lo.Filter(m.Spins, func(item *component.Spin, index int) bool {
		return item.IsFree
	})
	for _, free := range frees {
		if free.Jackpot != nil {
			r.N95++
			r.N96 += free.Gain
		}
		if free.FreeSpinParams.ReNum > 0 {
			r.N74++
		}
		Multiple = float64(free.Gain) / float64(m.First.Bet)
		if Multiple >= 0 && Multiple < 5 {
			r.N52++
			r.N53 += free.Gain
		} else if Multiple >= 5 && Multiple < 10 {
			r.N54++
			r.N55 += free.Gain

		} else if Multiple >= 10 && Multiple < 20 {
			r.N56++
			r.N57 += free.Gain
		} else if Multiple >= 20 && Multiple < 30 {
			r.N58++
			r.N59 += free.Gain
		} else if Multiple >= 30 && Multiple < 40 {
			r.N60++
			r.N61 += free.Gain
		} else if Multiple >= 40 && Multiple < 50 {
			r.N62++
			r.N63 += free.Gain
		} else if Multiple >= 50 && Multiple < 100 {
			r.N64++
			r.N65 += free.Gain
		} else if Multiple >= 100 && Multiple < 200 {
			r.N66++
			r.N67 += free.Gain
		} else if Multiple >= 200 && Multiple < 500 {
			r.N68++
			r.N69 += free.Gain
		} else if Multiple >= 500 && Multiple < 1000 {
			r.N70++
			r.N71 += free.Gain
		} else if Multiple >= 1000 {
			r.N72++
			r.N73 += free.Gain
		}
	}

	FreeSpinRes := lo.Filter(m.Spins, func(item *component.Spin, index int) bool {
		return item.TriggerMode == 1 && item.Gain > 0
	})
	for _, Fre := range FreeSpinRes {
		Multiple = float64(Fre.Gain) / float64(m.First.Bet)
		if Multiple >= 0 && Multiple < 5 {
			r.N75++
			r.N76 += Fre.Gain
		} else if Multiple >= 5 && Multiple < 10 {
			r.N77++
			r.N78 += Fre.Gain
		} else if Multiple >= 10 && Multiple < 20 {
			r.N79++
			r.N80 += Fre.Gain
		} else if Multiple >= 20 && Multiple < 30 {
			r.N81++
			r.N82 += Fre.Gain
		} else if Multiple >= 30 && Multiple < 40 {
			r.N83++
			r.N84 += Fre.Gain
		} else if Multiple >= 40 && Multiple < 50 {
			r.N85++
			r.N86 += Fre.Gain
		} else if Multiple >= 50 && Multiple < 100 {
			r.N87++
			r.N88 += Fre.Gain
		} else if Multiple >= 100 && Multiple < 200 {
			r.N89++
			r.N90 += Fre.Gain
		} else if Multiple >= 200 && Multiple < 500 {
			r.N91++
			r.N92 += Fre.Gain
		} else if Multiple >= 500 && Multiple < 1000 {
			r.N93++
			r.N94 += Fre.Gain
		} else if Multiple >= 1000 {
			r.N95++
			r.N96 += Fre.Gain
		}
	}

	AllRes := lo.Filter(m.Spins, func(item *component.Spin, index int) bool {
		return item.Gain > 0 && item.IsReSpin
	})
	for _, re := range AllRes {
		likeList := base.GetSpecialTags(re.InitDataList, "link_coin")
		for _, tag := range likeList {
			r.coinMultiplierMap[tag.Multiple]++
		}
	}
}
