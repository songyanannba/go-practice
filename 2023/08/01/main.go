package main

import (
	str2 "01/str"
	"fmt"
)

func main() {

	var str1 string
	str1 = "0:scatter&multiplier&high_1&high_2&high_3&high_4&low_1&low_2&low_3&low_4&low_5@0&5&8&42&82&122&162&202&242&282&322&362"

	// 使用类型转换将原生字符串类型转换为自定义字符串类型
	myStr := str2.PreStr(str1)

	fmt.Println("PreStr", myStr)

	str, s, err := myStr.ChangeStrToIntAndStr()
	//
	fmt.Println("===", str, s, err)

}
