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
	conn      *websocket.Conn
	send      chan []byte
	closeOnce sync.Once
}

// 游戏房间
type gameRoom struct {
	clients map[*client]bool
	mutex   sync.Mutex
}

// 向所有客户端广播消息
func (gr *gameRoom) broadcast(message []byte) {
	gr.mutex.Lock()
	defer gr.mutex.Unlock()
	for c := range gr.clients {
		select {
		case c.send <- message:
		default:
			gr.removeClient(c)
		}
	}
}

// 移除客户端
func (gr *gameRoom) removeClient(c *client) {
	gr.mutex.Lock()
	defer gr.mutex.Unlock()
	if _, ok := gr.clients[c]; ok {
		close(c.send)
		delete(gr.clients, c)
	}
}

// 处理 WebSocket 连接
func handleWebSocket(gr *gameRoom, w http.ResponseWriter, r *http.Request) {
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

	// 加入游戏房间
	gr.mutex.Lock()
	gr.clients[client] = true
	gr.mutex.Unlock()

	// 处理 WebSocket 连接
	go func() {
		defer func() {
			gr.removeClient(client)
			client.conn.Close()
		}()
		for {
			// 读取消息
			_, message, err := client.conn.ReadMessage()
			if err != nil {
				log.Println("Read:", err)
				return
			}

			// 打印消息
			log.Printf("Received: %s", message)

			// 广播消息给所有客户端
			gr.broadcast(message)
		}
	}()

	// 发送消息给客户端
	go func() {
		for {
			message, ok := <-client.send
			if !ok {
				return
			}
			err := client.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("Write:", err)
				gr.removeClient(client)
				return
			}
		}
	}()
}

func main() {

	// 创建游戏房间
	gameroom := &gameRoom{
		clients: make(map[*client]bool),
	}

	// 静态文件服务器
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// WebSocket 服务端
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(gameroom, w, r)
	})

	// 启动服务器
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
