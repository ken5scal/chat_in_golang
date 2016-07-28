package main

import "github.com/gorilla/websocket"

type client struct {
	socket *websocket.Conn
	send chan []byte // channel where message is sent
	room *room
}

// Client read message by using ReadMessage in WebSocket
func (c *client) read() {
	defer c.socket.Close() // Closing WebSocket
	for {
		_, msg, err := c.socket.ReadMessage()
		if  err != nil {
			return
		}
		c.room.forward <- msg
	}
}

// Receive messages from send channel and Write by using WriteMessage in WebSocket
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg);
		if  err!=nil {
			return
		}
	}
}