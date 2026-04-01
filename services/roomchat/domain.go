package roomchat

import "your_project/ws"

type ChatService struct {
	Hub *ws.Hub
}

func NewChatService(hub *ws.Hub) *ChatService {
	return &ChatService{Hub: hub}
}
