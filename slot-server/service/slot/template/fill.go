package template

func (s *SpinInfo) FillInitDisplay() {
	for x, row := range s.Display {
		for y, _ := range row {
			s.Display[x][y] = s.GetTemplateTag(x, y)
		}
	}
}
