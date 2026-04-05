// websocket/client.go
package ws

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Send chan []byte
	Hub  *Hub
}

// ReadPump - menggunakan MessageProcessor interface
func (c *Client) ReadPump(processor MessageProcessor) {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}

		// Process message via processor
		msgMap, err := processor.ProcessMessage(message)
		if err != nil {
			log.Printf("Process error: %v", err)
			continue
		}

		// Ignore typing indicator
		if msgType, ok := msgMap["type"]; ok && msgType == "typing" {
			continue
		}

		// Get receiver ID
		receiverID, ok := msgMap["receiver_id"].(string)
		if !ok {
			log.Printf("No receiver_id in message")
			continue
		}

		log.Printf("📨 Message from %v to %s", msgMap["sender_id"], receiverID)

		// Cari receiver di hub
		c.Hub.mu.RLock()
		receiver, exists := c.Hub.Clients[receiverID]
		c.Hub.mu.RUnlock()

		if exists {
			// Forward message ke receiver
			data, _ := json.Marshal(msgMap)
			select {
			case receiver.Send <- data:
				log.Printf("✅ Message forwarded to %s", receiverID)
			default:
				log.Printf("⚠️ Receiver %s channel full", receiverID)
			}
		} else {
			log.Printf("⚠️ Receiver %s offline", receiverID)
			// Kirim error balik ke sender
			errorMsg := map[string]interface{}{
				"type":    "error",
				"message": "receiver offline",
				"id":      msgMap["id"],
			}
			errorData, _ := json.Marshal(errorMsg)
			c.Send <- errorData
		}
	}
}

// WritePump mengirim pesan ke client
func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Send {
		if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			break
		}
	}
}
