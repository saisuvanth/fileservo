package app

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/saisuvanth/fileservo/middlewares"
)

func SetApp() error {

	err := LoadEnv()
	if err != nil {
		return err
	}

	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.HTTPErrorHandler = middlewares.AppErrorHandler()

	err = app.Start(os.Getenv("PORT"))
	if err != nil {
		return err
	}
	return nil
}
