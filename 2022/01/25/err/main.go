package main

import (
	"fmt"
	"strconv"
)

func fun1(){
	parseInt, err := strconv.ParseInt("11q" ,10 ,64)
	fmt.Println(parseInt)
	fmt.Println(err)

	defer func() {
		i2 := recover()
		if i2 != nil {
			fmt.Println("1111")
			fmt.Println(i2)
			fmt.Println("333")
		}
	} ()
	fmt.Println("2222")
	var i *int
	//i = new(int)
	*i = 3

	fmt.Println("hdkgfjhsdgf")
}

func fun2() {
	day := 9
	switch day {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 7:
		fmt.Println(7)
	default:
		fmt.Println(666666)
	}
}

func main() {

	//fun1()
	fun2()
}
