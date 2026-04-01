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
		ID:   userID,
		Conn: conn,
		Send: make(chan []byte),
	}

	h.hub.Register <- client

	go client.WritePump()
	go client.ReadPump(h.hub, h.service)

	return nil
}

func (h *domainHandler) GetChatHistory(c echo.Context) error {
	userID, _ := strconv.Atoi(c.QueryParam("user_id"))
	receiverID, _ := strconv.Atoi(c.QueryParam("receiver_id"))

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 20
	}

	data, err := h.service.GetHistory(userID, receiverID, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}
