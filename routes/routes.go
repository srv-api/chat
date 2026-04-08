// routes/routes.go
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
	"github.com/srv-api/middlewares/middlewares"
)

func New() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	JWT := middlewares.NewJWTService()

	// Database
	db := configs.InitDB()

	// Ambil credentials file dari environment variable
	credFile := os.Getenv("CredFile")

	// Cek apakah file credentials ada
	var fcmService *notification.FCMService
	if credFile != "" {
		// Cek apakah file ada
		if _, err := os.Stat(credFile); err == nil {
			service, err := notification.NewFCMService(credFile)
			if err != nil {
				log.Println("⚠️ FCM Service not initialized:", err)
			} else {
				fcmService = service
				log.Println("✅ FCM Service initialized")
			}
		} else {
			log.Println("⚠️ Firebase credentials file not found:", credFile)
		}
	} else {
		log.Println("⚠️ CredFile environment variable not set")
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

	usermerchant := e.Group("/users", middlewares.AuthorizeJWT(JWT))
	{
		usermerchant.POST("/fcm-token", handler.UpdateFCMToken)
	}

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"status":  "ok",
			"service": "chat-websocket",
			"fcm":     fcmService != nil,
		})
	})

	log.Println("🚀 Server running on :2369")
	log.Println("📡 WebSocket: ws://localhost:2369/ws?user_id=xxx")
	log.Println("📱 FCM Token: POST /users/fcm-token?user_id=xxx")

	return e
}
