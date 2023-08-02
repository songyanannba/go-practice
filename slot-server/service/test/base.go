package test

import (
	"fmt"
	"github.com/samber/lo"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/service/slot/component"
	"slot-server/service/test/public"
	"slot-server/service/test/test/unit1"
	"slot-server/service/test/test/unit2"
	"slot-server/service/test/test/unit3"
	"slot-server/service/test/test/unit4"
	"slot-server/service/test/test/unit5"
	"slot-server/service/test/test/unit6"
	"slot-server/service/test/test/unit8"
	"slot-server/utils/helper"
	"strconv"
)

type Repeat interface {
	Repeat(public.RunSlotTest, ...component.Option) ([]*component.Spin, error)
	Calculate([]*component.Spin)
	GetDetail() string
	GetBasicData() *public.BasicData
	GetResult(spins []*component.Spin) (int64, int64)
	GetReturnRatio() float64
}

// RepeatTest 指定次数测试
func RepeatTest(slotTest *business.SlotTests, run public.RunSlotTest, opts ...component.Option) error {
	var repeat Repeat
	switch run.SlotId {
	case 1:
		repeat = unit1.NewUnit(run)
	case 2:
		repeat = unit2.NewUnit(run)
	case 3:
		repeat = unit3.NewUnit(run)
	case 4:
		repeat = unit4.NewUnit(run)
	case 5:
		repeat = unit5.NewUnit(run)
	case 6:
		repeat = unit6.NewUnit(run)
	case 8:
		repeat = unit8.NewUnit()
	default:

	}

	defer public.EndProcessing(slotTest, repeat.GetBasicData())
	ch, errCh, _ := helper.Parallel[[]*component.Spin](run.Num, 1000, func() (spins []*component.Spin, err error) {
		for i := 0; i < 10; i++ {
			spins, err = repeat.Repeat(run, opts...)
			if err != nil {
				return nil, err
			}
			spin := spins[0]
			mul := 0
			if spin.Bet != 0 {
				mul = lo.SumBy(spins, func(spin *component.Spin) int {
					return spin.Gain
				}) / spin.Bet
			}
			if mul <= spin.Config.TopMul {
				break
			}
		}
		return
	})
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
			repeat.Calculate(v)
		}
	}
end:
	slotTest.Detail = repeat.GetDetail()
	return nil
}

func DeathTest(slotTest *business.SlotTests, run public.RunSlotTest, opts ...component.Option) error {
	var repeat Repeat
	switch run.SlotId {
	case 1:
		repeat = unit1.NewUnit(run)
	case 2:
		repeat = unit2.NewUnit(run)
	case 3:
		repeat = unit3.NewUnit(run)
	case 4:
		repeat = unit4.NewUnit(run)
	case 5:
		repeat = unit5.NewUnit(run)
	case 6:
		repeat = unit6.NewUnit(run)
	case 8:
		repeat = unit8.NewUnit()
	default:

	}
	defer public.EndProcessing(slotTest, repeat.GetBasicData())

	ch, errCh, _ := helper.Parallel[string](run.Num, 1000, func() (string, error) {
		return OnlyDeathTest(repeat, run, opts...)
	})
	count := 0
	a := 0
	str := ""
	for {
		a++
		select {
		case err := <-errCh:
			slotTest.Detail = err.Error()
			str += err.Error()
		case v, beforeClosed := <-ch:
			count++
			if run.Num >= 1000 && count%(run.Num/10) == 0 && count < run.Num-1 {
				go func() {
					slotTest.Detail = fmt.Sprintf("进度:%d/%d", count, run.Num)
					global.GVA_DB.Save(&slotTest)
				}()
			}
			if !beforeClosed {
				goto end
			}
			str += fmt.Sprintf("%d %s", count, v)
		}
	}
end:
	slotTest.Detail = str
	return nil

}

func OnlyDeathTest(repeat Repeat, run public.RunSlotTest, opts ...component.Option) (string, error) {
	holdCop := int64(run.Hold)

	count := 0
	for i := 0; i < 100000 && holdCop >= int64(run.Amount); i++ {
		count++
		spins, err := repeat.Repeat(run, opts...)
		if err != nil {
			return err.Error(), err
		}
		gain, bet := repeat.GetResult(spins)
		holdCop += gain
		holdCop -= bet
	}

	return strconv.Itoa(int(holdCop)) + " " + strconv.Itoa(count) + "\r\n", nil
}
