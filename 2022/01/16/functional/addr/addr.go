package main

import "fmt"

func addr() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		//fmt.Printf("hh %v\n" ,v)
		return sum
	}
}

type iAddr func(int) (int, iAddr)

func addr2(base int) iAddr {
	return func(v int) (int, iAddr) {
		return base + v, addr2(base + v)
	}
}

func main() {
/*	a := addr()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}*/

	a := addr2(0)
	for i := 0; i < 10; i++ {
		var s int
		s , a = a(i)
		fmt.Printf("0 + 1-orm + ... %d = %d\n" , i ,s)
	}
}
