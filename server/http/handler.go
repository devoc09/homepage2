package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func indexHandler() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.HTML(http.StatusOK, "public/index.html")
	}
}
