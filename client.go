package main

import "golang.org/x/net/websocket"

type client struct {
	socket *websocket.Conn
	send   chan []byte
	room   *room
}
