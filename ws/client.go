// ws/client.go
package ws

import (
	"encoding/json"
	"srv-api/chat/services/roomchat"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   int
	Conn *websocket.Conn
	Send chan []byte
}

func (c *Client) ReadPump(hub *Hub, service roomchat.ChatService) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		data, err := service.ProcessMessage(msg)
		if err != nil {
			continue
		}

		_ = service.SaveMessage(*data)

		out, _ := json.Marshal(data)
		hub.SendToUser(data.ReceiverID, out)
	}
}

func (c *Client) WritePump() {
	for msg := range c.Send {
		_ = c.Conn.WriteMessage(websocket.TextMessage, msg)
	}
}
