package eliminate

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"slot-server/global"
	"slot-server/service/slot/base"
	"slot-server/utils/helper"
	"strconv"
	"strings"
)

type Table struct {
	Row          int               // 行数
	Col          int               // 列数
	Tags         []*base.Tag       // 所有tag
	WildTags     []*base.Tag       // wild标签
	Scatter      *base.Tag         // scatter
	DoubleTag    *base.Tag         // double标签
	TagList      [][]*base.Tag     // 二维列表
	InitTable    [][]*base.Tag     // 初始表
	Target       *Target           // 目标
	Mul          float64           // 倍数
	SkillMul     float64           // 技能倍数
	PayTableList []*base.PayTable  // 所有的paytable
	AlterFlows   []*base.AlterFlow // 改变的流程

	RmCount int //移除的个数

	PayTableListMaps map[string]map[int]*base.PayTable //paytable 划线的map

	SlotId uint // 机器id
	RankId int

	TableWildPoint map[[2]int]*base.Tag //全局

	//模版生成属性
	ModelEvent *ModelEvent
	BaseRow    int // 游戏行数
	FlowsMap   string
	Scatters   []*base.Tag // 消散标签

	tagMap   map[string]*base.Tag // 所有标签的Map
	tagIdMap map[int]*base.Tag    // 所有标签的Map
	TagMul   float64              // 翻倍标签倍率记录

	SkillTags     map[[2]int]int //初始技能位置
	SkillSchedule []int          //技能条
	LevelEvents   []*LevelEvent
	Skill2Count   int
}

func (t *Table) SetTagMap(tags []*base.Tag) {
	t.tagIdMap = make(map[int]*base.Tag)
	t.tagMap = make(map[string]*base.Tag)
	for _, tag := range tags {
		t.tagMap[tag.Name] = tag
		t.tagIdMap[tag.Id] = tag
	}
}

func (t *Table) GetTagById(id int) *base.Tag {
	tag, ok := t.tagIdMap[id]
	if ok {
		return tag.Copy()
	}
	return &base.Tag{}
}

func (t *Table) PayTableListMap() map[string]map[int]*base.PayTable {
	if t.PayTableListMaps != nil {
		return t.PayTableListMaps
	}
	var ptm = make(map[string]map[int]*base.PayTable)
	for _, v := range t.PayTableList {
		nameKay := v.Tags[0].Name
		if ptm[nameKay] == nil {
			ptm[nameKay] = make(map[int]*base.PayTable)
		}
		ptm[nameKay][len(v.Tags)] = v
	}
	t.PayTableListMaps = ptm
	return ptm
}

// NeedFill 需要填充的空白Tag
func (t *Table) NeedFill() []*base.Tag {
	minRow := len(t.TagList) - t.Row
	return lo.Filter(helper.ListToArr(t.TagList), func(tag *base.Tag, i int) bool {
		if tag.X < minRow {
			return false
		}
		return tag.Name == ""
	})
}

// NeedFillEdge 获取空白边缘的标签
func (t *Table) NeedFillEdge() []*base.Tag {
	v := make(map[[2]int]bool, 0)
	tags := make([]*base.Tag, 0)
	needFillTags := t.NeedFill()
	for _, fillTag := range needFillTags {
		for _, tag := range t.GetAdjacent(fillTag.X, fillTag.Y) {
			if tag.Name != "" && tag.Name != "scatter" {
				if ok := v[[2]int{tag.X, tag.Y}]; !ok {
					tags = append(tags, tag)
					//tags = append(tags, tag.Copy())
				}
				v[[2]int{tag.X, tag.Y}] = true
			}
		}
	}
	//tags := make([]*base.Tag, 0)
	//for k := range v {
	//	tags = append(tags, t.TagList[k[0]][k[1]])
	//}
	return tags
}

// QueryTags  查询指定名称的标签
func (t *Table) QueryTags(name string) []*base.Tag {
	return lo.Filter(t.ToArr(), func(tag *base.Tag, i int) bool {
		return tag.Name == name
	})
}

// GetTag 获取tag
func (t *Table) GetTag(x, y int) *base.Tag {
	return t.TagList[x][y]
}

func (t *Table) NameGetTag(name string) *base.Tag {
	if name == "scatter" {
		return t.Scatter.Copy()
	}
	if name == "multiplier" {
		return t.DoubleTag.Copy()
	}

	for _, tag := range t.Tags {
		if tag.Name == name {
			return tag.Copy()
		}
	}

	for _, tag := range t.WildTags {
		if tag.Name == name {
			return tag.Copy()
		}
	}

	return &base.Tag{}
}

// SetTag 设置tag
func (t *Table) SetTag(x, y int, tag base.Tag) {
	tag.X = x
	tag.Y = y
	t.TagList[x][y] = &tag
}

// randTag 随机一个tag
func (t *Table) randTag() *base.Tag {
	return t.Tags[helper.RandInt(len(t.Tags))].Copy()
}

// Count 统计某个name的个数
func (t *Table) Count(name string) (count int) {
	for _, tag := range t.ToArr() {
		if tag.Name == name {
			count++
		}
	}
	return
}

// ToArr 转换为一维数组
func (t *Table) ToArr() []*base.Tag {
	return helper.ListToArr(t.TagList)
}
func (t *Table) WindowToArr() []*base.Tag {
	startIndex := len(t.TagList) - t.Row
	result := t.TagList[startIndex:]
	return helper.ListToArr(result)
}

// GetColumn 获取某一列
func (t *Table) GetColumn(Y int) []*base.Tag {
	var cols []*base.Tag
	for _, tags := range t.TagList {
		cols = append(cols, tags[Y])
	}
	return cols
}

// GetAdjacent 获取附近四个方向的tag
func (t *Table) GetAdjacent(x, y int) []*base.Tag {
	var adjacent []*base.Tag
	if x > 0 {
		adjacent = append(adjacent, t.TagList[x-1][y].Copy())
	}
	if x < t.Row-1 {
		adjacent = append(adjacent, t.TagList[x+1][y].Copy())
	}
	if y > 0 {
		adjacent = append(adjacent, t.TagList[x][y-1].Copy())
	}
	if y < t.Col-1 {
		adjacent = append(adjacent, t.TagList[x][y+1].Copy())
	}
	return adjacent
}

// GetBiasAdjacent 获取斜角方向
func (t *Table) GetBiasAdjacent(x, y int) []*base.Tag {
	var adjacent []*base.Tag
	if x > 0 && y > 0 {
		adjacent = append(adjacent, t.TagList[x-1][y-1].Copy())
	}
	if x > 0 && y < t.Col-1 {
		adjacent = append(adjacent, t.TagList[x-1][y+1].Copy())
	}

	if x < t.Row-1 && y > 0 {
		adjacent = append(adjacent, t.TagList[x+1][y-1].Copy())
	}
	if x < t.Row-1 && y < t.Col-1 {
		adjacent = append(adjacent, t.TagList[x+1][y+1].Copy())
	}

	return adjacent
}

// GetAllAdjacent 获取附近8个方向的tag
func (t *Table) GetAllAdjacent(x, y int) []*base.Tag {
	var adjacent []*base.Tag
	adjacent = append(adjacent, t.GetAdjacent(x, y)...)
	adjacent = append(adjacent, t.GetBiasAdjacent(x, y)...)
	return adjacent
}

// GetAdjacentOne  获取附近四个方向的tag随机取一个
func (t *Table) GetAdjacentOne(x, y int) *base.Tag {
	var adjacent []*base.Tag
	if x > 0 {
		adjacent = append(adjacent, t.TagList[x-1][y].Copy())
	}
	if x < t.Row-1 {
		adjacent = append(adjacent, t.TagList[x+1][y].Copy())
	}
	if y > 0 {
		adjacent = append(adjacent, t.TagList[x][y-1].Copy())
	}
	if y < t.Col-1 {
		adjacent = append(adjacent, t.TagList[x][y+1].Copy())
	}
	return adjacent[helper.RandInt(len(adjacent))]
}

// RandPosition 返回随机位置
func (t *Table) RandPosition(verify *Verify) [2]int {
	var x, y int
	count := 0
	for {
		if count > 100 {
			global.GVA_LOG.Error("随机位置超过100次" + t.PrintTable(""))
		}
		count++
		x = helper.RandInt(t.Row)
		y = helper.RandInt(t.Col)
		if !verify.GetVerify(x, y) {
			break
		}
	}
	return [2]int{
		x, y,
	}
}

func (t *Table) PrintTable(str string) string {
	str += ":\n"
	min := len(t.TagList) - t.Row
	for _, row := range t.TagList[min:] {
		for _, col := range row {
			str += fmt.Sprintf("%s\t", strconv.Itoa(col.X)+":"+strconv.Itoa(col.Y)+" "+helper.If(col.Name == "", "🀆", col.Name))
		}
		str += "\r\n"
	}
	//fmt.Println(str)
	return str + "\r\n"
}

func (t *Table) PrintModelTem(str string, windows bool) string {
	str += ":\n"
	min := 0
	if windows {
		min = len(t.InitTable) - t.Row
	}
	for _, row := range t.InitTable[min:] {
		for _, col := range row {
			name := col.Name
			if name == "" {
				name = "🀆"
			}
			if name == "multiplier" {
				name += "*" + strconv.Itoa(int(col.Multiple))
			}
			str += fmt.Sprintf("%s\t", name)
		}
		str += "\r\n"
	}
	return str + "\r\n"
}

func (t *Table) PrintTemplate() string {
	str := ""
	for _, row := range t.InitTable {
		for _, col := range row {
			name := strconv.Itoa(col.Id)
			if name == "" {
				name = "🀆"
			}
			if col.Name == "multiplier" {
				name += "*" + strconv.Itoa(int(col.Multiple))
			}
			str += fmt.Sprintf("%s ", name)
		}
		str += "\r\n"
	}
	return str
}

func (t *Table) TemplateGenerate(str string) [][]*base.Tag {
	var template [][]*base.Tag
	for _, row := range strings.Split(str, "\r\n") {
		if row == "" {
			continue
		}

		var rowList []*base.Tag
		cols := strings.Split(row, " ")
		for _, col := range cols {
			if col == "" {
				continue
			}

			if col == "🀆" {
				rowList = append(rowList, &base.Tag{
					Name: "",
				})
				continue
			}

			strs := strings.Split(col, "*")
			tagId, err := strconv.Atoi(strs[0])
			if err != nil {
				rowList = append(rowList, &base.Tag{
					Name: "",
				})
			}
			fillTag := t.GetTagById(tagId)

			if len(strs) == 2 {
				tagMul, err1 := strconv.Atoi(strs[1])
				if err1 != nil {
					fillTag.Multiple = 1
				}
				fillTag.Multiple = float64(tagMul)
			}
			rowList = append(rowList, fillTag)
		}
		template = append(template, rowList)
	}
	return template
}

func (t *Table) PrintModel(str string) string {
	str += ":\n"
	min := len(t.TagList) - t.Row
	for _, row := range t.TagList[min:] {
		for _, col := range row {
			name := col.Name
			if name == "" {
				name = "🀆"
			}
			if name == "multiplier" {
				name += "*" + strconv.Itoa(int(col.Multiple))
			}
			str += fmt.Sprintf("%s\t", name)
		}
		str += "\r\n"
	}
	return str + "\r\n"
}

func (t *Table) PrintList(tags []*base.Tag) {
	str := ":\n"
	for r, row := range t.TagList {
		for c, col := range row {

			tagNeeds := lo.Filter(tags, func(tag *base.Tag, i int) bool {
				return tag.X == r && tag.Y == c
			})
			if len(tagNeeds) > 0 {
				str += fmt.Sprintf("%s\t", strconv.Itoa(col.X)+":"+strconv.Itoa(col.Y)+helper.If(tagNeeds[0].Name == "", "🀆", tagNeeds[0].Name))
			} else {
				str += fmt.Sprintf("%s\t", strconv.Itoa(col.X)+":"+strconv.Itoa(col.Y)+"🀆")
			}

		}
		str += "\r\n"
	}
	fmt.Println(str)
}

func (t *Table) AddMul(mul float64) {
	t.Mul, _ = decimal.NewFromFloat(t.Mul).Add(decimal.NewFromFloat(mul)).Float64()
}
func (t *Table) AddMulDecimal(mul decimal.Decimal) {
	t.Mul, _ = decimal.NewFromFloat(t.Mul).Add(mul).Float64()
}

func (t *Table) MulMulDecimal(mul decimal.Decimal) {
	t.Mul, _ = decimal.NewFromFloat(t.Mul).Mul(mul).Float64()
}
func (t *Table) MulMul(mul float64) {
	t.Mul, _ = decimal.NewFromFloat(t.Mul).Mul(decimal.NewFromFloat(mul)).Float64()
}

func (t *Table) AddSkillMul(mul float64) {
	t.SkillMul, _ = decimal.NewFromFloat(t.SkillMul).Add(decimal.NewFromFloat(mul)).Float64()
}

// SetCoordinates 设置坐标
func (t *Table) SetCoordinates() {
	for i, tags := range t.TagList {
		for i2, tag := range tags {
			tag.X = i
			tag.Y = i2
		}
	}
}

// TableReset 初始化列表
func (t *Table) TableReset() {
	t.TagList = helper.NewTable[*base.Tag](t.Col, t.Row, func(x, y int) *base.Tag {
		return &base.Tag{X: x, Y: y, Name: ""}
	})
}

func (t *Table) TableModelReset() {
	t.TagList = helper.NewTable[*base.Tag](t.Col, t.BaseRow, func(x, y int) *base.Tag {
		return &base.Tag{X: x, Y: y, Name: ""}
	})
}

func (t *Table) Copy() *Table {
	newT := *t
	return &newT
}

// GetGraph 获取布局
func (t *Table) GetGraph() [][]*base.Tag {
	tagList := make([][]*base.Tag, len(t.TagList))
	for i, tags := range t.TagList {
		tagList[i] = make([]*base.Tag, t.Col)
		for i2, tag := range tags {
			tagList[i][i2] = tag.Copy()
		}
	}
	return tagList
}

func (t *Table) GetInformation() string {
	//MinMul      float64 // 最小倍数
	//MaxMul      float64 // 最大倍数
	//InitNum     int     // 初始个数
	//ScatterNum  int     // scatter次数
	mul := t.Mul
	return fmt.Sprintf(
		"最大倍率:%g,最小倍率:%g,初始个数:%d,scatter次数:%d,倍率:%g",
		t.Target.MaxMul,
		t.Target.MinMul,
		t.Target.InitNum,
		t.Target.ScatterNum,
		mul,
	)

}

//func (t *Table) GetTagListKeyNum() map[string]int {
//	// 现在只能在用到的时候重新计算
//	//todo 优化方向：每次重新布局的时候 填充数据
//	var m = make(map[string]int)
//	for _, tags := range t.TagList {
//		for _, tag := range tags {
//			m[tag.Name]++
//		}
//	}
//	return m
//}

func (t *Table) GetUpDownLeftRight(x, y int) []*base.Tag {
	var adjacent []*base.Tag
	//上
	if x > 0 {
		if t.TagList[x-1][y].Copy().Name == "" {
			adjacent = append(adjacent, t.TagList[x-1][y].Copy())
		}
	}

	//下
	if x < t.Row-1 {
		if t.TagList[x+1][y].Copy().Name == "" {
			adjacent = append(adjacent, t.TagList[x+1][y].Copy())
		}

	}

	//左
	if y > 0 {
		if t.TagList[x][y-1].Copy().Name == "" {
			adjacent = append(adjacent, t.TagList[x][y-1].Copy())
		}
	}
	//右
	if y < t.Col-1 {
		if t.TagList[x][y+1].Copy().Name == "" {
			adjacent = append(adjacent, t.TagList[x][y+1].Copy())
		}
	}

	return adjacent
}

func (t *Table) GetRoundTagsByCenterPoint(x, y int) []*base.Tag {
	var adjacent []*base.Tag
	//上
	if x > 0 {
		if t.TagList[x-1][y].Copy().Name != "" {
			adjacent = append(adjacent, t.TagList[x-1][y].Copy())
		}
	}
	//左上
	if x > 0 && y > 0 {
		if t.TagList[x-1][y-1].Copy().Name != "" {
			adjacent = append(adjacent, t.TagList[x-1][y-1].Copy())
		}
	}

	//右上
	if x > 0 && y < t.Col-1 {
		if t.TagList[x-1][y+1].Copy().Name != "" {
			adjacent = append(adjacent, t.TagList[x-1][y+1].Copy())
		}
	}

	//下
	if x < t.Row-1 {
		if t.TagList[x+1][y].Copy().Name != "" {
			adjacent = append(adjacent, t.TagList[x+1][y].Copy())
		}
	}

	//左下
	if x < t.Row-1 && y > 0 {
		if t.TagList[x+1][y-1].Copy().Name != "" {
			adjacent = append(adjacent, t.TagList[x+1][y-1].Copy())
		}
	}

	//右下
	if x < t.Row-1 && y < t.Col-1 {
		if t.TagList[x+1][y+1].Copy().Name != "" {
			adjacent = append(adjacent, t.TagList[x+1][y+1].Copy())
		}
	}

	//左
	if y > 0 {
		if t.TagList[x][y-1].Copy().Name != "" {
			adjacent = append(adjacent, t.TagList[x][y-1].Copy())
		}
	}

	//右
	if y < t.Col-1 {
		if t.TagList[x][y+1].Copy().Name != "" {
			adjacent = append(adjacent, t.TagList[x][y+1].Copy())
		}
	}
	return adjacent
}
