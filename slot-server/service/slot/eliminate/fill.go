package eliminate

import (
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/service/slot/base"
	"slot-server/utils/helper"
	"strconv"
)

// Fill 填充图形内空白内容,按倍率区分填充可消除内容还是不可消除内容
func (t *Table) Fill(lineList [][]*base.Tag, isFee bool, stop bool) (rTags []*base.Tag) {
	if stop {
		return t.FillAll()
	}
	//lineList 需要消除的集合
	//lineList 和赢钱的组合去匹配 ,返回集合中 每种要消除划线的集合 和对应的倍率
	_, mul := t.WinMatchList(lineList, isFee)
	if t.Target.Compare(decimal.NewFromFloat(t.Mul).Add(decimal.NewFromFloat(mul)).Sub(decimal.NewFromFloat(t.SkillMul))) {
		rTags = t.FillLine()
	} else {
		rTags = t.FillAll()
	}
	return rTags
}

// SpecifyFill 随机填充指定长度的标签
func (t *Table) SpecifyFill(Num int, tag *base.Tag, verify *Verify) []*base.Tag {
	var (
		result   []*base.Tag
		position [2]int
		randTag  base.Tag
		choTag   *base.Tag
		v        map[[2]int]bool
		countz   int
		count    int
	)

	for {
		if countz > 5 {
			global.GVA_LOG.Error("SpecifyFill error 100 times" + t.PrintTable("") + randTag.Name + strconv.Itoa(Num))
			break
		}
		countz++
		result = make([]*base.Tag, 0)
		v = make(map[[2]int]bool, 0)
		position = t.RandPosition(verify)

		choTag = SetSite(tag.Copy(), position[0], position[1])
		v[[2]int{choTag.X, choTag.Y}] = true
		result = append(result, choTag)

		for len(result) != Num {
			if count > 200 {
				global.GVA_LOG.Error("SpecifyFill error 50 times")
				break
			}
			count++
			tag := result[helper.RandInt(len(result))]
			adjacentTag := t.GetAdjacentOne(tag.X, tag.Y)

			choTag = SetSite(tag, adjacentTag.X, adjacentTag.Y)
			if v[[2]int{choTag.X, choTag.Y}] {
				continue
			}
			v[[2]int{choTag.X, choTag.Y}] = true
			result = append(result, choTag)
		}
		if len(result) == Num {
			break
		}
	}
	for ints, _ := range v {
		verify.SetVerify(ints[0], ints[1])
	}
	return result
}

// FillList 将数组填充进图形
func (t *Table) FillList(tags []*base.Tag) {
	for _, tag := range tags {
		t.TagList[tag.X][tag.Y] = tag.Copy()
	}
}

// FillLine 填充可以消除的组合
func (t *Table) FillLine() []*base.Tag {
	rTags := make([]*base.Tag, 0)
	//NeedFillEdge 获取消除标签 周边的标签
	mayLines := t.FindAllMayLine(t.NeedFillEdge())
	var (
		mayLine MayLine
	)
	mayLines = lo.Filter(mayLines, func(item MayLine, index int) bool {
		mul, _ := decimal.NewFromFloat(t.Mul).Add(decimal.NewFromFloat(item.Mul)).Float64()
		return mul <= t.Target.MaxMul && len(item.Tags) == t.Target.InitNum
	})
	pTbLines := t.FindPTLine()
	count := 0
	for len(pTbLines) != 0 {
		if count > 30 {
			break
		}
		count++
		if count > 10 || len(mayLines) == 0 {
			mayLine = pTbLines[helper.RandInt(len(pTbLines))]
		} else {
			mayLine = mayLines[helper.RandInt(len(mayLines))]
		}

		fillTags := make([]*base.Tag, 0)
		for _, tag := range mayLine.Tags {
			fillTag := &base.Tag{
				Id:   mayLine.Id,
				Name: mayLine.Name,
				X:    tag.X,
				Y:    tag.Y,
			}
			fillTags = append(fillTags, fillTag.Copy())
		}
		verify := NewVerify()
		verify.sites = fillTags
		t.AndFindLine(verify)
		mul, _ := t.WinMatchName(mayLine.Name, len(verify.verify)).Float64()
		nowMul, _ := decimal.NewFromFloat(mul).Add(decimal.NewFromFloat(t.Mul)).Float64()
		if nowMul <= t.Target.MaxMul && mul > 0 {
			break
		}
	}
	for _, tag := range mayLine.Tags {
		if t.TagList[tag.X][tag.Y].Name != "" {
			continue
		}
		fillTag := &base.Tag{
			Id:   mayLine.Id,
			Name: mayLine.Name,
			X:    tag.X,
			Y:    tag.Y,
		}
		t.TagList[tag.X][tag.Y] = fillTag

		rTags = append(rTags, fillTag.Copy())
	}

	rTags = append(rTags, t.FillSca(1)...)
	rTags = append(rTags, t.FillAll()...)
	return rTags
}

// FillAll 填充全部空白,使其不能被消除
func (t *Table) FillAll() []*base.Tag {
	rTags := t.FillSca(t.Target.ScatterNum)
	needFillTags := t.NeedFill()
	for _, tag := range needFillTags {
		fillTag := &base.Tag{}
		count := 0
		for {
			count++
			choiceTag := t.Tags[helper.RandInt(len(t.Tags))]
			fillTag = choiceTag.Copy()
			fillTag.X = tag.X
			fillTag.Y = tag.Y

			t.TagList[tag.X][tag.Y] = fillTag
			if count > 20 {
				break
			}
			verify := NewVerify()
			verify.SetSite(t.TagList[tag.X][tag.Y])
			getLine := t.FindLine(verify)
			if len(getLine) >= enum.GetLine {
				t.TagList[tag.X][tag.Y].Name = ""
				continue
			} else {
				break
			}
		}
		rTags = append(rTags, fillTag.Copy())
	}
	return rTags
}

// FillSca 随机填充scatter
func (t *Table) FillSca(num int) []*base.Tag {
	rTags := make([]*base.Tag, 0)
	scatterNum := len(t.QueryTags(enum.ScatterName))
	if scatterNum == t.Target.ScatterNum {
		return rTags
	}
	needFillTags := t.NeedFill()

	helper.SliceShuffle(needFillTags)

	for _, tag := range needFillTags {
		if len(rTags) >= num {
			break
		}
		scatterNum = len(t.QueryTags(enum.ScatterName))

		nowColSca := len(lo.Filter(t.ToArr(), func(tag1 *base.Tag, i int) bool {
			return tag1.Name == enum.ScatterName && tag1.Y == tag.Y
		}))

		if scatterNum < t.Target.ScatterNum && nowColSca == 0 {
			fillTag := t.Scatter.Copy()
			fillTag.X = tag.X
			fillTag.Y = tag.Y
			t.TagList[tag.X][tag.Y] = fillTag
			rTags = append(rTags, fillTag.Copy())
			continue
		}

	}
	return rTags
}

// InitFill 初始化填充划线
func (t *Table) InitFill() {
	if t.Target.InitNum < enum.InitFillTagNum {
		return
	}

	pTables := lo.Filter(t.PayTableList, func(item *base.PayTable, index int) bool {
		return item.Multiple <= t.Target.MaxMul && len(item.Tags) == t.Target.InitNum
	})
	pTable := pTables[helper.RandInt(len(pTables))]
	verify := NewVerify()
	fillTags := t.SpecifyFill(len(pTable.Tags), pTable.Tags[0].Copy(), verify)
	t.FillList(fillTags)
	if len(fillTags) == 0 {
		t.FillSca(t.Target.ScatterNum)

	} else {
		t.FillSca(helper.RandInt(t.Target.ScatterNum))
	}

}

func (t *Table) FillDebug(debug string) {
	if debug == "" {
		return
	}
	//[["scatter","","scatter","","scatter"],["","","","",""],["","","","",""]]
	var arr [][]string
	if err := global.Json.Unmarshal([]byte(debug), &arr); err != nil {
		global.GVA_LOG.Error("debug", zap.Any("err", err))
	}
	//arr = helper.ArrVertical(arr)
	for x, tags := range arr {
		for y, tag := range tags {
			if tag == "" {
				continue
			}
			fillTag := t.NameGetTag(tag)
			fillTag.X = x
			fillTag.Y = y
			t.TagList[x][y] = fillTag
		}
	}
}

func (t *Table) FillTest(tagList [][]*base.Tag) {
	if len(tagList) == 0 {
		return
	}
	t.TagList = tagList
}

// FillModelRand  随机填充
func (t *Table) FillModelRand() {
	scatterMap := make(map[int]bool)
	for i, tags := range t.TagList {
		for i2, _ := range tags {
			for {
				fillName := t.ModelEvent.M["TagsWeight"].(*base.ChangeTableStrEvent).Fetch()
				fillTag := t.NameGetTag(fillName)
				fillTag.X = i
				fillTag.Y = i2
				if fillName == "scatter" {
					if scatterMap[i2] {
						continue
					}
					scatterMap[i2] = true
				}
				if fillName == "multiplier" {
					mul := t.ModelEvent.M["MultiplierWeight"].(*base.ChangeTableEvent).Fetch()
					fillTag.Multiple = float64(mul)
				}
				t.TagList[i][i2] = fillTag
				break
			}

		}
	}
}

func (t *Table) FillModelWindows() []*base.Tag {
	tags := t.WindowToArr()
	tagsMap := lo.GroupBy(tags, func(tag *base.Tag) string {
		return tag.Name
	})

	rTags := make([]*base.Tag, 0)
	needFillTags := t.NeedFill()
	for _, tag := range needFillTags {
		fillTag := &base.Tag{}
		count := 0
		for {
			count++
			choiceTag := t.Tags[helper.RandInt(len(t.Tags))]
			fillTag = choiceTag.Copy()
			fillTag.X = tag.X
			fillTag.Y = tag.Y
			if len(tagsMap[fillTag.Name]) >= enum.NumLine-1 {
				continue
			} else {
				t.TagList[tag.X][tag.Y] = fillTag
				tagsMap[fillTag.Name] = append(tagsMap[fillTag.Name], fillTag)
				break
			}
		}
		rTags = append(rTags, fillTag.Copy())
	}
	return rTags
}

func (t *Table) FillModel(tags []*base.Tag) {
	if len(tags) == 0 {
		return
	}
	minRow := len(t.InitTable)
	maxRow := 0
	for _, tag := range tags {
		if tag.X < minRow {
			minRow = tag.X
		}
		if tag.X > maxRow {
			maxRow = tag.X
		}
	}
	listTag := helper.NewTable[*base.Tag](t.Col, maxRow-minRow+1, func(x, y int) *base.Tag {
		return &base.Tag{X: x, Y: y, Name: ""}
	})

	for _, tag := range tags {
		listTag[tag.X-minRow][tag.Y] = tag
	}
	t.InitTable = append(listTag, t.InitTable...)
}

func (t *Table) FillModelInit() {
	t.TableModelReset()
	t.FillModelRand()
	t.InitTable = helper.CopyListArr(t.TagList)
}

func (t *Table) FillModelScatter(num int) {
	t.TableReset()
	tags := t.FillModelWindows()
	fillMap := make(map[int]bool)
	for num != 0 {
		rand := helper.RandInt(len(tags))
		tag := tags[rand]
		if fillMap[tag.Y] {
			continue
		}
		fillMap[tag.Y] = true
		t.TagList[tag.X][tag.Y] = t.Scatter.Copy()
		num--
	}
}
