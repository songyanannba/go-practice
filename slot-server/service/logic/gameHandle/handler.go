package gameHandle

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"slot-server/service/slot/component"
)

type Handler interface {
	SetHandle(*Handle) // 设置handle

	GetSpin() *component.Spin // 获取spin主体

	GetSpins() []*component.Spin // 获取所有spin

	NeedJackpot() bool // 是否需要奖池逻辑

	Run() error // 运行

	GetAck() protoreflect.ProtoMessage // 获取ack

	GetTotalWin() int64 // 获取总赢

	GetPlayNum() (int64, int64) // 获取普通玩次数 和 免费玩次数

	GetSumMul() float64 // 获取总倍数
}
