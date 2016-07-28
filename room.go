package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"log"
)

type room struct {
	forward chan []byte // Que message to other clients
	join chan *client
	leave chan *client

	clients map[*client]bool
}

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join: make(chan *client),
		leave: make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (r *room) run() {
	for { // Endless Loop
		select {
		// When messages arrives at channel, corresponding case will be executed
		// They run concurrently, so change in r.clients do not happen simultaneously
		case client := <- r.join:
			r.clients[client] = true
		case client := <- r.leave:
			delete(r.clients, client)
			close(client.send) // difference between closing socket and channel?
		case msgInByte := <- r.forward:
			for client := range r.clients {
				select {
				case client.send <- msgInByte:
				default:
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}


const (
	socketBufferSize = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize: socketBufferSize,
	WriteBufferSize: messageBufferSize,
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil) // Retrieve WebSocket connection
	if err != nil {
		log.Fatal("ServerHTTP: ", err)
		return
	}

	client := &client{	// generate Client
		socket: socket,
		send: make(chan []byte, messageBufferSize),
		room: r,
	}

	r.join <- client
	defer func() {
		r.leave <- client
	}()
	go client.write()
	client.read()
}