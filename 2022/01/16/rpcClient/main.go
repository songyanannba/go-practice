package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"rpcClient/data"
)

func main() {

	addr := "127.0.0.1-orm:9999"

	dial, err := jsonrpc.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("[-] dial err:%s\n" ,err)
	}
	defer dial.Close()

	request := &data.CalcRequest{Left: 2 , Right: 7}
	response := &data.CalcResponse{}

	//发起请求
	err = dial.Call("Calc.Add", request, response)
	fmt.Println("hhhh")
	fmt.Println(response.Result)

}
