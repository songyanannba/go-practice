package slotTpl

import (
	"slot-server/model/business"
	"slot-server/service/slot"
	"slot-server/service/slot/component"
)

func Run(slotGenTpl *business.SlotGenTpl) error {
	h := NewTplHandle(slotGenTpl)
	err := h.ParseParams()
	if err != nil {
		return err
	}
	h.Exec()
	return nil
}

func (h *TplHandle) Exec() error {
	config, err := component.GetSlotConfig(h.Model.SlotId, false)
	if err != nil {
		return err
	}
	c := *config
	h.Config = &c
	for i := 0; i < h.Model.Num; i++ {
		h.Gen()
	}
	return nil
}

func (h *TplHandle) Gen() error {
	//h.Params.Height
	//h.Params.Width
	spin := &component.Spin{
		Options: nil,
		Bet:     0,
		Gain:    0,
		Config:  h.Config,
	}
	m, err := slot.GetMachine(spin)
	if err != nil {
		return err
	}
	err = m.Exec()
	if err != nil {
		return err
	}
	return nil
}
