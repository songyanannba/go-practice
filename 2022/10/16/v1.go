package main

import "fmt"

type Benz struct {

}

func(b *Benz)Run() {
	fmt.Println("奔驰")
}

type BMW struct {

}

func (m *BMW) Run() {
	fmt.Println("宝马")
}

type Zhang3 struct {

}

func (z *Zhang3)Drive(benz *Benz)  {
	fmt.Println("张三 开")
	benz.Run()
}

type Li4 struct {

}

func (l *Li4)Drive(bmw *BMW)  {
	fmt.Println("李四 开")
	bmw.Run()
}
