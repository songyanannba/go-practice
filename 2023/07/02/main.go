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

	IsWild bool
}

func GetTags(y int) []*Tag {

	tag1 := &Tag{
		Id:     0,
		Name:   "11",
		X:      0,
		Y:      y,
		IsWild: false,
	}

	tag2 := &Tag{
		Id:     1,
		Name:   "",
		X:      1,
		Y:      y,
		IsWild: true,
	}

	tag3 := &Tag{
		Id:     2,
		Name:   "",
		X:      2,
		Y:      y,
		IsWild: false,
	}

	tag4 := &Tag{
		Id:     3,
		Name:   "",
		X:      3,
		Y:      y,
		IsWild: false,
	}

	tag5 := &Tag{
		Id:     4,
		Name:   "",
		X:      4,
		Y:      y,
		IsWild: true,
	}

	tag6 := &Tag{
		Id:     5,
		Name:   "55",
		X:      5,
		Y:      y,
		IsWild: false,
	}

	var tags []*Tag

	tags = append(tags, tag1)
	tags = append(tags, tag2)
	tags = append(tags, tag3)
	tags = append(tags, tag4)
	tags = append(tags, tag5)
	tags = append(tags, tag6)

	return tags

}

func GetTagss() [][]*Tag {

	ta0 := GetTags(0)
	ta1 := GetTags(1)
	ta2 := GetTags(2)
	ta3 := GetTags(3)
	ta4 := GetTags(4)
	ta5 := GetTags(5)
	var tags [][]*Tag
	tags = append(tags, ta0)
	tags = append(tags, ta1)
	tags = append(tags, ta2)
	tags = append(tags, ta3)
	tags = append(tags, ta4)
	tags = append(tags, ta5)

	return tags
}

func Drop6() {

	tagss := GetTagss()
	tags1 := GetTagss()[3]

	sort.Slice(tags1, func(i, j int) bool {
		return tags1[i].X > tags1[j].X
	})

	var tags []*Tag
	for _, v := range tags1 {
		if v.Name == "" {
			tags = append(tags, v)
		}
	}

	for _, i2 := range tags {
		for i := i2.X; i > 0; i-- {
			//wild位置不能动
			if tagss[i][i2.Y].IsWild {
				continue
			}
			if tagss[i-1][i2.Y] == nil {
				continue
			}
			if tagss[i-1][i2.Y].IsWild {
				if i-2 < 0 || tagss[i-2][i2.Y] == nil {
					continue
				}
				tagss[i][i2.Y], tagss[i-2][i2.Y] = tagss[i-2][i2.Y], tagss[i][i2.Y]
			} else {
				tagss[i][i2.Y], tagss[i-1][i2.Y] = tagss[i-1][i2.Y], tagss[i][i2.Y]
			}
		}
	}

	fmt.Println("=", tagss)
	fmt.Println("=", tagss)

}
func main() {
	Drop6()
}
