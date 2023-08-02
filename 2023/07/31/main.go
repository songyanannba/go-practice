package main

import "fmt"

func main() {

	//a := map[int]int{}
	////a = make(map[int]int)
	//
	//a[1] = 2
	//a[2] = 2123
	//a[3] = 2123
	//
	//fmt.Println(a)
	//
	//b := map[string]string{}
	//
	//b["aa"] = "qq"
	//fmt.Println(b)

	// 声明map
	var a map[string]string

	// 在使用map前需要使用make分配空间
	a = make(map[string]string, 10)
	a["username"] = "张三"
	a["addr"] = "张家界"

	fmt.Println(a) // map[addr:张家界 username:张三]

}
