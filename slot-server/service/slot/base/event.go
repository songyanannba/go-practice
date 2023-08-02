package base

import (
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/utils/helper"
	"strconv"
	"strings"
)

type Event struct {
	M map[int]any
}

func NewEvent() *Event {
	return &Event{M: map[int]any{}}
}

func (e *Event) Add(typ int, params string) {
	var event any
	switch typ {
	// 目前换表和免费转事件参数一致
	case enum.SlotEvent1ChangeTable:
		event = ParseWeightData(params)
	default:
		event = ParseWeightData(params)
	}
	if event != nil {
		e.M[typ] = event
	}
}

func (e *Event) Get(typ int) (any, bool) {
	res, ok := e.M[typ]
	return res, ok
}

// ChangeTableEvent 数值换表事件
type ChangeTableEvent struct {
	Data   []int
	Weight []int
}

// ChangeTableStrEvent 字符换表事件
type ChangeTableStrEvent struct {
	Data   []string
	weight []int
}

// IntervalRatioEvent 倍率区间
type IntervalRatioEvent struct {
	Data   [][2]int
	weight []int
}

type Unit6LevelEvent struct {
	Collect   int               // 收集的数量
	CoreCount int               // 核心数量
	EmitEvent *ChangeTableEvent // 发射的数量
	WildEvent *ChangeTableEvent // wild的数量
}

// Fetch 根据权重获取一个值
func (e ChangeTableEvent) Fetch() int {
	k := helper.RandomLongWeight(e.Weight)
	return helper.SliceVal(e.Data, k)
}

func (s ChangeTableStrEvent) Fetch() string {
	k := helper.RandomLongWeight(s.weight)
	return helper.SliceVal(s.Data, k)
}

func (s ChangeTableStrEvent) GetSection(i int) int {
	if len(s.weight) > i+1 && i >= 0 {
		return s.weight[i+1] - s.weight[i]
	}
	return 0
}

func (s IntervalRatioEvent) Fetch() ([2]int, int) {
	k := helper.RandomLongWeight(s.weight)
	return helper.SliceVal(s.Data, k), k
}

func newChangeTableEvent(data, weight []int) *ChangeTableEvent {
	return &ChangeTableEvent{
		Data:   data,
		Weight: weight,
	}
}

// ParseWeightData 解析权重数据的格式
func ParseWeightData(s string) *ChangeTableEvent {
	if s == "" {
		return nil
	}
	tagStr, weightStr, _ := strings.Cut(s, "@")
	tags := helper.SplitInt[int](tagStr, "&")
	weights := helper.SplitInt[int](weightStr, "&")
	if len(weights)-1 != len(tags) {
		return nil
	}
	return newChangeTableEvent(tags, weights)
}

func ParseWeightDataStr(s string) *ChangeTableStrEvent {
	if s == "" {
		return nil
	}
	tagStr, weightStr, _ := strings.Cut(s, "@")
	tags := helper.SplitStr(tagStr, "&")
	weights := helper.SplitInt[int](weightStr, "&")
	if len(weights)-1 != len(tags) {
		return nil
	}
	return &ChangeTableStrEvent{
		Data:   tags,
		weight: weights,
	}
}
func (e *Event) Unit5NewEvent(events []*business.SlotEvent) {
	for i, event := range events {
		if i == 0 || i == 18 {
			e.M[i] = GetIntervalRatioEvent(event.Event1)
		} else {
			e.Add(i, event.Event1)
		}
	}
}

func (e *Event) Unit6NewEvent(events []*business.SlotEvent) {
	for i, event := range events {
		if i == 0 {
			e.M[i] = GetIntervalRatioEvent(event.Event1)
		} else {
			switch i {
			case 16, 17, 18, 19, 20:
				e.M[i] = Unit6ParseWeightDataStr(event.Event1)
			default:
				e.Add(i, event.Event1)
			}

		}
	}
}
func (e *Event) Unit8NewEvent(events []*business.SlotEvent) {
	for i, event := range events {
		e.M[i] = ParseWeightData(event.Event1)
	}
}

func Unit6ParseWeightDataStr(str string) *Unit6LevelEvent {
	result := &Unit6LevelEvent{}
	// 按@符号拆分成两个字符串
	split := strings.Split(str, "@")
	if len(split) != 6 {
		global.GVA_LOG.Info("输入字符串格式不正确")
		return nil
	}
	result.Collect = helper.Atoi(split[0])
	result.CoreCount = helper.Atoi(split[1])
	result.EmitEvent = ParseWeightData(split[2] + "@" + split[3])
	result.WildEvent = ParseWeightData(split[4] + "@" + split[5])
	return result
}

func GetIntervalRatioEvent(str string) *IntervalRatioEvent {
	// 按@符号拆分成两个字符串
	split := strings.Split(str, "@")
	if len(split) != 2 {
		global.GVA_LOG.Info("输入字符串格式不正确")
		return nil
	}

	arrStr := strings.Split(split[0], "&")
	var arr [][2]int

	for _, str := range arrStr {
		// 去除首尾的方括号
		str = strings.Trim(str, "[]")
		// 按逗号拆分
		values := strings.Split(str, ",")
		if len(values) != 2 {
			global.GVA_LOG.Info("输入字符串格式不正确")
			return nil
		}

		// 将字符串转换为整数
		start, err := strconv.Atoi(values[0])
		if err != nil {
			global.GVA_LOG.Info("输入字符串格式不正确")
			return nil
		}

		end, err := strconv.Atoi(values[1])
		if err != nil {
			global.GVA_LOG.Info("输入字符串格式不正确")
			return nil
		}

		// 添加到二维数组
		arr = append(arr, [2]int{start, end})
	}
	// 拆分后面的字符串按&符号
	numStr := strings.Split(split[1], "&")
	var nums []int

	// 遍历拆分的字符串进行转换
	for _, str := range numStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			global.GVA_LOG.Info("输入字符串格式不正确")
			return nil
		}

		// 添加到数组
		nums = append(nums, num)
	}
	return &IntervalRatioEvent{
		Data:   arr,
		weight: nums,
	}
}
