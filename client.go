package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	// Socket for client
	socket *websocket.Conn
	// Channel where message is send
	send_chan chan []byte
	room      *room
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.fwd_chan <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send_chan {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
