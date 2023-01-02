package main

import "fmt"

//定义自定义类型
func f1() {
	type Zi_ding_yi int
	var zdy Zi_ding_yi
	fmt.Println("zdy := ", zdy)
	fmt.Println("zdy := ", zdy+10)

	type User map[string]string
	user := make(User)
	user["ss"] = "s1"
	user["yy"] = "y1"
	fmt.Println(user)
	fmt.Printf("%#v\n", user)

	type Us struct {
		i    int
		t    int
		Name string
	}
	var uus Us
	uus.Name = "sssyyynnn"
	/*	fmt.Printf("%T\n" ,uus)
		fmt.Printf("%#v\n" ,uus)
		fmt.Printf("==\n" ,uus)
		fmt.Println("==+")
		fmt.Println(uus)*/
	fmt.Println(&uus)
	fmt.Printf("%#v\n", &uus)

}

func main() {

	f1() //定义自定义类型
}
