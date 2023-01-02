package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"recService/service"
)

func main() {

	addr := ":9999"

	rpc.Register(&service.Calc{})
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	log.Printf("[+] leson addr %s\n" ,addr)
	for {
		accept, err := listen.Accept()
		if err != nil {
			log.Printf("[-]client err code :%s\n"  ,err.Error())
			continue
		}
		log.Printf("[+]client  conn :%s\n"  ,accept.RemoteAddr())
		go jsonrpc.ServeConn(accept)
	}

}
