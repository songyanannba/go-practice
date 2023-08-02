package template

import (
	"slot-server/enum"
)

func (t *GenTemplate) Range(r *LogicResult) int {
	f := r.GainRatio
	if f > t.MaxRatio+enum.LargeScale || f < t.MinRatio-enum.LargeScale {
		return MaxRange
	}
	if r.ScatterRatio < t.MinScatter {
		return UpScatter
	}
	if r.ScatterRatio > t.MaxScatter {
		return DownScatter
	}
	if f > t.MaxRatio {
		return DownTrim
	}
	if f < t.MinRatio {
		return UpTrim
	}
	if f <= t.MaxRatio && f >= t.MinRatio && r.ScatterRatio >= t.MinScatter && r.ScatterRatio <= t.MaxScatter {
		return Ok
	}
	return MaxRange
}
