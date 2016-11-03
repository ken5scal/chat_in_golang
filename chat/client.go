package main

import (
	"github.com/gorilla/websocket"
	"time"
)

// client represents a single chatting user.
type client struct {
	// socket is the web socket for this client.
	socket   *websocket.Conn

	// send is a channel on which messages are sent.
	send     chan *message

	// room is the room this client is chatting in.
	room     *room

	userData map[string]interface{}
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			msg.Name = c.userData["name"].(string)
			if avatarURL, ok := c.userData["avatar_url"]; ok {
				msg.AvatarURL = avatarURL.(string)
			}
			c.room.forward <- msg
		} else {
			break
		}
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
}
