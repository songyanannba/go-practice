package template

import (
	"fmt"
	"github.com/samber/lo"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/service/slot/base"
	"slot-server/service/slot/component"
	"slot-server/utils"
	"slot-server/utils/helper"
	"strconv"
	"strings"
)

func NewGenTemplate(tem *business.SlotTemplateGen) (gen *GenTemplate, err error) {
	var (
		config *component.Config
		newMap map[int][]*Scale
	)

	config, err = component.GetSlotConfig(uint(tem.SlotId), false)
	if err != nil {
		return nil, err
	}
	gen = &GenTemplate{
		Config:        config,
		SlotId:        tem.SlotId,
		Type:          tem.Type,
		MinRatio:      tem.MinRatio,
		MaxRatio:      tem.MaxRatio,
		MinScatter:    tem.MinScatter,
		MaxScatter:    tem.MaxScatter,
		InitialWeight: map[int]map[string]int{},
		LargeScale:    map[int][]*WeightInterval{},
		Interval:      []*WeightInterval{},
		TrimDown:      map[int][]*Scale{},
		TrimUp:        map[int][]*Scale{},
		Template:      map[int][]*base.Tag{},
		Schedule:      tem.Schedule,
		SpecialWeight: map[int]any{},
	}

	err = gen.GetInitialWeight(tem.InitialWeight)
	if err != nil {
		return nil, err
	}
	err = gen.GetLargeScale(tem.LargeScale)
	if err != nil {
		return nil, err
	}
	newMap, err = gen.GetTrim(tem.TrimDown)
	if err != nil {
		return nil, err
	}
	gen.TrimDown = newMap
	newMap, err = gen.GetTrim(tem.TrimUp)
	if err != nil {
		return nil, err
	}
	gen.TrimUp = newMap
	err = gen.SetSpecialWeight(tem.SpecialConfig)
	if err != nil {
		return nil, err
	}
	err = gen.InitTem()
	if err != nil {
		return nil, err
	}
	return gen, nil

}

func GetColInfo(str string) (int, string, error) {

	colW := strings.Split(str, ":")
	if len(colW) != 2 {
		return 0, "", fmt.Errorf("initial weight format error")
	}
	colNum, err := strconv.Atoi(colW[0])
	if err != nil {
		return 0, "", err
	}
	return colNum, colW[1], nil
}

func GetColMap(str string) (map[int]string, error) {
	if str == "" {
		return map[int]string{}, nil
	}
	strs := utils.FormatCommand(str)
	strMap := map[int]string{}
	for _, str := range strs {
		if str == "" {
			continue
		}
		info, s, err := GetColInfo(str)
		if err != nil {
			return nil, err
		}
		strMap[info] = s
	}
	return strMap, nil
}

func GetInterval(str string) (map[string][2]int, error) {
	if str == "" {
		return map[string][2]int{}, nil
	}
	strMap := map[string][2]int{}
	info := strings.Split(str, ",")
	for _, s := range info {
		if s == "" {
			continue
		}
		nameInt := strings.Split(s, "=")
		if len(nameInt) != 2 {
			return nil, fmt.Errorf("interval format error")
		}
		name := nameInt[0]
		interval := strings.Split(nameInt[1], "-")
		if len(interval) != 2 {
			return nil, fmt.Errorf("interval format error")
		}
		min, err := strconv.Atoi(interval[0])
		if err != nil {
			min = 0
		}
		max, err := strconv.Atoi(interval[1])
		if err != nil {
			max = 0
		}
		if min >= max {
			return nil, fmt.Errorf("interval format error")
		}
		strMap[name] = [2]int{min, max}
	}
	return strMap, nil
}

func (t *GenTemplate) GetInitialWeight(inWei string) error {
	if inWei == "" {
		return nil
	}
	t.InitialWeight = map[int]map[string]int{}
	colMap, err := GetColMap(inWei)
	if err != nil {
		return err
	}
	for colNum, s := range colMap {
		t.InitialWeight[colNum] = map[string]int{}
		strEvn := base.ParseWeightDataStr(s)
		for i, datum := range strEvn.Data {
			t.InitialWeight[colNum][datum] = strEvn.GetSection(i)
		}
	}
	return nil
}

func (t *GenTemplate) GetLargeScale(inWei string) error {
	colMap, err := GetColMap(inWei)
	if err != nil {
		return err
	}

	for colNum, s := range colMap {
		var (
			tagCs []*WeightInterval
		)
		interval, err := GetInterval(s)
		if err != nil {
			return err
		}
		for name, interval := range interval {
			tagCs = append(tagCs, &WeightInterval{
				Tag:      t.Config.GetTag(name),
				MinCount: interval[0],
				MaxCount: interval[1],
			})
		}
		t.LargeScale[colNum] = tagCs
	}

	return nil
}

func (t *GenTemplate) GetTrim(inWei string) (map[int][]*Scale, error) {
	newMap := map[int][]*Scale{}
	colMap, err := GetColMap(inWei)
	if err != nil {
		return nil, err
	}
	for colNum, s := range colMap {
		trims := strings.Split(s, ",")
		var (
			tagCs []*Scale
		)
		for _, trim := range trims {
			if trim == "" {
				continue
			}
			twoTag := strings.Split(trim, "=>")
			if len(twoTag) != 2 {
				return nil, fmt.Errorf("trim format error")
			}
			tagCs = append(tagCs, &Scale{
				Tag:        t.Config.GetTag(twoTag[0]),
				ReplaceTag: t.Config.GetTag(twoTag[1]),
			})
		}
		newMap[colNum] = tagCs
	}
	return newMap, nil
}

func (t *GenTemplate) InitTem() error {

	for col, counts := range t.InitialWeight {
		sumCount := 0
		for _, count := range counts {
			sumCount += count
		}
		allTags := make([]*base.Tag, sumCount)
		for _, interval := range t.Interval {
			randInterval := helper.RandInt(interval.MaxCount-interval.MinCount) + interval.MinCount
			err := IntervalPlacement(allTags, interval.Tag, counts[interval.Tag.Name], randInterval)
			if err != nil {
				return err
			}
		}

		var tags []*base.Tag
		for name, count := range counts {
			_, b := lo.Find(t.Interval, func(item *WeightInterval) bool {
				return item.Tag.Name == name
			})
			if b {
				continue
			}
			fillTag := t.Config.GetTag(name)
			if fillTag.Name == "" || fillTag.Id == -1 {
				global.GVA_LOG.Error("tag name is empty " + name)
			}
			for i := 0; i < count; i++ {
				//fillTag := t.Config.GetTag(name)
				//if fillTag.Name == "" || fillTag.Id == -1 {
				//	global.GVA_LOG.Error("tag name is empty " + name)
				//}
				tags = append(tags, fillTag.Copy())
			}
		}
		helper.SliceShuffle(tags)
		emptyTags := lo.Filter(allTags, func(item *base.Tag, i int) bool {
			return item == nil || item.Name == "" || item.Id == -1
		})
		if len(emptyTags) != len(tags) {
			return fmt.Errorf("init template error" + fmt.Sprintf("all:%d , empty:%d , tags:%d", len(allTags), len(emptyTags), len(tags)))
		}
		TagsFill(allTags, tags)
		t.Template[col] = allTags
	}
	return nil
}

func (t *GenTemplate) SetSpecialWeight(weight string) error {
	if weight == "" {
		return nil
	}
	strs := utils.FormatCommand(weight)
	for _, str := range strs {
		if str == "" {
			continue
		}
		infos := strings.Split(str, ":")
		if len(infos) != 2 {
			return fmt.Errorf("special weight format error")
		}
		weightName := infos[0]
		weightStr := infos[1]
		switch weightName {
		case MulTagWeight:
			t.SpecialWeight[MulTagWeightId] = base.ParseWeightData(weightStr)
		case Interval:
			interval, err := GetInterval(weightStr)
			if err != nil {
				return err
			}
			for name, interval := range interval {
				t.Interval = append(t.Interval, &WeightInterval{
					Tag:      t.Config.GetTag(name),
					MinCount: interval[0],
					MaxCount: interval[1],
				})

			}
		}
	}
	return nil
}
