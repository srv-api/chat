package repository

import "srv-api/chat/entity"

type ChatRepository interface {
	Save(chat *entity.Chat) error
	GetChatHistory(userID, receiverID string, limit, offset int) ([]entity.Chat, error)
}
