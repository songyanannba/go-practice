package unit7

import (
	"fmt"
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
	m.Table = component.NewGraph(m.Spin, false) //7
	m.Table.FillSkillTags()                     //背景技能坐标
	m.Table.LeveEvent()                         //初始化 每个等级的倍率和标签个数

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

	// 1 创建本次流程的结构
	AlterFlow := base.NewAlterFlow(count) //7
	rand := helper.RandInt(3)
	isStop := rand == 0
	//2 AddList 本次是否有空标签 如果存在空标签
	//就去填充(根基倍率 填充有两个情况) 可以消除的情况 和 不可以消除的情况;
	AlterFlow.AddList = m.Table.Fill(m.Table.FindAllLine(), false, helper.If(!m.Level.IsEmpty(), helper.If(reSpin >= 2, true, isStop), false)) //7
	//isSkill := false
	//3 是否有技能 当没有空标签 但是有技能可以进入
	if len(AlterFlow.AddList) == 0 && len(m.Table.SkillSchedule) > 0 {
		//如果没有划线集合 且 重试次数大于0 重试逻辑
		//使用技能 改变table里面的标签
		schedule := m.Table.SkillSchedule[0] //取出第一个技能
		//AlterFlow.SkillID = schedule
		AlterFlow.EmitList = m.SkillSpin(schedule)
		//AlterFlow.EmitList = m.ReSpin(m.Level.DeQueue())
		//isSkill = true
	}

	//copy 初始表格 然后打印
	AlterFlow.InitList = m.Table.GetGraph()
	AlterFlow.FlowMap += m.Table.PrintTable("初始")

	//4 lineList 和赢钱的组合去匹配 , 返回集合中 每种要消除划线的集合 和对应的倍率
	//4.1 如果存在可以消除的组合 组合的名字赋值为空
	AlterFlow.RemoveList, AlterFlow.SumMul = m.Table.WinMatchList(m.Table.EliminatedTagNameSetEmpty(m.Table.FindAllLine()), false)
	// 匹配技能点
	for _, removeList := range AlterFlow.RemoveList {
		m.Table.RmCount += len(removeList.RemoveList)       //table总消除
		AlterFlow.RemoveCount += len(removeList.RemoveList) //单次流程消除
		//检查需要消除的标签 背后是否存在技能点 如果存在 记录到新的技能切片中 相当于队列
		for _, rl := range removeList.RemoveList {
			if _, ok := m.Table.SkillTags[[2]int{rl.X, rl.Y}]; ok {
				m.Table.SkillSchedule = append(m.Table.SkillSchedule, m.Table.SkillTags[[2]int{rl.X, rl.Y}])
			}
		}
	}

	AlterFlow.RankId = rank

	//rank = m.RankUp7(rank, m.Table.RmCount) //消除一定标签数量进入 下一个等级
	//if rank > AlterFlow.RankId {
	//	reSpin = reSpin + (rank - AlterFlow.RankId)
	//}

	AlterFlow.Gain = int(decimal.NewFromFloat(AlterFlow.SumMul).Mul(decimal.NewFromInt(int64(m.Spin.Bet))).IntPart())
	m.Table.AddMul(AlterFlow.SumMul) //7

	AlterFlow.FlowMap += m.Table.PrintTable("消除")
	AlterFlow.AfterDrop = m.Table.DropExistWild() //空标签往上走
	AlterFlow.FlowMap += m.Table.PrintTable("下落")

	m.Table.AlterFlows = append(m.Table.AlterFlows, AlterFlow)

	ss := m.Table.NeedFill()
	if len(ss) > 0 || len(m.Table.SkillSchedule) > 0 {
		m.PlayGame(count+1, rank, reSpin)
	} else {
		//既没有可填充的标签 也没有技能
		//检测历史消除的个数 和 等级规则进行匹配 如果符合规则 进入reSpin
		//CompareRank()
	}

}

func (m *Machine) CompareRank(NowRank int, rmTags int) (rank int) {

	if rmTags >= 114 && NowRank < enum.Rand1 {
		rank = enum.Rand1
	}

	if rmTags >= 116 && NowRank < enum.Rand2 {
		rank = enum.Rand2
	}

	if rmTags >= 120 && NowRank < enum.Rand3 {
		rank = enum.Rand3
	}

	if rmTags >= 125 && NowRank < enum.Rand4 {
		rank = enum.Rand4
	}

	if rmTags >= 132 && NowRank < enum.Rand5 {
		rank = enum.Rand5
	}

	return rank
}

func (m *Machine) SkillSpin(schedule int) [][]*base.Tag {
	//schedule := m.Table.SkillSchedule[0] //取出第一个技能
	m.Table.SkillSchedule = m.Table.SkillSchedule[1:]
	fmt.Println("schedule = ", schedule) //技能
	var emitList [][]*base.Tag

	switch schedule {
	case 1: //技能1：使用技能点随机选取一个标签，依次标签为中心点向上、下、左、右四个方向扩散标签至棋盘边缘，每个方向上的标签相同，标签种类随机
		m.Table.Skill1()
		emitList = m.Skill1()

	case 2: //技能2：随机生成一定个数wild标签
		m.Table.Skill2()
		rid := helper.RandInt(2) + 1
		m.Table.Skill2Count = rid
		//根据随机的点数 生成几个wild 并随机放到表格中
		emitList = m.FillWildTags(rid)

	case 3: //技能3：随机生成几组2*2的标签
		rid := helper.RandInt(2) + 1
		m.Table.Skill3()
		emitList = m.Skill3(rid)

	case 4: //技能4：随机生成3*3、4*4的标签
		m.Table.Skill4()
		m.Skill4Standby()

	case 5: //技能5：随机生成一定个数wild标签（比技能2数量多）
		m.Table.Skill5()
		rid := helper.RandInt(2) + 2
		if rid <= m.Table.Skill2Count {
			rid = m.Table.Skill2Count + 1
		}
		emitList = m.FillWildTags(rid)
	default:
		fmt.Println("SkillSpin err")
	}

	return emitList

}

func (m *Machine) Skill1() [][]*base.Tag {
	var wtags []*base.Tag
	var emitList [][]*base.Tag

	//把表 变成一维 并获取其中的一个坐标点
	arrs := helper.ListToArr(m.Table.TagList)
	arrsRand := helper.RandInt(len(arrs))
	tag := arrs[arrsRand]

	//依此标签 为中心点向上、下、左、右四个方向扩散标签至棋盘边缘，每个方向上的标签相同，标签种类随机
	fillTagUp := m.Table.Tags[helper.RandInt(len(m.Table.Tags))]
	fillTagDown := m.Table.Tags[helper.RandInt(len(m.Table.Tags))]

	//假设是wild 重新赋值 普通标签
	if fillTagUp.IsWild {
		fillTagUp = m.Table.NameGetTag("high_1")
	}
	if fillTagDown.IsWild {
		fillTagUp = m.Table.NameGetTag("low_1")
	}

	for x := 0; x < m.Table.Row; x++ {
		if x == tag.X {
			continue
		}
		if x < tag.X { //上边
			fTag := m.Table.TagList[x][tag.Y]
			fillTagU := fillTagUp.Copy()
			fillTagU.X = fTag.X
			fillTagU.Y = fTag.Y
			wtags = append(wtags, fillTagU)
			m.Table.TagList[x][tag.Y] = fillTagU
		} else { //下边
			fTag := m.Table.TagList[x][tag.Y]
			fillTagD := fillTagDown.Copy()
			fillTagD.X = fTag.X
			fillTagD.Y = fTag.Y
			wtags = append(wtags, fillTagD)
			m.Table.TagList[x][tag.Y] = fillTagD
		}
	}

	//左边 和 右边
	fillTagYR := m.Table.Tags[helper.RandInt(len(m.Table.Tags))]
	fillTagYl := m.Table.Tags[helper.RandInt(len(m.Table.Tags))]
	//假设是wild 重新赋值 普通标签
	if fillTagYR.IsWild {
		fillTagYR = m.Table.NameGetTag("high_2")
	}
	if fillTagYl.IsWild {
		fillTagYl = m.Table.NameGetTag("low_2")
	}

	for y := 0; y < m.Table.Col; y++ {
		if y == tag.Y {
			continue
		}
		fTag := m.Table.TagList[tag.X][y]
		if y < tag.Y { //左边
			fillTagL := fillTagYl.Copy()
			fillTagL.X = fTag.X
			fillTagL.Y = fTag.Y
			wtags = append(wtags, fillTagL)
			m.Table.TagList[tag.X][y] = fillTagL
		} else { //右边
			fillTagR := fillTagYR.Copy()
			fillTagR.X = fTag.X
			fillTagR.Y = fTag.Y
			wtags = append(wtags, fillTagR)
			m.Table.TagList[tag.X][y] = fillTagR
		}
	}
	emitList = append(emitList, wtags)
	return emitList
}

func (m *Machine) Skill4Standby() [][]*base.Tag {

	intss := [][2]int{[2]int{1, 1}, [2]int{1, 6}, [2]int{6, 1}, [2]int{6, 6}}

	var coreList []*base.Tag //中心点
	var emitList [][]*base.Tag
	//取指定数量的核心
	for i := 0; i < 2; i++ {
		randInt := helper.RandInt(len(intss))
		ints := intss[randInt]
		intss = append(intss[:randInt], intss[randInt+1:]...)
		coreList = append(coreList, m.Table.TagList[ints[0]][ints[1]].Copy())
	}
	if len(coreList) == 0 {
		return emitList
	}

	for k, tag := range coreList {
		var emits []*base.Tag
		var adjacent []*base.Tag
		if k == 0 {
			adjacent = m.Table.GetAllAdjacent(tag.X, tag.Y) //7
		} else {
			adjacent = m.Table.GetAllAdjacent(tag.X, tag.Y) // 7
			if tag.X == 1 && tag.Y == 1 {
				adjacent = append(adjacent, m.Table.TagList[tag.X][tag.Y+2].Copy())
				adjacent = append(adjacent, m.Table.TagList[tag.X-1][tag.Y+2].Copy())
				adjacent = append(adjacent, m.Table.TagList[tag.X+1][tag.Y+2].Copy())
				adjacent = append(adjacent, m.Table.TagList[tag.X+2][tag.Y+2].Copy())

				adjacent = append(adjacent, m.Table.TagList[tag.X+2][tag.Y+1].Copy())
				adjacent = append(adjacent, m.Table.TagList[tag.X+2][tag.Y-1].Copy())
				adjacent = append(adjacent, m.Table.TagList[tag.X+2][tag.Y].Copy())
			}
		}

		//上下左右
		mayFillTags := lo.Filter(adjacent, func(t *base.Tag, i int) bool {
			return !t.IsWild
		})

		for i := 0; i < len(mayFillTags); i++ {
			fillTag := mayFillTags[i].Copy()
			emits = append(emits, fillTag)
		}
		emitList = append(emitList, emits)
		for _, emit := range emits {
			fillTag := tag.Copy()
			fillTag.X = emit.X
			fillTag.Y = emit.Y
			m.Table.TagList[emit.X][emit.Y] = fillTag.Copy()
		}
	}

	return emitList
}

func (m *Machine) Skill3(rid int) [][]*base.Tag {
	var wtags []*base.Tag
	var allAdjacent []*base.Tag
	var emitList [][]*base.Tag
	arrs := helper.ListToArr(m.Table.TagList) //把表 变成一维

	for i := 0; i < rid; i++ {
		arrsRand := helper.RandInt(len(arrs))
		zy := helper.RandInt(2) //随机方向  0 左 1 右

		//1 获取一个点
		tag := arrs[arrsRand]
		wtags = append(wtags, tag)
		//2 获取周围2*2的点

		tagX1 := &base.Tag{}
		if tag.X == 0 { //只能向右
			tagX1 = m.Table.TagList[tag.X+1][tag.Y]
		} else if tag.X == 7 { //只能向左
			tagX1 = m.Table.TagList[tag.X-1][tag.Y]
		} else if zy == 1 {
			tagX1 = m.Table.TagList[tag.X+1][tag.Y]
		} else {
			tagX1 = m.Table.TagList[tag.X-1][tag.Y]
		}
		wtags = append(wtags, tagX1)

		//tag 和 tag1 确定了 x 方向
		//确定 y 方向
		tagY := &base.Tag{}
		tag1Y := &base.Tag{}

		if tag.Y == 0 {
			tagY = m.Table.TagList[tag.X][tag.Y+1]
			tag1Y = m.Table.TagList[tagX1.X][tag.Y+1]
		} else if tag.Y == 7 {
			tagY = m.Table.TagList[tag.X][tag.Y-1]
			tag1Y = m.Table.TagList[tagX1.X][tag.Y-1]
		} else if zy == 1 {
			tagY = m.Table.TagList[tag.X][tag.Y+1]
			tag1Y = m.Table.TagList[tagX1.X][tag.Y+1]
		} else {
			tagY = m.Table.TagList[tag.X][tag.Y-1]
			tag1Y = m.Table.TagList[tagX1.X][tag.Y-1]
		}
		wtags = append(wtags, tagY)
		wtags = append(wtags, tag1Y)

		for _, wt := range wtags {
			biasAdjacent := m.Table.GetAllAdjacent(wt.X, wt.Y)
			allAdjacent = append(allAdjacent, biasAdjacent...)
		}

		//重新填充arrs
		arrs := []*base.Tag{}
		for _, mtVals := range m.Table.TagList {
			for _, mtVal := range mtVals {
				for _, aa := range allAdjacent {
					if aa.X != mtVal.X && aa.Y != mtVal.Y {
						arrs = append(arrs, mtVal)
					}
				}
			}
		}

		emitList = append(emitList, wtags)
	}

	return emitList
}

func (m *Machine) FillWildTags(rid int) [][]*base.Tag {
	var wtags []*base.Tag
	var emitList [][]*base.Tag
	arrs := helper.ListToArr(m.Table.TagList) //把表 变成一维

	for i := 0; i < rid; i++ {
		arrsRand := helper.RandInt(len(arrs))
		tag := arrs[arrsRand]

		//给表格填充 wild 标签
		wildTga := m.Table.NameGetTag(enum.SlotWild)
		wildTga.X = tag.X
		wildTga.Y = tag.Y
		m.Table.TagList[tag.X][tag.Y] = wildTga.Copy()
		wtags = append(wtags, tag)

		arrs = append(arrs[:arrsRand], arrs[arrsRand:]...)
	}
	for _, v := range wtags {
		var emits []*base.Tag
		emits = append(emits, v)
		emitList = append(emitList, emits)
	}
	return emitList
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

func (m *Machine) GetCoreList(n int) []*base.Tag {
	var coreList []*base.Tag //中心点

	arrs := lo.Filter(helper.ListToArr(m.Table.TagList), func(tag1 *base.Tag, i int) bool {
		return tag1.X != 0 || tag1.Y != 0 || tag1.X != 7 || tag1.Y != 7
	})

	for i := 0; i < n; i++ {
		arrsRand := helper.RandInt(len(arrs))

		//1 随机获取一个点
		tag := arrs[arrsRand]
		coreList = append(coreList, tag.Copy())
		arrs = append(arrs[:arrsRand], arrs[arrsRand:]...)
	}

	return coreList
}

// analyzeCoreTags 一个中心点 周围 4*4 所有的可能
func (m *Machine) analyzeCoreTags(coreTag *base.Tag) []*base.Tag {

	var tagsList44 []*base.Tag
	var adjacentTags []*base.Tag

	// 这个中心点 3*3 所占的区域是
	adjacents := m.Table.GetAllAdjacent(coreTag.X, coreTag.Y)

	// 获取 4*4 所有的可能
	for _, adjacent := range adjacents {
		allAdjacents := m.Table.GetAllAdjacent(adjacent.X, adjacent.Y)
		tagsList44 = append(tagsList44, allAdjacents...)
	}
	//tagsList44 去重
	tagsList44Map := make(map[[2]int]*base.Tag)
	for _, val := range tagsList44 {
		if _, ok := tagsList44Map[[2]int{val.X, val.Y}]; !ok {
			tagsList44Map[[2]int{val.X, val.Y}] = val
		}
	}

	for _, ta := range tagsList44Map {
		adjacentTags = append(adjacentTags, ta)
	}

	return adjacentTags
}

//func (m *Machine) PlayGame(count, rank int, reSpin int) {
//	if count > enum.SlotMaxSpinNum {
//		global.GVA_LOG.Error(enum.SlotMaxSpinStr)
//		return
//	}
//
//	// 1 创建本次流程的结构
//	AlterFlow := base.NewAlterFlow(count) //7
//	rand := helper.RandInt(3)
//	isStop := rand == 0
//	//2 AddList 本次是否有空标签 如果存在空标签
//	//就去填充(根基倍率 填充有两个情况) 可以消除的情况 和 不可以消除的情况;
//	AlterFlow.AddList = m.Table.Fill(m.Table.FindAllLine(), false, helper.If(!m.Level.IsEmpty(), helper.If(reSpin >= 2, true, isStop), false)) //7
//
//	//3 是否进入 reSpin
//	//没有 空标签 并且有技能 可以进入
//	if len(AlterFlow.AddList) == 0 && len(m.Table.SkillSchedule) > 0 {
//		//如果没有划线集合 且 重试次数大于0 重试逻辑
//		m.ReSpin7()
//		AlterFlow.EmitList = m.ReSpin(m.Level.DeQueue())
//	}
//
//	AlterFlow.InitList = m.Table.GetGraph()
//	AlterFlow.FlowMap += m.Table.PrintTable("初始")
//
//	//4 lineList 和赢钱的组合去匹配 , 返回集合中 每种要消除划线的集合 和对应的倍率
//	//4.1 如果存在可以消除的组合 组合的名字赋值为空
//	AlterFlow.RemoveList, AlterFlow.SumMul = m.Table.WinMatchList(m.Table.EliminatedTagNameSetEmpty(m.Table.FindAllLine()), false)
//	// 匹配技能点
//	for _, removeList := range AlterFlow.RemoveList {
//		m.Table.RmCount += len(removeList.RemoveList)
//		//检查需要消除的标签 背后是否存在技能点 如果存在 记录到新的技能切片中 相当于队列
//		for _, rl := range removeList.RemoveList {
//			if _, ok := m.Table.SkillTags[[2]int{rl.X, rl.Y}]; ok {
//				m.Table.SkillSchedule = append(m.Table.SkillSchedule, m.Table.SkillTags[[2]int{rl.X, rl.Y}])
//			}
//		}
//	}
//	AlterFlow.RemoveCount = m.Table.RmCount
//	AlterFlow.RankId = rank
//	rank = m.RankUp7(rank, m.Table.RmCount) //消除一定标签数量进入 ReSpin 玩法
//	if rank > AlterFlow.RankId {            //
//		reSpin++
//	}
//
//	AlterFlow.Gain = int(decimal.NewFromFloat(AlterFlow.SumMul).Mul(decimal.NewFromInt(int64(m.Spin.Bet))).IntPart())
//	m.Table.AddMul(AlterFlow.SumMul) //7
//
//	AlterFlow.FlowMap += m.Table.PrintTable("消除")
//	AlterFlow.AfterDrop = m.Table.Drop6() //空标签往上走
//	AlterFlow.FlowMap += m.Table.PrintTable("下落")
//
//	m.Table.AlterFlows = append(m.Table.AlterFlows, AlterFlow)
//
//	ss := m.Table.NeedFill()
//	if len(ss) > 0 || len(m.Table.SkillSchedule) > 0 {
//		m.PlayGame(count+1, rank, reSpin)
//	}
//}
