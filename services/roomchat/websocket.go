package roomchat

import (
	"encoding/json"
	"srv-api/chat/dto"
	"srv-api/chat/entity"
	repository "srv-api/chat/repositories/roomchat"
)

type chatService struct {
	repo repository.ChatRepository
}

func NewChatService(repo repository.ChatRepository) ChatService {
	return &chatService{repo: repo}
}

func (s *chatService) ProcessMessage(msg []byte) (*dto.ChatMessage, error) {
	var data dto.ChatMessage
	err := json.Unmarshal(msg, &data)
	return &data, err
}

func (s *chatService) SaveMessage(data dto.ChatMessage) error {
	chat := entity.Chat{
		SenderID:   data.SenderID,
		ReceiverID: data.ReceiverID,
		Message:    data.Message,
	}

	return s.repo.Save(&chat)
}

func (s *chatService) GetHistory(userID, receiverID, page, limit int) ([]entity.Chat, error) {
	offset := (page - 1) * limit
	return s.repo.GetChatHistory(userID, receiverID, limit, offset)
}
