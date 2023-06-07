package main

import (
	"bytes"
	"errors"
)

var ErrWrongPacketType = errors.New("wrong packet type")

type Packet struct {
	Type   Type
	Length int
	Data   []byte
}

type Type byte

const (
	_ Type = iota
	// Handshake represents a handshake: request(client) <====> handshake response(server)
	Handshake = 0x01

	// HandshakeAck represents a handshake ack from client to server
	HandshakeAck = 0x02

	// Heartbeat represents a heartbeat
	Heartbeat = 0x03

	// Data represents a common data packet
	Data = 0x04

	// Kick represents a kick off packet
	Kick = 0x05 // disconnect message from server
)


type Decoder struct {
	buf  *bytes.Buffer
	size int  // last packet length
	typ  byte // last packet type
}

const (
	HeadLength    = 4
	MaxPacketSize = 64 * 1024
)


var ErrPacketSizeExcced = errors.New("codec: packet size exceed")

// NewDecoder returns a new decoder that used for decode network bytes slice.
func NewDecoder() *Decoder {
	return &Decoder{
		buf:  bytes.NewBuffer(nil),
		size: -1,
	}
}

func bytesToInt(b []byte) int {
	result := 0
	for _, v := range b {
		result = result<<8 + int(v)
	}
	return result
}

func (c *Decoder) forward() error {
	header := c.buf.Next(HeadLength)
	c.typ = header[0]
	if c.typ < Handshake || c.typ > Kick {
		return ErrPacketSizeExcced
	}
	c.size = bytesToInt(header[1:])

	// packet length limitation
	if c.size > MaxPacketSize {
		return ErrPacketSizeExcced
	}
	return nil
}

func (c *Decoder) Decode(data []byte) ([]*Packet, error) {
	c.buf.Write(data)

	var (
		packets []*Packet
		err     error
	)
	// check length
	if c.buf.Len() < HeadLength {
		return nil, err
	}

	// first time
	if c.size < 0 {
		if err = c.forward(); err != nil {
			return nil, err
		}
	}

	for c.size <= c.buf.Len() {
		p := &Packet{Type: Type(c.typ), Length: c.size, Data: c.buf.Next(c.size)}
		packets = append(packets, p)

		// more packet
		if c.buf.Len() < HeadLength {
			c.size = -1
			break
		}

		if err = c.forward(); err != nil {
			return nil, err

		}

	}

	return packets, nil
}