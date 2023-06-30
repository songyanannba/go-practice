package Rand

import (
	crand "crypto/rand"
	"math/big"
)

func RandInt(v int) int {
	if v <= 0 {
		return 0
	}
	bigInt, _ := crand.Int(crand.Reader, big.NewInt(int64(v)))
	return int(bigInt.Int64())
}
