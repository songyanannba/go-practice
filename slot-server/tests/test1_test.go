package tests

import (
	"fmt"
	"slot-server/utils/helper"
	"testing"
)

func TestCF(t *testing.T) {
	s := -3 % 100
	fmt.Println(s)
}

func TestRand(t *testing.T) {
	ssd := make(map[int]int)
	for i := 0; i < 1000000; i++ {
		ssd[helper.RandInt(1000)]++
	}
	fmt.Println(ssd)
}
