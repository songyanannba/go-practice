package public

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/service/logic/gameHandle/slotHandle"
	"slot-server/service/slot"
	"slot-server/service/slot/base"
	"slot-server/service/slot/component"
	"slot-server/utils"
	"slot-server/utils/helper"
	"strings"
)

// TestOnce 单次测试
func TestOnce(run RunSlotTest, opts ...component.Option) error {
	var (
		detail string
	)
	spin, err := component.NewSpin(run.SlotId, run.Amount, opts...)
	if err != nil {
		return err
	}
	if run.Raise == 1 {
		raise := decimal.NewFromInt(int64(run.Amount)).Mul(decimal.NewFromFloat(spin.Config.Raise)).IntPart()
		spin.Raise = raise
		opts = append(opts, component.WithRaise(raise))
	}

	if run.Type == enum.SlotTestType4Result {
		strs := utils.FormatCommand(run.Result)
		if len(strs) != spin.Config.Index*spin.Config.Row {
			return errors.New("结果标签数量错误")
		}
		var arr = make([][]*base.Tag, spin.Config.Row)
		for i := 0; i < spin.Config.Row; i++ {
			for j := 0; j < spin.Config.Index; j++ {
				tagName := strings.TrimSpace(strs[i*spin.Config.Index+j])
				arr[i] = append(arr[i], spin.Config.GetTag(tagName))
			}
		}
		opts = append(opts, component.WithSetResult(helper.ArrVertical(arr)))

	}

	if run.IsMustFree > 0 {
		opts = append(opts, component.WithIsMustFree())
	}

	if run.IsMustRes > 0 {
		opts = append(opts, component.WithIsMustRes())
	}

	m, err := slot.Play(run.SlotId, run.Amount, opts...)
	spin = m.GetSpin()
	if err != nil {
		return err
	}

	detail, err = global.Json.MarshalToString(spin.Dump())
	if err != nil {
		return err
	}
	sta := spin.GateData()
	fmt.Print(sta)

	slotTest := business.SlotTests{
		Type:     uint8(run.Type),
		SlotId:   run.SlotId,
		Hold:     0,
		Amount:   run.Amount,
		Win:      spin.Gain,
		MaxNum:   1,
		RunNum:   1,
		Detail:   detail,
		Status:   enum.CommonStatusFinish,
		Bet:      spin.Bet,
		Raise:    int(spin.Raise) + int(spin.BuyFreeCoin) + int(spin.BuyReCoin),
		GameType: enum.SlotSpinType1Normal,
		TestId:   0,
		Rank:     spin.Rank,
		GameData: spin.GateData(),
	}
	err = global.GVA_DB.Create(&slotTest).Error

	var slotTests []business.SlotTests

	if spin.FreeSpinParams.FreeNum > 0 || spin.FreeSpinParams.ReNum > 0 {
		mergeSpin, err := slotHandle.RunMergeSpin(0, spin, &business.SlotUserSpin{})

		if err != nil {
			return err
		}
		for _, reSpin := range mergeSpin.Spins {
			detail, err = global.Json.MarshalToString(reSpin.Dump())

			slotTests = append(slotTests, business.SlotTests{
				Type:     uint8(run.Type),
				SlotId:   run.SlotId,
				Hold:     0,
				Amount:   run.Amount,
				Win:      reSpin.Gain,
				MaxNum:   1,
				RunNum:   1,
				Detail:   detail,
				Status:   enum.CommonStatusFinish,
				Bet:      reSpin.Bet,
				Raise:    0,
				GameType: reSpin.Type(),
				TestId:   int(slotTest.ID),
				Rank:     reSpin.Rank,
				GameData: reSpin.GateData(),
			})

		}

		if len(slotTests) > 0 {
			err = global.GVA_DB.Create(&slotTests).Error
		}
	}

	return err
}
