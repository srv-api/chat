// handler/websocket_handler.go
package handler

import (
	"strconv"
	"your_project/ws"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) HandleWebSocket(c echo.Context) error {
	userIDStr := c.QueryParam("user_id")
	userID, _ := strconv.Atoi(userIDStr)

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	client := &ws.Client{
		ID:   userID,
		Conn: conn,
		Send: make(chan []byte),
	}

	h.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump(h.Hub)

	return nil
}
