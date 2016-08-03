package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"github.com/ken5scal/chat/trace"
	"github.com/stretchr/objx"
)

type room struct {
	forward chan *message // Que message to other clients
	join chan *client	// channel for client who attempts to join the room
	leave chan *client // channel for client who attempts to leave the room
	clients map[*client]bool // all client in room
	tracer trace.Tracer // receives vent log on chat room
}

func newRoom() *room {
	return &room{
		forward: make(chan *message),
		join: make(chan *client),
		leave: make(chan *client),
		clients: make(map[*client]bool),
		tracer: trace.Off(),
	}
}

func (r *room) run() {
	for { // Endless Loop
		select {
		// When messages arrives at channel, corresponding case will be executed
		// They run concurrently, so change in r.clients do not happen simultaneously
		case client := <- r.join:
			r.clients[client] = true
			r.tracer.Trace("New Client joined")
		case client := <- r.leave:
			delete(r.clients, client)
			close(client.send) // difference between closing socket and channel?
			r.tracer.Trace("Client left")
		case msg := <- r.forward:
			r.tracer.Trace("Received message: ", msg.Message)
			for client := range r.clients {
				select {
				case client.send <- msg:
					r.tracer.Trace(" -- has been sent to client")
				default:
					delete(r.clients, client)
					close(client.send)
					r.tracer.Trace(" -- failed sending")
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

	authCookie, err := req.Cookie("auth")
	if err != nil {
		log.Fatal("Failed fetching cookie: ", err)
		return
	}

	client := &client{	// generate Client
		socket: socket,
		send: make(chan *message, messageBufferSize),
		room: r,
		userData: objx.MustFromBase64(authCookie.Value),
	}

	r.join <- client
	defer func() {
		r.leave <- client
	}()
	go client.write()
	client.read()
}