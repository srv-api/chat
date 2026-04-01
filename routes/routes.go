package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-api/chat/configs"
)

var (
	DB = configs.InitDB()
)

func New() *echo.Echo {

	e := echo.New()
	return e
}
