package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	sum := md5.Sum([]byte("11"))
	fmt.Println("sum = " ,sum)
	fmt.Println("sum = " , fmt.Sprintf("sum =  %x " , sum))
}