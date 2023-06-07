package main

import (
	"fmt"
)

const (
	name  string  = "q"
)

type UUser struct {
	name string
	age int
}

func a(u *UUser) {
	u.name = "sss"
	u.age = 20

	//fmt.Println("aa == " , u)
}

func ss(a []string) {
	a[0] = "M"
	a = append(a , "c")
	a = append(a , "d")
	fmt.Println("slice == " , a)
}

func main() {

	u := UUser{
		name: "syn",
		age: 18,
	}

	//fmt.Println("11" , u)
	a(&u)
	//fmt.Println("22" , u)

	slc := []string{
		"a","b",
	}

	fmt.Println("11" , slc)
	ss(slc)
	fmt.Println("33" , slc)

}
