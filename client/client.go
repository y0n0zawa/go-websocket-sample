package client

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Socket        *websocket.Conn
	Send          chan []byte
	OnChatMessage func(msg []byte)
}

func (c *Client) Read() {
	for {
		if _, msg, err := c.Socket.ReadMessage(); err == nil {
			c.OnChatMessage(msg)
		} else {
			break
		}
	}
	c.Socket.Close()
}

func (c *Client) Write() {
	for msg := range c.Send {
		if err := c.Socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.Socket.Close()
}
