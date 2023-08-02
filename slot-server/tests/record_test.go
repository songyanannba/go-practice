package tests

import (
	"fmt"
	"slot-server/core"
	"slot-server/pbs"
	"slot-server/service/logic/gameHandle"
	"testing"
)

func TestRecordMenu(t *testing.T) {
	core.BaseInit()
	ack := &pbs.RecordMenuAck{}
	req := &pbs.RecordMenuReq{
		Date:   "2023-06",
		GameId: 2,
	}

	err := gameHandle.RecordMenu(1, req, ack)
	if err != nil {
		return
	}
	fmt.Print(ack)
}
