package backend

import (
	"go.uber.org/zap"
	"log"
	"slot-server/global"
	"slot-server/service/cluster/backend/codec"
	"slot-server/service/cluster/backend/message"
	"slot-server/service/cluster/backend/packet"
	"slot-server/utils/helper"
	"sync"

	"google.golang.org/protobuf/proto"
)

var (
	hsd []byte // handshake data
	had []byte // handshake ack data
	hbd []byte // heartbeat data
)

func init() {
	var err error
	hsd, err = codec.Encode(packet.Handshake, nil)
	if err != nil {
		panic(err)
	}

	had, err = codec.Encode(packet.HandshakeAck, nil)
	if err != nil {
		panic(err)
	}

	hbd, err = codec.Encode(packet.Heartbeat, nil)
	if err != nil {
		panic(err)
	}
}

type (

	// Callback represents the callback type which will be called
	// when the correspond events is occurred.
	Callback func(data interface{})

	// Connector is a tiny Nano client
	Connector struct {
		conn  *helper.WebsocketClientManager // low-level connection
		codec *codec.Decoder                 // decoder
		//die    chan struct{}                  // connector close channel
		chSend chan []byte // send queue
		mid    uint64      // message id

		// events handler
		muEvents sync.RWMutex
		events   map[string]Callback

		// response handler
		muResponses sync.RWMutex
		responses   map[uint64]Callback

		connectedCallback func() // connected callback
		first             bool
	}
)

// NewConnector create a new Connector
func NewConnector() *Connector {
	return &Connector{
		//die:       make(chan struct{}),
		codec:     codec.NewDecoder(),
		chSend:    make(chan []byte, 64),
		mid:       1,
		events:    map[string]Callback{},
		responses: map[uint64]Callback{},
	}
}

// Start connect to the server and send/recv between the c/s
func (c *Connector) Start(url, path, scheme string) {
	c.first = true
	c.conn = helper.NewWsClientManager(url, path, scheme, 10, func() {
		// send handshake packet
		c.send(hsd)
	})
	c.conn.Start()
	// read and process network message
	go c.read()
	return
}

// OnConnected set the callback which will be called when the client connected to the server
func (c *Connector) OnConnected(callback func()) {
	c.connectedCallback = callback
}

// Request send a request to server and register a callbck for the response
func (c *Connector) Request(route string, v proto.Message, callback Callback) error {
	data, err := serialize(v)
	if err != nil {
		return err
	}

	msg := &message.Message{
		Type:  message.Request,
		Route: route,
		ID:    c.mid,
		Data:  data,
	}

	c.setResponseHandler(c.mid, callback)
	if err := c.sendMessage(msg); err != nil {
		c.setResponseHandler(c.mid, nil)
		return err
	}

	return nil
}

// Notify send a notification to server
func (c *Connector) Notify(route string, v proto.Message) error {
	data, err := serialize(v)
	if err != nil {
		return err
	}

	msg := &message.Message{
		Type:  message.Notify,
		Route: route,
		Data:  data,
	}
	return c.sendMessage(msg)
}

// On add the callback for the event
func (c *Connector) On(event string, callback Callback) {
	c.muEvents.Lock()
	defer c.muEvents.Unlock()

	c.events[event] = callback
}

// Close the connection, and shutdown the benchmark
func (c *Connector) Close() {
	c.conn.Close()
	//close(c.die)
}

func (c *Connector) eventHandler(event string) (Callback, bool) {
	c.muEvents.RLock()
	defer c.muEvents.RUnlock()

	cb, ok := c.events[event]
	return cb, ok
}

func (c *Connector) responseHandler(mid uint64) (Callback, bool) {
	c.muResponses.RLock()
	defer c.muResponses.RUnlock()

	cb, ok := c.responses[mid]
	return cb, ok
}

func (c *Connector) setResponseHandler(mid uint64, cb Callback) {
	c.muResponses.Lock()
	defer c.muResponses.Unlock()

	if cb == nil {
		delete(c.responses, mid)
	} else {
		c.responses[mid] = cb
	}
}

func (c *Connector) sendMessage(msg *message.Message) error {
	data, err := msg.Encode()
	if err != nil {
		return err
	}

	//log.Printf("%+v",msg)

	payload, err := codec.Encode(packet.Data, data)
	if err != nil {
		return err
	}

	c.mid++
	c.send(payload)

	return nil
}

//
//func (c *Connector) write() {
//	defer close(c.chSend)
//
//	for {
//		select {
//		case data := <-c.chSend:
//			//if _, err := c.conn.Write(data); err != nil {
//			//	log.Println(err.Error())
//			//	c.Close()
//			//}
//			c.conn.Send(data)
//		case <-c.die:
//			return
//		}
//	}
//}

func (c *Connector) send(data []byte) {
	c.conn.WriteMsgChan <- data
}

func (c *Connector) read() {
	buf := make([]byte, 2048)
	defer c.Close()
	for {
		buf = <-c.conn.ReadMsgChan
		packets, err := c.codec.Decode(buf)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			continue
		}
		for i := range packets {
			p := packets[i]
			c.processPacket(p)
		}
	}
}

func (c *Connector) processPacket(p *packet.Packet) {
	switch p.Type {
	case packet.Handshake:
		c.send(had)
		if c.first {
			c.first = false
			c.connectedCallback()
		}
	case packet.Data:
		msg, err := message.Decode(p.Data)
		if err != nil {
			global.GVA_LOG.Error("message decode error", zap.Any("err", err))
			return
		}
		c.processMessage(msg)

	case packet.Heartbeat:
		c.send(hbd)
	case packet.Kick:
		c.Close()
	}
}

func (c *Connector) processMessage(msg *message.Message) {
	switch msg.Type {
	case message.Push:
		cb, ok := c.eventHandler(msg.Route)
		if !ok {
			log.Println("event handler not found", msg.Route)
			return
		}

		cb(msg.Data)

	case message.Response:
		cb, ok := c.responseHandler(msg.ID)
		if !ok {
			log.Println("response handler not found", msg.ID)
			return
		}

		cb(msg.Data)
		c.setResponseHandler(msg.ID, nil)
	}
}

func serialize(v proto.Message) ([]byte, error) {
	data, err := proto.Marshal(v)
	if err != nil {
		return nil, err
	}
	return data, nil
}
