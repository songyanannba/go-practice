package component

import (
	"slot-server/service/slot/base"
)

type JackpotData struct {
	Id uint
	//Start   int
	//Inc     float64
	End     float64
	Combine []string
}

func NewJackpotData(id uint, start int, inc float64, end float64, combine []string) *JackpotData {
	return &JackpotData{
		Id: id,
		//Start:   start,
		//Inc:     inc,
		End:     end,
		Combine: combine,
	}
}

// Match 判断是否匹配
func (j JackpotData) Match(tags []*base.Tag) (*Line, bool) {
	jLen := len(j.Combine)
	tLen := len(tags)
	line := Line{}
	if jLen == 0 || tLen == 0 || jLen > tLen {
		return nil, false
	}
	for k, tag := range j.Combine {
		// jackpot 不判断百搭逻辑 直接比较即可
		if tag != tags[k].Name {
			return nil, false
		}
		line.Tags = append(line.Tags, tags[k])
	}
	return &line, true
}

// Contain 判断是否包含
func (j JackpotData) Contain(tags []*base.Tag) bool {
	jLen := len(j.Combine)
	tLen := len(tags)
	if jLen == 0 || tLen == 0 || jLen > tLen {
		return false
	}
	jStep := 0
	for k, tag := range tags {
		// jackpot 不判断百搭逻辑 直接比较即可
		if tag.Name == j.Combine[jStep] {
			jStep++
		} else {
			// 长度不足时直接返回
			if jLen > tLen-k-1 {
				return false
			}
			jStep = 0
		}
		if jStep >= jLen {
			return true
		}
	}
	return false
}

func (s *Spin) GetJackpotId() uint {
	if s.Jackpot != nil {
		return s.Jackpot.Id
	}
	return 0
}
