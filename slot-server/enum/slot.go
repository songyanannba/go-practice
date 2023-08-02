package enum

const (
	SlotMaxSpinNum     = 10000 // slot最大转动
	SlotMaxSpinStr     = "超过最大次数10000"
	SlotMaxFreeSpinErr = "免费玩次数超过1000次"
)

// Slot 第六台等级
const (
	Rand1 = iota + 1
	Rand2
	Rand3
	Rand4
	Rand5
)

const (
	EmptyTagName = "" //空标签
	SlotWild     = "wild"
	SlotWild1    = "wild_1"
)

const (
	CoreTagMinNum = 4 //中心点周围 最少填充的tag个数
	CoreTagMaxNum = 8 //中心点周围 最多填充的tag个数

	GetLine = 5 //默认匹配连续标签的最少个数
	NumLine = 8 //默认匹配随机标签的最少个数

	QuantityTagNum = 8 //第八台逻辑 最少8个标签才能连消

	IsMustFreeScatterNum = 4 // 购买免费转填充4个 Scatter
	ScatterNum4Datum     = 15

	MinFreeScatterNum = 3 // 免费转 3个 Scatter
	ScatterNum3Datum  = 5

	Slot2FreeNum = 8 //第二台获取免费转的次数

	SameTagLen = 3 //在第二台 和 第三台用

	InitFillTagNum = 5 //初始化 填充tag 的数量
)

// 机台ID
const (
	SlotId1 = iota + 1
	SlotId2
	SlotId3
	SlotId4
	SlotId5 //机台5
	SlotId6 //机台6
	SlotId7
	SlotId8
)

//debugType

const (
	SlotDebugType1 = iota + 1
)

const (
	NormalSpin = iota + 1 //普通转
	FreeSpin              //免费转
	ReSpin                //重转
	RaiseSpin             //加注转
)
