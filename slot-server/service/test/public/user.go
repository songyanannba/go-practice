package public

import (
	"errors"
	"fmt"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/pbs"
	"slot-server/service/logic"
	"slot-server/service/logic/gameHandle"
	"slot-server/service/slot/component"
	"slot-server/service/test/local"
	"slot-server/utils/helper"
	"strconv"
)

func TestUser(run RunSlotTest, opts ...component.Option) (err error) {
	lt := local.NewLocalTest()

	// 商户登录
	//ack, err := lt.MerchantLogin()
	//if err != nil {
	//	return err
	//}
	//ack = lt.Entity.Ack.(*pbs.MerchantLoginAck)
	//lt.Head = &pbs.ReqHead{Token: ack.Token, Platform: "test"}

	lt.Head = &pbs.ReqHead{Token: "123456", Platform: "test"}
	lt.Session.Set(enum.SessionDataTestSpinHandle, true)

	// spin
	spinAck, matchSpinAck, err := lt.Spin(&pbs.Spin{Opt: &pbs.SpinOpt{
		GameId: int32(run.SlotId), Bet: int64(run.Amount), Raise: helper.If(run.Raise == 1, true, false),
	}})
	handle, ok := lt.Session.Value("handle").(*gameHandle.Handle)
	if !ok {
		return errors.New("handle is nil")
	}
	detail := "gameRecordId: " + strconv.Itoa(int(handle.Record.ID)) + " txnId: " + strconv.Itoa(int(handle.Txn.ID)) + "\n\n"
	if err != nil {
		detail += err.Error()
	} else if spinAck != nil {
		detail += logic.DumpSpinAck(spinAck)
	} else if matchSpinAck != nil {
		detail += fmt.Sprintf("%+v", matchSpinAck)
	}

	spin := handle.GetSpin()

	slotTest := business.SlotTests{
		Type:     uint8(run.Type),
		SlotId:   run.SlotId,
		Hold:     0,
		Amount:   run.Amount,
		Win:      int(handle.GetTotalWin()),
		MaxNum:   1,
		RunNum:   1,
		Detail:   detail,
		Status:   enum.CommonStatusFinish,
		Bet:      run.Amount,
		Raise:    int(handle.Raise) + int(spin.BuyFreeCoin) + int(spin.BuyReCoin),
		GameType: enum.SlotSpinType1Normal,
		TestId:   0,
		Rank:     0,
		GameData: "",
	}
	err = global.GVA_DB.Create(&slotTest).Error

	return
}
