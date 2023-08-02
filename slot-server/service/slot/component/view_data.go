package component

import (
	"slot-server/service/slot/base"
	"slot-server/utils/helper"
	"strings"
)

type ViewData struct {
	Data         string    `json:"data"`
	Bet          int       `json:"bet"`          // 压注
	Gain         int       `json:"gain"`         // 赢钱
	PayTableMuls []float64 `json:"payTableMuls"` // payTable倍数
	WildMuls     []float64 `json:"wildMuls"`     // 百搭倍数
	JackpotMul   float64   `json:"jackpotMul"`   // 奖池倍数
	FreeSpin     int       `json:"freeSpin"`     // 免费转次数
}

func (s *Spin) DumpPayTable() string {
	var arr []string
	for _, table := range s.PayTables {
		arr = append(arr, table.Dump())
	}
	return strings.Join(arr, "\n")
}

func (s *Spin) Dump() *ViewData {
	var view = &ViewData{}
	view.Data = s.DumpData()
	view.Bet = s.Bet
	view.Gain = s.Gain
	for _, table := range s.PayTables {
		view.PayTableMuls = append(view.PayTableMuls, table.Multiple)
	}
	for _, wild := range s.WildList {
		view.WildMuls = append(view.WildMuls, wild.Multiple)
	}
	if s.Jackpot != nil {
		view.JackpotMul = s.Jackpot.End
	}
	return view
}

func (s *Spin) DumpData() string {
	return dataStr(s.InitDataList, func(tag *base.Tag) string {
		return tag.Dump()
	}, ";")
}

func dataStr[T any](arr [][]T, f func(T) string, sep string) string {
	arr = helper.ArrVertical(arr)
	str := ""
	for i, row := range arr {
		for ii, col := range row {
			str += f(col)
			if ii < len(row)-1 {
				str += ","
			}
		}
		if i < len(arr)-1 {
			str += sep
		}
	}
	return str
}
