package component

func (s *Spin) SumNextRank() {
	s.NextRank = s.Rank
	// 赢
	if s.Gain > 0 {
		// respin
		if !s.IsFree {
			// 次数+1
			s.FreeSpinParams.ReNum = 1
		}
		// 已达到最大
		if s.Rank >= 6 {
			return
		}
		// 增加
		s.NextRank++
		return
	}

	// 输
	if s.IsFree {
		// 是免费转则保持rank
		return
	}
	if s.Raise > 0 {
		// 有加注则重置为第三档
		s.NextRank = 2
		return
	}
	// 重置为0
	s.NextRank = 0
	return
}
