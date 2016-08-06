package main

import (
	"github.com/gorilla/websocket"
	"time"
)

type client struct {
	socket *websocket.Conn
	send chan *message // channel where message is sent
	room *room
	userData map[string]interface{}
}

// Client read message by using ReadMessage in WebSocket
func (c *client) read() {
	defer c.socket.Close() // Closing WebSocket
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			msg.Name = c.userData["name"].(string)
			msg.AvatarURL, _ = c.room.avatar.GetAvatarURL(c)
			if avatarURL, ok := c.userData["avatar_url"]; ok {
				msg.AvatarURL = avatarURL.(string)
			}
			c.room.forward <- msg
		} else {
			break
		}
		//_, msg, err := c.socket.ReadMessage()
		//if  err != nil {
		//	return
		//}
		//c.room.forward <- msg
	}
}

// Receive messages from send channel and Write by using WriteMessage in WebSocket
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err!=nil {
			return
		}
	}
}