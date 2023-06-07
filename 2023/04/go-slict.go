package main

import "fmt"

func intToBytes(n int) []byte {
	buf := make([]byte, 3)
	buf[0] = byte((n >> 16) & 0xFF)
	buf[1] = byte((n >> 8) & 0xFF)
	buf[2] = byte(n & 0xFF)
	return buf
}

func main() {

	fmt.Println(intToBytes(1))
	fmt.Println(intToBytes(2))
	fmt.Println( string(intToBytes(1)) )


	return

	var a []string
	var b  = []string{}

	//fmt.Printf("%T,%V \n", a,a)
	fmt.Printf("%T %T\n", a,b)
	fmt.Printf("%#v %#v\n", a,b)

	fmt.Println(len(b))
	b = append(b ,"aaa")
	fmt.Println(b)
	fmt.Println(len(a))
	a = append(a ,"bbb")
	fmt.Println(b)
}
