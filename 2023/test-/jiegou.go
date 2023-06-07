package main

type zuobiao struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Tag struct {
	Id       int `json:"-"`
	Name     string
	Include  []string `json:"-"`
	Multiple int      `json:"-"`

	X int `json:"-"`
	Y int `json:"-"`

	IsLine     bool `json:"-"`
	IsPayTable bool `json:"-"`
	IsWild     bool `json:"-"`
	IsSingle   bool `json:"-"`
	IsJackpot  bool `json:"-"`
	ISLock     bool `json:"-"` // 是否锁定
}

type ViewData struct {
	Data         string    `json:"data"`
	Bet          int       `json:"bet"`          // 压注
	Gain         int       `json:"gain"`         // 赢钱
	PayTableMuls []float64 `json:"payTableMuls"` // payTable倍数
	WildMuls     []int     `json:"wildMuls"`     // 百搭倍数
	JackpotMul   int       `json:"jackpotMul"`   // 奖池倍数
	FreeSpin     int       `json:"freeSpin"`     // 免费转次数
}

type Spin struct {
	InitDataList [][]*Tag `json:"-"` // 初始数据
	ResDataList  [][]*Tag // 结果数据
	WildList     []*Tag   // 百搭数据
	SingleList   []*Tag   // 单出数据
}


func GetSpin() *Spin {
	return &Spin{
		InitDataList: nil,
		ResDataList:  nil,
		WildList:     nil,
		SingleList:   nil,
	}
}