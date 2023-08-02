package enum

const (
	ScatterName    = "scatter"
	MultiplierName = "multiplier"
)

const (
	TriggerFreeGet = 4
	TriggerFreeAdd = 3
)

const (
	FreeGetNum = 15
	FreeAddNum = 5
)

const (
	LineLength = 8
)

const (
	LargeScale = 0.1
	Trim       = 0.05
	Ok         = 0.02
)

const (
	Unit8NormalMulWeight = iota //普通转权重
	Unit8RaiseMulWeight         //加注转权重
	Unit8FreeMulWeight          //免费转权重
)
