// routes/routes.go
package routes

import (
	h "srv-api/chat/handlers/roomchat"
	s "srv-api/chat/services/roomchat"
	"srv-api/chat/ws"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize WebSocket Hub
	hub := ws.NewHub()
	go hub.Run()

	// Initialize Service (tanpa repository)
	service := s.NewChatService()

	// Initialize Handler
	handler := h.NewRoomChatHandler(hub, service)

	// Routes - Hanya WebSocket
	e.GET("/ws", handler.HandleWebSocket)

	// Health check (opsional)
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"status":  "ok",
			"service": "chat-websocket",
		})
	})

	return e
}
