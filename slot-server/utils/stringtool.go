package utils

import (
	"context"
	"errors"
	"math"
	"math/rand"
	"regexp"
	"slot-server/global"
	"strconv"
	"strings"
	"time"
)

func Str2IntSlice(s string) (intArr []int64, err error) {
	strArr := FormatCommand(s)
	for _, v := range strArr {
		if v == "" {
			return []int64{}, errors.New("不能为空字符串")
		}
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return []int64{}, err
		}
		intArr = append(intArr, id)
	}
	return
}

func FormatCommand(s string) []string {
	s = FormatCommandStr(s)
	arr := strings.Split(s, " ")
	return arr
}

func FormatCommandStr(s string) string {
	// 所有制表符替换为单个空格
	reg := regexp.MustCompile("\\s+")
	s = reg.ReplaceAllString(s, " ")
	s = strings.Trim(s, " ")
	return s
}

func ArrStr2Lower(arr []string) []string {
	var lower []string
	for _, s := range arr {
		lower = append(lower, strings.ToLower(s))
	}
	return lower
}

func IntJoin(i []int64, sep string) string {
	var (
		s   string
		end = len(i) - 1
	)
	for k, v := range i {
		s += strconv.FormatInt(v, 10)
		if end != k {
			s += sep
		}
	}
	return s
}

// RandomString 在数字、大写字母、小写字母范围内生成num位的随机字符串
func RandomString(length int) string {
	// 48 ~ 57 数字
	// 65 ~ 90 A ~ Z
	// 97 ~ 122 a ~ z
	// 一共62个字符，在0~61进行随机，小于10时，在数字范围随机，
	// 小于36在大写范围内随机，其他在小写范围随机
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		t := rand.Intn(62)
		if t < 10 {
			result = append(result, strconv.Itoa(rand.Intn(10)))
		} else if t < 36 {
			result = append(result, string(rune(rand.Intn(26)+65)))
		} else {
			result = append(result, string(rune(rand.Intn(26)+97)))
		}
	}
	return strings.Join(result, "")
}

// Same 判断字符存再是否包含重复字符
func Same(s string) bool {
	if strings.Contains(s[:len(s)-1], s[len(s)-1:]) {
		return true
	}
	return false
}

// Cut 截断字符串
func Cut(s, sep string) (before, after string, found bool) {
	if i := strings.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}

// Mid 获取中间字符串
func Mid(s, before, after string) (mid string, found bool) {
	if before == "" || after == "" {
		return "", false
	}
	if a := strings.Index(s, before); a >= 0 {
		length := len(before)
		b := strings.Index(s[a+length:], after)
		if b >= 0 {
			b += a + length
			return s[a+length : b], true
		}
	}
	return "", false
}

// RemoveMid 去除中间字符串
func RemoveMid(s, before, after string) (mid string, found bool) {
	if before == "" || after == "" {
		return "", false
	}
	if a := strings.Index(s, before); a >= 0 {
		length := len(before)
		b := strings.Index(s[a+length:], after)
		if b >= 0 {
			b += a + length
			return s[a+length : b], true
		}
	}
	return "", false
}

// NoOrderMid 没有前后的Mid
func NoOrderMid(s, before, after string) (mid string, found bool) {
	if before == "" || after == "" {
		return "", false
	}
	if a := strings.Index(s, after); a >= 0 {
		if after == before {
			length := len(before)
			b := strings.Index(s[a+length:], after)
			if b >= 0 {
				b += a + length
				return s[a+length : b], true
			}
		} else {
			if b := strings.Index(s, before); b >= 0 {
				if a < b {
					return s[a+len(after) : b], true
				} else if a == b {
					b = strings.Index(s[a+1:], after)
					if b >= 0 {
						b += a + 1
						return s[a+len(after) : b], true
					}
				} else {
					return s[b+len(before) : a], true
				}
			}
		}
	}
	return "", false
}

var base = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func Base62Encode(num int) string {
	if num == 0 {
		return string(base[0])
	}

	encoded := ""
	baseLength := len(base)

	for num > 0 {
		remainder := num % baseLength
		encoded = string(base[remainder]) + encoded
		num = num / baseLength
	}

	return encoded
}

func Base62Decode(str string) int {
	baseLength := len(base)
	strLength := len(str)
	decoded := 0

	for i, char := range str {
		power := strLength - i - 1
		index := strings.IndexRune(base, char)
		decoded += index * int(math.Pow(float64(baseLength), float64(power)))
	}

	return decoded
}

func BuildMicrosecondOrderNo() string {
	orderNo := ""
	now := time.Now()
	// 年字符和时分秒字符
	yearStr := now.Format("2006")[2:]
	hourMinuteStr := now.Format("1504")
	secondStr := now.Format("05")

	year, _ := strconv.Atoi(yearStr)
	hourMinute, _ := strconv.Atoi(hourMinuteStr)
	second, _ := strconv.Atoi(secondStr)
	//月日字符
	day := now.Format("0201")
	// 微秒
	microsecond := now.Nanosecond() / 1000

	// 年和时分秒转base62
	base62Year := Base62Encode(year)
	base62HourMinute := Base62Encode(hourMinute)
	base62Second := Base62Encode(second)
	base62Nanosecond := Base62Encode(microsecond)
	date := base62Year + base62HourMinute + base62Second + base62Nanosecond
	// 月日 + 拼接年时分
	orderNo += day + date

	return orderNo
}

func BuildOrderNo() string {
	var (
		ok                    bool
		orderNo               string
		microsecondOrderNoKey = "MicrosecondOrderNo:"
	)
	orderNo = BuildMicrosecondOrderNo()
	ok = global.GVA_REDIS.SetNX(context.Background(), microsecondOrderNoKey+orderNo, 1, 1).Val()
	// 只有明确为nil值错误才重试 防止redis挂掉无限重试
	if !ok {
		orderNo = BuildOrderNo()
	}
	return orderNo
}

// Substr substr()
func Substr(str string, start uint, length int) string {
	return ""
}

func Unique(arr []int) []int {
	result := make([]int, 0)
	m := make(map[int]bool) //map的值不重要
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
	}
	return result
}
