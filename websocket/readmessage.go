package websocket

import (
	"golang.org/x/net/websocket"
	"ztaylor.me/cast"
)

// ReadMessage reads a Message from the socket API
func ReadMessage(conn *websocket.Conn) (*Message, error) {
	s, msg := "", &Message{}
	if err := websocket.Message.Receive(conn, &s); err != nil {
		return nil, err
	} else if err := cast.DecodeJSON(cast.NewBuffer(s), msg); err != nil {
		return nil, err
	}
	return msg, nil
}

// ReadMessageChan creates a goroutine monitor using ReadMessage
func ReadMessageChan(conn *websocket.Conn) chan *Message {
	msgs := make(chan *Message)
	go func() {
		for {
			if msg, err := ReadMessage(conn); err == nil {
				msgs <- msg
			} else if err == cast.EOF {
				break
			}
		}
		close(msgs)
	}()
	return msgs
}

// drainMessageChan waits to drink every message and then return
func drainMessageChan(msgs <-chan *Message) {
	for {
		_, ok := <-msgs
		if !ok {
			return
		}
	}
}
