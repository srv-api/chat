package roomchat

import (
	"net/http"
	"srv-api/chat/dto" // Alias

	// Alias
	"srv-api/chat/ws"

	res "github.com/srv-api/util/s/response"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

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
	go client.ReadPump(h.service, h.fcmService, h.fcmRepo)

	return nil
}

func (h *domainHandler) UpdateFCMToken(c echo.Context) error {
	var req dto.FCMTokenRequest

	useridToken, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if req.FCMToken == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "fcm_token is required",
		})
	}

	if req.DeviceType == "" {
		req.DeviceType = "android"
	}

	err := h.fcmRepo.SaveOrUpdateToken(useridToken, req.FCMToken, req.DeviceType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to save token",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "FCM token saved",
	})
}
