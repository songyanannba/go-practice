package unit8

import (
	"fmt"
	"github.com/shopspring/decimal"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/service/slot/base"
	"slot-server/service/slot/component"
	"slot-server/service/slot/eliminate"
	"slot-server/service/slot/template"
	"slot-server/service/slot/template/flow"
	"slot-server/utils/helper"
)

// Machine 划线 + 特殊免费玩
type Machine struct {
	Spin        *component.Spin    `json:"-"`
	Spins       []*component.Spin  `json:"-"`
	SpinInfo    *template.SpinInfo `json:"-"`
	BaseSpin    *component.Spin    `json:"-"`
	multipliers []*base.Tag
	config      *template.Config
	Table       *eliminate.Table
}

func (m *Machine) GetInitData() {}

func (m *Machine) GetResData() {}

func (m *Machine) SumGain() {}

func (m *Machine) GetSpin() *component.Spin {
	return m.Spin
}

func (m *Machine) GetSpins() []*component.Spin {
	return m.Spins
}

func NewMachine(s *component.Spin) *Machine {
	config := &template.Config{
		Event:        s.Config.Event,
		SlotId:       int(s.Config.SlotId),
		PayTableList: s.Config.PayTableList,
		Template:     map[int][]*base.Tag{},
		Col:          s.Config.Index,
		Row:          s.Config.Row,
		TagMapById:   map[int]*base.Tag{},
		TagMapByName: map[string]*base.Tag{},
	}
	config.InitConfig(s.Config.GetAllTagQuote())

	return &Machine{Spin: &component.Spin{},
		Spins:    []*component.Spin{},
		BaseSpin: s,
		config:   config,
	}
}

func (m *Machine) NewSpin(isFree bool, getFreeNum int, id, parentId int) *component.Spin {
	newSpin := &component.Spin{}
	Option := &component.Options{
		IsTest: m.BaseSpin.Options.IsTest,
		IsFree: isFree,
	}
	newSpin.Options = Option
	newSpin.Bet = m.BaseSpin.Bet

	newSpin.FreeSpinParams = component.FreeSpinParams{
		FreeNum: getFreeNum,
	}
	newSpin.SpinInfo = m.SpinInfo.Copy()

	newSpin.Gain = m.SpinInfo.GetWin(m.BaseSpin.Bet)
	newSpin.Id = id
	newSpin.ParentId = parentId
	newSpin.Config = m.BaseSpin.Config
	if m.BaseSpin.IsMustFree && id == 0 {
		m.Spin.BuyFreeCoin = decimal.NewFromInt(int64(m.BaseSpin.Bet)).Mul(decimal.NewFromFloat(m.BaseSpin.Config.BuyFee)).IntPart()
	}
	return newSpin
}

func (m *Machine) Exec() error {

	m.SpinInfo = m.NewGameInfo(enum.NormalSpin)

	err := m.PlayGame(0)
	if err != nil {
		return err
	} //根据模版去消除数据的玩法

	scatters := m.SpinInfo.FindTagsByName(enum.ScatterName)
	scatterLine := m.SpinInfo.GetWinLine(scatters)
	m.SpinInfo.Scatter = scatterLine

	datum := 0
	if len(scatters) >= enum.TriggerFreeGet {
		datum = enum.FreeGetNum
	}

	multipliers := m.SpinInfo.FindTagsByName(enum.MultiplierName)

	m.SpinInfo.Multiplier = multipliers

	m.Spin = m.NewSpin(false, datum, 0, 0)
	verify := base.Verify{
		Count: 0,
	}
	for i := 0; i < datum; i++ {
		if verify.Count >= enum.SlotMaxSpinNum {
			//global.GVA_LOG.Error(enum.SlotMaxFreeSpinErr)
			return fmt.Errorf(enum.SlotMaxFreeSpinErr)
		}
		err := m.FreeSpinExec(&verify, 0)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Machine) FreeSpinExec(verify *base.Verify, parentId int) error {
	verify.Count++

	m.SpinInfo = m.NewGameInfo(enum.FreeSpin)

	err := m.PlayGame(0)
	if err != nil {
		return err
	}

	scatters := m.SpinInfo.FindTagsByName(enum.ScatterName)
	scatterLine := m.SpinInfo.GetWinLine(scatters)
	m.SpinInfo.Scatter = scatterLine

	datum := 0
	if len(scatters) >= enum.TriggerFreeAdd {
		datum = enum.FreeAddNum
	}

	multipliers := m.SpinInfo.FindTagsByName(enum.MultiplierName)
	m.multipliers = append(m.multipliers, multipliers...)
	m.SpinInfo.Multiplier = helper.CopyList(m.multipliers)

	m.Spins = append(m.Spins, m.NewSpin(true, datum, 0, 0))
	for i := 0; i < datum; i++ {
		if verify.Count >= enum.SlotMaxSpinNum {
			//global.GVA_LOG.Error(enum.SlotMaxFreeSpinErr)
			return fmt.Errorf(enum.SlotMaxFreeSpinErr)
		}
		err := m.FreeSpinExec(verify, 0)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Machine) PlayGame(count int) error {

	SpinFlow := flow.NewSpinFlow(count)
	SpinFlow.FlowMap += m.SpinInfo.PrintTable("初始")
	SpinFlow.InitList = helper.CopyListArr(m.SpinInfo.Display)
	//获取划线
	lines := m.SpinInfo.FindCountLine(enum.LineLength)

	//获取划线赢钱
	SpinFlow.AddOmitList(m.SpinInfo.GetWinLines(lines)...)
	//删除标签
	m.SpinInfo.DeleteTagList(lines)
	SpinFlow.FlowMap += m.SpinInfo.PrintTable("删除")
	//掉落标签
	m.SpinInfo.Drop()
	SpinFlow.FlowMap += m.SpinInfo.PrintTable("掉落")
	if len(m.SpinInfo.GetEmptyTags()) > 0 {
		global.GVA_LOG.Error("模版缺失还有标签没有填充")
	}
	m.SpinInfo.SpinFlow = append(m.SpinInfo.SpinFlow, SpinFlow)
	lines = m.SpinInfo.FindCountLine(enum.LineLength)
	if len(lines) == 0 {
		return nil
	}
	if count > enum.SlotMaxSpinNum {
		//global.GVA_LOG.Error(enum.SlotMaxSpinStr)
		return fmt.Errorf(enum.SlotMaxSpinStr)
	}
	return m.PlayGame(count + 1)
}

func (m *Machine) NewGameInfo(spinType int) *template.SpinInfo {
	config := &template.Config{
		Event:        m.config.Event,
		SlotId:       m.config.SlotId,
		PayTableList: m.config.PayTableList,
		Template:     m.BaseSpin.Config.Template[spinType],
		Col:          m.config.Col,
		Row:          m.config.Row,
		GameType:     spinType,
		TagMapById:   m.config.TagMapById,
		TagMapByName: m.config.TagMapByName,
	}
	info := template.NewGameInfo(config)
	return info
}

//模版测试逻辑
