package unit5

import (
	"fmt"
	"github.com/shopspring/decimal"
	"slot-server/global"
	"slot-server/service/slot/base"
	"slot-server/service/slot/component"
	"slot-server/service/slot/eliminate"
)

// Machine 划线 + 特殊免费玩
type Machine struct {
	Spin     *component.Spin   `json:"-"`
	Spins    []*component.Spin `json:"-"`
	BaseSpin *component.Spin   `json:"-"`
	Table    *eliminate.Table  `json:"-"`
}

func NewMachine(spin *component.Spin) *Machine {
	return &Machine{Spin: &component.Spin{},
		Spins:    []*component.Spin{},
		BaseSpin: spin}
}

func (m *Machine) GetSpin() *component.Spin {
	return m.Spin
}

func (m *Machine) NewSpin(isFree bool, getFreeNum int, id, parentId int) *component.Spin {
	newSpin := &component.Spin{}
	Option := *m.BaseSpin.Options
	newSpin.Options = &Option
	newSpin.IsFree = isFree
	newSpin.Bet = m.BaseSpin.Bet
	newSpin.FreeSpinParams = component.FreeSpinParams{
		FreeNum: getFreeNum,
	}
	newSpin.Table = m.Table.Copy()
	newSpin.Gain = int(decimal.NewFromFloat(m.Table.Mul).Mul(decimal.NewFromInt(int64(m.BaseSpin.Bet))).IntPart())
	newSpin.Id = id
	newSpin.ParentId = parentId
	return newSpin
}

func (m *Machine) Exec() error {
	m.Table = component.NewGraph(m.BaseSpin, false)
	m.PlayGame(0, false)
	scatterNum := len(m.Table.QueryTags("scatter"))
	datum := 0
	if scatterNum >= 3 {
		datum = m.BaseSpin.Config.Event.M[17].(*base.ChangeTableEvent).Weight[scatterNum-3]
	}

	m.Spin = m.NewSpin(false, datum, 0, 0)
	m.Spin.Config = m.BaseSpin.Config
	if m.BaseSpin.IsMustFree {
		m.Spin.BuyFreeCoin = decimal.NewFromInt(int64(m.BaseSpin.Bet)).Mul(decimal.NewFromFloat(m.BaseSpin.Config.BuyFee)).IntPart()
	}
	verify := base.Verify{
		Count: 0,
	}
	for i := 0; i < datum; i++ {
		if verify.Count >= 1000 {
			//global.GVA_LOG.Error("免费玩次数超过1000次")
			return fmt.Errorf("免费玩次数超过1000次")
		}
		m.FreeSpinExec(&verify, 0)
	}
	return nil
}

func (m *Machine) FreeSpinExec(verify *base.Verify, parentId int) {
	verify.Count++
	m.Table = component.NewGraph(m.BaseSpin, true)
	m.PlayGame(0, true)
	scatterNum := len(m.Table.QueryTags("scatter"))

	datum := 0
	if scatterNum >= 3 {
		datum = m.BaseSpin.Config.Event.M[21].(*base.ChangeTableEvent).Weight[scatterNum-3]
	}
	m.Spins = append(m.Spins, m.NewSpin(true, datum, verify.Count, parentId))
	parentId = verify.Count
	for i := 0; i < datum; i++ {
		if verify.Count >= 1000 {
			global.GVA_LOG.Error("免费玩次数超过1000次")
			return
		}
		m.FreeSpinExec(verify, parentId)
	}
}

func (m *Machine) GetInitData() {

}

func (m *Machine) GetResData() {}

func (m *Machine) SumGain() {}

func (m *Machine) GetSpins() []*component.Spin {
	return m.Spins
}

func (m *Machine) PlayGame(count int, isFree bool) {
	if count > 10000 {
		global.GVA_LOG.Error("超过最大次数10000")
		return
	}

	AlterFlow := base.NewAlterFlow(count)

	//获取当前 要消除的划线集合
	lineList := m.Table.FindAllLine()
	//
	AlterFlow.AddList = m.Table.Fill(lineList, isFree, false)

	AlterFlow.InitList = m.Table.GetGraph()
	AlterFlow.FlowMap += m.Table.PrintTable("初始")

	//lineList 和赢钱的组合去匹配 ,返回集合中 每种要消除划线的集合 和对应的倍率
	AlterFlow.RemoveList, AlterFlow.SumMul = m.Table.WinMatchList(m.Table.EliminatedTagNameSetEmpty(m.Table.FindAllLine()), isFree)
	AlterFlow.Gain = int(decimal.NewFromFloat(AlterFlow.SumMul).Mul(decimal.NewFromInt(int64(m.BaseSpin.Bet))).IntPart())
	m.Table.AddMul(AlterFlow.SumMul)

	AlterFlow.FlowMap += m.Table.PrintTable("消除")

	AlterFlow.AfterDrop = m.Table.Drop()
	AlterFlow.FlowMap += m.Table.PrintTable("下落")

	m.Table.AlterFlows = append(m.Table.AlterFlows, AlterFlow)

	if len(m.Table.NeedFill()) > 0 {
		m.PlayGame(count+1, isFree)
	}
}
