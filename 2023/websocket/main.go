package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)


func GetPack()  ([]byte, error) {
	type UserMessage struct {
		luYou   string `json:"name"`
		Name    string `json:"name"`
		Content string `json:"content"`
		Route      string // route for locating service
	}

	userMessage := UserMessage{
		luYou:   "room.Message1",
		Name:    "room",
		Content: "Message1",
		Route: "room.join",
	}
	marshal, _ := json.Marshal(userMessage)
	//p1 := &Packet{Type: Handshake, Data: marshal, Length: len(marshal)}

	return Encode(Handshake, marshal)


/*	d1 := NewDecoder()
	packets, err := d1.Decode(pp1)

	if !reflect.DeepEqual(p1, packets[0]) {
		log.Println("expect: %v, got: %v", p1, packets[0])
	}
	if err != nil {
		log.Println(err.Error())
	}*/

	//return p1
}

func Encode(typ Type, data []byte) ([]byte, error) {
	if typ < Handshake || typ > Kick {
		return nil, ErrWrongPacketType
	}

	p := &Packet{Type: typ, Length: len(data)}
	buf := make([]byte, p.Length+HeadLength)
	buf[0] = byte(p.Type)

	copy(buf[1:HeadLength], intToBytes(p.Length))
	copy(buf[HeadLength:], data)

	return buf, nil
}

// Encode packet data length to bytes(Big end)
func intToBytes(n int) []byte {
	buf := make([]byte, 3)
	buf[0] = byte((n >> 16) & 0xFF)
	buf[1] = byte((n >> 8) & 0xFF)
	buf[2] = byte(n & 0xFF)
	return buf
}




func main() {
	u := url.URL{Scheme: "ws", Host: "0.0.0.0:34590", Path: "/nano"}
	log.Printf("client1 connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial server:", err)
	}
	defer c.Close()

	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("client read err:", err)
				return
			}
			log.Printf("client recv msg: %s", message)
		}
	}()
	//room.message {"name":"guest1682496232864","content":"123"}

	/*type UserMessage struct {
		luYou   string `json:"name"`
		Name    string `json:"name"`
		Content string `json:"content"`
	}

	userMessage := UserMessage{
		luYou:   "room.Message1",
		Name:    "room",
		Content: "Message1",
	}

	marshal, _ := json.Marshal(userMessage)*/
	packet ,_ := GetPack()
	err = c.WriteMessage(websocket.TextMessage, packet)

	if err != nil {
		log.Println("client write:", err)
		return
	}

	for  {
		
	}

	/*for {
		select {
		case <-time.Tick(time.Second * 3):
			fmt.Println("123")
			err := c.WriteMessage(websocket.TextMessage, marshal)
			if err != nil {
				log.Println("client write:", err)
				return
			}
		}
	}*/
}
