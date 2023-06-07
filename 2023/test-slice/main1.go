package main

import (
	"fmt"
	"time"
)

func ArrVertical[T any](arr [][]T) [][]T {
	if len(arr) == 0 {
		return arr
	}
	var res [][]T
	for i := 0; i < len(arr[0]); i++ {
		var row []T
		for j := 0; j < len(arr); j++ {
			row = append(row, arr[j][i])
		}
		res = append(res, row)
	}
	return res
}


func main() {

	var aa  [][]int

	aa = [][]int {
		{1,2,3,4,5},
		{6,7,8,9,10},
		{11,12,13,14,15},
		{21,22,23,24,25},
		{16,17,18,19,20},
		{26,27,28,29,30},
	}


	//llen := 1000000
	llen := 10000000

	tt1 := time.Now().UnixMilli()
	for i := 0; i<llen; i++ {
		ArrVertical(aa)
	}
	tt2 := time.Now().UnixMilli()

	fmt.Println("tt1 = " , tt1)
	fmt.Println("tt2 = " , tt2)
	fmt.Println("tt2-tt1" , tt2-tt1)



}