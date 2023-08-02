package eliminate

type Count struct {
	count      int
	superiorId int
	spinNum    int
}

func NewCount() *Count {
	return &Count{
		count:      0,
		superiorId: 0,
		spinNum:    0,
	}
}

func (c *Count) GetId() int {
	return c.count
}

func (c *Count) InitSpin() {
	c.spinNum = 0
	c.superiorId = c.count
}

func (c *Count) Add() {
	c.count++
	c.spinNum++
}

func (c *Count) GetSupId() int {
	return c.superiorId
}

func (c *Count) Compare(max int) bool {
	return c.count >= max
}

func (c *Count) GetSpinNum() int {
	return c.spinNum
}
