package main

import "fmt"

type table struct {
	Id   int
	Name string
}

func (t *table) Copy() *table {
	nt := *t
	return &nt
}

type Spin struct {
	stb *table
}

type Machine struct {
	tb *table
}

func (m *Machine) aa() {
	t := m.tb.Copy()
	t.Name = "我变了"
	m.tb = t
}

func main() {

	t := &table{
		Id:   1,
		Name: "我是一",
	}

	//t2 := t.Copy()
	//t2.Name = "我是谁"
	//fmt.Println(t2)

	m := &Machine{
		tb: t,
	}

	m.aa()

	ss := &Spin{
		stb: m.tb,
	}

	fmt.Println("t = ", ss)

}
