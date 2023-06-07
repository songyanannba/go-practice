package main

import (
	"fmt"
	"math"
)

func Solution1(A []int) int {
	// Implement your solution here
	m := map[int]struct{}{}
	for i := 0 ; i < len(A) ; i++ {
		m[A[i]] = struct{}{}
	}

	for i := 1 ; i <= len(A) ; i++ {
		if _, ok := m[i] ; !ok {
			return i
		}
	}

	return len(A) + 1
}

func reverse(N int)  {
	enable_print := 0
	for N != 0 {
		if enable_print < math.MinInt32/10 || enable_print > math.MaxInt32/10 {
			fmt.Println(0)
		}
		digit := N % 10
		N /= 10
		enable_print = enable_print*10 + digit
	}
	fmt.Println(enable_print)
}

func Solution(N int) {
	var enable_print int;
	enable_print = N % 10;
	for N > 0 {
		if enable_print == 0 && N % 10 != 0 {
			enable_print = 1;
		} else if enable_print == 1 {
			fmt.Print(N % 10);
		}
		N = N / 10;
	}
	fmt.Print(N);
}


func main() {

	Solution(123)


}
