package roomchat

import (
	"net/http"
	"strconv"

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
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (h *domainHandler) HandleWebSocket(c echo.Context) error {
	userID, _ := strconv.Atoi(c.QueryParam("user_id"))

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	client := &ws.Client{
		ID:   strconv.Itoa(userID),
		Conn: conn,
		Send: make(chan []byte),
	}

	h.hub.Register <- client

	go client.WritePump()
	go client.ReadPump(h.hub, h.service)

	return nil
}
