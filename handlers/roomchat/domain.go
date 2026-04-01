package handler

import (
	s "github.com/srv-api/chat/services/roomchat"
	"github.com/srv-api/chat/ws"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	HandleWebSocket(c echo.Context) error
}

type domainHandler struct {
	hub     *ws.Hub
	service s.ChatService
}

func NewRoomChatHandler(hub *ws.Hub, service s.ChatService) DomainHandler {
	return &domainHandler{
		hub:     hub,
		service: service,
	}
}
