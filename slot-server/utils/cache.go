package utils

import (
	"bytes"
	"context"
	"encoding/gob"
	"github.com/bytedance/sonic"
	"slot-server/global"
	"strings"
	"time"
)

// 数据位置定义
const (
	PlaceHash   = "hash"
	PlaceString = "string"
	PlaceSet    = "set"
	PlaceList   = "list"
	PlaceLock   = "lock"
)

func GetCacheKey(place string, keys ...string) string {
	return place + ":" + strings.Join(keys, ":")
}

func GetCache(key string, to interface{}) error {
	res, err := global.GVA_REDIS.Get(context.Background(), key).Bytes()
	if err != nil {
		return err
	}
	return JsonDecode(res, to)
}

func DelCache(key ...string) error {
	err := global.GVA_REDIS.Del(context.Background(), key...).Err()
	return err
}

func SetCache(key string, data interface{}, expiration time.Duration) error {
	res, err := JsonEncode(data)
	if err != nil {
		return err
	}
	err = global.GVA_REDIS.Set(context.Background(), key, res, expiration).Err()
	return err
}

// GobEncode 用gob进行数据编码
func GobEncode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// GobDecode 用gob进行数据解码
func GobDecode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}

// JsonEncode 用sonic进行数据编码
func JsonEncode(data interface{}) ([]byte, error) {
	return sonic.ConfigFastest.Marshal(&data)
}

// JsonDecode 用sonic进行数据解码
func JsonDecode(data []byte, to interface{}) error {
	return sonic.ConfigFastest.Unmarshal(data, to)
}
