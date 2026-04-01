package history

import (
	"srv-api/chat/entity"
	r "srv-api/chat/repositories/history"

	m "github.com/srv-api/middlewares/middlewares"
)

type HistoryService interface {
	GetHistory(userID, receiverID string, page, limit int) ([]entity.Chat, error)
}

type historyService struct {
	repo r.HistoryRepository
	jwt  m.JWTService
}

func NewHistoryService(repo r.HistoryRepository, jwtS m.JWTService) HistoryService {
	return &historyService{
		repo: repo,
		jwt:  jwtS,
	}
}
