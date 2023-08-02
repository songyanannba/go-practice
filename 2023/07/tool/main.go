package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("===")

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}
	fmt.Println(dir)

	// 读取文件内容
	file, err := os.Open("/GenTemp.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	//// 读取文件内容到字节切片
	//data, err := ioutil.ReadAll(file)
	//if err != nil {
	//	fmt.Println("Error reading file:", err)
	//	return
	//}
	//
	//var slotTemplateGen SlotTemplateGen
	//
	//// 解析JSON数据并赋值到结构体
	//err = json.Unmarshal(data, &slotTemplateGen)
	//if err != nil {
	//	fmt.Println("Error unmarshaling JSON:", err)
	//	return
	//}
	//
	//// 打印解析后的结果
	//fmt.Println("Name:", slotTemplateGen)

}
