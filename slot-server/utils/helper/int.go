package helper

func AbsoluteValue(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}
