package template

import (
	"slot-server/service/slot/base"
	"slot-server/service/slot/component"
)

type GenTemplate struct {
	Config        *component.Config
	SlotId        int
	Type          uint8
	MinRatio      float64
	MaxRatio      float64
	MinScatter    float64
	MaxScatter    float64
	InitialWeight map[int]map[string]int // 列=>(标签=>数量); 意思每一列对应多少不同标签种类 每个标签根据权重有多少个;
	LargeScale    map[int][]*WeightInterval
	Interval      []*WeightInterval
	TrimDown      map[int][]*Scale
	TrimUp        map[int][]*Scale
	Template      map[int][]*base.Tag //key 代表列; val 代表列生成的标签个数（个数根据标签权重计算:标签权重的后一个位置减去前一个位置）
	Schedule      string
	SpecialWeight map[int]any
}

type TagCount struct {
	Tag   *base.Tag
	Count int
}

type WeightInterval struct {
	Tag      *base.Tag
	MinCount int
	MaxCount int
}

type Scale struct {
	Tag        *base.Tag
	ReplaceTag *base.Tag
}
