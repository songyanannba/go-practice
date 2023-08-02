package eliminate

import (
	"slot-server/service/slot/base"
)

type SpecialTag struct {
	Scatter base.Tag // scatter
	//Wild    component.Tag // wild
	//Mul     component.Tag // 翻倍
}

func SetScatterName(name string) func(s *SpecialTag) {
	return func(s *SpecialTag) {
		s.Scatter.Name = name
	}
}
