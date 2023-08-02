package eliminate

import (
	"slot-server/service/slot/base"
	"sort"
)

func (t *Table) Drop() [][]*base.Tag {
	tags := t.NeedFill()

	sort.Slice(tags, func(i, j int) bool {
		return tags[i].X < tags[j].X
	})

	for _, i2 := range tags {
		for i := 0; i < i2.X; i++ {
			if t.TagList[i][i2.Y].IsWild || t.TagList[i][i2.Y].Name == "" {
				continue
			}
			t.TagList[i2.X][i2.Y], t.TagList[i][i2.Y] = t.TagList[i][i2.Y].Copy(), t.TagList[i2.X][i2.Y].Copy()

		}
	}
	t.SetCoordinates()
	return t.GetGraph()
}

func (t *Table) DropExistWild() [][]*base.Tag {
	tags := t.NeedFill()
	sort.Slice(tags, func(i, j int) bool {
		return tags[i].X < tags[j].X
	})

	for _, i2 := range tags {
		for i := i2.X; i > 0; i-- {
			//wild位置不能动
			if t.TagList[i][i2.Y].IsWild {
				continue
			}
			if t.TagList[i-1][i2.Y] == nil {
				continue
			}

			if t.TagList[i-1][i2.Y].IsWild {
				if i-2 < 0 || t.TagList[i-2][i2.Y] == nil || t.TagList[i-2][i2.Y].IsWild {
					continue
				}
				t.TagList[i][i2.Y], t.TagList[i-2][i2.Y] = t.TagList[i-2][i2.Y], t.TagList[i][i2.Y]
			} else {
				t.TagList[i][i2.Y], t.TagList[i-1][i2.Y] = t.TagList[i-1][i2.Y], t.TagList[i][i2.Y]
			}

		}
	}
	t.SetCoordinates()
	return t.GetGraph()
}
