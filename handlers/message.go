package handlers

import (
	"tgator/middleware"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type CreateMessageBind struct {
	Raw string `json:"raw"`
}

func CreateMessage(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	bind := CreateMessageBind{}
	if err := cc.Bind(&bind); err != nil {
		return echo.ErrBadRequest
	}

	if bind.Raw == "" {
		return echo.ErrBadRequest
	}

	cc.Queries.CreateMessage(
		cc.Request().Context(),
		pgtype.Text{String: bind.Raw, Valid: true},
	)

	return nil
}

func GetMessages(c echo.Context) error {
	return nil
}
