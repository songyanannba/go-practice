package main

import (
	"fmt"
	"strconv"

	//"runtime/debug"

)

//func MoveFile(sourcePath string) error {
//
//	getwd, _ := os.Getwd()
//	destGetwd := getwd + "/uploads/slot/"
//
//	/*getwds := getwd + "/uploads/file/"
//	destGetwd := getwd + "/uploads/slot/"
//	sourceFile := getwds + "image1.jpg"
//	destinationFile := destGetwd + "image1.jpg"*/
//	//MoveFile(sourceFile, destinationFile)
//
//	mkdirErr := os.MkdirAll(destGetwd, os.ModePerm)
//	if mkdirErr != nil {
//		fmt.Println("123 === ", mkdirErr)
//	}
//
//	source, err := os.Open(getwd+sourcePath) //open the source file
//	if err != nil {
//		panic(err)
//	}
//	defer source.Close()
//
//	out, createErr := os.Create(destinationFile)
//	if createErr != nil {
//		fmt.Println("createErr", createErr)
//	}
//	defer out.Close() // 创建文件 defer 关闭
//
//	_, copyErr := io.Copy(out, source) // 传输（拷贝）文件
//	if copyErr != nil {
//		fmt.Println("copyErr", copyErr)
//	}
//
//	err = os.Remove(sourceFile)
//	if err != nil {
//		fmt.Println("Remove", err)
//	}
//
//}

func main() {

	atoi, _ := strconv.Atoi("234")
	ato, _ := strconv.Atoi("dj-")

	fmt.Println(atoi,ato)

	//MoveFile("uploads/file/image1.jpg")
}
