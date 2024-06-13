package routes

import (
	"github.com/labstack/echo/v4"
	"tgator/handlers"
)

func AddV1(e *echo.Echo) {
	group := e.Group("/api/v1")

	// message
	prefix := "/message"
	group.POST(prefix+"/create", handlers.CreateMessage)
	group.GET(prefix+"/items", handlers.GetMessages)
}
