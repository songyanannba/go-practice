package main

import "fmt"

type Dog struct {
	name string
}

func (d Dog) Call() {
	fmt.Println("call", d.name)
}

func (d *Dog) SetName(name string) {
	d.name = name
	fmt.Println("set name")
}

func main() {
	d1 := Dog{}
	d1.Call()
	d1.SetName("a")
	d1.Call()

	d2 := &Dog{}
	d2.Call()
	d2.SetName("dog")

	d1.Call()
	d2.Call()

}
