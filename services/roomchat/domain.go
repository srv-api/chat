package roomchat

import (
	"srv-api/chat/dto"
)

type ChatService interface {
	ProcessMessage(msg []byte) (*dto.ChatMessage, error)
	SaveMessage(data dto.ChatMessage) error
}
