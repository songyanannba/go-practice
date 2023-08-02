package component

type FreeSpinParams struct {
	WildNum      int // 获得的百搭数量
	WildInterval int // 百搭区间

	Count   int // 此次转 获得的免费转标签数量统计
	FreeNum int // 此次转 获得实际免费转数量
	ReNum   int // 此次转 获得实际Respin数量

}

func (s *Spin) KeepSpin() bool {
	return s.FreeSpinParams.FreeNum > 0 || s.FreeSpinParams.ReNum > 0
}
