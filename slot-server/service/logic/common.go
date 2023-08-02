package logic

import (
	"fmt"
	. "slot-server/pbs"
	"slot-server/service/slot/base"
	"slot-server/service/slot/component"
	"slot-server/utils/helper"
)

// NewAckHead 新建回应头部
func NewAckHead(head *ReqHead) *AckHead {
	if head == nil {
		return &AckHead{
			Code: Code_Ok,
		}
	}
	return &AckHead{
		Uid:  head.Uid,
		Code: Code_Ok,
	}
}

type AckHeads interface {
	GetHead() *AckHead
}

type ReqHeads interface {
	GetHead() *ReqHead
}

func DumpSpinAck(ack *SpinAck) string {
	conf, _ := component.GetSlotConfig(uint(ack.Opt.GameId), false)
	return fmt.Sprintf(
		"Head: %+v\n"+
			"Opt: %+v\n"+
			"TotalWin: %+v\n"+
			"TxnId: %d\n"+
			"StepList: \n%+v\n",
		ack.Head,
		ack.Opt,
		ack.TotalWin,
		ack.TxnId,
		DumpSpinSteps(ack.StepList, conf),
	)
}

func DumpSpinSteps(steps []*SpinStep, conf *component.Config) string {
	m := conf.GetTagIdMap()
	s := "- - - - -\n"
	for _, step := range steps {
		s += fmt.Sprintf(
			"Id: %+v, "+
				"Pid: %+v, "+
				"Type: %+v, "+
				"Win: %+v, "+
				"JackpotId: %+v\n"+
				"CardList: \n%s"+
				"LineList: \n%s",
			step.Id,
			step.Pid,
			step.Type,
			step.Win,
			step.JackpotId,
			DumpCardList(m, step.CardList, true),
			DumpCardList(m, step.LineList, false),
		)
	}
	s += "- - - - -\n"
	return s
}

func DumpCardList(m map[int]base.Tag, cardList []*Cards, flip bool) string {
	var (
		list [][]*Card
		wins []int64
	)
	for _, cards := range cardList {
		list = append(list, cards.Cards)
		wins = append(wins, cards.Amount)
	}
	if flip {
		list = helper.ArrVertical(list)
	}

	s := "- - -\n"
	for i, cards := range list {
		if !flip {
			s += "win: " + fmt.Sprintf("%d", wins[i]) + "\n"
		}
		s += "["
		for _, card := range cards {
			s += m[int(card.CardId)].Name + " " + fmt.Sprintf("{%+v},", card)
		}
		if len(cards) > 0 {
			s = s[:len(s)-1]
		}
		s += "]\n"
	}
	if len(list) == 0 {
		s += "nil\n"
	}
	s += "- - -\n"
	return s
}
