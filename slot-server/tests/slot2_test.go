package tests

import (
	"errors"
	"fmt"
	"slot-server/core"
	"slot-server/enum"
	"slot-server/pbs"
	"slot-server/service/logic"
	"slot-server/service/logic/gameHandle"
	"slot-server/service/test/local"
	"slot-server/utils/helper"
	"strconv"
	"testing"
)

type RunSlotTest struct {
	Type       int    `json:"type"`
	SlotId     uint   `json:"slotId"`
	Num        int    `json:"num"`
	Hold       int    `json:"hold"` //持有金币
	Amount     int    `json:"amount"`
	Opts       []int  `json:"opts"`
	Raise      int    `json:"raise"`
	Result     string `json:"result"`
	IsMustFree int    `json:"isMustFree"` // 购买免费次数
	IsMustRes  int    `json:"isMustRes"`  // 购买Respin次数
}

func TestRandSlot2(t *testing.T) {
	core.BaseInit()
	run := RunSlotTest{
		SlotId: 2,
		Amount: 100,
	}

	RandSlot2(run)
}

func RandSlot2(run RunSlotTest) (err error) {
	lt := local.NewLocalTest()

	lt.Head = &pbs.ReqHead{Token: "123456", Platform: "test"}
	lt.Session.Set(enum.SessionDataTestSpinHandle, true)

	// spin
	spinAck, matchSpinAck, err := lt.Spin(&pbs.Spin{
		Opt: &pbs.SpinOpt{
			GameId: int32(run.SlotId),
			Bet:    int64(run.Amount),
			Raise:  helper.If(run.Raise == 1, true, false),
		},
	})

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

	handle.GetSpin()
	return
}
