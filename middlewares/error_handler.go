package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AppErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		status_code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			status_code = he.Code
		}
		c.Logger().Error(err)
		c.JSON(status_code, map[string]interface{}{
			"message": err.Error(),
		})
	}
}
