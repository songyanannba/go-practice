package unit6

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/service/slot"
	"slot-server/service/slot/base"
	"slot-server/service/slot/component"
	"slot-server/service/test/public"
	"slot-server/utils"
	"slot-server/utils/helper"
	"strconv"
	"strings"
)

type Unit6 struct {
	Result unit6Result
}

type unit6Result struct {
	*public.BasicData
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
	N45 int //普通转连续消除10+次的次数:::::
	N46 int //普通转连续消除10+次的赢钱数:::::

	N47 int //进度1达到次数
	N48 int //进度2达到次数
	N49 int //进度3达到次数
	N50 int //进度4达到次数
	N51 int //进度5达到次数
	N52 int //最大赢钱倍数
}

func NewUnit(run public.RunSlotTest) *Unit6 {
	return &Unit6{Result: unit6Result{BasicData: &public.BasicData{}}}
}

func (u *Unit6) Repeat(run public.RunSlotTest, opts ...component.Option) ([]*component.Spin, error) {
	m, err := slot.Play(run.SlotId, run.Amount, opts...)
	if err != nil {
		return nil, err
	}
	spins := []*component.Spin{m.GetSpin()}
	spins = append(spins, m.GetSpins()...)
	return spins, nil
}

func (u *Unit6) Calculate(spins []*component.Spin) { //6
	public.AnalyticData(u.Result.BasicData, spins)

	spin := spins[0]

	rId := spin.Table.AlterFlows[len(spin.Table.AlterFlows)-1].RankId

	//普通转消除次数和赢钱
	{
		switch len(spin.Table.AlterFlows) {
		case 2:
			u.Result.N27++
			u.Result.N28 += spin.Gain
		case 3:
			u.Result.N29++
			u.Result.N30 += spin.Gain
		case 4:
			u.Result.N31++
			u.Result.N32 += spin.Gain
		case 5:
			u.Result.N33++
			u.Result.N34 += spin.Gain
		case 6:
			u.Result.N35++
			u.Result.N36 += spin.Gain
		case 7:
			u.Result.N37++
			u.Result.N38 += spin.Gain
		case 8:
			u.Result.N39++
			u.Result.N40 += spin.Gain
		case 9:
			u.Result.N41++
			u.Result.N42 += spin.Gain
		case 10:
			u.Result.N43++
			u.Result.N44 += spin.Gain
		case 11:
			u.Result.N45++
			u.Result.N46 += spin.Gain

		default:
			u.Result.N45++
			u.Result.N46 += spin.Gain
		}
	}

	{
		switch rId {
		case 1:
			u.Result.N47++
		case 2:
			u.Result.N48++
		case 3:
			u.Result.N49++
		case 4:
			u.Result.N50++
		case 5:
			u.Result.N51++
			//default:
			//	u.Result.N37++
			//	u.Result.N38 += spin.Gain
		}
	}

	var MaxBet int
	for _, v := range spin.Table.AlterFlows {
		if int(decimal.NewFromFloat(v.SumMul).IntPart()) > MaxBet {
			MaxBet = int(decimal.NewFromFloat(v.SumMul).IntPart())
		}
	}
	u.Result.N52 = MaxBet
}

func (u *Unit6) GetDetail() string {
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
进度1达到次数:::: ` + strconv.Itoa(u.Result.N47) + `
进度2达到次数:::: ` + strconv.Itoa(u.Result.N48) + `
进度3达到次数:::: ` + strconv.Itoa(u.Result.N49) + `
进度4达到次数:::: ` + strconv.Itoa(u.Result.N50) + `
进度5达到次数:::: ` + strconv.Itoa(u.Result.N51) + `
最大赢钱倍数::::  ` + strconv.Itoa(u.Result.N52)
	return str
}

func (u *Unit6) GetBasicData() *public.BasicData {
	return u.Result.BasicData
}

func (u *Unit6) GetResult(spins []*component.Spin) (int64, int64) {
	gain := lo.SumBy(spins, func(spin *component.Spin) int64 {
		return int64(spin.Gain)
	})
	bet := int64(spins[0].Bet) + int64(spins[0].Raise) + int64(spins[0].BuyFreeCoin) + int64(spins[0].BuyReCoin)
	return gain, bet
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
	for _, flow := range spin.Table.AlterFlows {
		slotTest := business.SlotTests{
			Type:     uint8(run.Type),
			SlotId:   run.SlotId,
			Hold:     0,
			Amount:   run.Amount,
			Win:      flow.Gain,
			MaxNum:   flow.RemoveCount,
			RunNum:   1,
			Detail:   flow.String(),
			Status:   enum.CommonStatusFinish,
			Bet:      spin.Bet,
			Raise:    helper.If(flow.Id == 0, int(spin.Raise)+int(spin.BuyFreeCoin)+int(spin.BuyReCoin), 0),
			GameType: helper.If(flow.Id == 0, enum.SlotSpinType1Normal, enum.SlotSpinType3Respin),
			TestId:   helper.If(flow.Id == 0, 0, int(mainSlotTest.ID)),
			Rank:     flow.RankId,
			GameData: helper.If(flow.Id == 0, "整局信息:"+spin.Table.GetInformation()+"\r\n 本次转信息: "+flow.GetInformation(), flow.GetInformation()),
		}
		if flow.Id == 0 {
			mainSlotTest = slotTest
			err = global.GVA_DB.Create(&mainSlotTest).Error
		} else {
			slotTests = append(slotTests, slotTest)
		}
	}
	for _, c := range spins {
		for _, flow := range c.Table.AlterFlows {
			slotTest := business.SlotTests{
				Type:     uint8(run.Type),
				SlotId:   run.SlotId,
				Hold:     0,
				Amount:   run.Amount,
				Win:      flow.Gain,
				MaxNum:   1,
				RunNum:   1,
				Detail:   flow.String(),
				Status:   enum.CommonStatusFinish,
				Bet:      spin.Bet,
				Raise:    0,
				GameType: helper.If(flow.Id == 0, enum.SlotSpinType3Respin, enum.SlotSpinType4FsRs),
				TestId:   int(mainSlotTest.ID),
				Rank:     spin.Rank,
				GameData: helper.If(flow.Id == 0, "整局信息:"+c.Table.GetInformation()+"\r\n 本次转信息: "+flow.GetInformation(), flow.GetInformation()),
			}
			slotTests = append(slotTests, slotTest)
		}
	}

	if len(slotTests) > 0 {
		err = global.GVA_DB.Create(&slotTests).Error
	}

	return err
}

func (u *Unit6) GetReturnRatio() float64 {
	f, b := decimal.NewFromInt(int64(u.Result.N1)).Div(decimal.NewFromInt(int64(u.Result.N2))).Float64()
	if b {
		return f
	}
	return 0
}
