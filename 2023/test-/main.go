package main

import (
	"fmt"
)

func (s *Spin) Dump() *ViewData {
	var view = &ViewData{}
	view.Data = dataStr(s.InitDataList, func(tag *Tag) string {
		return tag.Dump()
	}, ";")

	return view
}

func main() {
	var tts = [][]*Tag{}
	var tags []*Tag

	t1 := &Tag{
		Id:         15,
		Name:       "high_3",
		Include:    nil,
		Multiple:   0,
		X:          1,
		Y:          2,
		IsLine:     false,
		IsPayTable: false,
		IsWild:     false,
		IsSingle:   false,
		IsJackpot:  false,
		ISLock:     false,
	}

	t2 := &Tag{
		Id:         16,
		Name:       "high_3",
		Include:    nil,
		Multiple:   0,
		X:          2,
		Y:          2,
		IsLine:     false,
		IsPayTable: false,
		IsWild:     false,
		IsSingle:   false,
		IsJackpot:  true,
		ISLock:     false,
	}

	t3 := &Tag{
		Id:         25,
		Name:       "wild_1",
		Include:    []string{"19"},
		Multiple:   0,
		X:          4,
		Y:          4,
		IsLine:     true,
		IsPayTable: false,
		IsWild:     true,
		IsSingle:   false,
		IsJackpot:  false,
		ISLock:     false,
	}

	tags = append(tags, t1)
	tags = append(tags, t2)
	tags = append(tags, t3)
	tts = append(tts, tags)

	//dump := t1.Dump()
	//fmt.Println("dump == " , dump)

	str := dataStr(tts, func(tag *Tag) string {
		return tag.Dump()
	}, ";")

	fmt.Println("dump str == ", str)



	bin := DecToBin("64")
	fmt.Println("bin 64= ", bin)

	bin66 := DecToBin("66")
	fmt.Println("bin 66= ", bin66)

	bin104 := DecToBin("104")
	fmt.Println("bin 104= ", bin104)


	attrs := SplitInt[int](bin, "")
	arr := make([]int, 5)
	copy(arr, attrs)
	fmt.Println(arr)



}
