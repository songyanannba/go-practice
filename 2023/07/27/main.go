package main

import "fmt"

func main() {

	pp := make(map[int]int)

	pp[5] = 5
	pp[6] = 6
	pp[7] = 7
	pp[9] = 9
	pp[11] = 11

	for i := 10; i >= 5; i-- {
		if _, okk := pp[i]; okk {

			fmt.Println("===", i)
			break
		}
	}
}
