package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// 定义处理请求的函数
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}

	// 注册处理函数
	http.HandleFunc("/", handler)

	// 启动HTTP服务器
	http.ListenAndServe(":8888", nil)
}
