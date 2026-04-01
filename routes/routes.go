package routes

import (
	"srv-api/chat/configs"
	h_chat "srv-api/chat/handlers/roomchat"
	r_chat "srv-api/chat/repositories/roomchat"
	s_chat "srv-api/chat/services/roomchat"

	h_history "srv-api/chat/handlers/history"
	r_history "srv-api/chat/repositories/history"
	s_history "srv-api/chat/services/history"

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
	h := h_chat.NewRoomChatHandler(hub, service)

	historyR := r_history.NewHistoryRepository(DB)
	historyS := s_history.NewHistoryService(historyR)
	historyH := h_history.NewHistoryHandler(historyS)

	e.GET("/ws", h.HandleWebSocket)

	history := e.Group("/chat")
	{
		history.GET("/history", historyH.GetChatHistory)
	}

	return e
}
