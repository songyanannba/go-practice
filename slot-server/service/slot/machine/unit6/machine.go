package unit6

import (
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"math/rand"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/service/slot/base"
	"slot-server/service/slot/component"
	"slot-server/service/slot/eliminate"
	"slot-server/utils/helper"
)

// Machine 划线 + 特殊免费玩
type Machine struct {
	Spin  *component.Spin  `json:"-"`
	Table *eliminate.Table `json:"-"`
	Level component.Queue  `json:"-"`
}

func NewMachine(spin *component.Spin) *Machine {
	return &Machine{Spin: spin}
}

func (m *Machine) GetSpin() *component.Spin {
	return m.Spin
}

func (m *Machine) Exec() error {
	m.Table = component.NewGraph(m.Spin, false)
	m.PlayGame(0, 0, 0)

	m.Spin.Table = m.Table.Copy()
	m.Spin.Gain = int(decimal.NewFromFloat(m.Table.Mul).Mul(decimal.NewFromInt(int64(m.Spin.Bet))).IntPart())
	return nil
}

func (m *Machine) GetInitData() {

}

func (m *Machine) GetResData() {}

func (m *Machine) SumGain() {}

func (m *Machine) GetSpins() []*component.Spin {
	return []*component.Spin{}
}

func (m *Machine) PlayGame(count, rank int, reSpin int) {
	if count > enum.SlotMaxSpinNum {
		global.GVA_LOG.Error(enum.SlotMaxSpinStr)
		return
	}
	if !m.Level.IsEmpty() {
		reSpin++
	} else {
		reSpin = 0
	}
	AlterFlow := base.NewAlterFlow(count)

	//获取当前 要消除的划线集合
	lineList := m.Table.FindAllLine() //6
	//
	rand := helper.RandInt(3)
	isStop := rand == 0
	AlterFlow.AddList = m.Table.Fill(lineList, false, helper.If(!m.Level.IsEmpty(), helper.If(reSpin >= 2, true, isStop), false)) //6
	isSkill := false
	if len(AlterFlow.AddList) == 0 && !m.Level.IsEmpty() {
		//如果没有划线集合 且 重试次数大于0 重试逻辑
		AlterFlow.EmitList = m.ReSpin(m.Level.DeQueue())
		isSkill = true
	}

	AlterFlow.InitList = m.Table.GetGraph()
	AlterFlow.FlowMap += m.Table.PrintTable("初始")

	//lineList 和赢钱的组合去匹配 ,返回集合中 每种要消除划线的集合 和对应的倍率
	AlterFlow.RemoveList, AlterFlow.SumMul = m.Table.WinMatchList(m.Table.EliminatedTagNameSetEmpty(m.Table.FindAllLine()), false)
	for _, e := range AlterFlow.RemoveList {
		m.Table.RmCount += len(e.RemoveList)
	}
	AlterFlow.RemoveCount = m.Table.RmCount
	AlterFlow.RankId = rank
	rank = m.RankUp(rank, m.Table.RmCount)

	AlterFlow.Gain = int(decimal.NewFromFloat(AlterFlow.SumMul).Mul(decimal.NewFromInt(int64(m.Spin.Bet))).IntPart())
	m.Table.AddMul(AlterFlow.SumMul) //6
	if isSkill {
		m.Table.AddSkillMul(AlterFlow.SumMul)
	}

	AlterFlow.FlowMap += m.Table.PrintTable("消除")
	m.WildBeat()
	AlterFlow.AfterDrop = m.Table.DropExistWild()
	AlterFlow.FlowMap += m.Table.PrintTable("下落")

	m.Table.AlterFlows = append(m.Table.AlterFlows, AlterFlow)
	//fmt.Print(AlterFlow.FlowMap)
	ss := m.Table.NeedFill()
	if len(ss) > 0 || !m.Level.IsEmpty() {
		m.PlayGame(count+1, rank, reSpin)
	}
}

func (m *Machine) RankUp(NowRank int, rmTags int) (rank int) {
	rank1 := m.Spin.Config.Event.M[16].(*base.Unit6LevelEvent)
	rank2 := m.Spin.Config.Event.M[17].(*base.Unit6LevelEvent)
	rank3 := m.Spin.Config.Event.M[18].(*base.Unit6LevelEvent)
	rank4 := m.Spin.Config.Event.M[19].(*base.Unit6LevelEvent)
	rank5 := m.Spin.Config.Event.M[20].(*base.Unit6LevelEvent)
	rank = NowRank

	if rmTags >= rank1.Collect && NowRank < enum.Rand1 {
		m.Level.EnQueue(rank1)
		rank = enum.Rand1
	}

	if rmTags >= rank2.Collect && NowRank < enum.Rand2 {
		m.Level.EnQueue(rank2)
		rank = enum.Rand2
	}

	if rmTags >= rank3.Collect && NowRank < enum.Rand3 {
		m.Level.EnQueue(rank3)
		rank = enum.Rand3
	}

	if rmTags >= rank4.Collect && NowRank < enum.Rand4 {
		m.Level.EnQueue(rank4)
		rank = enum.Rand4
	}

	if rmTags >= rank5.Collect && NowRank < enum.Rand5 {
		m.Level.EnQueue(rank5)
		rank = enum.Rand5
	}

	return rank

}

func (m *Machine) ReSpin(level *component.Level) [][]*base.Tag {
	intss := [][2]int{[2]int{1, 1}, [2]int{1, 4}, [2]int{4, 1}, [2]int{4, 4}}
	var coreList []*base.Tag //中心点
	var emitList [][]*base.Tag
	//取指定数量的核心
	for i := 0; i < level.CoreCount; i++ {
		randInt := helper.RandInt(len(intss))
		ints := intss[randInt]
		intss = append(intss[:randInt], intss[randInt+1:]...)
		coreList = append(coreList, m.Table.TagList[ints[0]][ints[1]].Copy())
	}
	if len(coreList) == 0 {
		return emitList
	}
	//m.Table.RankId++
	//填充wild

	//分配指定数量的扩散
	randoms := RandomAllocation(level.CoreCount, level.EmitCount)
	balance := 0
	for i, tag := range coreList {
		var emits []*base.Tag
		adjacent := m.Table.GetAdjacent(tag.X, tag.Y)
		num := randoms[i] + balance
		//上下左右
		mayFillTags := lo.Filter(adjacent, func(t *base.Tag, i int) bool {
			return !t.IsWild
		})
		for i := 0; num > 0 && i < len(mayFillTags); i++ {
			fillTag := mayFillTags[i].Copy()
			emits = append(emits, fillTag)
			num--
		}
		//斜对角
		if num > 0 {
			biasAdjacent := m.Table.GetBiasAdjacent(tag.X, tag.Y)
			mayFillTags = lo.Filter(biasAdjacent, func(t *base.Tag, i int) bool {
				return !t.IsWild
			})
			for i := 0; num > 0 && i < len(mayFillTags); i++ {
				fillTag := mayFillTags[i].Copy()
				emits = append(emits, fillTag)
				num--
			}
		}
		emitList = append(emitList, emits)
		//这里分配不完的部分
		balance += num
		for _, emit := range emits {
			fillTag := tag.Copy()
			fillTag.X = emit.X
			fillTag.Y = emit.Y
			m.Table.TagList[emit.X][emit.Y] = fillTag.Copy()
		}
	}
	//填充 wild 标签
	if level.WildMul > 0 {

		wildTga := m.Table.NameGetTag(enum.SlotWild)
		wildTga.Multiple = float64(level.WildMul)
		var mayFillInTags []*base.Tag  //中心点非边框tags
		var mayFillBoxTags []*base.Tag //中心点周围边框tags
		for _, emits := range emitList {

			rand.Shuffle(len(emits), func(i, j int) {
				emits[i], emits[j] = emits[j], emits[i]
			})
			for _, emit := range emits {
				//先过滤掉中心点和周围的wild
				mayFillTags := lo.Filter(m.Table.GetAdjacent(emit.X, emit.Y), func(t *base.Tag, i int) bool {
					//return !t.IsWild && t.Name != coreList[randCore].Name
					if t.IsWild {
						return false
					}
					for _, vv := range coreList {
						if t.Name == vv.Name {
							return false
						}
					}
					return true
				})

				//如果是不是四条边的标签 优先级高
				fillInTags := lo.Filter(mayFillTags, func(t *base.Tag, i int) bool {
					return t.X != 0 && t.Y != 0 && t.X != 5 && t.Y != 5
				})
				mayFillInTags = append(mayFillInTags, fillInTags...)

				//如果是四条边的标签 优先级低
				fillBoxTags := lo.Filter(mayFillTags, func(t *base.Tag, i int) bool {
					return t.X == 0 || t.Y == 0 || t.X == 5 || t.Y == 5
				})
				mayFillBoxTags = append(mayFillBoxTags, fillBoxTags...)
			}
		}

		if len(mayFillInTags) > 0 {
			fillTag := mayFillInTags[helper.RandInt(len(mayFillInTags))].Copy()
			wildTga.X = fillTag.X
			wildTga.Y = fillTag.Y
			m.Table.TagList[fillTag.X][fillTag.Y] = wildTga.Copy()
			return emitList
		}

		if len(mayFillBoxTags) > 0 {
			fillTag := mayFillBoxTags[helper.RandInt(len(mayFillBoxTags))].Copy()
			wildTga.X = fillTag.X
			wildTga.Y = fillTag.Y
			m.Table.TagList[fillTag.X][fillTag.Y] = wildTga.Copy()
			return emitList
		}

		canFillTag := lo.Filter(m.Table.ToArr(), func(item *base.Tag, i int) bool {
			if item.IsWild {
				return false
			}
			if item.X == 1 && item.Y == 1 { //排除中心点
				return false
			}
			if item.X == 1 && item.Y == 4 { //排除中心点
				return false
			}
			if item.X == 4 && item.Y == 1 { //排除中心点
				return false
			}
			if item.X == 4 && item.Y == 4 { //排除中心点
				return false
			}
			if item.X == 0 && item.Y == 0 && item.X == 5 && item.Y == 5 {
				return false
			}
			return true
		})
		if len(canFillTag) > 0 {
			fillTag := canFillTag[helper.RandInt(len(canFillTag))].Copy()
			wildTga.X = fillTag.X
			wildTga.Y = fillTag.Y
			m.Table.TagList[fillTag.X][fillTag.Y] = wildTga.Copy()
			return emitList
		}
	}

	return emitList
}

func RandomAllocation(n int, num int) []int {

	arr := make([]int, n)
	nm := enum.CoreTagMinNum //每个中心点 最少4个标签
	// 首先为数组的每个元素分配一个值1
	for i := 0; i < n; i++ {
		arr[i] = nm
		num = num - nm
	}

	// 随机选择一个元素并增加它的值，直到我们达到所需的总和
	for num > 0 {
		i := helper.RandInt(n)           // 随机选择一个元素的索引
		if arr[i] < enum.CoreTagMaxNum { // 确保元素的值不大于8
			arr[i]++
			num--
		}
	}
	return arr
}

// WildBeat 移动wild
func (m *Machine) WildBeat() {
	wildTags := lo.Filter(m.Table.ToArr(), func(t *base.Tag, i int) bool {
		return t.IsWild
	})
	//intss := [][2]int{[2]int{1, 1}, [2]int{1, 4}, [2]int{4, 1}, [2]int{4, 4}}
	for _, tag := range wildTags {
		//result1: wild标签周围为空 并且不是 中心点的标签
		result1 := lo.Filter(m.Table.GetUpDownLeftRight(tag.X, tag.Y), func(item *base.Tag, index int) bool {
			if item.Name != "" {
				return false
			}
			if item.X == 1 && item.Y == 1 { //排除中心点
				return false
			}
			if item.X == 1 && item.Y == 4 { //排除中心点
				return false
			}
			if item.X == 4 && item.Y == 1 { //排除中心点
				return false
			}
			if item.X == 4 && item.Y == 4 { //排除中心点
				return false
			}
			return true
		})
		if len(result1) > 0 {
			position := result1[helper.RandInt(len(result1))]

			//交换标签坐标
			m.Table.TagList[position.X][position.Y], m.Table.TagList[tag.X][tag.Y] = m.Table.TagList[tag.X][tag.Y], m.Table.TagList[position.X][position.Y]
			m.Table.TagList[position.X][position.Y].X, m.Table.TagList[position.X][position.Y].Y = position.X, position.Y
			m.Table.TagList[tag.X][tag.Y].X, m.Table.TagList[tag.X][tag.Y].Y = tag.X, tag.Y
		}
	}
}
