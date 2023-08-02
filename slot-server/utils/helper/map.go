package helper

import (
	"golang.org/x/exp/slices"
)

type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// SortMapByFunc 自定义map排序
func SortMapByFunc[K comparable, V any](m map[K]V, sortFunc func(a, b Pair[K, V]) bool) []Pair[K, V] {
	p := make([]Pair[K, V], len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair[K, V]{Key: k, Value: v}
		i++
	}
	slices.SortFunc(p, sortFunc)
	return p
}

func ArrToMap[V comparable](arr []V) map[V]bool {
	var m = make(map[V]bool, len(arr))
	for _, v := range arr {
		m[v] = true
	}
	return m
}
