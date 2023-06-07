package main

import "fmt"

type Tg struct {
	Id       int
	X int
	Y int
}

func (t *Tg) Copy() *Tg {
	c := *t
	fmt.Println(" t ", t)
	return &c
}


func (t *Tg) Tset1() *Tg {
	c := *t
	c.Id = 1
	c.Y = 89
	return &c
}


func GetTg() {
	var t Tg
	t.X = 2
	tg := t.Copy()
	fmt.Println(tg)
	t.X = 2
	//t.Y = 3
	/*fmt.Println("t ==" , t)
	fmt.Println("t %p ==" , t)

	tg := t.Copy()
	fmt.Println("tg =="  ,&tg)

	tset1 := t.Tset1()
	fmt.Println("tset1 =="  ,&tset1, tset1)*/
}

func main()  {
	GetTg()
}