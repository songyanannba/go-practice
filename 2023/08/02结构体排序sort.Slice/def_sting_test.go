package main

import (
	"fmt"
	"testing"
)

func TestSplitStr(t *testing.T) {

	var str1 string
	str1 = "0:scatter&multiplier&high_1&high_2&high_3&high_4&low_1&low_2&low_3&low_4&low_5@0&5&8&42&82&122&162&202&242&282&322&362"

	myStr := DefString(str1)
	str, s, err := myStr.SplitStr(":")

	fmt.Println("str, s, err", str, s, err)

}

func TestSplitStrToMap(t *testing.T) {

	var str1 string
	//str1 = "0:scatter&multiplier&high_1&high_2&high_3&high_4&low_1&low_2&low_3&low_4&low_5@0&5&8&42&82&122&162&202&242&282&322&362\t\n1:scatter&multiplier&high_1&high_2&high_3&high_4&low_1&low_2&low_3&low_4&low_5@0&5&8&42&82&122&162&202&242&282&322&362"

	str1 = "1+po"
	myStr := DefString(str1)
	colMap, err := myStr.StrToMap("+")

	fmt.Println("colMap, err", colMap, err)

}
