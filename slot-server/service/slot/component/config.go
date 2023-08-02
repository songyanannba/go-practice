package component

import (
	"fmt"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/service/slot/base"
	"slot-server/utils"
	"slot-server/utils/helper"
	"strings"
)

var cache = local_cache.NewCache()

// Config 请勿在初始化后修改config中的数据 该数据为机台的全局数据
type Config struct {
	SlotId uint // 机器id

	Status uint8       // 状态
	Index  int         // 列数
	Row    int         // 行数
	BetMap base.BetMap // 押注区间

	Raise  float64 // 加注倍率
	BuyFee float64 // 购买费用
	BuyRes float64 // 购买资源

	Reel         []*Reel              // 滚轮数据
	Coords       [][]Coordinate       // 划线坐标
	tagMap       map[string]*base.Tag // 所有标签的Map
	tagIdMap     map[int]*base.Tag    // 所有标签的Map
	PayTableList []*base.PayTable     // 赢钱组合
	JackpotList  []*JackpotData       // 奖池
	Event        *base.Event          // 特殊事件
	Fakes        *Fakes               // 假数据

	place     []int                   //后台配置指定的排布索引
	freePlace []int                   //后台配置指定的排布索引
	Debugs    []*business.DebugConfig // 调试配置

	TopMul int // 最高倍数

	Template map[int]map[int][]*base.Tag // 模板数据 类型=>索引=>标签
}

// GetSlotConfig 获取slot配置 从缓存中获取，如果没有则从数据库中获取
func GetSlotConfig(slotId uint, demo bool) (*Config, error) {
	isDemo := helper.If(demo, enum.Yes, enum.No)
	c, exist := GetConfigByCache(slotId, isDemo)
	if exist {
		return c, nil
	}

	rawData, err := NewDbRawDataBySlotId(slotId)
	if err != nil {
		return nil, err
	}

	conf, confDemo := rawData.NewSlotConfig()
	SetConfigCache(slotId, conf, enum.No)
	SetConfigCache(slotId, confDemo, enum.Yes)
	if demo {
		return confDemo, nil
	}

	return conf, nil
}

func GetConfigCacheKey(slotId uint, isDemo int) string {
	return fmt.Sprintf("slot_config_%d_%d", slotId, isDemo)
}

func GetConfigByCache(slotId uint, isDemo int) (*Config, bool) {
	res, exist := cache.Get(GetConfigCacheKey(slotId, isDemo))
	if exist {
		return res.(*Config), true
	}
	return nil, false
}

func SetConfigCache(slotId uint, c *Config, isDemo int) {
	cache.Set(GetConfigCacheKey(slotId, isDemo), c, local_cache.NoExpire)
}

func DeleteConfigCache(slotId uint) {
	cache.Delete(GetConfigCacheKey(slotId, enum.No))
	cache.Delete(GetConfigCacheKey(slotId, enum.Yes))
}

func FlushConfigCache() {
	cache.Flush()
}

func (c *Config) GetTag(tagName string) *base.Tag {
	tag, ok := c.tagMap[tagName]
	if ok {
		return tag.Copy()
	}
	return base.NewTag(-1, "", 1, 0)
}

func (c *Config) GetTagById(id int) *base.Tag {
	tag, ok := c.tagIdMap[id]
	if ok {
		return tag.Copy()
	}
	return base.NewTag(-1, "", 1, 0)
}

func (c *Config) GetTags(tagNames ...string) []*base.Tag {
	var tags []*base.Tag
	for _, tagName := range tagNames {
		tag, ok := c.tagMap[tagName]
		if !ok {
			tag = base.NewTag(-1, tagName, 1, 0)
		}
		tags = append(tags, tag.Copy())
	}
	return tags
}

func (c *Config) GetTagsAndExclude(tagNames ...string) []*base.Tag {
	var tags []*base.Tag
	for _, tag := range c.tagMap {
		if helper.InArr(tag.Name, tagNames) {
			continue
		}
		tags = append(tags, tag.Copy())
	}
	return tags
}

func (c *Config) GetAllTag() []base.Tag {
	var tags []base.Tag
	for _, tag := range c.tagMap {
		tags = append(tags, *tag)
	}
	return tags
}

func (c *Config) GetAllTagQuote() []*base.Tag {
	var tags []*base.Tag
	for _, tag := range c.tagMap {
		tags = append(tags, tag.Copy())
	}
	return tags
}

func (c *Config) GetTagIdMap() map[int]base.Tag {
	m := map[int]base.Tag{}
	for _, tag := range c.tagMap {
		m[tag.Id] = *tag
	}
	return m
}

func (c *Config) GetTagsLayout(str string) []*base.Tag {
	var tags []*base.Tag
	strs := strings.Split(utils.FormatCommandStr(str), ",")
	for _, s := range strs {
		if s == "" {
			continue
		}
		tag := c.GetTag(s)
		if tag.Name == "" || tag.Id == -1 {
			global.GVA_LOG.Error("GetTagsLayout err name:" + s)
		}
		tags = append(tags, tag.Copy())
	}
	return tags
}
