package history

import (
	"srv-api/chat/entity"

	"gorm.io/gorm"
)

type HistoryRepository interface {
	GetChatHistory(userID, receiverID, limit, offset int) ([]entity.Chat, error)
}

type historyRepository struct {
	DB *gorm.DB
}

func NewHistoryRepository(DB *gorm.DB) HistoryRepository {
	return &historyRepository{
		DB: DB,
	}
}
