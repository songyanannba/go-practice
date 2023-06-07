package main

import "fmt"

type test1 struct {
}

func (t *test1) hhh() {
	fmt.Println(21)
}

type test2 struct {
	test1
}

type testGroup struct {
	test2
}

type test3 struct {
	testGroup
}

func main() {
	var tt test3
	tt.hhh()
}
