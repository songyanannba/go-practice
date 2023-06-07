package main

import (
	"fmt"
	"sync"
)

type  New func() interface{}

type Pool struct {
	mutex sync.Mutex
	Obj []interface{}
	New New
}

func NewPool(size int ,new New)  *Pool {
	obj := make([]interface{} ,size)
	for i := 0 ; i < size ;i++ {
		obj[i] = new()
	}

	return &Pool{
		Obj:  obj ,
		New:   new,
	}
}

func (p *Pool)Get() interface{} {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if len(p.Obj) >0 {
		o := p.Obj[0]
		p.Obj = p.Obj[1:]
		return o
	} else {
		return p.New()
	}
}


func (p *Pool)Put (obj interface{}) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.Obj = append(p.Obj , obj)
}


func main() {

	p := NewPool(3 , func() interface{} {
		fmt.Println("new 11")
		return 3
	})
	fmt.Println(  "p1",p)
	x := p.Get()
	fmt.Println(x)
	fmt.Println(  "p2",p)
	p.Put(x)
	fmt.Println(  "p3",p)

}
