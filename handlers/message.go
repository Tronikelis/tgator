package handlers

import (
	"io"
	"net/http"
	"tgator/binds"
	"tgator/db/sqlc"
	"tgator/dtos"
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

	err = cc.Queries.CreateMessage(
		cc.Request().Context(),
		pgtype.Text{String: bodyStr, Valid: true},
	)

	if err != nil {
		return echo.ErrBadGateway
	}

	return nil
}

func GetMessages(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	paginationBind := binds.PaginationBind{}

	if err := c.Bind(&paginationBind); err != nil {
		return err
	}

	if paginationBind.Limit == 0 {
		paginationBind.Limit = 50
	}

	paginationDto := dtos.PaginationDTO[sqlc.Message]{
		Limit:  paginationBind.Limit,
		Offset: paginationBind.Limit * paginationBind.Page,
	}

	messages, err := cc.Queries.GetMessagesDesc(c.Request().Context(), sqlc.GetMessagesDescParams{
		Limit:  paginationDto.Limit,
		Offset: paginationDto.Offset,
	})

	if err != nil {
		return err
	}

	paginationDto.Data = messages

	return c.JSON(http.StatusOK, paginationDto)
}
