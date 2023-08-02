package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"slot-server/global"
	"slot-server/utils/conver"
	"strings"
	"time"
)

type Gurl struct {
	Url      string
	Data     interface{}       // post数据
	BodyType BodyType          // post数据类型
	Param    map[string]string // query参数
	Method   string
	Header   map[string]string
	Cookie   map[string]string
	option   Option
	client   http.Client //custom client

	Payload   []byte   // post数据的字符串形式
	Response  Response // 响应结果
	ResBody   []byte   // 响应数据体
	Err       error    // 错误信息
	Consuming int64    // 请求耗时(微妙)
}

type Option struct {
	Timeout    time.Duration // request timeout
	Proxy      string
	SkipVerify bool // skip ssl verify

	Before func(*Gurl)
	After  func(*Gurl)
}

type Response *http.Response

type BodyType int

const (
	_ BodyType = iota
	TEXT
	FORM
	JSON
	XML
)

// NewGurl New request
func NewGurl(method, url string) *Gurl {
	return &Gurl{Url: url, Method: method}
}

// Set options
func (g *Gurl) Set(option Option) *Gurl {
	g.option = option
	return g
}

// SetData Data Set data and the type, default is JSON
func (g *Gurl) SetData(data interface{}, body ...BodyType) *Gurl {
	g.Data = data
	if len(body) > 0 {
		g.BodyType = body[0]
	} else {
		g.BodyType = JSON
	}
	return g
}

// SetParam Param Set params
func (g *Gurl) SetParam(param map[string]string) *Gurl {
	g.Param = param
	return g
}

// SetHeader Header Set headers
func (g *Gurl) SetHeader(header map[string]interface{}) *Gurl {
	headers := map[string]string{}
	for k, v := range header {
		headers[k] = conver.StringMust(v)
	}
	g.Header = headers
	return g
}

// SetCookie Cookie Set cookies
func (g *Gurl) SetCookie(cookie map[string]string) *Gurl {
	g.Cookie = cookie
	return g
}

// Client Set custom client
func (g *Gurl) Client(client http.Client) *Gurl {
	g.client = client
	return g
}

// Combined urls and parameters
func (g *Gurl) urlWithParam() (err error) {
	if g.Param == nil {
		return
	}

	var u *url.URL
	if u, err = url.Parse(g.Url); err != nil {
		return
	}

	q := u.Query()
	for k, v := range g.Param {
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()
	g.Url = u.String()

	return
}

func (g *Gurl) Request() (response Response, err error) {
	if g.Url == "" {
		return nil, errors.New("no url")
	} else {
		if err = g.urlWithParam(); err != nil {
			return
		}
	}
	if g.Method == "" {
		return nil, errors.New("no method")
	} else {
		g.Method = strings.ToUpper(g.Method)
	}

	if g.Data != nil && g.Method != "GET" {
		switch g.BodyType {
		case TEXT:
			g.Payload = []byte(g.Data.(string))
		case FORM:
			formData := ""
			data, ok := g.Data.(map[string]interface{})
			if !ok {
				return
			}
			for k, v := range data {
				formData += k + "=" + url.QueryEscape(conver.StringMust(v)) + "&"
			}
			formData = strings.TrimRight(formData, "&")
			g.Payload = []byte(formData)
		case JSON:
			if g.Payload, err = global.Json.Marshal(g.Data); err != nil {
				return nil, err
			}
		case XML:
			if g.Payload, err = xml.Marshal(g.Data); err != nil {
				return nil, err
			}
		}
	}
	request, err := http.NewRequest(g.Method, g.Url, bytes.NewReader(g.Payload))
	if err != nil {
		return
	}

	switch g.BodyType {
	case FORM:
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	case JSON:
		request.Header.Set("Content-Type", "application/json")
	}

	if g.Header != nil {
		for k, v := range g.Header {
			request.Header.Set(k, v)
		}
	}

	// options
	opt := g.option

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: opt.SkipVerify},
	}

	//使用代理
	if opt.Proxy != "" {
		proxy, err := url.Parse(opt.Proxy)
		if err == nil {
			tr.Proxy = http.ProxyURL(proxy)
		}
	}

	g.client.Transport = tr
	g.client.Timeout = opt.Timeout

	response, err = g.client.Do(request)
	return
}

func (g *Gurl) Do() ([]byte, error) {
	s := time.Now()
	g.Response, g.Err = g.Request()
	g.Consuming = time.Now().Sub(s).Microseconds()
	if g.Err != nil {
		return nil, g.Err
	}
	if g.Response.StatusCode != 200 {
		return nil, errors.New(g.Response.Status)
	}
	g.ResBody, g.Err = io.ReadAll(g.Response.Body)
	global.GVA_LOG.WithOptions(zap.AddCallerSkip(1)).Info(
		fmt.Sprintf("\ngurl Do data {url: %s %s} %dμs \n{request: %s}\n{response: %s}",
			g.Method, g.Url, g.Consuming, g.Payload, string(g.ResBody)))
	return g.ResBody, g.Err
}
