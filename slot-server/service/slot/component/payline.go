package component

import (
	"slot-server/service/slot/base"
	"slot-server/utils/helper"
	"strings"
)

// ParseCoordinate 解析payline坐标
func ParseCoordinate(position string) (cs []Coordinate) {
	posArr := strings.Split(position, "&")
	for _, pos := range posArr {
		x, y, _ := strings.Cut(pos, "_")
		cs = append(cs, Coordinate{
			X: helper.Atoi(x) - 1,
			Y: helper.Atoi(y) - 1,
		})
	}
	return
}

// 通过坐标在初始数据中获取tag
func (s *Spin) getCoordinateTagByInitData(x, y int) *base.Tag {
	if len(s.InitDataList) <= x || len(s.InitDataList[x]) <= y {
		return &base.Tag{}
	}
	tag := s.InitDataList[x][y]
	return tag
}

// GetPaylineByCoords 根据坐标获取payline结果
func (s *Spin) GetPaylineByCoords() {
	// 遍历所有坐标
	for _, coords := range s.Config.Coords {
		var tags []*base.Tag
		for _, coord := range coords {
			// 从初始数据中获取结果Tag
			tag := s.getCoordinateTagByInitData(coord.X, coord.Y)
			tag.IsLine = true
			tags = append(tags, tag)
		}
		s.ResDataList = append(s.ResDataList, tags)
	}
	// 获取划线中的百搭与单出
	for _, tags := range s.InitDataList {
		for _, tag := range tags {
			if tag.IsLine {
				// 判断是否为百搭
				if tag.IsWild {
					s.WildList = append(s.WildList, tag)
				}
				if tag.IsSingle {
					s.SingleList = append(s.SingleList, tag)
				}
			}
		}
	}
	return
}

// GetAllPayline 获取所有的payline结果
func (s *Spin) GetAllPayline() {
	for _, tags := range s.InitDataList {
		for _, tag := range tags {
			tag.IsLine = true
			// 判断是否为百搭
			if tag.IsWild {
				s.WildList = append(s.WildList, tag)
			}
			if tag.IsSingle {
				s.SingleList = append(s.SingleList, tag)
			}
		}
	}
	s.ResDataList = helper.Product(s.InitDataList)
	return
}

// GetSamePayline 获取相同图标的payline结果 l 长度
func (s *Spin) GetSamePayline(l int) {
	s.ResDataList = base.MatchSameTagList(s.InitDataList, l)
	return
}
