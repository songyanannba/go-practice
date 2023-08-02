package component

import (
	"fmt"
	"slot-server/enum"
	"slot-server/service/slot/base"
	"slot-server/utils/helper"
	"strconv"
	"strings"
)

func (s Spin) GetPayTableIdStr() string {
	var arr []string
	for _, v := range s.PayTables {
		arr = append(arr, strconv.Itoa(int(v.Id)))
	}
	return strings.Join(arr, ",")
}

func (s Spin) PayTablesString(typ string, slotId uint, tagId int) string {
	var arr []string
	if len(s.PayTables) > 0 {
		for _, payTable := range s.PayTables {
			for k, tag := range payTable.Tags {
				id := tag.Id
				if slotId == 2 {
					if k != 0 && tag.IsWild == true {
						//第二台 需要覆盖
						id = tagId
					}
				}
				arr = append(arr, tag.GetRecordStr(id, typ))
			}
		}
	}
	return strings.Join(arr, ";") + "#"
}

func (s Spin) WildAndSingleString(t string) string {
	var (
		str     string
		lenList = 0
		list    []*base.Tag
	)
	if t == enum.WildStrType {
		lenList = len(s.WildList)
		list = s.WildList
	} else if t == enum.SingStrType {
		lenList = len(s.SingleList)
		list = s.SingleList
	} else {
		return ""
	}

	if len(list) <= 0 {
		return ""
	}

	for k, v := range list {
		str += v.GetRecordStr(0, t) + helper.If(lenList == k+1, "#", ";")
	}
	return str
}

func (s Spin) JackpotTagString(t string) string {
	var arr []string
	if s.Jackpot != nil {
		for kj, vj := range s.Jackpot.Combine {
			endStrj := ";"
			if len(s.Jackpot.Combine) == kj+1 {
				endStrj = "#"
			}
			arr = append(arr, t+","+fmt.Sprintf("%d,%d,1,0,0", s.Config.GetTag(vj).Id, kj)+endStrj)
		}
	}
	return strings.Join(arr, ";") + "#"
}
