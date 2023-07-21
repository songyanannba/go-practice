package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SliceVal[T any](arr []T, key int) T {
	if len(arr) > key {
		return arr[key]
	}
	var a T
	return a
}

func SplitInt[V int](s string, sep string) []V {
	arr := strings.Split(s, sep)
	var intArr []V

	for _, s := range arr {
		v, _ := strconv.Atoi(s)
		intArr = append(intArr, V(v))
	}
	return intArr
}

func main() {
	//str := "scatter&low_1&low_2&high_4&low_2&low_2&low_3&high_1&low_3&scatter&low_1&low_1&low_3&high_3&low_2&low_2&low_4&low_4&wild_1&low_1&low_1&scatter&low_2&low_2&low_2&low_3&high_4&low_3&link_coin&link_coin&link_coin&low_3&low_4&low_4&high_2&low_1&low_3&low_3&scatter&low_2&low_4&high_3&low_1&low_1&link_coin&link_coin&link_coin&low_4&low_4&wild_1&low_1&low_1&low_2&high_4&low_2&low_2&low_3&scatter&low_3&link_coin&link_coin&link_coin&low_1&low_1&low_3&high_3&low_2&low_2&low_4&scatter&low_1&low_1&low_4&low_4&high_2&low_1&low_2&scatter&low_2&low_2&low_3&high_4&low_3&low_3&wild_1&low_4&low_4&scatter&low_1&low_3&low_3&scatter&low_2&low_4&high_3&low_1&low_1&low_4&low_4"
	//split := strings.Split(str, "&")
	//fmt.Println(split)

	str1 := "wild_100&wild_10&wild_5&wild_2&wild_1&7_high&bar_7&bar_high&bar_med&bar_low&null@1&1&1&1&1&519&1321&1771&2321&2971&3721&6325"

	before, after, found := strings.Cut(str1, "@")
	fmt.Println(before, after, found)

	split := strings.Split(before, "&")
	splitInt := SplitInt(after, "&")

	fmt.Println(split, splitInt)

}
