package tests

import (
	"fmt"
	"testing"
)

type Test struct {
	Name  string
	Nge   int
	Test1 *Test1
}
type Test1 struct {
	Name string
	Nge  int
	Test *Test
}

func TestName(t *testing.T) {
	test := &Test{Name: "test", Nge: 1}
	test1 := &Test1{Name: "test1", Nge: 2}
	test.Test1 = test1
	test1.Test = test
	fmt.Printf("%+v\n", test)
}
