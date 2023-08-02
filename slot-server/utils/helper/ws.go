package helper

import (
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/url"
	"slot-server/global"
	"sync"
	"time"
)

type WebsocketClientManager struct {
	Conn           *websocket.Conn
	Addr           *string
	Path           string
	WriteMsgChan   chan []byte
	ReadMsgChan    chan []byte
	IsAlive        bool
	Timeout        int
	Scheme         string
	Wg             *sync.WaitGroup
	ReconnectClose chan struct{}
	ReadClose      chan struct{}
	WriteClose     chan struct{}
	Fn             func()
}

func NewWsClientManager(addr, path, scheme string, timeout int, fn func()) *WebsocketClientManager {
	wg := sync.WaitGroup{}
	return &WebsocketClientManager{
		Addr:           &addr,
		Path:           path,
		WriteMsgChan:   make(chan []byte, 2048),
		ReadMsgChan:    make(chan []byte, 2048),
		IsAlive:        false,
		Timeout:        timeout,
		ReconnectClose: make(chan struct{}, 1),
		ReadClose:      make(chan struct{}, 1),
		WriteClose:     make(chan struct{}, 1),
		Scheme:         scheme,
		Wg:             &wg,
		Fn:             fn,
	}
}

func (wsc *WebsocketClientManager) Dail() error {
	var err error
	u := url.URL{Scheme: wsc.Scheme, Host: *wsc.Addr, Path: wsc.Path}
	wsc.Conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		global.GVA_LOG.Error("Dail to err:", zap.Error(err), zap.String("url", u.String()))
		return err
	}
	wsc.IsAlive = true
	global.GVA_LOG.Infof("Connecting to %s 成功！！！", u.String())
	return nil
}

// sendMsgThread 发送消息
func (wsc *WebsocketClientManager) sendMsgThread() {
	go func() {
		defer wsc.Wg.Done()
		wsc.Wg.Add(1)
		global.GVA_LOG.Info("Write service start")
		for {
			select {
			case msg := <-wsc.WriteMsgChan:
				err := wsc.Conn.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					global.GVA_LOG.Error("Write err:", zap.Error(err))
					continue
				}
				//if global.GVA_CONFIG.System.Mode != enum.Prod {
				//global.GVA_LOG.Info("Write", zap.String("message", string(msg)))
				//}
			case <-wsc.WriteClose:
				global.GVA_LOG.Info("Write channel closure")
				return
			}
		}
	}()
}

// 读取消息
func (wsc *WebsocketClientManager) readMsgThread() {
	go func() {
		defer wsc.Wg.Done()
		wsc.Wg.Add(1)
		global.GVA_LOG.Info("Read service start")
		for {
			select {
			default:
				_, message, err := wsc.Conn.ReadMessage()
				if err != nil {
					global.GVA_LOG.Error("Read err:", zap.Error(err))
					wsc.IsAlive = false
					wsc.WriteClose <- struct{}{}
					return
				}
				wsc.ReadMsgChan <- message
				//global.GVA_LOG.Info("Read", zap.String("message", string(message)))
			case <-wsc.ReadClose:
				global.GVA_LOG.Info("Read channel closure")
				return
			}
		}
	}()
}

// Start 开启服务并重连
func (wsc *WebsocketClientManager) Start() {
	go func() {
		defer wsc.Wg.Done()
		wsc.Wg.Add(1)
		tick := time.NewTicker(time.Second * 10)
		for {
			select {
			case <-wsc.ReconnectClose:
				global.GVA_LOG.Info("Reconnection service closed")
				return
			case <-tick.C:
				if wsc.IsAlive == false {
					err := wsc.Dail()
					if err != nil {
						global.GVA_LOG.Error("Dail err:", zap.Error(err))
						continue
					}
					wsc.Fn()
					wsc.sendMsgThread()
					wsc.readMsgThread()
				}
			}
		}
	}()
}

func (wsc *WebsocketClientManager) Send(msg []byte) {
	wsc.WriteMsgChan <- msg
}

func (wsc *WebsocketClientManager) Close() {
	defer wsc.Conn.Close()
	wsc.ReconnectClose <- struct{}{}
	wsc.ReadClose <- struct{}{}
	wsc.WriteClose <- struct{}{}
	wsc.Wg.Wait()
}
