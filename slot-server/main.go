package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	_ "net/http/pprof"
	"os"
	"slot-server/core"
	"slot-server/model/business"
	"slot-server/service/template"
)

func main() {
	core.BaseInit()
	ReadJsonTOTempStruct()
	//fmt.Println("slotTemplateGen", slotTemplateGen)
	template.CreateTemplate(&slotTemplateGen)
	fmt.Println("end======================end")
}

var slotTemplateGen business.SlotTemplateGen

func ReadJsonTOTempStruct() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}
	fmt.Println(dir)

	// 读取文件内容
	file, err := os.Open(dir + "/GenTemp.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 读取文件内容到字节切片
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	// 解析JSON数据并赋值到结构体
	err = json.Unmarshal(data, &slotTemplateGen)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
}
