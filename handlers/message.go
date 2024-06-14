package handlers

import (
	"tgator/binds"
	"tgator/middleware"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func CreateMessage(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	messageBind := binds.MessageBind{}
	if err := cc.Bind(&messageBind); err != nil {
		return echo.ErrBadRequest
	}

	if messageBind.Raw == "" {
		return echo.ErrBadRequest
	}

	cc.Queries.CreateMessage(
		cc.Request().Context(),
		pgtype.Text{String: messageBind.Raw, Valid: true},
	)

	return nil
}

func GetMessages(c echo.Context) error {
	return nil
}
