package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)


type Int interface {
	int8 | int | int32 | int64 | uint8 | uint16 | uint32 | uint | uint64
}

func BinToDec(s string) string {
	v, _ := strconv.ParseInt(s, 2, 64)
	//fmt.Println("vv == ", v)
	return strconv.FormatInt(v, 10)
}

func DecToBin(s string) string {
	v, _ := strconv.ParseInt(s, 10, 64)
	return strconv.FormatInt(v, 2)
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

func (t Tag) Dump(typ ...int) string {
	var name string
	if len(typ) > 0 {
		name = t.Name
	} else {
		name = strconv.Itoa(t.Id) + If(t.Multiple > 1, "@"+strconv.Itoa(t.Multiple), "")
	}
	strconv.ParseInt(name, 10, 64)
	attr := ""
	attr += If(t.IsLine, "1", "0")
	attr += If(t.IsPayTable, "1", "0")
	attr += If(t.IsWild, "1", "0")
	attr += If(t.IsSingle, "1", "0")
	attr += If(t.IsJackpot, "1", "0")
	attr += If(t.ISLock, "1", "0")
	if attr != "" {
		name += "|" + BinToDec("1"+attr)
	}

	if attr != "" {
		zb := zuobiao{
			X: 11,
			Y: 22,
		}
		marshal, _ := json.Marshal(zb)
		fmt.Println("---",string(marshal))
		name += "#" + string(marshal)
	}

	return name
}

func dataStr[T any](arr [][]T, f func(T) string, sep string) string {
	fmt.Println("dataStr", arr)
	arr = ArrVertical(arr)
	str := ""
	for i, row := range arr {
		for ii, col := range row {
			str += f(col)
			if ii < len(row)-1 {
				str += ","
			}
		}
		if i < len(arr)-1 {
			str += sep
		}
	}
	return str
}



func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}


func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func SplitInt[V Int](s string, sep string) []V {
	arr := strings.Split(s, sep)
	var intArr []V
	for _, s := range arr {
		v, _ := strconv.Atoi(s)
		intArr = append(intArr, V(v))
	}
	return intArr
}
