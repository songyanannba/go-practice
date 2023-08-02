package unit3

import (
	"modernc.org/mathutil"
	"slot-server/enum"
	"slot-server/service/slot/base"
	"slot-server/service/slot/component"
	"slot-server/utils/helper"
)

// Machine 划线 + 特殊免费玩
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
	var (
		which = m.Spin.Rank + 1
		typ   = enum.SlotReelType1Normal
	)
	if m.Spin.Raise > 0 {
		if m.Spin.Rank < 2 {
			m.Spin.Rank = 2
			which = 3
		}
		which += 7
	}
	m.Spin.Which = which
	if m.Spin.IsFree {
		typ = enum.SlotReelType2FS
	}
	m.Spin.GetInitDataByReel(typ, which)
	return
}

func (m *Machine) GetResData() {
	m.Spin.GetPaylineByCoords()
	return
}

func (m *Machine) payTableMatch() {
	s := m.Spin
	// 每条划线可能匹配到多个payTable 取最大的payTable
	s.PayTableMatch()
}

func (m *Machine) SumGain() {
	m.payTableMatch()
	var (
		s   = m.Spin
		mul float64
	)
	s.ResDataList = [][]*base.Tag{}

	// 计算payTable
	if len(s.PayTables) > 0 {
		var mulArr []float64
		for _, table := range s.PayTables {
			mulArr = append(mulArr, table.Multiple)
			s.ResDataList = append(s.ResDataList, table.Tags)
		}
		mul = helper.FloatSum(mulArr...)
		s.Gain = int(helper.IntMulFloatToInt(s.Bet, mul))
	}
	// 计算scatter
	scatter := base.MatchSameTagList(s.InitDataList, enum.SameTagLen, s.Config.GetTag(enum.ScatterName))
	if len(scatter) > 0 {
		s.FreeSpinParams.Count++
		s.WinLines = append(s.WinLines, &component.Line{
			Tags: scatter[0],
			Win:  int64(s.Bet * 3),
		})
		s.Gain += s.Bet * 3
	}
	s.FreeSpinParams.FreeNum += s.FreeSpinParams.Count * 10

	// 开启了加注 并且不是免费玩 则额外+2次免费玩 因为加注后直接首次进入免费转直接就是rank2
	if s.Raise > 0 && s.FreeSpinParams.Count > 0 && !s.IsFree {
		s.FreeSpinParams.FreeNum += 2
	}

	// 计算freeSpin 此处为scatter出现则乘三倍
	s.SumNextRank()
	if s.IsFree && s.NextRank > s.Rank {
		if s.NextRank == 2 || s.NextRank == 4 || s.NextRank == 6 {
			s.FreeSpinParams.FreeNum += 2
		}
	}

	// 合并相同name的winLine
	m.Spin.WinLineMergeSame()
}

func getTagsMaxLenByName(tagList [][]*base.Tag) map[string]int {
	var res = map[string]int{}
	if len(tagList) == 0 {
		return res
	}
	for _, tags := range tagList {
		name := GetTagsName(tags)
		if name == "" {
			continue
		}
		res[name] = mathutil.Max(res[name], len(tags))
	}
	return res
}

// GetTagsName 获取该组标签排除百搭标签的名称
func GetTagsName(tags []*base.Tag) string {
	return base.GetTagsName(tags, enum.SlotWild1)
}
