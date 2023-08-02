package helper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Search 字符串数组查找
func Search(arr []string, val string) int {
	for i, s := range arr {
		if s == val {
			return i
		}
	}
	return -1
}

// StrIndexBySplit 分割字符串并返回其中一个
func StrIndexBySplit(s string, sep string, index int) string {
	return SliceVal(strings.Split(s, sep), index)
}

// ReplaceAt 替换@符号后的文字
func ReplaceAt(text string) string {
	return regexp.MustCompile("((@.*? )|(@.*))").ReplaceAllStringFunc(text, func(s string) string {
		if s[len(s)-1:] == " " {
			s = "@[" + strings.Repeat("*", len(s)-1) + "] "
		} else {
			s = "@[" + strings.Repeat("*", len(s)-1) + "]"
		}
		return s
	})
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

// IsAlphaNumeric 判断字符串是否只包含数字或字母或下划线
func IsAlphaNumeric(s string) bool {
	for _, c := range s {
		if (c < '0' || c > '9') && (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') && c != '_' {
			return false
		}
	}
	return true
}

func BinToDec(s string) string {
	v, _ := strconv.ParseInt(s, 2, 64)
	return strconv.FormatInt(v, 10)
}

func DecToBin(s string) string {
	v, _ := strconv.ParseInt(s, 10, 64)
	return strconv.FormatInt(v, 2)
}

func Unit(x int) string {
	if x > 100000000 {
		return fmt.Sprintf("%.2f亿", float64(x)/100000000)
	}
	if x > 10000 {
		return fmt.Sprintf("%.2f万", float64(x)/10000)
	}
	return fmt.Sprintf("%d", x)
}

func StrReverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	return string(r)
}

func IntReverse(i int) int {
	n, _ := strconv.Atoi(StrReverse(strconv.Itoa(i)))
	return n
}

func LeaveOutStr(s string, length int, tail ...string) string {
	if len(s) <= length {
		return s
	}
	return s[:length] + strings.Join(tail, "")
}
