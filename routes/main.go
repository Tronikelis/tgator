package routes

import (
	"github.com/labstack/echo/v4"
	"tgator/handlers"
)

func AddV1(e *echo.Echo) {
	group := e.Group("/api/v1")

	// message
	prefix := "/messages"
	group.POST(prefix, handlers.CreateMessage)
	group.GET(prefix, handlers.GetMessages)

	// source
	prefix = "/sources"
	group.POST(prefix, handlers.CreateSource)
	group.GET(prefix, handlers.GetSources)
	group.GET(prefix+"/:id", handlers.GetSource)
	group.GET(prefix+"/:id/messages", handlers.GetSourceMessages)
}
