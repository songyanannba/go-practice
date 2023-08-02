package eliminate

import (
	"slot-server/service/slot/base"
)

type Verify struct {
	site   base.Tag
	verify map[[2]int]bool
	count  int
	sites  []*base.Tag
}

func NewVerify() *Verify {
	return &Verify{
		verify: map[[2]int]bool{},
		sites:  make([]*base.Tag, 0),
	}
}
func (v *Verify) ResetCount() {
	v.count = 0
	v.sites = make([]*base.Tag, 0)
}

func (v *Verify) SetSite(tag *base.Tag) {
	if tag.IsWild {
		nowTag := tag.Copy()
		nowTag.Name = v.site.Name
		v.site = *nowTag
	} else {
		v.site = *tag
	}

	v.verify[[2]int{v.site.X, v.site.Y}] = true

}

func (v *Verify) GetSite() base.Tag {
	return v.site
}

func (v *Verify) SetVerify(x, y int) {
	v.verify[[2]int{x, y}] = true
}

func (v *Verify) GetVerify(x, y int) bool {
	return v.verify[[2]int{x, y}]
}

func (v *Verify) Add() {
	v.count++
}

func (v *Verify) ResetVerify() {
	v.verify = map[[2]int]bool{}
}
func (v *Verify) ResetVerifyBlank(table *Table) {
	for ints, _ := range v.verify {
		if table.TagList[ints[0]][ints[1]].Name == "" || table.TagList[ints[0]][ints[1]].IsWild {
			delete(v.verify, ints)
		}
	}
}
