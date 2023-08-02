package base

type Verify struct {
	Site   Tag
	Verify map[[2]int]bool
	Count  int
}

type PlayFlow struct {
	Type  GameFlow `json:"type" from:"type"`
	Table string   `json:"table" from:"table"`
}
type GameFlow int

const (
	GameEliminate  GameFlow = iota // 消除
	GameFall                       // 掉落
	GameCustomFill                 // 划线填充
	GameAutoFill                   // 填充
	GameOver                       // 完成
	GameInit
)
