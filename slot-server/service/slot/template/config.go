package template

import "slot-server/service/slot/base"

type Config struct {
	SlotId       int                 // 游戏id
	PayTableList []*base.PayTable    // 赔付表
	Template     map[int][]*base.Tag // 模版 (代表列数据)
	TagMapByName map[string]*base.Tag
	TagMapById   map[int]*base.Tag
	Row          int         // 行数
	Col          int         // 列数
	GameType     int         // 游戏类型
	Event        *base.Event // 事件
}

func (c *Config) GetTagByName(name string) *base.Tag {
	return c.TagMapByName[name].Copy()
}

func (c *Config) GetTagById(id int) *base.Tag {
	return c.TagMapById[id].Copy()
}

func (c *Config) InitConfig(tags []*base.Tag) {
	for i, tag := range tags {
		c.TagMapByName[tag.Name] = tag
		c.TagMapById[i] = tag
	}
}
