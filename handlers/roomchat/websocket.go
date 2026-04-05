package roomchat

import (
	"net/http"

	"srv-api/chat/services/roomchat"
	"srv-api/chat/ws"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type domainHandler struct {
	hub     *ws.Hub
	service roomchat.ChatService
}

func NewRoomChatHandler(hub *ws.Hub, service roomchat.ChatService) DomainHandler {
	return &domainHandler{
		hub:     hub,
		service: service,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *domainHandler) HandleWebSocket(c echo.Context) error {
	userID := c.QueryParam("user_id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "user_id is required",
		})
	}

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to upgrade connection",
		})
	}

	client := &ws.Client{
		ID:   userID,
		Conn: conn,
		Send: make(chan []byte, 256),
		Hub:  h.hub,
	}

	h.hub.Register <- client

	go client.WritePump()
	go client.ReadPump(h.service) // ✅ Sekarang sesuai dengan signature

	return nil
}
