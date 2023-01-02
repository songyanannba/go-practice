package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	addr := "0.0.0.0:9999"
	listener, err := net.Listen("tcp", addr)
	defer listener.Close()

	if err != nil {
		fmt.Println("net listen : " ,err)
		os.Exit(-1)
		return
	}

	fmt.Println("listen : " ,addr)

	for  {
		client, err := listener.Accept()
		defer client.Close()

		if err != nil {
			fmt.Println("listener accept " , err)
			continue
		}

		client.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
		fmt.Println("client write : ")
	}



}

