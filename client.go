package main

import "golang.org/x/net/websocket"

type client struct {
	// Socket for client
	socket *websocket.Conn
	// Channel where message is send
	send_chan chan []byte
	room      *room
}
