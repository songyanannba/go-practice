package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	defer conn.Close()

	if err != nil {
		fmt.Println("net dial err ", err)
		return
	}

	//write, _ := conn.Write([]byte("hello world"))
	bytes := make([]byte, 1024)
	read, err := conn.Read(bytes)
	if err != nil {
		fmt.Println("conn read err ", err)
		return
	}

	fmt.Println("read int", read)
	fmt.Println(string(bytes))

}
