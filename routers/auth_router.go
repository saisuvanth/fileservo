package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/saisuvanth/fileservo/handlers/auth"
)

func InitAuthRoutes(router *echo.Echo) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", auth.Register)
	}
}
