package main

import (
	"fmt"
	"strings"
)

func main() {
	type StringReader struct {
		data []string
		step int
	}

	segments := []string{"zg", "日本語"}

	fmt.Println(segments)

	want := strings.Join(segments, "+")


	fmt.Println(want)
}
