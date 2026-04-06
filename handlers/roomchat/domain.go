package roomchat

import "github.com/labstack/echo/v4"

type DomainHandler interface {
	HandleWebSocket(c echo.Context) error
	UpdateFCMToken(c echo.Context) error
}
