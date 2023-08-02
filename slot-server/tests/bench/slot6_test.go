package bench

import (
	"slot-server/core"
	"slot-server/service/slot"
	"testing"
)

// 0>
// 需要注释 viper.go的这两行代码
// type在测试过程中 不能有相同的
//svType := flag.String("type", "gate", "the server type")
//global.SvName = *svType

// 1>
//BenchmarkSlot6-8     1363	    930152 ns/op
//函数名后面的-8，表示运行时对应的 GOMAXPROCS 的值；
//接着的 1363 表示运行 for 循环的次数，也就是调用被测试代码的次数，也就是在b.N的范围内执行的次数；
//最后的 930152 ns/op表示每次需要花费 112.9 纳秒；

// 2>
//编写测试代码：b.N是基准测试框架提供的，表示循环的次数，因为需要反复调用测试代码来评估性能。
//b.N 的值会以1, 2, 5, 10, 20, 50, …这样的规律递增下去直到运行时间大于1秒钟，由于程序判断运行时间稳定才会停止运行，
//所以千万不要在loop循环里面使用一个变化的值作为函数的参数。

// go test -v -bench=. -benchtime=3s
//benchtime 修改迷人时间

func BenchmarkSlot6(b *testing.B) {
	core.BaseInit()
	for i := 0; i < b.N; i++ {
		slot.Play(6, 100) // 6
	}

	//1秒结果 BenchmarkSlot6-8            2132            568349 ns/op
	//3秒结果 BenchmarkSlot6-8            6583            524807 ns/op
}
