package helper

import (
	"testing"
)

func TestGetColumnInArr(t *testing.T) {
	type test struct {
		Name string
		Age  int
	}
	var (
		arr = []test{
			{
				Name: "a",
				Age:  1,
			},
			{
				Name: "b",
				Age:  2,
			},
			{
				Name: "c",
				Age:  3,
			},
		}
	)
	res := GetColumnInArr(arr, func(data test) string {
		return data.Name
	})
	t.Log(res)
}

func TestGetDistinctColumnInArr(t *testing.T) {
	type test struct {
		Name string
		Age  int
	}
	var (
		arr = []test{
			{
				Name: "a",
				Age:  1,
			},
			{
				Name: "a",
				Age:  5,
			},
			{
				Name: "b",
				Age:  2,
			},
			{
				Name: "c",
				Age:  3,
			},
		}
	)
	res := DistinctByFunc(arr, func(data test) string {
		return data.Name
	})
	t.Log(res)
}
