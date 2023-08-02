package slotHandle

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"slot-server/pbs"
	"slot-server/service/logic"
	"slot-server/service/logic/gameHandle"
	"slot-server/service/slot"
	"slot-server/service/slot/component"
	"slot-server/utils/helper"
)

type Handle struct {
	*gameHandle.Handle
	Spin      *component.Spin
	MergeSpin *MergeSpin
}

func NewHandle() *Handle {
	return &Handle{}
}

func (h *Handle) SetHandle(gameHandle *gameHandle.Handle) {
	h.Handle = gameHandle
}

func (h *Handle) GetSpin() *component.Spin {
	return h.Spin
}

func (h *Handle) GetSpins() []*component.Spin {
	spins := []*component.Spin{h.Spin}
	if h.MergeSpin != nil {
		spins = append(spins, h.MergeSpin.Spins...)
	}
	return spins
}

func (h *Handle) NeedJackpot() bool {
	if h.SlotId == 1 {
		return true
	}
	return false
}

func (h *Handle) GetAck() protoreflect.ProtoMessage {
	ack := sumAck(h.GetSpins())
	ack.Opt = h.Req.Opt
	ack.Head = logic.NewAckHead(h.Req.Head)
	ack.TxnId = int32(h.Txn.ID)
	ack.BeforeAmount = h.Bill.OriginalAmount
	ack.AfterAmount = h.Bill.AfterAmount
	return ack
}

func (h *Handle) Run() error {
	h.Options = append(h.Options, component.SetDebugConfig(h.User.ID))
	m, err := slot.Play(h.SlotId, int(h.Bet), h.Options...)
	if err != nil {
		return err
	}
	// 获取第一次转结果
	h.Spin = m.GetSpin()
	h.Spin.Id = 1
	// 重新赋值防止超出倍率重转逻辑影响结果
	h.MergeSpin = nil
	// 如果出现免费转或respin则继续转
	if h.Spin.KeepSpin() {
		h.MergeSpin, err = RunMergeSpin(h.User.ID, h.Spin, h.UserSpin)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *Handle) GetTotalWin() int64 {
	return helper.SumByFunc(h.GetSpins(), func(spin *component.Spin) int64 {
		return int64(spin.Gain)
	})
}

func (h *Handle) GetPlayNum() (int64, int64) {
	if h.MergeSpin == nil {
		return 1, 0
	}
	return int64(1 + h.MergeSpin.ReSpinFinishNum), int64(h.MergeSpin.FreeSpinFinishNum)
}

func SumStep(spin *component.Spin) *pbs.SpinStep {
	ack := pbs.SpinStep{}
	ack.Win = int64(spin.Gain)
	//ack.FreeNum = int32(freeNum)
	ack.CardList = []*pbs.Cards{}
	ack.Type = int32(spin.Type())
	ack.Id = int32(spin.Id)
	ack.Pid = int32(spin.ParentId)
	ack.Which = int32(spin.Which)
	if spin.Jackpot != nil {
		ack.JackpotId = int32(spin.Jackpot.Id)
	}
	for _, tags := range spin.InitDataList {
		cards := pbs.Cards{}
		for _, tag := range tags {
			cards.Cards = append(cards.Cards, tag.ToCard())
		}
		ack.CardList = append(ack.CardList, &cards)
	}
	for _, line := range spin.WinLines {
		cards := pbs.Cards{Amount: line.Win}
		for _, tag := range line.Tags {
			cards.Cards = append(cards.Cards, tag.ToCard())
		}
		ack.LineList = append(ack.LineList, &cards)
	}
	return &ack
}

func sumAck(spins []*component.Spin) (ack *pbs.SpinAck) {
	ack = &pbs.SpinAck{}
	for _, spin := range spins {
		ack.StepList = append(ack.StepList, SumStep(spin))
		ack.TotalWin += int64(spin.Gain)
	}
	return
}
