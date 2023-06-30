package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IntervalRatioEvent6 struct {
	TotalFillNum map[int]int
	CenterPoint  map[int]int
	Data         [][2]int
	weight       []int
}

func main() {
	strs := []string{
		"10@1@4&8@1&2&3@1&2&3@1@1@1@1",
		"20@2@8&12@1&2&3@1&2&3@1@2@3@4",
		"40@2@8&12@1&2&3@1&2&3@1@2@3@4",
	}

	intervalRatioEvent6 := IntervalRatioEvent6{
		TotalFillNum: make(map[int]int),
		CenterPoint:  make(map[int]int),
	}

	for _, str := range strs {
		splits := strings.Split(str, "@")
		for k, v := range splits {
			if k == 0 {
				num, _ := strconv.Atoi(v)
				intervalRatioEvent6.TotalFillNum[len(intervalRatioEvent6.TotalFillNum)+1] = num
			}
			if k == 1 {
				num1, _ := strconv.Atoi(v)
				intervalRatioEvent6.CenterPoint[len(intervalRatioEvent6.CenterPoint)+1] = num1
			}
		}

	}

	fmt.Println("=")

}
