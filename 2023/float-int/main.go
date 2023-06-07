package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

func f2i(f float64) int {
	i, _ := strconv.Atoi(fmt.Sprintf("%1.0f", f ))
	fmt.Println("i")
	return i
}

func main() {

	x := 0.1
	fmt.Println("==== ",   strconv.FormatFloat(x, 'f', -1, 64))

	//x := 2.0
	y := 2
	g := 2

	g *= y
	fmt.Println(g)

	d  := int(decimal.NewFromFloat(x).Mul(decimal.NewFromFloat(float64(y))).IntPart())

	fmt.Println("d =" ,d)
}


