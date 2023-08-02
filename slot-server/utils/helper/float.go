package helper

import (
	"github.com/shopspring/decimal"
	"math"
	"modernc.org/mathutil"
	"strconv"
	"strings"
)

// StrSplitToFloat 分割字符串 为浮点数组
func StrSplitToFloat(data string, sep string) (res []float64) {
	arr := strings.Split(data, sep)
	for _, v := range arr {
		vv, _ := strconv.ParseFloat(v, 10)
		res = append(res, vv)
	}
	return res
}

// FloatToStr 浮点转字符串并舍去多余的0
func FloatToStr(v float64) string {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return "0"
	}
	return strconv.FormatFloat(Decimal(v, 2), 'f', -2, 64)
}

// FloatToRate 浮点转百分比
func FloatToRate(v float64) string {
	return FloatToStr(v*100) + "%"
}

// Decimal 截取两位小数
func Decimal(num float64, round int32) float64 {
	if math.IsNaN(num) || math.IsInf(num, 0) {
		num = 0
	}
	v, _ := decimal.NewFromFloat(num).Round(round).Float64()
	return v
}

// FloatArrSub 浮点数数组 相减
func FloatArrSub[T float64 | int](o []T, sub []T) []T {
	min := mathutil.Min(len(o), len(sub))
	if min == 0 {
		return o
	}
	var arr = make([]T, len(o))
	for i := 0; i < min; i++ {
		arr[i] = o[i] - sub[i]
	}
	if min < len(o) {
		copy(arr[min:], o[min:])
	}
	return arr
}
