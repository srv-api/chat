// ws/client.go
package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   int
	Conn *websocket.Conn
	Send chan []byte
}

func (c *Client) ReadPump(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}

		hub.Broadcast <- message
	}
}

func (c *Client) WritePump() {
	for msg := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}
