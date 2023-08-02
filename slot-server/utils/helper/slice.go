package helper

import (
	"math/rand"
	"slot-server/utils/conver"
	"strconv"
	"strings"
	"time"
)

func SliceDel[T any](arr []T, i int) []T {
	if i > len(arr)-1 {
		return arr
	}
	return append(arr[:i], arr[i+1:]...)
}

// SliceVal 获取切片中的值 不存在则返回零值
func SliceVal[T any](arr []T, key int) T {
	if len(arr) > key {
		return arr[key]
	}
	var a T
	return a
}

func SliceKeyExist[T any](arr []T, key int) bool {
	if len(arr) > key {
		return true
	}
	return false
}

// RangeOverlap 两个区间是否存在重叠
func RangeOverlap(arr1, arr2 []int) bool {
	var (
		a1 = SliceVal(arr1, 0)
		a2 = SliceVal(arr1, 1)

		b1 = SliceVal(arr2, 0)
		b2 = SliceVal(arr2, 1)
	)
	if max(a1, b1) <= min(a2, b2) {
		return true
	}
	return false
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

// InArr 某个值是否在数组中
func InArr[V comparable](v V, sl []V) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// CaseInsensitiveInArr 某个值是否在数组中(不区分大小写)
func CaseInsensitiveInArr(v string, sl []string) bool {
	v = strings.ToLower(v)
	for _, vv := range sl {
		if strings.ToLower(vv) == v {
			return true
		}
	}
	return false
}

// GetColumnInArr 获取切片中的某个结构值
func GetColumnInArr[T any, V any](arr []T, getVal func(T) V) []V {
	var (
		res []V
	)
	for _, data := range arr {
		res = append(res, getVal(data))
	}
	return res
}

func Distinct[V comparable](arr []V) []V {
	var (
		res []V
		m   = map[V]bool{}
	)
	for _, v := range arr {
		m[v] = true
	}
	for k, _ := range m {
		res = append(res, k)
	}
	return res
}

// DistinctByFunc 获取切片中的某个结构值并去重
func DistinctByFunc[T any, V comparable](arr []T, getVal func(T) V) []V {
	var (
		res []V
		m   = map[V]bool{}
	)
	for _, v := range arr {
		m[getVal(v)] = true
	}
	for k, _ := range m {
		res = append(res, k)
	}
	return res
}

// CountDistinct 获取切片中的某个结构值的去重个数
func CountDistinct[T any, V comparable](arr []T, getVal func(T) V) int {
	var (
		m = map[V]bool{}
	)
	for _, v := range arr {
		m[getVal(v)] = true
	}
	return len(m)
}

// ArrIntersection 切片交集
func ArrIntersection[V comparable](a, b []V) (c []V) {
	m := make(map[V]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

// ArrDifference 切片差集
func ArrDifference[V comparable](a, b []V) (c []V) {
	m := make(map[V]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; !ok {
			c = append(c, item)
		}
	}
	return
}

func ArrByFunc[T any, V comparable](arr []T, getVal func(T) V) []V {
	var (
		res []V
	)
	for _, v := range arr {
		res = append(res, getVal(v))
	}
	return res
}

func IntArr(arr []string) []int {
	var intArr []int
	for _, s := range arr {
		v, _ := strconv.Atoi(s)
		intArr = append(intArr, v)
	}
	return intArr
}

func StringArr[V int | int64 | uint | uint64](arr []V) []string {
	var stringArr []string
	for _, s := range arr {
		stringArr = append(stringArr, strconv.FormatInt(int64(s), 10))
	}
	return stringArr
}

// SplitInt 字符串分割为整型数组
func SplitInt[V Int](s string, sep string) []V {
	arr := strings.Split(s, sep)
	var intArr []V
	for _, s := range arr {
		v, _ := strconv.Atoi(s)
		intArr = append(intArr, V(v))
	}
	return intArr
}

// SplitStr 字符串分割为整型数组
func SplitStr(s string, sep string) []string {
	arr := strings.Split(s, sep)
	var intArr []string
	for _, s := range arr {
		intArr = append(intArr, s)
	}
	return intArr
}

// SliceByRange 截取数组中的一段，长度不够时从头开始
func SliceByRange[T any](arr []T, start, length int) []T {
	if len(arr) == 0 {
		return nil
	}
	if start < 0 {
		start = len(arr) + start
	}
	if start >= len(arr) {
		start %= len(arr)
	}
	if length > len(arr) {
		length = len(arr)
	}
	if start+length > len(arr) {
		return append(arr[start:], arr[:length-(len(arr)-start)]...)
	}
	return arr[start : start+length]
}

func ArrVertical[T any](arr [][]T) [][]T {
	if len(arr) == 0 {
		return arr
	}
	var res [][]T
	for i := 0; i < len(arr[0]); i++ {
		var row []T
		for j := 0; j < len(arr); j++ {
			row = append(row, arr[j][i])
		}
		res = append(res, row)
	}
	return res
}

// Sprint2DArr 返回二维数组的字符串
func Sprint2DArr[T any](arr [][]T, sep string) string {
	s := ""
	for _, vv := range arr {
		s += sep
		for _, v := range vv {
			s += " " + conver.StringMust(v) + " " + sep
		}
		s += "\n"
	}
	return s
}

// Format2DArr 将二维数组格式化
func Format2DArr[T any, V any](arr [][]T, fn func(T) V) [][]V {
	var res [][]V
	for _, v := range arr {
		res = append(res, FormatArr(v, fn))
	}
	return res
}

// FormatArr 将一维数组格式化
func FormatArr[T any, V any](arr []T, f func(T) V) []V {
	var res []V
	for _, v := range arr {
		res = append(res, f(v))
	}
	return res
}

// Product 计算笛卡尔积
func Product[T any](sets [][]T) [][]T {
	lens := func(i int) int { return len(sets[i]) }
	var product [][]T
	for ix := make([]int, len(sets)); ix[0] < lens(0); nextIndex(ix, lens) {
		var r []T
		for j, k := range ix {
			r = append(r, sets[j][k])
		}
		product = append(product, r)
	}
	return product
}

func nextIndex(ix []int, lens func(i int) int) {
	for j := len(ix) - 1; j >= 0; j-- {
		ix[j]++
		if j == 0 || ix[j] < lens(j) {
			return
		}
		ix[j] = 0
	}
}

func ArrMap[T any, V any](arr []T, f func(T) V) []V {
	var res []V
	for _, v := range arr {
		res = append(res, f(v))
	}
	return res
}

// SliceShuffle 切片乱序
func SliceShuffle[T any](slice []T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// Apart 拆分数组
func Apart[T any](arr []T, fn func(T) bool) (a, b []T) {
	for _, v := range arr {
		if fn(v) {
			a = append(a, v)
		} else {
			b = append(b, v)
		}
	}
	return
}

// NewTable 创建二维数组
func NewTable[T any](col, row int, fn func(x, y int) T) [][]T {
	list := make([][]T, row)
	for x := 0; x < row; x++ {
		list[x] = make([]T, col)
		for y := 0; y < col; y++ {
			list[x][y] = fn(x, y)
		}
	}
	return list
}
