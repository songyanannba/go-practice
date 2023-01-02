package main

import (
	"fmt"
	"sort"
)

type User struct {
	Id   int
	Name string
}

func main() {
	list := [][2]int{{1, 3}, {5, 9}, {4, 5}, {6, 2}, {5, 8}}
	fmt.Println(list)
	sort.Slice(list, func(i, j int) bool {
		return list[i][0] < list[j][0]
	})

	fmt.Println(list)

	user := []User{{2, "hh"}, {1, "aa"}, {5, "tt"}, {3, "kl"}}

	sort.Slice(user, func(i, j int) bool {
		return user[i].Id < user[j].Id
	})

	fmt.Println(user)
}
