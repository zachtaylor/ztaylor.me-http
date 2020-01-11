package websocket // import "ztaylor.me/http/websocket"

import (
	"golang.org/x/net/websocket"
	"ztaylor.me/cast"
	"ztaylor.me/http/session"
)

// T is a websocket connection
type T struct {
	ID      string
	Session *session.T
	conn    *websocket.Conn
	send    chan []byte
	recv    <-chan *Message
	done    chan bool
}

// New creates an initialied orphan websocket
func New(conn *websocket.Conn) *T {
	return &T{
		conn: conn,
		send: make(chan []byte),
		recv: ReadMessageChan(conn),
		done: make(chan bool),
	}
}

func (t *T) String() string {
	return "websocket.T{" + t.conn.Request().RemoteAddr + "}"
}

// SendChan is a writable channel that queues writes to the websocket API
func (t *T) SendChan() chan<- []byte {
	return t.send
}

// ReceiveChan returns the
func (t *T) ReceiveChan() <-chan *Message {
	return t.recv
}

// DoneChan is an observable channel that closes when the socket has been closed
func (t *T) DoneChan() <-chan bool {
	return t.done
}

// Close closes the observable channel
func (t *T) Close() {
	if t.done != nil {
		close(t.send)
		t.send = nil
		// close(t.recv) // closed elsewhere
		t.recv = nil
		close(t.done)
		t.done = nil
	}
}

// Message is a macro for SendMessage(NewMessage)
func (t *T) Message(uri string, json cast.JSON) {
	t.SendMessage(NewMessage(uri, json))
}

// SendMessage is a macro for Send(m.json bytes)
func (t *T) SendMessage(m *Message) {
	t.Send(cast.BytesS(m.JSON().String()))
}

// Send starts a goroutine to push to send chan
func (t *T) Send(buff []byte) {
	go func() {
		if t.send != nil {
			t.send <- buff
		}
	}()
}
