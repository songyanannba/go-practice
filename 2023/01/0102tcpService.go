package main

import (
	"bufio"
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

		reader := bufio.NewReader(client)
		writer := bufio.NewWriter(client)

		for  {
			writeString, err := writer.WriteString(time.Now().Format("2006-01-02 15:04:05") + "\n")
			writer.Flush()
			if err != nil {
				fmt.Println(" bufio  writeString err " ,err)
				break
			}
			fmt.Println(" bufio  writeString : " ,writeString)


			readString, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("bufio reader err : ", err)
				break
			}
			fmt.Println("bufio readString  === " ,readString)

			time.Sleep(time.Second * 1)
		}

		//client.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
		fmt.Println("client write : ")
	}



}

