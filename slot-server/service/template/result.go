package template

import (
	"github.com/samber/lo"
	"slot-server/enum"
)

type RunResult struct {
	Gain           int
	ScatterTrigger int
	SpinCount      int
}

type LogicResult struct {
	GainRatio    float64
	ScatterRatio float64
}

func GetLogicResult(res []*RunResult, gameType int) *LogicResult {
	sumGain := lo.SumBy(res, func(item *RunResult) int {
		return item.Gain
	})

	sumCount := lo.SumBy(res, func(item *RunResult) int {
		return item.SpinCount
	})

	sumScatter := lo.SumBy(res, func(item *RunResult) int {
		return item.ScatterTrigger
	})

	gainRatio := 0.0
	switch gameType {
	case enum.NormalSpin:
		gainRatio = float64(sumGain) / float64(sumCount) / float64(100)
	case enum.FreeSpin:
		gainRatio = float64(sumGain) / float64(len(res)) / float64(100)
	default:
		gainRatio = float64(sumGain) / float64(sumCount) / float64(100)
	}
	return &LogicResult{
		ScatterRatio: float64(sumScatter) / float64(sumCount),
		GainRatio:    gainRatio,
	}
}
