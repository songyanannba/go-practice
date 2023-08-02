package eliminateHandle

import (
	"go.uber.org/zap"
	"google.golang.org/protobuf/reflect/protoreflect"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/pbs"
	"slot-server/service/logic"
	"slot-server/service/logic/gameHandle"
	"slot-server/service/slot"
	"slot-server/service/slot/component"
	"slot-server/utils/helper"
)

type Handle struct {
	*gameHandle.Handle
	Spin  *component.Spin
	Spins []*component.Spin
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
	return append([]*component.Spin{h.Spin}, h.Spins...)
}

func (h *Handle) NeedJackpot() bool {
	return false
}

func (h *Handle) GetAck() protoreflect.ProtoMessage {
	ack := h.SumAck()
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
	h.Spin = m.GetSpin()
	h.Spins = m.GetSpins()
	return nil
}

func (h *Handle) GetTotalWin() int64 {
	return helper.SumByFunc(h.GetSpins(), func(spin *component.Spin) int64 {
		return int64(spin.Gain)
	})
}

// GetPlayNum 获取普通玩和免费玩的次数
func (h *Handle) GetPlayNum() (int64, int64) {
	// TODO 获取普通玩和免费玩的次数
	return 0, 0
}

func (h *Handle) SumAck() *pbs.MatchSpinAck {
	ack := &pbs.MatchSpinAck{
		Opt:        h.Req.Opt,
		TotalWin:   0,
		TotalBet:   int64(h.Spin.Bet),
		TotalRaise: int64(h.Spin.Raise),
		Steps:      make([]*pbs.MatchSpinStep, 0),
	}

	spins := []*component.Spin{h.Spin}
	spins = append(spins, h.Spins...)
	for i, spin := range spins {
		step := &pbs.MatchSpinStep{
			Id:   int32(spin.Id),
			Pid:  int32(spin.ParentId),
			Type: helper.If(i == 0, int32(enum.NormalSpin), int32(enum.FreeSpin)),
			//FreeNum:  int32(spin.FreeSpinParams.FreeNum),
			SumGain:  int64(spin.Gain),
			InitList: make([]*pbs.Tags, 0),
			Flows:    make([]*pbs.StepFlow, 0),
		}

		initTable := spin.Table.InitTable
		for _, rows := range initTable {
			tags := &pbs.Tags{}
			tags.Tags = make([]*pbs.Tag, 0)
			for _, col := range rows {
				tags.Tags = append(tags.Tags, &pbs.Tag{
					TagId:    int32(col.Id),
					X:        int32(col.X),
					Y:        int32(col.Y),
					Multiple: float32(col.Multiple),
					IsWild:   col.IsWild,
				})
			}
			step.InitList = append(step.InitList, tags)
		}
		for _, flow := range spin.Table.AlterFlows {
			stepFlow := &pbs.StepFlow{
				Index:      int64(flow.Id),
				Gain:       int64(flow.Gain),
				RemoveList: make([]*pbs.Tags, 0),
				AddList: &pbs.Tags{
					Tags: make([]*pbs.Tag, 0),
				},
			}
			for _, elim := range flow.RemoveList {
				tags := &pbs.Tags{}
				tags.Tags = make([]*pbs.Tag, 0)
				for _, col := range elim.RemoveList {
					tags.Tags = append(tags.Tags, &pbs.Tag{
						TagId:    int32(col.Id),
						X:        int32(col.X),
						Y:        int32(col.Y),
						Multiple: float32(col.Multiple),
						IsWild:   col.IsWild,
					})
				}
				tags.Amount = helper.IntMulFloatToInt(spin.Bet, elim.Mul)
				stepFlow.RemoveList = append(stepFlow.RemoveList, tags)
			}
			for _, add := range flow.AddList {
				stepFlow.AddList.Tags = append(stepFlow.AddList.Tags, &pbs.Tag{
					TagId:    int32(add.Id),
					X:        int32(add.X),
					Y:        int32(add.Y),
					Multiple: float32(add.Multiple),
					IsWild:   add.IsWild,
				})
			}
			step.Flows = append(step.Flows, stepFlow)
		}
		ack.Steps = append(ack.Steps, step)
	}
	//fmt.Print(h.Spin.Table.InitTable)
	go WriteDB(h.GetSpins())
	return ack
}

func WriteDB(MSpins []*component.Spin) {
	global.GVA_LOG.Info("WriteDB")
	var err error
	var mainSlotTest business.SlotTests
	var slotTests []business.SlotTests
	spin := MSpins[0]
	spins := MSpins[1:]
	for _, flow := range spin.Table.AlterFlows {
		slotTest := business.SlotTests{
			Type:     uint8(5),
			SlotId:   5,
			Hold:     0,
			Amount:   spin.Bet,
			Win:      flow.Gain,
			MaxNum:   1,
			RunNum:   1,
			Detail:   flow.String(),
			Status:   enum.CommonStatusFinish,
			Bet:      spin.Bet,
			Raise:    helper.If(flow.Id == 1, int(spin.Raise)+int(spin.BuyFreeCoin)+int(spin.BuyReCoin), 0),
			GameType: helper.If(flow.Id == 1, enum.SlotSpinType1Normal, enum.SlotSpinType3Respin),
			TestId:   helper.If(flow.Id == 1, 0, int(mainSlotTest.ID)),
			Rank:     spin.Rank,
			GameData: helper.If(flow.Id == 1, "整局信息:"+spin.Table.GetInformation()+"\r\n 本次转信息: "+flow.GetInformation(), flow.GetInformation()),
		}
		if flow.Id == 0 {
			mainSlotTest = slotTest
			err = global.GVA_DB.Create(&mainSlotTest).Error
			if err != nil {
				global.GVA_LOG.Error("WriteDB err", zap.Any("err", err))
			}
		} else {
			slotTests = append(slotTests, slotTest)
		}
	}
	for _, c := range spins {
		for _, flow := range c.Table.AlterFlows {
			slotTest := business.SlotTests{
				Type:     uint8(5),
				SlotId:   5,
				Hold:     0,
				Amount:   spin.Bet,
				Win:      flow.Gain,
				MaxNum:   1,
				RunNum:   1,
				Detail:   flow.String(),
				Status:   enum.CommonStatusFinish,
				Bet:      spin.Bet,
				Raise:    0,
				GameType: helper.If(flow.Id == 1, enum.SlotSpinType2Fs, enum.SlotSpinType4FsRs),
				TestId:   int(mainSlotTest.ID),
				Rank:     spin.Rank,
				GameData: helper.If(flow.Id == 1, "整局信息:"+c.Table.GetInformation()+"\r\n 本次转信息: "+flow.GetInformation(), flow.GetInformation()),
			}
			slotTests = append(slotTests, slotTest)
		}
	}

	if len(slotTests) > 0 {
		err = global.GVA_DB.Create(&slotTests).Error
		if err != nil {
			global.GVA_LOG.Error("WriteDB err", zap.Any("err", err))
		}
	}
}
