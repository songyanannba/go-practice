package component

import "slot-server/service/slot/base"

type Queue struct {
	first *node
	last  *node
	n     int
}

type node struct {
	item *Level
	next *node
}

type Level struct {
	CoreCount int // 核心数量
	EmitCount int // 发射数量
	WildMul   int // 万能数量

}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) IsEmpty() bool {
	return q.n == 0
}

func (q *Queue) Size() int {
	return q.n
}

func (q *Queue) EnQueue(event *base.Unit6LevelEvent) {
	oldLast := q.last
	q.last = &node{}
	q.last.item = GetLevel(event)
	q.last.next = nil
	if q.IsEmpty() {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
	q.n++
}

func (q *Queue) DeQueue() *Level {
	if q.IsEmpty() {
		return nil
	}
	item := q.first.item
	q.first = q.first.next
	if q.IsEmpty() {
		q.last = nil
	}
	q.n--
	return item
}

func GetLevel(event *base.Unit6LevelEvent) *Level {
	level := &Level{}
	level.CoreCount = event.CoreCount
	level.EmitCount = event.EmitEvent.Fetch()
	level.WildMul = event.WildEvent.Fetch()
	return level
}
