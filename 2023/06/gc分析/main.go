package main

import (
	"fmt"
)

type Tag struct {
	Id       int
	Name     string
	X int
	Y int
}

type Table struct {
	Row          int               // 行数
	Col          int               // 列数
	Tags         []Tag        // 所有tag
	Scatter      *Tag         // scatter
}

type Verify struct {
	site   Tag
	verify map[[2]int]bool
	count  int
	sites  []*Tag
}
func (v *Verify) SetVerify(x, y int) {
	v.verify[[2]int{x, y}] = true
}

func NewVerify() *Verify {
	return &Verify{
		verify: map[[2]int]bool{},
		sites:  make([]*Tag, 0),
	}
}

func (t *Table) InitFill() {

	verify := NewVerify()
	fillTags := t.SpecifyFill(verify)

	fmt.Println("==",verify)
	fmt.Println(fillTags)

}

func (t *Table) SpecifyFill(verify *Verify) []Tag {

	var  v    map[[2]int]bool
	v = make(map[[2]int]bool, 0)
	v[[2]int{1,2}] = true
	v[[2]int{2,4}] = true
	v[[2]int{4,6}] = true

	for ints, _ := range v {
		verify.SetVerify(ints[0], ints[1])
	}
	return nil
}

func main() {

	var t Table
	t.InitFill()
}