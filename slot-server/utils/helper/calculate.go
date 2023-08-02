package helper

import "github.com/shopspring/decimal"

func Float64Add(fls ...float64) float64 {
	sum := decimal.NewFromFloat(0)
	for _, fl := range fls {
		sum = sum.Add(decimal.NewFromFloat(fl))
	}
	fl, ok := sum.Float64()
	if ok {
		return fl
	}
	return 0
}

func Float64Mul(fls ...float64) float64 {
	sum := decimal.NewFromFloat(1)
	for _, fl := range fls {
		sum = sum.Mul(decimal.NewFromFloat(fl))
	}
	fl, ok := sum.Float64()
	if ok {
		return fl
	}
	return 0
}
