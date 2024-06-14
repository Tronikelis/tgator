package handlers

import (
	"tgator/middleware"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func CreateMessage(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	cc.Queries.CreateMessage(
		cc.Request().Context(),
		pgtype.Text{String: "wat'supz", Valid: true},
	)

	return nil
}

func GetMessages(c echo.Context) error {
	return nil
}
