package main

import (
	"fmt"
	"os"
)

//	文件操作
//1 读
//2 写

//文件路径
//绝对路径 ：文件从根目录查找
//相对路径 ：与程序位置有关系

func f1() {

	path := "/Users/songyanan/GolandProjects/go-2022/xue-xi/base-随便练习/2022/10/29/03文件/user.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open err ", err)
		return
	}
	//fmt.Println("open = " ,file)
	//fmt.Printf("%T = " ,file)

	context := make([]byte, 10)

	n, err := file.Read(context)
	fmt.Println("read err = ", err)
	fmt.Println("n = ", n)
	fmt.Println("context", context)
	fmt.Println("context1", string(context[:n]))
	fmt.Println("context2 = ", string(context))

	_, err = file.Read(context)
	fmt.Println("read11 err = ", err)
	file.Close()

}

func main() {
	path := "user2.txt"
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("create file err:", err)
		return
	}

	file.Write([]byte("123aaaAAA"))

	file.Close()

}
