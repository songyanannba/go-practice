package main

import (
	"fmt"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	// [1]
	//rand.Seed(time.Now().Unix()) //打开此处得到正确的随机数
	for i := range b {
		// [2]
		//rand.Seed(time.Now().Unix()) //打开此处每次得到相同的随机数
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	//打开【2】得到类似 [71 71 71 71 71 71 71 71 71 71]的相同结果
	//打开【1】得到类似 [75 111 79 115 101 73 120 66 71 83]的正常结果
	fmt.Println(b)
	return string(b)
}

func main() {
	s := randomString(100)
	fmt.Println(s)
}
