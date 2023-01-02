package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	defer conn.Close()

	if err != nil {
		fmt.Println("net dial err ", err)
		os.Exit(-1)
		return
	}

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	readString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("ReadString err " , err)
	}
	fmt.Println("readString : ", readString)

	writeString, err := writer.WriteString("i am nb \n")
	writer.Flush()
	if err != nil {
		fmt.Println("WriteString err :" ,err)
	}

	fmt.Println("writeString = " ,writeString)




/*	bytes := make([]byte, 1024)
	read, err := conn.Read(bytes)
	if err != nil {
		fmt.Println("conn read err ", err)
		return
	}

	fmt.Println("read int", read)
	fmt.Println(string(bytes))*/

}
