package main

import (
	"fmt"
	"time"
)

func ArrVertical1[T any](arr [][]T) [][]T {
	if len(arr) == 0 {
		return arr
	}
	var res [][]T

	lieLen := len(arr[0]) //多少个数组
	//fmt.Println("lieLen = ", lieLen)

	//var row map[int][]T
	row := make(map[int][]T ,10)
	for j := 0; j < lieLen ; j++ {
		var r []T
		row[j] = r
	}

	for j := 0; j < len(arr); j++ {

		for i := 0 ; i < lieLen; i++ {
			//fmt.Println("---",   arr[j][i])
			//res = append(res, ts01[i])
			row[i] = append(row[i], arr[j][i])
		}
	}

	for _ , vv :=  range row {
		res = append(res, vv)
	}


	return res
}

func main()  {

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

	tt3 := time.Now().UnixMilli()
	for i := 0; i<llen; i++ {
		ArrVertical1(aa)
	}
	tt4 := time.Now().UnixMilli()

	fmt.Println("tt3 = " , tt3)
	fmt.Println("tt4 = " , tt4)
	fmt.Println("tt4-tt3" , tt4-tt3)


}