package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io"
	"net/http"
	"runtime"
	"slot-server/global"
	"strings"
	"sync"
)

// PanicRecover 捕获恐慌
func PanicRecover() {
	if err := recover(); err != nil {
		global.GVA_LOG.Error("Analysis panic " + fmt.Sprintf("%v", err))
		global.GVA_LOG.Error(Stack())
		return
	}
}

func Stack() string {
	buf := make([]byte, 10000)
	n := runtime.Stack(buf, false)
	buf = buf[:n]

	s := string(buf)

	// skip nano frames lines
	const skip = 7
	count := 0
	index := strings.IndexFunc(s, func(c rune) bool {
		if c != '\n' {
			return false
		}
		count++
		return count == skip
	})
	return s[index+1:]
}

func ContentDecode(data []byte, v any, fields ...string) (err error) {
	j, err := simplejson.NewJson(data)
	if err != nil {
		return
	}
	for _, field := range fields {
		j = j.Get(field)
	}
	bytes, err := j.MarshalJSON()
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &v)
	if err != nil {
		return
	}
	return
}

func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

func MergeMap[T comparable, V any](maps ...map[T]V) map[T]V {
	m := map[T]V{}
	for _, map1 := range maps {
		for k, v := range map1 {
			m[k] = v
		}
	}
	return m
}

// RequestInputs 获取并打印所有请求参数
func RequestInputs(c *gin.Context) (map[string]interface{}, error) {
	const defaultMemory = 32 << 20
	contentType := c.ContentType()

	var (
		dataMap  = make(map[string]interface{})
		queryMap = make(map[string]interface{})
		postMap  = make(map[string]interface{})
	)

	// @see gin@v1.7.7/binding/query.go ==> func (queryBinding) Bind(req *http.Request, obj interface{})
	for k := range c.Request.URL.Query() {
		queryMap[k] = c.Query(k)
	}

	if "application/json" == contentType {
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		// @see gin@v1.7.7/binding/json.go ==> func (jsonBinding) Bind(req *http.Request, obj interface{})
		if c.Request != nil && c.Request.Body != nil {
			if err := json.NewDecoder(c.Request.Body).Decode(&postMap); err != nil {
				return nil, err
			}
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	} else if "multipart/form-data" == contentType {
		// @see gin@v1.7.7/binding/form.go ==> func (formMultipartBinding) Bind(req *http.Request, obj interface{})
		if err := c.Request.ParseMultipartForm(defaultMemory); err != nil {
			return nil, err
		}
		for k, v := range c.Request.PostForm {
			if len(v) > 1 {
				postMap[k] = v
			} else if len(v) == 1 {
				postMap[k] = v[0]
			}
		}
	} else {
		// ParseForm 解析 URL 中的查询字符串，并将解析结果更新到 r.Form 字段
		// 对于 POST 或 PUT 请求，ParseForm 还会将 body 当作表单解析，
		// 并将结果既更新到 r.PostForm 也更新到 r.Form。解析结果中，
		// POST 或 PUT 请求主体要优先于 URL 查询字符串（同名变量，主体的值在查询字符串的值前面）
		// @see gin@v1.7.7/binding/form.go ==> func (formBinding) Bind(req *http.Request, obj interface{})
		if err := c.Request.ParseForm(); err != nil {
			return nil, err
		}
		if err := c.Request.ParseMultipartForm(defaultMemory); err != nil {
			if err != http.ErrNotMultipart {
				return nil, err
			}
		}
		for k, v := range c.Request.PostForm {
			if len(v) > 1 {
				postMap[k] = v
			} else if len(v) == 1 {
				postMap[k] = v[0]
			}
		}
	}

	var mu sync.RWMutex
	for k, v := range queryMap {
		mu.Lock()
		dataMap[k] = v
		mu.Unlock()
	}
	for k, v := range postMap {
		mu.Lock()
		dataMap[k] = v
		mu.Unlock()
	}
	global.GVA_LOG.Info("Accept request",
		zap.String("FullPath", c.FullPath()),
		zap.Any("ContentType", contentType),
		zap.Any("Inputs", dataMap))
	return dataMap, nil
}

// PerfectData 使用arr2中的数据完善arr1中的数据
func PerfectData[T any, V comparable, T2 any](arr []T, getVal func(T) V, db *gorm.DB, arr2Field string, getVal2 func(T2) V, f func(T, T2) T) []T {
	vals := DistinctByFunc(arr, getVal)
	var arr2 []T2
	db.Where(arr2Field+" in ?", vals).Find(&arr2)

	var (
		arr2Map = map[V]T2{}
	)
	for _, t := range arr2 {
		v := getVal2(t)
		arr2Map[v] = t
	}

	for i, t := range arr {
		v := getVal(t)
		arr[i] = f(t, arr2Map[v])
	}
	return arr
}

func Sum[V Int](arr ...V) V {
	var sum V
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumByFunc[V any, V2 Int](arr []V, f func(V) V2) V2 {
	var sum V2
	for _, v := range arr {
		sum += f(v)
	}
	return sum
}

type CopyFunc[T any] interface {
	Copy() T
}

func ListToArr[T CopyFunc[T]](list [][]T) []T {
	var arr []T
	for _, v := range list {
		for _, t := range v {
			arr = append(arr, t.Copy())
		}
	}
	return arr
}

func CopyList[T CopyFunc[T]](list []T) []T {
	var arr []T
	for _, t := range list {
		arr = append(arr, t.Copy())
	}
	return arr

}

func CopyListArr[T CopyFunc[T]](list [][]T) [][]T {
	var arr [][]T
	for _, t := range list {
		arr = append(arr, CopyList(t))
	}
	return arr

}
func NewValue[T any](t *T) *T {
	newT := *t
	return &newT
}
