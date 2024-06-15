package handlers

import (
	"io"
	"tgator/middleware"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func CreateMessage(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	defer c.Request().Body.Close()

	bodyStr := string(body)

	if bodyStr == "" {
		return echo.ErrBadRequest
	}

	if err := cc.Queries.CreateMessage(
		cc.Request().Context(),
		pgtype.Text{String: bodyStr, Valid: true},
	); err != nil {
		return echo.ErrBadGateway
	}

	return nil
}

func GetMessages(c echo.Context) error {
	return nil
}
