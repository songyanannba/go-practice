package tests

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"slot-server/core"
	"slot-server/service/slot/base"
	"slot-server/service/slot/component"
	goUnit8 "slot-server/service/slot/machine/unit8"
	"slot-server/service/test"
	"slot-server/service/test/test/unit8"
	"slot-server/utils/helper"
	"strconv"
	"testing"
	"time"
)

// go test -v -run TestTemp temp_test.go
func TestTemp(t *testing.T) {

	s := "USD:1,2,3,4,5,6,7,8,9.1,10,11,12,13,14, 15ADS:1212,1313,22.0"

	spew.Dump(base.NewBetMap(s))
}

func BenchmarkTemp(t *testing.B) {
	for i := 0; i < 10; i++ {
		t.Run("1", BenchmarkSub)
	}
}

func BenchmarkSub(t *testing.B) {
	time.Sleep(1 * time.Second)
	t.Log(1)
}

func TestModelRun(t *testing.T) {
	core.BaseInit()
	count := 0
	for {
		count++
		s := GetTable()
		f := GoRun(s)
		fmt.Printf("第%d次, 返回率%f\n", count, f)
		table := s.Config.Template[1]
		str := ""
		for _, tags := range table {
			for _, tag := range tags {
				name := strconv.Itoa(tag.Id)
				if tag.Multiple > 0 {
					name += fmt.Sprintf("*%g", tag.Multiple)
				}
				str += name + "\t"
			}
			str += ";\n"
		}
		fmt.Println(str)
		if f > 0.96 && f < 0.98 {
			fmt.Println("结果合格")
		}
		if count > 0 {
			break
		}
	}
}

func GetTable() *component.Spin {
	opts := []component.Option{component.WithTest(), component.WithNeedSpecify(true)}
	TagsWeight := "scatter&multiplier&high_1&high_2&high_3&high_4&low_1&low_2&low_3&low_4&low_5@0&0&5&20&35&50&70&90&110&130&150&170"
	SpecialConfig := "3&4&5@1&2&2&2"
	s, _ := component.NewSpin(8, 100, opts...)
	strEvn := base.ParseWeightDataStr(TagsWeight)
	speEnv := base.ParseWeightData(SpecialConfig)

	strMap := map[int]map[string]int{}

	for i := 0; i < 6; i++ {
		strMap[i] = map[string]int{}
	}

	table := helper.NewTable(6, 170, func(x, y int) *base.Tag {
		for {
			fillName := strEvn.Fetch()
			if strMap[y][fillName] < strEvn.GetSection(fillName) {
				fillTag := s.Config.GetTag(fillName)
				fillTag.X = x
				fillTag.Y = y
				if fillName == "multiplier" {
					mul := speEnv.Fetch()
					fillTag.Multiple = float64(mul)
				}
				strMap[y][fillName]++
				return fillTag.Copy()
			} else {
				continue
			}
		}
	})

	s.Config.Template[1] = table
	return s
}

func GoRun(spin *component.Spin) float64 {
	var repeat test.Repeat
	repeat = unit8.NewUnit()
	ch, errCh, _ := helper.Parallel[[]*component.Spin](10000000, 1000, func() (spins []*component.Spin, err error) {
		machine := goUnit8.NewMachine(spin)
		err = machine.Exec()
		if err != nil {
			return nil, err
		}
		spins = []*component.Spin{machine.GetSpin()}
		spins = append(spins, machine.GetSpins()...)
		return
	})
	count := 0
	a := 0
	for {
		a++
		select {
		case err := <-errCh:
			fmt.Println(err)
			return 0
		case v, beforeClosed := <-ch:
			count++
			if !beforeClosed {
				goto end
			}
			repeat.Calculate(v)
		}
	}
end:
	return repeat.GetReturnRatio()
}
