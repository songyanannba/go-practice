package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocket 升级器
var upgrader = websocket.Upgrader{}

// 客户端连接
type client struct {
	conn           *websocket.Conn
	send           chan []byte
	lastActiveTime time.Time // 最后活跃时间
}

// 聊天室
type chatRoom struct {
	clients map[*client]bool
	mutex   sync.Mutex
}

// 心跳检测间隔
const heartbeatInterval = 10 * time.Second

// 最长允许的空闲时间
const maxIdleTime = 60 * time.Second

// 向所有客户端广播消息
func (cr *chatRoom) broadcast(message []byte) {
	cr.mutex.Lock()
	defer cr.mutex.Unlock()
	for c := range cr.clients {
		select {
		case c.send <- message:
		default:
			close(c.send)
			delete(cr.clients, c)
		}
	}
}

// 处理 WebSocket 连接
func handleWebSocket(cr *chatRoom, w http.ResponseWriter, r *http.Request) {
	// 升级为 WebSocket 连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade:", err)
		return
	}

	// 创建客户端连接
	client := &client{
		conn:           conn,
		send:           make(chan []byte, 256),
		lastActiveTime: time.Now(),
	}

	// 加入聊天室
	cr.mutex.Lock()
	cr.clients[client] = true
	cr.mutex.Unlock()

	// 处理 WebSocket 连接
	go func() {
		defer func() {
			cr.mutex.Lock()
			defer cr.mutex.Unlock()
			delete(cr.clients, client)
			client.conn.Close()
		}()
		for {
			// 读取消息
			_, message, err := client.conn.ReadMessage()
			if err != nil {
				log.Println("Read:", err)
				break
			}

			// 打印消息
			log.Printf("Received: %s", message)

			// 更新最后活跃时间
			client.lastActiveTime = time.Now()

			// 广播消息给所有客户端
			cr.broadcast(message)
		}
	}()

	// 发送消息给客户端
	go func() {
		for {
			message, ok := <-client.send
			if !ok {
				break
			}
			err := client.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("Write:", err)
				break
			}
		}
	}()

	// 心跳检测
	go func() {
		for {
			time.Sleep(heartbeatInterval)
			if time.Since(client.lastActiveTime) > maxIdleTime {
				log.Printf("Client idle timeout: %s", client.conn.RemoteAddr())
				client.conn.Close()
				break
			}
			err := client.conn.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				log.Println("Write Ping:", err)
				break
			}
		}
	}()
}

func main() {
	// 创建聊天室
	chatroom := &chatRoom{
		clients: make(map[*client]bool),
	}

}