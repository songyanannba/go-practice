package flow

import (
	"fmt"
	"github.com/samber/lo"
	"slot-server/global"
	"slot-server/service/slot/base"
	"strconv"
)

type SpinFlow struct {
	Id       int
	InitList [][]*base.Tag // 初始列表
	OmitList []*WinLine    // 消除列表
	AddList  []*base.Tag   // 填充列表
	FlowMap  string        // 流程图
	Gain     float64       // 获得
	SumMul   float64       // 总倍数
	EmitList [][]*base.Tag // 发射列表
}

func NewSpinFlow(id int) SpinFlow {
	return SpinFlow{
		Id:       id,
		InitList: make([][]*base.Tag, 0),
		OmitList: make([]*WinLine, 0),
		AddList:  make([]*base.Tag, 0),
		Gain:     0,
		SumMul:   0,
	}
}

func (s *SpinFlow) AddOmitList(lines ...*WinLine) {
	s.OmitList = append(s.OmitList, lines...)
	fls := lo.FilterMap(lines, func(line *WinLine, i int) (float64, bool) {
		return line.Mul, true
	})
	sum := lo.SumBy(fls, func(item float64) float64 {
		return item
	})
	s.SumMul += sum
}

type View struct {
	Data string `json:"data"`
	Text string `json:"text"`
}

func (s *SpinFlow) String() string {
	data := ""
	for x, tags := range s.InitList {
		for y, i3 := range tags {
			i, mul, emit := s.QueryElimination(x, y)
			data += strconv.Itoa(i) + "|" + strconv.Itoa(i3.Id) + "|" + strconv.Itoa(int(mul)) + "|" + strconv.Itoa(int(emit)) + ","
		}
		data = data[:len(data)-1]
		data += ";"
	}
	data = data[:len(data)-1]
	view := &View{
		Data: data,
		Text: s.FlowMap,
	}
	detail, err := global.Json.MarshalToString(view)
	if err != nil {
		return ""
	}
	return detail
}

func (s *SpinFlow) QueryElimination(x, y int) (int, int, int) {
	var a1, a2, a3 int
	a2 = int(s.InitList[x][y].Multiple)
	for _, eliminate := range s.OmitList {
		for _, tag := range eliminate.Tags {
			if tag.X == x && tag.Y == y {
				a1 = 1
				break
			}
		}
		if a1 != 0 {
			break
		}
	}
	for _, tags := range s.EmitList {
		for _, tag := range tags {
			if tag.X == x && tag.Y == y {
				a3 = 1
				break
			}
		}
	}

	return a1, a2, a3
}

func (s *SpinFlow) GetInformation() string {
	str := ""
	str += fmt.Sprintf("编号:%d, 赢钱:%g, 倍率:%g \n", s.Id, s.Gain, s.SumMul)
	for _, eliminate := range s.OmitList {
		str += fmt.Sprintf("消除 %d 个 %s 倍率: %g \n", len(eliminate.Tags), eliminate.Tags[0].Name, eliminate.Mul)
	}
	for _, add := range s.AddList {
		str += fmt.Sprintf("增加 %s  %d:%d \n", add.Name, add.X, add.Y)
	}
	str += fmt.Sprintf("扩散位置 %v 个 \n", s.EmitList)
	return str
}
