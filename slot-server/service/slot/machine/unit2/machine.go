package unit2

import (
	"github.com/shopspring/decimal"
	"slot-server/enum"
	"slot-server/service/slot/base"
	"slot-server/service/slot/component"
	"slot-server/utils/helper"
)

// Machine 雷神 多线slot
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
	if m.Spin.IsFree {
		if v, ok := m.Spin.Config.Event.Get(enum.SlotEvent1ChangeTable); ok && v != nil {
			m.Spin.FreeSpinParams.WildNum = v.(*base.ChangeTableEvent).Fetch()
			// 根据换表事件中的权重获取需要替换的wild数量
			if m.Spin.FreeSpinParams.WildNum <= 3 {
				m.Spin.FreeSpinParams.WildInterval = 1
			} else if m.Spin.FreeSpinParams.WildNum >= 7 {
				m.Spin.FreeSpinParams.WildInterval = 3
			} else {
				m.Spin.FreeSpinParams.WildInterval = 2
			}

			m.Spin.GetInitDataByReel(enum.SlotReelType2FS, m.Spin.FreeSpinParams.WildInterval)
			// 获取中间三列填充wild
			data := m.Spin.InitDataList
			middle := append(helper.SliceVal(data, 1), helper.SliceVal(data, 2)...)
			middle = append(middle, helper.SliceVal(data, 3)...)
			res := excludeScatter(middle...)
			wild1 := m.Spin.Config.GetTag(enum.SlotWild1)
			helper.SliceShuffle(res)
			for i, v := range res {
				v.IsWild = true
				v.Include = append(v.Include, wild1.Include...)
				if i+1 >= m.Spin.FreeSpinParams.WildNum {
					break
				}
			}
		}
	} else {
		m.Spin.GetInitDataByReel(enum.SlotReelType1Normal, 1)
	}
	return
}

func (m *Machine) GetResData() {
	m.Spin.GetSamePayline(enum.SameTagLen)
	return
}

func (m *Machine) payTableMatch() {
	// 每条划线可能匹配到多个payTable 取最大的payTable
	m.Spin.PayTableMatch()
}

func (m *Machine) SumGain() {
	m.payTableMatch()
	// 计算payTable的赢钱
	if len(m.Spin.PayTables) > 0 {
		var mulArr []float64
		for _, table := range m.Spin.PayTables {
			// 判断是否为免费转标签
			if GetTagsName(table.Tags) == enum.ScatterName {
				m.Spin.FreeSpinParams.Count++
			}
			mulArr = append(mulArr, table.Multiple)
		}
		m.Spin.FreeSpinParams.FreeNum = m.Spin.FreeSpinParams.Count * enum.Slot2FreeNum
		mul := helper.FloatSum(mulArr...)
		m.Spin.Gain = int(decimal.NewFromInt(int64(m.Spin.Bet)).Mul(decimal.NewFromFloat(mul)).IntPart())
	}

	// 合并相同name的winLine
	m.Spin.WinLineMergeSame()
}

// GetTagsName 获取该组标签排除百搭标签的名称
func GetTagsName(tags []*base.Tag) string {
	return base.GetTagsName(tags, enum.SlotWild1)
}

func excludeScatter(tags ...*base.Tag) []*base.Tag {
	var res []*base.Tag
	for _, tag := range tags {
		if tag.Name != enum.ScatterName && tag.Name != enum.SlotWild1 {
			res = append(res, tag)
		}
	}
	return res
}
