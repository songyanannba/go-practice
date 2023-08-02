package unit1

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"slot-server/service/slot"
	"slot-server/service/slot/component"
	"slot-server/service/test/public"
)

type Unit1 struct {
	Result unit1Result
}

type unit1Result struct {
	*public.BasicData
	N27 int              //jackpot_1赢钱次数:::::
	N28 int              //jackpot_1赢钱总数:::::
	N29 int              //jackpot_2赢钱次数:::::
	N30 int              //jackpot_2赢钱总数:::::
	N31 int              //jackpot_3赢钱次数:::::
	N32 int              //jackpot_3赢钱总数:::::
	N33 int              //jackpot_4赢钱次数:::::
	N34 int              //jackpot_4赢钱总数:::::
	N35 []map[string]int //统计数据
}

func NewUnit(run public.RunSlotTest) *Unit1 {
	return &Unit1{Result: unit1Result{BasicData: &public.BasicData{}}}
}

func (u *Unit1) Repeat(run public.RunSlotTest, opts ...component.Option) ([]*component.Spin, error) {
	m, _ := slot.Play(run.SlotId, run.Amount, opts...)
	spin := m.GetSpin()
	return []*component.Spin{spin}, nil

}

func (u *Unit1) Calculate(spins []*component.Spin) {
	spin := spins[0]
	public.AnalyticData(u.Result.BasicData, spins)
	if spin.Jackpot != nil {
		if spin.Jackpot.Id == 1 {
			u.Result.N27++
			u.Result.N28 += spin.Gain
		} else if spin.Jackpot.Id == 2 {
			u.Result.N29++
			u.Result.N30 += spin.Gain
		} else if spin.Jackpot.Id == 3 {
			u.Result.N31++
			u.Result.N32 += spin.Gain
		} else if spin.Jackpot.Id == 4 {
			u.Result.N33++
			u.Result.N34 += spin.Gain
		}
	}
	//初始化标签统计
	for i := 0; len(u.Result.N35) < spin.Config.Index; i++ {
		var tags = make(map[string]int)
		for _, v := range spin.Config.GetAllTag() {
			tags[v.Name] = 0
		}
		u.Result.N35 = append(u.Result.N35, tags)
	}

	for i, v := range spin.InitDataList[1] {
		tags := u.Result.N35[i]
		tags[v.Name]++
	}
}

func (u *Unit1) GetDetail() string {
	str := public.GetInitResult(u.Result.BasicData)
	str += fmt.Sprintf("jackpot_1赢钱次数:::::%d\n"+
		"jackpot_1赢钱总数:::::%d\n"+
		"jackpot_2赢钱次数:::::%d\n"+
		"jackpot_2赢钱总数:::::%d\n"+
		"jackpot_3赢钱次数:::::%d\n"+
		"jackpot_3赢钱总数:::::%d\n"+
		"jackpot_4赢钱次数:::::%d\n"+
		"jackpot_4赢钱总数:::::%d\n",
		u.Result.N27, u.Result.N28, u.Result.N29, u.Result.N30, u.Result.N31, u.Result.N32, u.Result.N33, u.Result.N34)

	//拼接统计数据
	for i, v := range u.Result.N35 {
		str += fmt.Sprintf("第 %d 列标签出现次数统计:\n", i+1)
		for vv, ii := range v {
			str += fmt.Sprintf("%s:::::%d\n", vv, ii)
		}
	}
	str += fmt.Sprintf("最大倍率:::::%g\n", u.Result.N0)
	return str
}

func (u *Unit1) GetBasicData() *public.BasicData {
	return u.Result.BasicData
}

func (u *Unit1) GetResult(spins []*component.Spin) (int64, int64) {
	gain := lo.SumBy(spins, func(spin *component.Spin) int64 {
		return int64(spin.Gain)
	})
	bet := int64(spins[0].Bet) + int64(spins[0].Raise) + int64(spins[0].BuyFreeCoin) + int64(spins[0].BuyReCoin)
	return gain, bet
}

func (u *Unit1) GetReturnRatio() float64 {
	f, b := decimal.NewFromInt(int64(u.Result.N1)).Div(decimal.NewFromInt(int64(u.Result.N2))).Float64()
	if b {
		return f
	}
	return 0
}
