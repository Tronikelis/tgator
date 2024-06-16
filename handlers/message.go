package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/netip"
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

	remoteAddrIp, err := netip.ParseAddr(remoteAddr)
	if err != nil {
		return err
	}

	source, err := cc.Queries.GetSourceByIp(ctx, remoteAddrIp)
	if err == pgx.ErrNoRows {
		source, err = cc.Queries.CreateSource(ctx, remoteAddrIp)
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
	if paginationBind.Limit == 0 {
		paginationBind.Limit = 50
	}

	paginationDto := dtos.PaginationDTO[sqlc.GetMessagesDescRow]{
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
