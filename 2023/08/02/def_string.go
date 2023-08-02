package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type DefString string

func (ds *DefString) FormatCommand() []string {
	//command := FormatCommandStr(string(*ds))
	arr := strings.Split(string(*ds), " ")
	return arr
}

func (ds *DefString) SplitStr(sep string) (int, string, error) {
	colW := strings.Split(string(*ds), sep)
	if len(colW) != 2 {
		return 0, "", fmt.Errorf("format error")
	}
	colNum, err := strconv.Atoi(colW[0])
	if err != nil {
		return 0, "", err
	}
	return colNum, colW[1], nil
}

// StrToMap 字符串被分割的第一部分 可以转换成整形
func (ds *DefString) StrToMap(sep string) (map[int]string, error) {
	strMap := map[int]string{}
	if len(*ds) == 0 {
		return strMap, nil
	}
	strS := ds.FormatCommand()
	for _, str := range strS {
		if str == "" {
			continue
		}
		defString := DefString(str)
		info, s, err := defString.SplitStr(sep)
		if err != nil {
			return strMap, err
		}
		strMap[info] = s
	}
	return strMap, nil
}
