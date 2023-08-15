package main

import (
	"fmt"
	"sort"
)

type Tag struct {
	Id   int
	Name string
	X    int
	Y    int
}

func main() {

	tags := []*Tag{}
	tag1 := &Tag{
		Id:   0,
		Name: "11",
		X:    4,
		Y:    0,
	}
	tag2 := &Tag{
		Id:   1,
		Name: "22",
		X:    1,
		Y:    0,
	}
	tag3 := &Tag{
		Id:   2,
		Name: "33",
		X:    4,
		Y:    1,
	}
	tags = append(tags, tag1)
	tags = append(tags, tag2)
	tags = append(tags, tag3)

	fmt.Printf("%v", tags)

	sort.Slice(tags, func(i, j int) bool {
		return tags[i].X < tags[j].X
	})

	fmt.Printf("%v", tags)
}
