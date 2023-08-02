package unit1

import (
	"slot-server/enum"
	"slot-server/service/slot/base"
	"slot-server/service/slot/component"
	"slot-server/utils/helper"
)

// Machine 单线slot 3*3
type Machine struct {
	Spin *component.Spin
}

func NewMachine(spin *component.Spin) *Machine {
	return &Machine{spin}
}

func (m *Machine) GetSpin() *component.Spin {
	return m.Spin
}

func (m *Machine) GetSpins() []*component.Spin {
	return []*component.Spin{m.Spin}
}

func (m *Machine) Exec() error {

	if len(m.Spin.IsSetResult) > 0 {
		m.Spin.InitDataList = m.Spin.IsSetResult
	} else if len(m.Spin.InitDataList) == 0 {
		m.GetInitData()
	}

	m.GetResData()
	m.SumGain()
	return nil
}

func (m *Machine) GetInitData() {
	which := 1
	if v, ok := m.Spin.Config.Event.Get(enum.SlotEvent1ChangeTable); ok && v != nil {
		which = v.(*base.ChangeTableEvent).Fetch()
	}
	m.Spin.GetInitDataByReel(enum.SlotReelType1Normal, which)
	return
}

func (m *Machine) GetResData() {
	m.Spin.GetPaylineByCoords()
	return
}

func (m *Machine) SumGain() {
	s := m.Spin
	if s.JackpotMatch() != nil {
		return
	}
	// 计算payTable
	var (
		table = m.Spin.PayTableOnceMatch()
	)
	if table != nil {
		s.Gain = int(helper.FloatMul(float64(s.Bet), table.Multiple))
		return
	}
	SumSingle(s)
}

func SumSingle(s *component.Spin) {
	if len(s.SingleList) == 0 {
		return
	}
	for _, single := range s.SingleList {
		if len(s.WinLines) == 0 {
			s.WinLines = append(s.WinLines, &component.Line{
				Name: "",
				Tags: helper.SliceVal(s.ResDataList, 0),
				Win:  helper.IntMulFloatToInt(s.Bet, single.Multiple),
			})
		} else {
			s.WinLines[0].Win = helper.IntMulFloatToInt(s.WinLines[0].Win, single.Multiple)
		}
	}
	s.Gain = int(s.WinLines[0].Win)
}
