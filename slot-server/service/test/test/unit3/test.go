package unit3

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"math"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/service/logic/gameHandle/slotHandle"
	"slot-server/service/slot"
	"slot-server/service/slot/component"
	"slot-server/service/test/public"
	"strconv"
)

type unit3TestResult struct {
	public.BasicData
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
}

func TestSlot(slotTest *business.SlotTests, slotId uint, Num int, Amount int, opts ...component.Option) error {
	normalResult := unit3TestResult{}
	fillResult := unit3TestResult{}
	UserSpin := &business.SlotUserSpin{
		UserId: 0,
		SlotId: slotId,
	}
	c, _ := component.GetSlotConfig(slotId, false)
	raise := decimal.NewFromInt(int64(Amount)).Mul(decimal.NewFromFloat(c.Raise)).IntPart()
	defer public.EndProcessing(slotTest, &normalResult.BasicData)
	for i := 0; i < Num; i++ {
		if i%10000 == 0 && i < Num-1 {
			go func() {
				slotTest.Detail = fmt.Sprintf("进度:%d/%d", i, Num)
				global.GVA_DB.Save(&slotTest)
			}()
		}
		m, _ := slot.Play(slotId, Amount, opts...)
		spin := m.GetSpin()
		mergeSpin, err := slotHandle.RunMergeSpin(0, spin, UserSpin)
		if err != nil {
			slotTest.Detail = err.Error()
			return err
		}

		normalResult.Record(mergeSpin)

		m1, _ := slot.Play(slotId, Amount, append(opts, component.WithRaise(raise))...)
		spin1 := m1.GetSpin()

		mergeSpin1, err := slotHandle.RunMergeSpin(0, spin1, UserSpin)
		if err != nil {
			slotTest.Detail = err.Error()
			return err
		}
		fillResult.Record(mergeSpin1)
	}
	slotTest.Bet = Amount
	slotTest.Raise = int(raise)

	if true {
		slotTest.Detail =
			`总次数::::: ` + strconv.Itoa(normalResult.N79) + `
总赢钱:::: ` + strconv.Itoa(normalResult.N1) + `
总押注消耗钱::::: ` + strconv.Itoa(normalResult.N2) + `
总转动次数::::: ` + strconv.Itoa(normalResult.N3) + `
总赢钱次数::::: ` + strconv.Itoa(normalResult.N4) + `
普通转划线赢钱0-5倍的次数::::: ` + strconv.Itoa(normalResult.N5) + `
普通转划线赢钱0-5倍的总钱数::::: ` + strconv.Itoa(normalResult.N6) + `
普通转划线赢钱5-10倍的次数::::: ` + strconv.Itoa(normalResult.N7) + `
普通转划线赢钱5-10倍的总钱数:::::  ` + strconv.Itoa(normalResult.N8) + `
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
普通转进度条为0时的赢钱次数::::: ` + strconv.Itoa(normalResult.N27) + `
普通转进度条为0时的赢钱数::::: ` + strconv.Itoa(normalResult.N28) + `
普通转进度条为1时的赢钱次数::::: ` + strconv.Itoa(normalResult.N29) + `
普通转进度条为1时的赢钱数::::: ` + strconv.Itoa(normalResult.N30) + `
普通转进度条为2时的赢钱次数::::: ` + strconv.Itoa(normalResult.N31) + `
普通转进度条为2时的赢钱数::::: ` + strconv.Itoa(normalResult.N32) + `
普通转进度条为3时的赢钱次数::::: ` + strconv.Itoa(normalResult.N33) + `
普通转进度条为3时的赢钱数::::: ` + strconv.Itoa(normalResult.N34) + `
普通转进度条为4时的赢钱次数::::: ` + strconv.Itoa(normalResult.N35) + `
普通转进度条为4时的赢钱数::::: ` + strconv.Itoa(normalResult.N36) + `
普通转进度条为5时的赢钱次数::::: ` + strconv.Itoa(normalResult.N37) + `
普通转进度条为5时的赢钱数::::: ` + strconv.Itoa(normalResult.N38) + `
普通转进度条为6时的赢钱次数::::: ` + strconv.Itoa(normalResult.N39) + `
普通转进度条为6时的赢钱数::::: ` + strconv.Itoa(normalResult.N40) + `
Freespin划线赢钱0-5倍的次数::::: ` + strconv.Itoa(normalResult.N41) + `
Freespin划线赢钱0-5倍的总钱数::::: ` + strconv.Itoa(normalResult.N42) + `
Freespin划线赢钱5-10倍的次数:::::  ` + strconv.Itoa(normalResult.N43) + `
Freespin划线赢钱5-10倍的总钱数::::: ` + strconv.Itoa(normalResult.N44) + `
Freespin划线赢钱10-20倍的次数::::: ` + strconv.Itoa(normalResult.N45) + `
Freespin划线赢钱10-20倍的总钱数::::: ` + strconv.Itoa(normalResult.N46) + `
Freespin划线赢钱20-30倍的次数::::: ` + strconv.Itoa(normalResult.N47) + `
Freespin划线赢钱20-30倍的总钱数::::: ` + strconv.Itoa(normalResult.N48) + `
Freespin划线赢钱30-40倍的次数::::: ` + strconv.Itoa(normalResult.N49) + `
Freespin划线赢钱30-40倍的总钱数::::: ` + strconv.Itoa(normalResult.N50) + `
Freespin划线赢钱40-50倍的次数::::: ` + strconv.Itoa(normalResult.N51) + `
Freespin划线赢钱40-50倍的总钱数::::: ` + strconv.Itoa(normalResult.N52) + `
Freespin划线赢钱50-100倍的次数::::: ` + strconv.Itoa(normalResult.N53) + `
Freespin划线赢钱50-100倍的总钱数::::: ` + strconv.Itoa(normalResult.N54) + `
Freespin划线赢钱100-200倍的次数::::: ` + strconv.Itoa(normalResult.N55) + `
Freespin划线赢钱100-200倍的总钱数::::: ` + strconv.Itoa(normalResult.N56) + ` 
Freespin划线赢钱200-500倍的次数::::: ` + strconv.Itoa(normalResult.N57) + `
Freespin划线赢钱200-500倍的总钱数::::: ` + strconv.Itoa(normalResult.N58) + `
Freespin划线赢钱500-1000倍的次数::::: ` + strconv.Itoa(normalResult.N59) + `
Freespin划线赢钱500-1000倍的总钱数::::: ` + strconv.Itoa(normalResult.N60) + `
Freespin划线赢钱1000以上倍的次数::::: ` + strconv.Itoa(normalResult.N61) + `
Freespin划线赢钱1000以上倍的总钱数::::: ` + strconv.Itoa(normalResult.N62) + `
Freespin进度条为0时的赢钱次数::::: ` + strconv.Itoa(normalResult.N63) + `
Freespin进度条为0时的赢钱数::::: ` + strconv.Itoa(normalResult.N64) + `
Freespin进度条为1时的赢钱次数::::: ` + strconv.Itoa(normalResult.N65) + `
Freespin进度条为1时的赢钱数::::: ` + strconv.Itoa(normalResult.N66) + `
Freespin进度条为2时的赢钱次数::::: ` + strconv.Itoa(normalResult.N67) + `
Freespin进度条为2时的赢钱数::::: ` + strconv.Itoa(normalResult.N68) + `
Freespin进度条为3时的赢钱次数::::: ` + strconv.Itoa(normalResult.N69) + `
Freespin进度条为3时的赢钱数::::: ` + strconv.Itoa(normalResult.N70) + `
Freespin进度条为4时的赢钱次数::::: ` + strconv.Itoa(normalResult.N71) + `
Freespin进度条为4时的赢钱数::::: ` + strconv.Itoa(normalResult.N72) + `
Freespin进度条为5时的赢钱次数::::: ` + strconv.Itoa(normalResult.N73) + `
Freespin进度条为5时的赢钱数::::: ` + strconv.Itoa(normalResult.N74) + `
Freespin进度条为6时的赢钱次数::::: ` + strconv.Itoa(normalResult.N75) + `
Freespin进度条为6时的赢钱数::::: ` + strconv.Itoa(normalResult.N76) + `
普通转触发Freespin玩法次数::::: ` + strconv.Itoa(normalResult.N77) + `
Freespin玩法总spin转动次数:::::` + strconv.Itoa(normalResult.N78) + `
平方和:::::` + fmt.Sprintf("%v", normalResult.N80) + `
加注玩法:
总次数::::: ` + strconv.Itoa(fillResult.N79) + `
总赢钱:::: ` + strconv.Itoa(fillResult.N1) + `
总押注消耗钱::::: ` + strconv.Itoa(fillResult.N2) + `
总转动次数::::: ` + strconv.Itoa(fillResult.N3) + `
总赢钱次数::::: ` + strconv.Itoa(fillResult.N4) + `
普通转划线赢钱0-5倍的次数::::: ` + strconv.Itoa(fillResult.N5) + `
普通转划线赢钱0-5倍的总钱数::::: ` + strconv.Itoa(fillResult.N6) + `
普通转划线赢钱5-10倍的次数::::: ` + strconv.Itoa(fillResult.N7) + `
普通转划线赢钱5-10倍的总钱数:::::  ` + strconv.Itoa(fillResult.N8) + `
普通转划线赢钱10-20倍的次数::::: ` + strconv.Itoa(fillResult.N9) + `
普通转划线赢钱10-20倍的总钱数::::: ` + strconv.Itoa(fillResult.N10) + `
普通转划线赢钱20-30倍的次数::::: ` + strconv.Itoa(fillResult.N11) + `
普通转划线赢钱20-30倍的总钱数::::: ` + strconv.Itoa(fillResult.N12) + `
普通转划线赢钱30-40倍的次数::::: ` + strconv.Itoa(fillResult.N13) + `
普通转划线赢钱30-40倍的总钱数::::: ` + strconv.Itoa(fillResult.N14) + `
普通转划线赢钱40-50倍的次数::::: ` + strconv.Itoa(fillResult.N15) + `
普通转划线赢钱40-50倍的总钱数::::: ` + strconv.Itoa(fillResult.N16) + `
普通转划线赢钱50-100倍的次数::::: ` + strconv.Itoa(fillResult.N17) + `
普通转划线赢钱50-100倍的总钱数::::: ` + strconv.Itoa(fillResult.N18) + `
普通转划线赢钱100-200倍的次数::::: ` + strconv.Itoa(fillResult.N19) + `
普通转划线赢钱100-200倍的总钱数::::: ` + strconv.Itoa(fillResult.N20) + `
普通转划线赢钱200-500倍的次数::::: ` + strconv.Itoa(fillResult.N21) + `
普通转划线赢钱200-500倍的总钱数::::: ` + strconv.Itoa(fillResult.N22) + `
普通转划线赢钱500-1000倍的次数::::: ` + strconv.Itoa(fillResult.N23) + `
普通转划线赢钱500-1000倍的总钱数::::: ` + strconv.Itoa(fillResult.N24) + `
普通转划线赢钱1000以上倍的次数::::: ` + strconv.Itoa(fillResult.N25) + `
普通转划线赢钱1000以上倍的总钱数::::: ` + strconv.Itoa(fillResult.N26) + `
普通转进度条为0时的赢钱次数::::: ` + strconv.Itoa(fillResult.N27) + `
普通转进度条为0时的赢钱数::::: ` + strconv.Itoa(fillResult.N28) + `
普通转进度条为1时的赢钱次数::::: ` + strconv.Itoa(fillResult.N29) + `
普通转进度条为1时的赢钱数::::: ` + strconv.Itoa(fillResult.N30) + `
普通转进度条为2时的赢钱次数::::: ` + strconv.Itoa(fillResult.N31) + `
普通转进度条为2时的赢钱数::::: ` + strconv.Itoa(fillResult.N32) + `
普通转进度条为3时的赢钱次数::::: ` + strconv.Itoa(fillResult.N33) + `
普通转进度条为3时的赢钱数::::: ` + strconv.Itoa(fillResult.N34) + `
普通转进度条为4时的赢钱次数::::: ` + strconv.Itoa(fillResult.N35) + `
普通转进度条为4时的赢钱数::::: ` + strconv.Itoa(fillResult.N36) + `
普通转进度条为5时的赢钱次数::::: ` + strconv.Itoa(fillResult.N37) + `
普通转进度条为5时的赢钱数::::: ` + strconv.Itoa(fillResult.N38) + `
普通转进度条为6时的赢钱次数::::: ` + strconv.Itoa(fillResult.N39) + `
普通转进度条为6时的赢钱数::::: ` + strconv.Itoa(fillResult.N40) + `
Freespin划线赢钱0-5倍的次数::::: ` + strconv.Itoa(fillResult.N41) + `
Freespin划线赢钱0-5倍的总钱数::::: ` + strconv.Itoa(fillResult.N42) + `
Freespin划线赢钱5-10倍的次数:::::  ` + strconv.Itoa(fillResult.N43) + `
Freespin划线赢钱5-10倍的总钱数::::: ` + strconv.Itoa(fillResult.N44) + `
Freespin划线赢钱10-20倍的次数::::: ` + strconv.Itoa(fillResult.N45) + `
Freespin划线赢钱10-20倍的总钱数::::: ` + strconv.Itoa(fillResult.N46) + `
Freespin划线赢钱20-30倍的次数::::: ` + strconv.Itoa(fillResult.N47) + `
Freespin划线赢钱20-30倍的总钱数::::: ` + strconv.Itoa(fillResult.N48) + `
Freespin划线赢钱30-40倍的次数::::: ` + strconv.Itoa(fillResult.N49) + `
Freespin划线赢钱30-40倍的总钱数::::: ` + strconv.Itoa(fillResult.N50) + `
Freespin划线赢钱40-50倍的次数::::: ` + strconv.Itoa(fillResult.N51) + `
Freespin划线赢钱40-50倍的总钱数::::: ` + strconv.Itoa(fillResult.N52) + `
Freespin划线赢钱50-100倍的次数::::: ` + strconv.Itoa(fillResult.N53) + `
Freespin划线赢钱50-100倍的总钱数::::: ` + strconv.Itoa(fillResult.N54) + `
Freespin划线赢钱100-200倍的次数::::: ` + strconv.Itoa(fillResult.N55) + `
Freespin划线赢钱100-200倍的总钱数::::: ` + strconv.Itoa(fillResult.N56) + ` 
Freespin划线赢钱200-500倍的次数::::: ` + strconv.Itoa(fillResult.N57) + `
Freespin划线赢钱200-500倍的总钱数::::: ` + strconv.Itoa(fillResult.N58) + `
Freespin划线赢钱500-1000倍的次数::::: ` + strconv.Itoa(fillResult.N59) + `
Freespin划线赢钱500-1000倍的总钱数::::: ` + strconv.Itoa(fillResult.N60) + `
Freespin划线赢钱1000以上倍的次数::::: ` + strconv.Itoa(fillResult.N61) + `
Freespin划线赢钱1000以上倍的总钱数::::: ` + strconv.Itoa(fillResult.N62) + `
Freespin进度条为0时的赢钱次数::::: ` + strconv.Itoa(fillResult.N63) + `
Freespin进度条为0时的赢钱数::::: ` + strconv.Itoa(fillResult.N64) + `
Freespin进度条为1时的赢钱次数::::: ` + strconv.Itoa(fillResult.N65) + `
Freespin进度条为1时的赢钱数::::: ` + strconv.Itoa(fillResult.N66) + `
Freespin进度条为2时的赢钱次数::::: ` + strconv.Itoa(fillResult.N67) + `
Freespin进度条为2时的赢钱数::::: ` + strconv.Itoa(fillResult.N68) + `
Freespin进度条为3时的赢钱次数::::: ` + strconv.Itoa(fillResult.N69) + `
Freespin进度条为3时的赢钱数::::: ` + strconv.Itoa(fillResult.N70) + `
Freespin进度条为4时的赢钱次数::::: ` + strconv.Itoa(fillResult.N71) + `
Freespin进度条为4时的赢钱数::::: ` + strconv.Itoa(fillResult.N72) + `
Freespin进度条为5时的赢钱次数::::: ` + strconv.Itoa(fillResult.N73) + `
Freespin进度条为5时的赢钱数::::: ` + strconv.Itoa(fillResult.N74) + `
Freespin进度条为6时的赢钱次数::::: ` + strconv.Itoa(fillResult.N75) + `
Freespin进度条为6时的赢钱数::::: ` + strconv.Itoa(fillResult.N76) + `
普通转触发Freespin玩法次数::::: ` + strconv.Itoa(fillResult.N77) + `
Freespin玩法总spin转动次数:::::` + strconv.Itoa(fillResult.N78) + `
平方和:::::` + fmt.Sprintf("%v", fillResult.N80)
	}
	return nil
}

func (r *unit3TestResult) Record(mSpin *slotHandle.MergeSpin) {
	noFreespin := lo.Filter(mSpin.Spins, func(item *component.Spin, index int) bool {
		return !item.IsFree
	})
	sumSpin := lo.SumBy(noFreespin, func(i *component.Spin) int {
		return i.Gain
	})
	sumZSpin := lo.SumBy(mSpin.Spins, func(i *component.Spin) int {
		return i.Gain
	})
	sumSpin += mSpin.First.Gain
	r.N1 += sumZSpin + mSpin.First.Gain
	r.N2 += mSpin.First.Bet
	r.N3++
	r.N79++
	r.N79 += len(mSpin.Spins)
	npf := math.Pow(float64(sumZSpin+mSpin.First.Gain), 2)

	r.N80 += npf
	Multiple := float64(sumSpin) / float64(mSpin.First.Bet)
	if Multiple >= 0 && Multiple < 5 {
		r.N5++
		r.N6 += sumSpin
	} else if Multiple >= 5 && Multiple < 10 {
		r.N7++
		r.N8 += sumSpin
	} else if Multiple >= 10 && Multiple < 20 {
		r.N9++
		r.N10 += sumSpin
	} else if Multiple >= 20 && Multiple < 30 {
		r.N11++
		r.N12 += sumSpin
	} else if Multiple >= 30 && Multiple < 40 {
		r.N13++
		r.N14 += sumSpin
	} else if Multiple >= 40 && Multiple < 50 {
		r.N15++
		r.N16 += sumSpin
	} else if Multiple >= 50 && Multiple < 100 {
		r.N17++
		r.N18 += sumSpin
	} else if Multiple >= 100 && Multiple < 200 {
		r.N19++
		r.N20 += sumSpin
	} else if Multiple >= 200 && Multiple < 500 {
		r.N21++
		r.N22 += sumSpin
	} else if Multiple >= 500 && Multiple < 1000 {
		r.N23++
		r.N24 += sumSpin
	} else if Multiple >= 1000 {
		r.N25++
		r.N26 += sumSpin
	}
	if sumSpin > 0 {
		r.N4++
		ss := append(mSpin.Spins, mSpin.First)
		spins := lo.Filter(ss, func(spin *component.Spin, index int) bool {
			return (!spin.IsFree) && spin.Gain > 0
		})

		for _, v := range spins {
			if v.Gain == 0 {
				continue
			}
			switch v.Rank {
			case 0:
				r.N27++
				r.N28 += v.Gain
			case 1:
				r.N29++
				r.N30 += v.Gain
			case 2:
				r.N31++
				r.N32 += v.Gain
			case 3:
				r.N33++
				r.N34 += v.Gain
			case 4:
				r.N35++
				r.N36 += v.Gain
			case 5:
				r.N37++
				r.N38 += v.Gain
			case 6:
				r.N39++
				r.N40 += v.Gain
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
				r.N41++
				r.N42 += v.Gain
			} else if freeMultiple >= 5 && freeMultiple < 10 {
				r.N43++
				r.N44 += v.Gain
			} else if freeMultiple >= 10 && freeMultiple < 20 {
				r.N45++
				r.N46 += v.Gain
			} else if freeMultiple >= 20 && freeMultiple < 30 {
				r.N47++
				r.N48 += v.Gain
			} else if freeMultiple >= 30 && freeMultiple < 40 {
				r.N49++
				r.N50 += v.Gain
			} else if freeMultiple >= 40 && freeMultiple < 50 {
				r.N51++
				r.N52 += v.Gain
			} else if freeMultiple >= 50 && freeMultiple < 100 {
				r.N53++
				r.N54 += v.Gain
			} else if freeMultiple >= 100 && freeMultiple < 200 {
				r.N55++
				r.N56 += v.Gain
			} else if freeMultiple >= 200 && freeMultiple < 500 {
				r.N57++
				r.N58 += v.Gain
			} else if freeMultiple >= 500 && freeMultiple < 1000 {
				r.N59++
				r.N60 += v.Gain
			} else if freeMultiple >= 1000 {
				r.N61++
				r.N62 += v.Gain
			}
			if v.Gain > 0 {
				switch v.Rank {
				case 0:
					r.N63++
					r.N64 += v.Gain
				case 1:
					r.N65++
					r.N66 += v.Gain
				case 2:
					r.N67++
					r.N68 += v.Gain
				case 3:
					r.N69++
					r.N70 += v.Gain
				case 4:
					r.N71++
					r.N72 += v.Gain
				case 5:
					r.N73++
					r.N74 += v.Gain
				case 6:
					r.N75++
					r.N76 += v.Gain
				}
			}

		}
		r.N77 += len(lo.Filter(spins, func(spin *component.Spin, index int) bool {
			return spin.FreeSpinParams.Count > 0
		}))
		r.N78 += len(lo.Filter(ss, func(spin *component.Spin, index int) bool {
			return spin.IsFree
		}))
	}
}

func TestDeath(slotTest *business.SlotTests, slotId uint, Num int, Amount int, Hold int, opts ...component.Option) error {
	defer public.EndProcessing(slotTest, &public.BasicData{N1: 0, N3: 0})
	str := ""
	for a := 0; a < Num; a++ {
		count := 0
		HoldCop := Hold
		AmountCop := Amount
		for i := 0; i < 1000000 && HoldCop >= AmountCop; i++ {
			HoldCop -= AmountCop
			count++
			m, err := slot.Play(slotId, AmountCop, opts...)
			if err != nil {
				global.GVA_LOG.Error("死亡测试错误:" + err.Error())
				return err
			}
			spin := m.GetSpin()
			mergeSpin1, err := slotHandle.RunMergeSpin(0, spin, &business.SlotUserSpin{
				UserId: 0,
				SlotId: slotId,
			})
			if err != nil {
				slotTest.Detail = err.Error()
				return err
			}
			sumGain := lo.SumBy(mergeSpin1.Spins, func(spin *component.Spin) int {
				return spin.Gain
			})
			HoldCop += spin.Gain + sumGain
		}
		str += "编号:" + strconv.Itoa(a+1) + " 余额:" + strconv.Itoa(HoldCop) + " 死亡次数:" + strconv.Itoa(count) + "\r\n"
	}
	slotTest.Detail = str
	return nil
}
