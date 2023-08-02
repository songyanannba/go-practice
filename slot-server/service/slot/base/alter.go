package base

import (
	"fmt"
	"slot-server/global"
	"strconv"
)

type AlterFlow struct {
	Id          int
	InitList    [][]*Tag
	RemoveList  []*Eliminate
	AfterDrop   [][]*Tag
	AddList     []*Tag
	FlowMap     string
	SumMul      float64
	Gain        int
	RankId      int
	RemoveCount int
	EmitList    [][]*Tag
}

func NewAlterFlow(count int) *AlterFlow {
	return &AlterFlow{
		Id:          count,
		InitList:    make([][]*Tag, 0),
		RemoveList:  make([]*Eliminate, 0),
		AfterDrop:   make([][]*Tag, 0),
		AddList:     make([]*Tag, 0),
		FlowMap:     "",
		SumMul:      float64(0),
		Gain:        0,
		RankId:      0,
		RemoveCount: 0,
	}
}

type View struct {
	Data string `json:"data"`
	Text string `json:"text"`
}

func (a *AlterFlow) String() string {
	data := ""
	for x, tags := range a.InitList {
		for y, i3 := range tags {
			i, mul, emit := a.QueryElimination(x, y)
			data += strconv.Itoa(i) + "|" + strconv.Itoa(i3.Id) + "|" + strconv.Itoa(int(mul)) + "|" + strconv.Itoa(int(emit)) + ","
		}
		data = data[:len(data)-1]
		data += ";"
	}
	data = data[:len(data)-1]
	view := &View{
		Data: data,
		Text: a.FlowMap,
	}
	detail, err := global.Json.MarshalToString(view)
	if err != nil {
		return ""
	}
	return detail
}

type Eliminate struct {
	RemoveList []*Tag
	Mul        float64
}

func (e *Eliminate) String() string {
	return fmt.Sprintf("消除 %d 个 %s 倍率: %g \n", len(e.RemoveList), e.RemoveList[0].Name, e.Mul)
}

func (a *AlterFlow) QueryElimination(x, y int) (int, int, int) {
	var a1, a2, a3 int
	a2 = int(a.InitList[x][y].Multiple)
	for _, eliminate := range a.RemoveList {
		for _, tag := range eliminate.RemoveList {
			if tag.X == x && tag.Y == y {
				a1 = 1
				break
			}
		}
		if a1 != 0 {
			break
		}
	}
	for _, tags := range a.EmitList {
		for _, tag := range tags {
			if tag.X == x && tag.Y == y {
				a3 = 1
				break
			}
		}
	}

	return a1, a2, a3
}

func (a *AlterFlow) GetInformation() string {
	mul := a.SumMul
	str := ""
	str += fmt.Sprintf("编号:%d, 赢钱:%d, 倍率:%g \n", a.Id, a.Gain, mul)
	for _, eliminate := range a.RemoveList {
		mul = eliminate.Mul
		str += fmt.Sprintf("消除 %d 个 %s 倍率: %g \n", len(eliminate.RemoveList), eliminate.RemoveList[0].Name, mul)
	}
	for _, add := range a.AddList {
		str += fmt.Sprintf("增加 %s  %d:%d \n", add.Name, add.X, add.Y)
	}
	str += fmt.Sprintf("扩散位置 %v 个 \n", a.EmitList)
	return str
}

func (a *AlterFlow) RemoveListCount() {
	for _, eliminate := range a.RemoveList {
		a.RemoveCount += len(eliminate.RemoveList)
	}
}
