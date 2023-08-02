package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/fatih/structs"
	"slot-server/utils/conver"
	"sort"
	"strings"
)

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

func Md5Sign(params map[string]interface{}, secret string) string {
	var (
		sortedKeys []string
		s          string
		arr        []string
	)
	for k, v := range params {
		params[k] = conver.StringMust(v)
		if k == "sign" || params[k].(string) == "" {
			continue
		}
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	for _, k := range sortedKeys {
		arr = append(arr, k+"="+params[k].(string))
	}
	s = strings.Join(arr, "&") + secret
	//global.GVA_LOG.Info("sign the original string: " + s)
	return MD5V([]byte(s))
}

func StructMd5Sign(params any, secret string) string {
	return Md5Sign(structs.Map(params), secret)
}
