package template

import (
	"fmt"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/service/slot/base"
	"slot-server/utils/helper"
)

//第八台游戏逻辑数个数

// GetTemplateTag 从模版获取Tag
func (s *SpinInfo) GetTemplateTag(x, y int) *base.Tag {
	defer func() {
		if err := recover(); err != nil {
			global.GVA_LOG.Error(fmt.Sprintf("x:%d,y:%d,TemX:%d", x, y, s.templateRowMap[y]))
		}
	}()

	// 获取模版（列）对应的x
	temX := s.GetTemX(y)

	// 获取模版对应的tag
	fillTag := s.config.Template[y][temX]

	// 如果模版对应的tag为空,则报错
	if fillTag == nil || fillTag.Name == "" {
		global.GVA_LOG.Error(fmt.Sprintf("x:%d,y:%d,TemX:%d", x, y, temX))
	}
	//第八台游戏逻辑特殊配置
	if fillTag.Name == enum.MultiplierName && s.config.SlotId == enum.SlotId8 {
		fillTag.Multiple = float64(s.config.Event.M[s.config.GameType-1].(*base.ChangeTableEvent).Fetch())
	}

	// 设置tag的坐标
	fillTag.X = x
	fillTag.Y = y

	return fillTag.Copy()
}

// GetTemX 获取模版Y对应的X值并修改
// y 是列的意思; 获取每一列的一个值（相当于x轴坐标） 然后x轴向上移动一步
func (s *SpinInfo) GetTemX(y int) int {
	temX := s.templateRowMap[y]
	if temX < 0 {
		temX = len(s.config.Template[y]) + temX
	}
	s.templateRowMap[y] = temX
	s.templateRowMap[y]--
	return temX
}

// GetAllTags 获取当前显示窗口所有Tag
func (s *SpinInfo) GetAllTags() []*base.Tag {
	return helper.ListToArr(s.Display)
}

// FindTagsByName 查找指定名称的tag
func (s *SpinInfo) FindTagsByName(names ...string) []*base.Tag {
	mapByName := s.GetTagMapByName()
	tags := make([]*base.Tag, 0)
	for _, name := range names {
		tags = append(tags, mapByName[name]...)
	}
	return tags
}

// FindCountLine 查找当前窗口满足指定数量的tag
func (s *SpinInfo) FindCountLine(number int) [][]*base.Tag {
	tagsMap := s.GetTagMapByName()
	tagList := make([][]*base.Tag, 0)
	for tagName, tags := range tagsMap {
		if len(tags) >= number && tagName != "" && tagName != "scatter" && tagName != "multiplier" {
			tagList = append(tagList, helper.CopyList(tags))
		}
	}
	return tagList
}
