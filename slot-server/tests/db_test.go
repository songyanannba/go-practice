package tests

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"slot-server/core"
	"slot-server/model/business"
	"slot-server/pbs"
	"slot-server/utils/helper"
	"strconv"
	"testing"
)

func TestDb(t *testing.T) {
	core.BaseInit()
	arr, _ := business.GetList[*business.SlotRecord](10963)
	a := pbs.MatchSpinAck{}
	err := proto.Unmarshal(arr[0].Ack, &a)
	for _, step := range a.Steps {
		PrintTable(step.InitList)
		for _, flow := range step.Flows {
			for _, tag := range flow.AddList.Tags {
				step.InitList[tag.X].Tags[tag.Y].TagId = tag.TagId
			}
			PrintTable(step.InitList)

			for _, tags := range flow.RemoveList {
				for _, tag := range tags.Tags {
					step.InitList[tag.X].Tags[tag.Y].TagId = -1
				}
			}
			PrintTable(step.InitList)

			for i := len(step.InitList) - 1; i > 0; i-- {
				for i2, tag := range step.InitList[i].Tags {
					if tag.TagId == -1 {
						for i3 := i; i3 >= 0; i3-- {
							if step.InitList[i3].Tags[i2].TagId != -1 {
								step.InitList[i].Tags[i2], step.InitList[i3].Tags[i2] = step.InitList[i3].Tags[i2], step.InitList[i].Tags[i2]
							}
						}

					}
				}
			}

			for i, tags := range step.InitList {
				for i2, tag := range tags.Tags {
					tag.X = int32(i)
					tag.Y = int32(i2)
				}
			}
			PrintTable(step.InitList)
		}
	}

	fmt.Println(err)
}

func PrintTable(tagList []*pbs.Tags) {
	str := ""
	for _, row := range tagList {
		for _, col := range row.Tags {
			str += fmt.Sprintf("%s\t", strconv.Itoa(int(col.X))+":"+strconv.Itoa(int(col.Y))+" "+helper.If(col.TagId == -1, "🀆", strconv.Itoa(int(col.TagId))))
		}
		str += "\r\n"
	}
	fmt.Println(str)
	//return str + "\r\n"
}

func Drop(tagList []*pbs.Tags) {

}

func TestRecord(t *testing.T) {
	core.BaseInit()
	arr, _ := business.GetList[*business.SlotRecord](3468, 3483)

	var acks []*pbs.SpinAck
	for _, record := range arr {
		ack := pbs.SpinAck{}
		err := proto.Unmarshal(record.Ack, &ack)
		acks = append(acks, &ack)
		t.Log(err)
		//totalWin := int64(0)
		//s := ""
		//for _, step := range a.StepList {
		//	totalWin += step.Win
		//	s += strconv.Itoa(int(step.Win)) + " + "
		//}
		//s += " = " + strconv.Itoa(int(totalWin))
		//t.Log(s, a, err, totalWin, a.TotalWin)
	}
	//t.Log(acks)
}
