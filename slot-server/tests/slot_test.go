package tests

import (
	"slot-server/core"
	"slot-server/service/slot"
	"slot-server/service/slot/base"
	"slot-server/service/slot/component"
	"slot-server/service/slot/machine/unit5"
	"testing"
)

func TestSlot(t *testing.T) {
	core.BaseInit()
	m, _ := slot.Play(3, 100)
	t.Log(m)
}

func winMatchInit() (*unit5.Machine, [][]*base.Tag) {
	core.BaseInit()
	s, _ := component.NewSpin(5, 100)
	m := unit5.NewMachine(s)
	m.Table = component.NewGraph(m.BaseSpin, false)
	return m, m.Table.FindAllLine()
}

// 赢钱匹配基准测试
func BenchmarkWinMatch(b *testing.B) {
	m, lineList := winMatchInit()

	b.Run("MatchWinMul", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, tags := range lineList {
				m.Table.MatchWinMul(tags)
			}
		}
	})

	b.Run("WinMatch", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, tags := range lineList {
				m.Table.WinMatch(tags)
			}
		}
	})
}
