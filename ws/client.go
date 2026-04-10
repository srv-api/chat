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

// Definisikan interface yang dibutuhkan di sini (agar tidak circular import)
type FCMServiceInterface interface {
	SendToDevice(userFCMToken string, data map[string]interface{}) error
}

type FCMRepositoryInterface interface {
	GetTokenByUserID(userID string) (string, error)
}

func (c *Client) ReadPump(processor MessageProcessor, fcmService FCMServiceInterface, fcmRepo FCMRepositoryInterface) {
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

		senderID, _ := msgMap["sender_id"].(string)
		messageText, _ := msgMap["message"].(string)
		senderName, _ := msgMap["sender_name"].(string)

		log.Printf("📨 Message from %s to %s: %s", senderID, receiverID, messageText)

		// Cari receiver di hub
		c.Hub.mu.RLock()
		receiver, exists := c.Hub.Clients[receiverID]
		c.Hub.mu.RUnlock()

		if exists {
			// Receiver ONLINE → kirim via WebSocket
			data, _ := json.Marshal(msgMap)
			select {
			case receiver.Send <- data:
				log.Printf("✅ Message forwarded via WebSocket to %s", receiverID)
			default:
				log.Printf("⚠️ Receiver %s channel full", receiverID)
			}
		} else {
			// Receiver OFFLINE → kirim via FCM
			log.Printf("⚠️ Receiver %s offline, sending FCM notification", receiverID)

			if fcmRepo == nil || fcmService == nil {
				log.Printf("FCM not available")
				errorMsg := map[string]interface{}{
					"type":    "error",
					"message": "receiver offline",
					"id":      msgMap["id"],
				}
				errorData, _ := json.Marshal(errorMsg)
				c.Send <- errorData
				continue
			}

			fcmToken, err := fcmRepo.GetTokenByUserID(receiverID)
			if err != nil || fcmToken == "" {
				log.Printf("No FCM token for user %s", receiverID)
				errorMsg := map[string]interface{}{
					"type":    "error",
					"message": "receiver offline",
					"id":      msgMap["id"],
				}
				errorData, _ := json.Marshal(errorMsg)
				c.Send <- errorData
				continue
			}

			// Kirim push notification
			err = fcmService.SendToDevice(fcmToken, map[string]interface{}{
				"title":       senderName,
				"body":        messageText,
				"receiver_id": receiverID,
				"sender_id":   senderID,
				"sender_name": senderName,
				"type":        "chat",
			})
			if err != nil {
				log.Printf("Failed to send FCM: %v", err)
			} else {
				log.Printf("✅ Push notification sent to %s", receiverID)
			}
		}
	}
}

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
