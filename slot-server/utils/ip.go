package utils

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/lonng/nano/session"
	"net"
	"slot-server/global"
	"strconv"
	"time"
)

type Ip struct {
	Ip          string  `json:"ip"`
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Message     string  `json:"message"`
	Country     string  `json:"country"`     // 国家/地区名称 China
	CountryCode string  `json:"countryCode"` // 两个字母的国家代码 CN
	Region      string  `json:"region"`      // 地区/州短码 HB
	RegionName  string  `json:"regionName"`  // 地区/州 Hubei
	City        string  `json:"city"`        // 城市 Wuhan
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"` // 时区 Asia/Shanghai
	Isp         string  `json:"isp"`      // ISP name
	Org         string  `json:"org"`      // Organization name
	As          string  `json:"as"`
}

func GetIp(ip string) (ipInfo Ip, err error) {
	info, ok := global.BlackCache.Get(ip)
	if ok {
		return info.(Ip), nil
	}
	url := "https://pro.ip-api.com/json/" + ip + "?key=ndqP5jzUqKX4XQ9"
	res, err := NewGurl("GET", url).Do()
	if err != nil {
		return
	}
	err = jsoniter.Unmarshal(res, &ipInfo)
	if err != nil {
		return
	}
	if ipInfo.Status != "success" {
		err = errors.New("status :" + ipInfo.Status + " msg: " + ipInfo.Message)
		return
	}

	global.BlackCache.Set(ip, ipInfo, 120*time.Second)
	return
}

func RemoteIpString(s *session.Session) string {
	addr := s.NetworkEntity().RemoteAddr()
	host, _, _ := SpliteAddress(addr.String())
	return host
}

// SpliteAddress 将普通地址格式(host:port)拆分
func SpliteAddress(addr string) (host string, port int, err error) {
	var portStr string
	host, portStr, err = net.SplitHostPort(addr)
	if err != nil {
		return "", 0, err
	}
	port, err = strconv.Atoi(portStr)
	if err != nil {
		return "", 0, err
	}
	return
}
