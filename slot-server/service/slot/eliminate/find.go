package eliminate

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/service/slot/base"
)

func (t *Table) FindLine(verify *Verify) []*base.Tag {
	verify.Add()

	tags := make([]*base.Tag, 0)
	if verify.count > 100 {
		global.GVA_LOG.Error(fmt.Sprintf("超过100次: %+v", verify.site))
		return tags
	}
	tags = append(tags, verify.site.Copy())
	nearbys := t.GetAdjacent(verify.site.X, verify.site.Y)

	for _, tag := range nearbys {
		if tag.Name == "" || tag.Name == "scatter" || verify.GetVerify(tag.X, tag.Y) {
			continue
		}
		if tag.Name == verify.site.Name || tag.IsWild {
			verify.SetSite(tag)
			getTags := t.FindLine(verify)
			tags = append(tags, getTags...)
		}
	}
	return tags
}

func (t *Table) AndFindLine(verify *Verify) {
	for _, i2 := range verify.sites {
		verify.SetSite(i2)
	}
	for _, i2 := range verify.sites {
		verify.SetSite(i2)
		t.FindLine(verify)
	}
}

func (t *Table) FindAllLine() [][]*base.Tag {
	tagList := make([][]*base.Tag, 0)
	verify := NewVerify()
	for _, tags := range t.TagList {
		for _, tag := range tags {
			if tag.Name != "" && !verify.GetVerify(tag.X, tag.Y) && !tag.IsWild {
				verify.ResetVerifyBlank(t)
				verify.SetSite(tag)
				verify.ResetCount()
				getTags := t.FindLine(verify)
				if len(getTags) >= enum.GetLine {
					tagList = append(tagList, getTags)
				}
			}
		}
	}
	return tagList
}

type MayLine struct {
	Id   int
	Name string
	Tags []*base.Tag
	Mul  float64
	Long int
}

func (t *Table) FindAllMayLine(tags []*base.Tag) []MayLine {
	mayLines := make([]MayLine, 0)
	verify := NewVerify()
	for _, tag := range tags {
		if tag.Name != "" && !verify.GetVerify(tag.X, tag.Y) {
			verify.SetSite(tag)
			verify.ResetCount()
			verify.ResetVerifyBlank(t)
			//通过 空白周边的标签去 寻找是否有消除的可能
			getTags := t.FindMayLine(verify)

			for i := 5; i <= len(getTags); i++ {
				mul, _ := t.MatchWinMul(getTags).Float64()
				mayLines = append(mayLines, MayLine{
					Id:   tag.Id,
					Name: tag.Name,
					Tags: getTags[0:i],
					Mul:  mul,
					Long: i,
				})
			}
		}
	}

	return mayLines
}

func (t *Table) FindPTLine() []MayLine {
	mayLines := make([]MayLine, 0)
	needTags := t.NeedFill()
	if len(needTags) == 0 {
		return mayLines
	}
	verify := NewVerify()
	verify.SetSite(needTags[0])
	blanks := t.FindMayLine(verify)
	pTables := lo.Filter(t.PayTableList, func(item *base.PayTable, index int) bool {
		mul, _ := decimal.NewFromFloat(t.Mul).Add(decimal.NewFromFloat(item.Multiple)).Float64()
		return mul <= t.Target.MaxMul && len(item.Tags) <= len(blanks)
	})

	for _, table := range pTables {
		mayLines = append(mayLines, MayLine{
			Id:   table.Tags[0].Id,
			Name: table.Tags[0].Name,
			Tags: blanks[0:len(table.Tags)],
			Mul:  table.Multiple,
		})
	}
	return mayLines
}

func (t *Table) FindMayLine(verify *Verify) []*base.Tag {
	verify.Add()

	result := make([]*base.Tag, 0)
	if verify.count > 100 {
		global.GVA_LOG.Error(fmt.Sprintf("超过100次: %+v", verify.site))
		return result
	}
	tags := make([]*base.Tag, 0)
	tags = append(tags, verify.site.Copy())
	nearbys := t.GetAdjacent(verify.site.X, verify.site.Y)

	for _, tag := range nearbys {
		if tag.Name == "scatter" || verify.GetVerify(tag.X, tag.Y) {
			continue
		}
		if tag.Name == verify.site.Name || tag.Name == "" {
			verify.SetSite(tag)
			getTags := t.FindMayLine(verify)
			tags = append(tags, getTags...)
		}
	}
	return tags
}

func (t *Table) FindQuantityTags(num int) [][]*base.Tag {
	tags := t.WindowToArr()
	tagsMap := lo.GroupBy(tags, func(item *base.Tag) string {
		return item.Name
	})
	tagList := make([][]*base.Tag, 0)
	for tagName, tags := range tagsMap {
		if len(tags) >= num && tagName != "" && tagName != "scatter" && tagName != "multiplier" {
			tagList = append(tagList, tags)
		}
	}
	return tagList
}

func (t *Table) FindTagsUseName(name string) []*base.Tag {
	tags := t.ToArr()
	tagsMap := lo.GroupBy(tags, func(item *base.Tag) string {
		return item.Name
	})
	return tagsMap[name]
}

func (t *Table) FindTagsUseNameInWindows(name string) []*base.Tag {
	tags := t.WindowToArr()
	tagsMap := lo.GroupBy(tags, func(item *base.Tag) string {
		return item.Name
	})
	return tagsMap[name]
}
