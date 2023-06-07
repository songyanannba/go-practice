package main


type Tag struct {
	Id       int
	Name     string
	X int
	Y int
}
var TagList [][]*Tag

func NewTable[T any](col, row int, fu func(x, y int) T) [][]T {
	list := make([][]T, row)
	for i := 0; i < row; i++ {
		list[i] = make([]T, col)
		for j := 0; j < col; j++ {
			list[i][j] = fu(i, j)
		}
	}
	return list
}

func TestTag () [][]*Tag {
	TagList = NewTable[*Tag](3,3 , func(x,y int) *Tag{
		return &Tag{
			X:x,
			Y:y,
			Name:"",
		}
	})
	return TagList
}


func test1[Key comparable ,Val any](m map[Key]Val) []Val {
	s := make([]Val, 0)
	for _,v := range m {
		s = append(s, v)
	}
	return s
}