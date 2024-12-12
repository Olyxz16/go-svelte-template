package handlers

import (
	"github.com/labstack/echo/v4"
)


func APIModeResponse(c echo.Context) error {
    return c.HTML(500, "Api mode is active. Static pages are not rendered.") 
}