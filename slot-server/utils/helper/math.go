package helper

import "github.com/shopspring/decimal"

func FloatSum(num ...float64) float64 {
	if num == nil {
		return 0
	}
	var sum decimal.Decimal
	for _, v := range num {
		sum = sum.Add(decimal.NewFromFloat(v))
	}
	f, _ := sum.Float64()
	return f
}

func FloatMul(num ...float64) float64 {
	if len(num) == 0 {
		return 0
	}
	var sum = decimal.NewFromFloat(num[0])
	for i := 1; i < len(num); i++ {
		sum = sum.Mul(decimal.NewFromFloat(num[i]))
	}
	f, _ := sum.Float64()
	return f
}

func FloatsMul[T Number](n T, n2 T, other ...T) float64 {
	sum := decimal.NewFromFloat(float64(n)).
		Mul(decimal.NewFromFloat(float64(n2)))
	for _, v := range other {
		sum = sum.Mul(decimal.NewFromFloat(float64(v)))
	}
	f, _ := sum.Float64()
	return f
}

func IntMulFloatToInt[T Int](n T, n2 ...float64) int64 {
	sum := decimal.NewFromInt(int64(n))
	for _, v := range n2 {
		sum = sum.Mul(decimal.NewFromFloat(v))
	}
	return sum.IntPart()
}

func IntMulFloatToFloat[T Int](n T, n2 ...float64) float64 {
	sum := decimal.NewFromInt(int64(n))
	for _, v := range n2 {
		sum = sum.Mul(decimal.NewFromFloat(v))
	}
	f, _ := sum.Float64()
	return f
}

func Abs[T Signed](num T) T {
	if num < 0 {
		return -num
	}
	return num
}

// NearKey 匹配最近值的key nums为降序
func NearKey[T Signed](nums []T, num T) []int {
	minDiff := nums[0] - num
	keys := []int{0}
	for i := 1; i < len(nums); i++ {
		diff := Abs(nums[i] - num)
		if diff < minDiff {
			minDiff = diff
			keys = []int{i}
		} else if diff == minDiff {
			keys = append(keys, i)
		} else {
			break
		}
	}
	return keys
}

// NearVal 匹配最近值 nums为降序
func NearVal[T Signed](nums []T, num T) T {
	minDiff := nums[0] - num
	for i := 1; i < len(nums); i++ {
		diff := Abs(nums[i] - num)
		if diff > minDiff {
			return nums[i-1]
		} else {
			minDiff = diff
		}
	}
	return nums[len(nums)-1]
}

func Mul100(num float64) int64 {
	return decimal.NewFromFloat(num).Mul(decimal.NewFromInt(100)).IntPart()
}

func Div100(num int64) float64 {
	v, _ := decimal.NewFromInt(num).Div(decimal.NewFromInt(100)).Float64()
	return v
}

func Range(n1, n2 int) []int {
	numbers := make([]int, n2-n1+1)

	for i := range numbers {
		numbers[i] = n1 + i
	}
	return numbers
}
