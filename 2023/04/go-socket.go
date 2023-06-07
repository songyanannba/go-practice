package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// WebSocket 升级器
var upgrader = websocket.Upgrader{}

// 客户端连接
type client struct {
	conn *websocket.Conn
	send chan []byte
}

// 聊天室
type chatRoom struct {
	clients map[*client]bool
	mutex   sync.Mutex
}

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
		conn: conn,
		send: make(chan []byte, 256),
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
}

func main() {
	// 创建聊天室
	chatroom := &chatRoom{
		clients: make(map[*client]bool),
	}

	// 静态文件服务器
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// WebSocket 处理器
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(chatroom, w, r)
	})

	// 启动 HTTP 服务器
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

