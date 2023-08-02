package component

import (
	. "slot-server/model/business"
	"slot-server/utils/helper"
	"strings"
)

func ParseReel(s []*SlotReelData) *Reel {
	reels := newReel()
	for _, v := range s {
		reels.Normal[v.Which] = newReelData(v.ReelData, v.WeightData)
		reels.Fs[v.Which] = newReelData(v.FsReelData, v.FsWeight)
	}
	return reels
}

func parseWeight(s string) ([]string, []int) {
	if s == "" {
		return nil, nil
	}
	tagStr, weightStr, _ := strings.Cut(s, "@")
	tags := strings.Split(tagStr, "&")
	weights := helper.SplitInt[int](weightStr, "&")
	if len(weights)-1 != len(tags) {
		return nil, nil
	}
	return tags, weights
}

type Reel struct {
	Normal map[int]*ReelData
	Fs     map[int]*ReelData
}

func newReel() *Reel {
	return &Reel{
		Normal: make(map[int]*ReelData),
		Fs:     make(map[int]*ReelData),
	}
}

func (r *Reel) GetReelData(typ, which int) *ReelData {
	if typ == 1 {
		return r.Normal[which]
	}
	return r.Fs[which]
}

type ReelData struct {
	Data      []string // 排布数据
	WeightTag []string // 权重标签
	Weight    []int    // 权重
}

func newReelData(data, weightData string) *ReelData {
	weightTag, normalWeight := parseWeight(weightData)
	return &ReelData{
		Data:      strings.Split(data, "&"),
		WeightTag: weightTag,
		Weight:    normalWeight,
	}
}

// Fetch 根据配置获取一组数据 type:1 普通 2 免费 | which:取哪组配置 | offset p偏移量 | length: 长度 | place 自定义排布索引
func (r *Reel) Fetch(typ, which int, offset, length int, place int) []string {
	reelData := r.GetReelData(typ, which)
	if reelData == nil {
		return nil
	}
	start := 0
	if place > 0 {
		// 如果有自定义排布索引 则从该索引开始取
		start = place - 1
	} else if len(reelData.WeightTag) > 0 && len(reelData.Weight) == len(reelData.WeightTag)+1 {
		// 如果排布数据有权重标签 则根据权重标签随机取一个标签
		k := helper.RandomLongWeight(reelData.Weight)
		tagName := helper.SliceVal(reelData.WeightTag, k)
		var keys []int
		for key, name := range reelData.Data {
			if name == tagName {
				keys = append(keys, key)
			}
		}
		start = keys[helper.RandInt(len(keys))]
	} else {
		// 否则直接取随机排布
		if len(reelData.Data) == 0 {
			reelData.Data = []string{"null"}
		} else {
			start = helper.RandInt(len(reelData.Data))
		}
	}
	return helper.SliceByRange(reelData.Data, start+offset, length)
}

func (s *Spin) GetInitDataByReel(typ, which int) {
	// 如果有假数据 则使用假数据
	places := helper.If(s.IsFree, s.Config.freePlace, s.Config.place)
	if !s.IsTest && s.PlayNum > 0 && s.Demo {
		fake := s.Config.Fakes.GetFake(typ, s.PlayNum)
		if fake != nil {
			places = fake.Position
		}
	}
	// 根据滚轮配置初始化数据
	for k, reel := range s.Config.Reel {
		var (
			offset = 0 // 偏移量仅在单线下使用 用于指定权重标签的位置
		)
		if len(s.Config.Coords) == 1 {
			// y坐标 - 行数
			coord := s.Config.Coords[0]
			if helper.SliceKeyExist(coord, k) {
				offset = coord[k].Y - s.Config.Row + 1
			}
		}
		reelTag := reel.Fetch(typ, which, offset, s.Config.Row, helper.SliceVal(places, k))
		s.AddInitData(k, reelTag)
		if k >= s.Config.Index-1 {
			break
		}
	}
}
