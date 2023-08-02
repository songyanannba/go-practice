package template

//第八台游戏逻辑

// DropDisplay 当前窗口空白位置上浮
func (s *SpinInfo) DropDisplay() {
	tags := s.GetEmptyTags()
	for _, tag := range tags {
		for x := tag.X; x > 0; x-- {
			//和上面的交换位置
			s.Display[x][tag.Y], s.Display[x-1][tag.Y] = s.Display[x-1][tag.Y], s.Display[x][tag.Y]
		}
	}
	s.SetAllLocation()
}

func (s *SpinInfo) Drop() {
	s.DropDisplay()
	for x := len(s.Display) - 1; x >= 0; x-- {
		row := s.Display[x]
		for y, tag := range row {
			if tag.Name == "" {
				s.Display[x][y] = s.GetTemplateTag(x, y)
			}
		}
	}
}
