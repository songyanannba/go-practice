package main

import (
	"fmt"
	"net"
)

func GetLocalIp() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("ipnet.IP.To4 == ", ipnet.IP)
				return ipnet.IP, nil
			}
		}
	}
	return nil, nil
}

func main() {
	GetLocalIp() //获取本地IP
}
