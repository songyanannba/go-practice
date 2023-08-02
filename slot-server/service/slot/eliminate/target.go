package eliminate

import (
	"github.com/shopspring/decimal"
	"slot-server/service/slot/base"
)

type Target struct {
	MinMul      float64 // 最小倍数
	MaxMul      float64 // 最大倍数
	InitNum     int     // 初始个数
	ScatterNum  int     // scatter次数
	MulNumEvent *base.ChangeTableEvent
}

func (t *Target) Compare(value decimal.Decimal) bool {
	v, _ := value.Float64()
	//return v < t.MinMul && v < t.MaxMul
	return v < t.MinMul
}
