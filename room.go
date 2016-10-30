package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	//channel that olds message pending to transfer
	fwd_chan chan []byte

	// Channel for clients waiting to join
	join_chan chan *client

	// Channel for clients waiting or leaving
	leave_chan chan *client

	clients map[*client]bool
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket:    socket,
		send_chan: make(chan []byte, messageBufferSize),
		room:      r,
	}

	r.join_chan <- client
	defer func() { r.leave_chan <- client }()

	go client.write()
	client.read()
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join_chan:
			// Join
			r.clients[client] = true
		case client := <-r.leave_chan:
			// Leave
			delete(r.clients, client)
			close(client.send_chan)
		case msg := <-r.fwd_chan:
			// Forwad a message to all clients within the room
			for client := range r.clients {
				select {
				case client.send_chan <- msg:
					// By sending message to channel, client.write() will be executed.
				default:
					delete(r.clients, client)
					close(client.send_chan)
				}
			}
		}
	}
}
