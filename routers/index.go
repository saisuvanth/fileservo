package routers

import "github.com/labstack/echo/v4"

func SetupRouter(router *echo.Echo) {
	InitAuthRoutes(router)
}
