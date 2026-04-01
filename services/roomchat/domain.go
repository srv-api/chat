package roomchat

import (
	"srv-api/chat/dto"
	"srv-api/chat/entity"
)

type ChatService interface {
	ProcessMessage(msg []byte) (*dto.ChatMessage, error)
	SaveMessage(data dto.ChatMessage) error
	GetHistory(userID, receiverID, page, limit int) ([]entity.Chat, error)
}
