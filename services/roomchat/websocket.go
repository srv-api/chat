// services/roomchat/service.go
package roomchat

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type chatService struct{}

func NewChatService() ChatService {
	return &chatService{}
}

func (s *chatService) ProcessMessage(msg []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal(msg, &data)
	if err != nil {
		return nil, err
	}

	// Set ID jika tidak ada
	if _, ok := data["id"]; !ok || data["id"] == "" {
		data["id"] = uuid.New().String()
	}

	// Set timestamp jika tidak ada
	if _, ok := data["created_at"]; !ok || data["created_at"] == "" {
		data["created_at"] = time.Now().Format(time.RFC3339)
	}

	// Set default type
	if _, ok := data["type"]; !ok || data["type"] == "" {
		data["type"] = "chat"
	}

	return data, nil
}
