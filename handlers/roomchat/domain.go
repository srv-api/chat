package roomchat

import (
	repoNotif "srv-api/chat/repositories/notification" // Alias
	serNotif "srv-api/chat/services/notification"
	chatService "srv-api/chat/services/roomchat"
	"srv-api/chat/ws"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	HandleWebSocket(c echo.Context) error
	UpdateFCMToken(c echo.Context) error
}

type domainHandler struct {
	hub        *ws.Hub
	service    chatService.ChatService
	fcmService serNotif.FcmService     // ✅ langsung pakai notification
	fcmRepo    repoNotif.FCMRepository // ✅ pakai repoNotif
}

func NewRoomChatHandler(
	hub *ws.Hub,
	service chatService.ChatService,
	fcmService serNotif.FcmService,
	fcmRepo repoNotif.FCMRepository,
) DomainHandler {
	return &domainHandler{
		hub:        hub,
		service:    service,
		fcmService: fcmService,
		fcmRepo:    fcmRepo,
	}
}
