package routes

import (
	"log"
	"os"
	"srv-api/chat/configs"
	"srv-api/chat/handlers/roomchat"
	repNotification "srv-api/chat/repositories/notification"
	"srv-api/chat/services/notification"
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

	// Database
	db := configs.InitDB() // ✅ pakai configs
	credFile := os.Getenv("CredFile")

	// FCM Service (optional - jika file tidak ada, jalankan tanpa FCM)
	fcmService, err := notification.NewFCMService(credFile)
	if err != nil {
		log.Println("⚠️ FCM Service not initialized:", err)
		log.Println("   Push notifications will not work")
		fcmService = nil
	} else {
		log.Println("✅ FCM Service initialized")
	}

	// Repository
	fcmRepo := repNotification.NewFCMRepository(db)

	// WebSocket Hub
	hub := ws.NewHub()
	go hub.Run()

	// Service
	chatService := s.NewChatService()

	// Handler
	handler := roomchat.NewRoomChatHandler(hub, chatService, fcmService, fcmRepo)

	// Routes
	e.GET("/ws", handler.HandleWebSocket)
	e.POST("/users/fcm-token", handler.UpdateFCMToken)

	// Health check
	e.GET("/health", func(c echo.Context) error {
		fcmStatus := false
		if fcmService != nil {
			fcmStatus = true
		}
		return c.JSON(200, map[string]interface{}{
			"status":  "ok",
			"service": "chat-websocket",
			"fcm":     fcmStatus,
		})
	})

	log.Println("🚀 Server running on :2369")
	log.Println("📡 WebSocket: ws://localhost:2369/ws?user_id=xxx")
	log.Println("📱 FCM Token: POST /users/fcm-token?user_id=xxx")

	return e
}
