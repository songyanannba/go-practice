package str

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type PreStr string

func (s *PreStr) ChangeStrToIntAndStr() (int, string, error) {
	colW := strings.Split(string(*s), ":")
	if len(colW) != 2 {
		return 0, "", fmt.Errorf("initial weight format error")
	}
	colNum, err := strconv.Atoi(colW[0])
	if err != nil {
		return 0, "", err
	}
	return colNum, colW[1], nil
}

func (s *PreStr) GetColMap() (map[int]string, error) {
	if *s == "" {
		return map[int]string{}, nil
	}
	strs := FormatCommand(string(*s))
	strMap := map[int]string{}
	for _, str := range strs {
		if str == "" {
			continue
		}
		preStr := PreStr(str)
		info, s, err := preStr.ChangeStrToIntAndStr()
		if err != nil {
			return nil, err
		}
		strMap[info] = s
	}
	return strMap, nil
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
