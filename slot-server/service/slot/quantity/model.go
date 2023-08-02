package quantity

type ModelResult struct {
	InitList     string
	Multiple     float64
	RmCount      int // 消除次数
	RmTag        int // 消除标签个数
	WildCount    int // 万能符号个数
	ScatterCount int // 消散符号个数
}
