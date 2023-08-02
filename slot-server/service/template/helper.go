package template

import (
	"errors"
	"fmt"
	"slot-server/service/slot/base"
)

// IntervalPlacement 按照间隔放置
func IntervalPlacement(tags []*base.Tag, fillTag *base.Tag, num, interval int) error {
	if len(tags) < num*interval+num {
		return errors.New(fmt.Sprintf("tags length is %d, but num is %d, interval is %d", len(tags), num, interval))
	}
	nowInterval := interval
	nowNum := 0
	for i, tag := range tags {
		if nowInterval >= interval && (tag == nil || tag.Name == "" || tag.Id == -1) {
			tags[i] = fillTag.Copy()
			nowInterval = 0
			nowNum++
		} else {
			nowInterval++
		}
		if nowNum >= num {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Need to fill %d has filled %d", num, nowNum))

}

func TagsFill(tags []*base.Tag, fillTags []*base.Tag) {
	fillIndex := 0
	for i, tag := range tags {
		if tag == nil || tag.Name == "" || tag.Id == -1 {
			tags[i] = fillTags[fillIndex].Copy()
			fillIndex++
		}
	}
}
