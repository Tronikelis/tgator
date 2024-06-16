package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"tgator/binds"
	"tgator/db/sqlc"
	"tgator/dtos"
	"tgator/middleware"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func CreateMessage(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	ctx := c.Request().Context()

	remoteAddr := c.Request().RemoteAddr
	remoteAddr = strings.Split(remoteAddr, ":")[0]

	source, err := cc.Queries.GetSourceByIp(ctx, remoteAddr)
	if err == pgx.ErrNoRows {
		source, err = cc.Queries.CreateSource(ctx, remoteAddr)
	}
	if err != nil {
		return err
	}

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	defer c.Request().Body.Close()

	bodyStr := string(body)

	if bodyStr == "" {
		return echo.ErrBadRequest
	}

	params := sqlc.CreateMessageParams{
		Raw:      pgtype.Text{String: bodyStr, Valid: true},
		SourceID: source.ID,
	}

	if json.Valid(body) {
		params.RawJsonb = body
	}

	message, err := cc.Queries.CreateMessage(
		ctx,
		params,
	)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, message)
}

func GetMessages(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	paginationBind := binds.PaginationBind{}
	if err := c.Bind(&paginationBind); err != nil {
		return err
	}

	paginationDto := dtos.PaginationDTO[sqlc.GetMessagesDescRow]{
		Limit:  paginationBind.Limit(),
		Offset: paginationBind.Offset(),
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
