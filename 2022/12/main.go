package main

import (
	"fmt"
	"reflect"
)

//反射
func f1()  {
	var i int = 1
	var t reflect.Type = reflect.TypeOf(i)
	fmt.Println(t)
 
	var s string = "qwer"
	var str reflect.Type = reflect.TypeOf(s)
	fmt.Println(str.Name())

}

func main() {

	f1()

}
