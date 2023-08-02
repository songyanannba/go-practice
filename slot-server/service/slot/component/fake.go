package component

import "slot-server/model/business"

type Fake struct {
	Type     int
	SlotId   int
	Num      uint
	Position []int
	Which    uint8
}

type Fakes struct {
	common map[uint]*Fake
	fs     map[uint]*Fake
}

func newFakes(fakeList []*business.SlotFake) *Fakes {
	var fakes = &Fakes{
		common: make(map[uint]*Fake),
		fs:     make(map[uint]*Fake),
	}
	for _, fake := range fakeList {
		f := &Fake{
			Type:     fake.Type,
			SlotId:   fake.SlotId,
			Num:      fake.Num,
			Position: fake.ParsePosition(),
			Which:    fake.Which,
		}
		if fake.Type == 1 {
			fakes.common[fake.Num] = f
		} else {
			fakes.fs[fake.Num] = f
		}
	}
	return fakes
}

func (f *Fakes) GetFake(typ int, num uint) *Fake {
	if typ == 1 {
		return f.common[num]
	}
	return f.fs[num]
}
