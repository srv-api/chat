package history

import (
	"srv-api/chat/entity"
	r "srv-api/chat/repositories/history"
)

type HistoryService interface {
	GetHistory(userID, receiverID, page, limit int) ([]entity.Chat, error)
}

type historyService struct {
	repo r.HistoryRepository
}

func NewHistoryService(repo r.HistoryRepository) HistoryService {
	return &historyService{repo: repo}
}
