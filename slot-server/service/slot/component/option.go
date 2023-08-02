package component

import (
	"slot-server/global"
	"slot-server/service/slot/base"
)

type Options struct {
	// 从父级继承时无需变动的参数
	IsTest           bool  // 是否测试
	Raise            int64 // 额外加注
	JackpotStartTime int64 // 奖池开始时间

	// 从父级继承时需要变动的参数
	IsFree   bool // 是否免费转
	IsReSpin bool // 是否respin

	TriggerMode int // 触发模式 0 普通转 1 freeSpin

	PlayNum     uint  // 转动次数 用于假数据
	BuyFreeCoin int64 // 购买免费次数
	BuyReCoin   int64 // 购买免费次数

	Rank     int // 进度
	NextRank int // 下一次进度

	IsSetResult [][]*base.Tag // 是否设置结果
	NeedSpecify bool          // 是否需要指定结果

	TagsLock   []*base.Tag // 锁定的tag
	IsMustFree bool        // 购买免费Free
	IsMustRes  bool        // 购买Res

	Demo bool // 是否是demo

	FreeNum int // 剩余免费次数
	ResNum  int //  剩余Respin次数

	DebugConfig string

	Spin *Spin `json:"-"`
}

func (s Options) String() string {
	toString, err := global.Json.MarshalToString(s)
	if err != nil {
		return ""
	}
	return toString
}

type Option func(*Options)

func GetOptions(opts ...Option) *Options {
	o := &Options{}
	for _, option := range opts {
		option(o)
	}
	return o
}

//func WithFreeSpin(n int) Option {
//	return func(o *Options) {
//		o.IsFree = true
//		o.FreeNum = n
//	}
//}

func WithTest() Option {
	return func(o *Options) {
		o.IsTest = true
	}
}

func WithJackpotStartTime(t int64) Option {
	return func(o *Options) {
		o.JackpotStartTime = t
	}
}

func WithPlayNum(n uint) Option {
	return func(o *Options) {
		o.PlayNum = n
	}
}

func WithRaise(n int64) Option {
	return func(o *Options) {
		o.Raise = n
	}
}

func WithRank(n int) Option {
	return func(o *Options) {
		o.Rank = n
	}
}

func WithReSpin() Option {
	return func(o *Options) {
		o.IsReSpin = true
	}
}

func WithFreeSpin() Option {
	return func(o *Options) {
		o.IsFree = true
	}
}

func WithSetResult(tags [][]*base.Tag) Option {
	return func(o *Options) {
		o.IsSetResult = tags
	}
}

func WithNeedSpecify(ok bool) Option {
	return func(o *Options) {
		o.NeedSpecify = ok
	}
}

func WithTagsLock(tags []*base.Tag) Option {
	return func(o *Options) {
		o.TagsLock = make([]*base.Tag, 0)
		for _, tag := range tags {
			o.TagsLock = append(o.TagsLock, &base.Tag{
				Id:       tag.Id,
				Name:     tag.Name,
				X:        tag.X,
				Y:        tag.Y,
				Multiple: tag.Multiple,
				ISLock:   tag.ISLock,
			})
		}
	}
}

func WithTriggerMode(mode int) Option {
	return func(o *Options) {
		o.TriggerMode = mode
	}
}

func WithIsMustFree() Option {
	return func(o *Options) {
		o.IsMustFree = true
	}
}

func WithIsMustRes() Option {
	return func(o *Options) {
		o.IsMustRes = true
	}
}

func WithDemo() Option {
	return func(o *Options) {
		o.Demo = true
	}
}

func SpinCop(spin *Spin) *Spin {
	newSpin := *spin
	Option := *spin.Options
	newSpin.Options = &Option
	return &newSpin
}

func SetFreeNum(int2 int) Option {
	return func(o *Options) {
		o.FreeNum = int2
	}
}

func SetResNum(int2 int) Option {
	return func(o *Options) {
		o.ResNum = int2
	}
}

func SetDebugConfig(usId uint) Option {
	return func(o *Options) {
		s := o.Spin
		s.SetDebugInitData(usId)
	}
}
