package main

import (
	"fmt"
	"time"
)

/*func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i //或者直接写成return
}
func main() {
	fmt.Println("return:", b())
}*/


func c() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i
}


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


func main() {

	var aa  [][]int

	aa = [][]int {
		{1,2},
		{3,4},
		{5,6},
		{7,8},
		{9,10},
	}


	llen := 1000000

	tt3 := time.Now().Nanosecond()
	fmt.Println("tt3 = " , tt3)
	for i := 0; i<llen; i++ {
		ArrVertical1(aa)
	}
	tt4 := time.Now().Nanosecond()

	fmt.Println("tt4 = " , tt4)
	fmt.Println("tt4-tt3" , tt4-tt3)



	tt1 := time.Now().Nanosecond()
	for i := 0; i<llen; i++ {
		ArrVertical(aa)
	}
	tt2 := time.Now().Nanosecond()
	fmt.Println("tt1 = " , tt1)
	fmt.Println("tt2 = " , tt2)
	fmt.Println("tt2-tt1" , tt2-tt1)



}