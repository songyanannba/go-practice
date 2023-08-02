// 自动生成模板Configs
package business

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm"
	"slot-server/enum"
	"slot-server/global"
	"time"
)

// Configs 结构体
type Configs struct {
	global.GVA_MODEL
	Name   string `json:"name" form:"name" gorm:"index;column:name;comment:name;size:30;"`
	Value  string `json:"value" form:"value" gorm:"column:value;type:text;comment:value;"`
	Status uint8  `json:"status" form:"status" gorm:"column:status;default:1;comment:状态;size:8;"`
}

// TableName Configs 表名
func (Configs) TableName() string {
	return "x_configs"
}

func getCacheKey(name string) string {
	return "x_config:" + name
}

func GetXConfigCacheByName(name string) (c *Configs, err error) {
	c = &Configs{}
	cacheKey := getCacheKey(name)
	json := jsoniter.ConfigFastest
	res, err := global.GVA_REDIS.Get(context.Background(), cacheKey).Bytes()
	if err == nil {
		err = json.Unmarshal(res, c)
		return
	}
	var jsonRes []byte
	err = global.GVA_DB.Select("id", "name", "value", "status").First(&c, "name = ?", name).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			jsonRes, err = json.Marshal(map[string]any{"id": 0, "name": name, "value": "", "status": enum.No})
			global.GVA_REDIS.Set(context.Background(), cacheKey, string(jsonRes), 10*time.Hour)
		}
		return
	}
	jsonRes, err = json.Marshal(map[string]any{"id": c.ID, "name": c.Name, "value": c.Value, "status": c.Status})
	if err != nil {
		return
	}
	global.GVA_REDIS.Set(context.Background(), cacheKey, string(jsonRes), 10*time.Hour)
	return
}

func SetXConfigCacheByName(name string, value string) (err error) {
	err = global.GVA_REDIS.Set(context.Background(), getCacheKey(name), value, 10*time.Hour).Err()
	return
}

func DelXConfigCacheByName(name string) (err error) {
	err = global.GVA_REDIS.Del(context.Background(), getCacheKey(name)).Err()
	return
}

func DelXConfigCacheById(id uint) (err error) {
	name := ""
	global.GVA_DB.Model(&Configs{}).Where("id", id).Pluck("name", &name)
	err = global.GVA_REDIS.Del(context.Background(), getCacheKey(name)).Err()
	return
}

func DelAllXConfigCache() (err error) {
	keys := global.GVA_REDIS.Keys(context.Background(), getCacheKey("")).Val()
	err = global.GVA_REDIS.Del(context.Background(), keys...).Err()
	return
}
