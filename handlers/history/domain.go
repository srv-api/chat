package history

import (
	s "srv-api/chat/services/history"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	GetChatHistory(c echo.Context) error
}

type domainHandler struct {
	serviceHistory s.HistoryService
}

func NewHistoryHandler(service s.HistoryService) DomainHandler {
	return &domainHandler{
		serviceHistory: service,
	}
}
