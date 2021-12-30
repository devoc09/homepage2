package http

import "github.com/labstack/echo/v4"

func indexHandler() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.File("public/index.html")
	}
}
