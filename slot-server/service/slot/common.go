package slot

import (
	"errors"
	"slot-server/enum"
	"slot-server/service/slot/component"
	"slot-server/service/slot/machine/unit1"
	"slot-server/service/slot/machine/unit2"
	"slot-server/service/slot/machine/unit3"
	"slot-server/service/slot/machine/unit4"
	"slot-server/service/slot/machine/unit5"
	"slot-server/service/slot/machine/unit6"
	"slot-server/service/slot/machine/unit7"
	"slot-server/service/slot/machine/unit8"
)

type Machine interface {
	GetSpin() *component.Spin
	Exec() error
	GetInitData()
	GetResData()
	SumGain()
	GetSpins() []*component.Spin
}

func Play(slotId uint, amount int, options ...component.Option) (m Machine, err error) {
	var s *component.Spin
	s, err = component.NewSpin(slotId, amount, options...)
	if err != nil {
		return nil, err
	}
	return RunSpin(s)
}

func RunSpin(s *component.Spin) (m Machine, err error) {
	m, err = GetMachine(s)
	if err != nil {
		return
	}
	err = m.Exec()
	if err != nil {
		return nil, err
	}
	return m, nil
}

func GetMachine(s *component.Spin) (m Machine, err error) {
	switch s.Config.SlotId {
	case enum.SlotId1:
		m = unit1.NewMachine(s)
	case enum.SlotId2:
		m = unit2.NewMachine(s)
	case enum.SlotId3:
		m = unit3.NewMachine(s)
	case enum.SlotId4:
		m = unit4.NewMachine(s)
	case enum.SlotId5:
		m = unit5.NewMachine(s)
	case enum.SlotId6:
		m = unit6.NewMachine(s)
	case enum.SlotId7:
		m = unit7.NewMachine(s)
	case enum.SlotId8:
		m = unit8.NewMachine(s)
	default:
		return nil, errors.New("slotId not found")
	}
	return m, nil
}
