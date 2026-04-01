package roomchat

import "github.com/labstack/echo/v4"

type DomainHandler interface {
	HandleWebSocket(c echo.Context) error
	GetChatHistory(c echo.Context) error
}
