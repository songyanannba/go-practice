package template

import (
	"fmt"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/utils/helper"
)

func (t *GenTemplate) AdjLarge() error {
	for i, intervals := range t.LargeScale {
		for _, interval := range intervals {
			if interval.MaxCount-interval.MinCount >= 0 {
				num := helper.RandInt(interval.MaxCount-interval.MinCount) + interval.MinCount
				t.InitialWeight[i][interval.Tag.Name] = num
			}
		}
	}
	return t.InitTem()
}

func (t *GenTemplate) AdjTrimDown() error {
	for i, scales := range t.TrimDown {
		for _, scale := range scales {
			t.InitialWeight[i][scale.Tag.Name]--
			t.InitialWeight[i][scale.ReplaceTag.Name]++
		}
	}
	return t.InitTem()
}

func (t *GenTemplate) AdjTrimUp() error {
	for i, scales := range t.TrimUp {
		for _, scale := range scales {
			t.InitialWeight[i][scale.Tag.Name]--
			t.InitialWeight[i][scale.ReplaceTag.Name]++
		}
	}
	return t.InitTem()

}

func (t *GenTemplate) AdjScatterUp() error {
	rand := helper.RandInt(len(t.InitialWeight))
	t.InitialWeight[rand][enum.ScatterName]++
	return t.InitTem()
}

func (t *GenTemplate) AdjScatterDown() error {
	rand := helper.RandInt(len(t.InitialWeight))
	t.InitialWeight[rand][enum.ScatterName]--
	return t.InitTem()
}

func (t *GenTemplate) CreateTem(tem *business.SlotTemplateGen) {
	var items []*business.SlotTemplate
	for i, tags := range t.Template {
		item := &business.SlotTemplate{
			SlotId: tem.SlotId,
			Type:   tem.Type,
			Column: i,
			GenId:  int(tem.ID),
		}
		temStr := ""
		for _, tag := range tags {
			temStr += fmt.Sprintf("%s,", tag.Name)
		}
		temStr += "\n"
		item.Layout = temStr
		items = append(items, item)
	}
	global.GVA_DB.Create(&items)
}
