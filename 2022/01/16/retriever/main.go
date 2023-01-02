package main

import (
	"fmt"
	"go-2022/xue-xi/base-zhishi/2022/01/16/retriever/mck"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("www.bai.com")
}

func main() {

	var r Retriever
	 r = mck.Retriever{"mck ..hhh"}
	 s := "www"
	get := r.Get(s)
	fmt.Println(r)
	fmt.Println(get)
}