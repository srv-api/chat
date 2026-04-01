package routes

import (
	"srv-api/chat/configs"
	h_chat "srv-api/chat/handlers/roomchat"
	r_chat "srv-api/chat/repositories/roomchat"
	s_chat "srv-api/chat/services/roomchat"

	"srv-api/chat/ws"

	"github.com/labstack/echo/v4"
)

var (
	DB = configs.InitDB()
)

func New() *echo.Echo {

	e := echo.New()

	hub := ws.NewHub()
	go hub.Run()

	repo := r_chat.NewChatRepository(DB)
	service := s_chat.NewChatService(repo)

	h := h_chat.NewRoomChatHandler(hub, service) // pakai variabel h

	e.GET("/ws", h.HandleWebSocket)
	e.GET("/chat/history", h.GetChatHistory)

	return e
}
