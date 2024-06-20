package handlers

import (
	"database/sql"
	"io"
	"net/http"
	"strings"
	"tgator/binds"
	"tgator/db"
	"tgator/dtos"
	"tgator/middleware"
	"tgator/models"

	"github.com/doug-martin/goqu/v9"
	"github.com/labstack/echo/v4"
)

func CreateMessage(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	remoteAddr := c.Request().RemoteAddr
	remoteAddr = strings.Split(remoteAddr, ":")[0]

	query, params, err := cc.DB.PG.From("sources").Where(goqu.C("ip").Eq(remoteAddr)).ToSQL()
	if err != nil {
		return err
	}

	source, err := db.QueryOne[models.SourceModel](cc.DB, cc.ReqCtx(), query, params...)

	if err == sql.ErrNoRows {
		query, params, err = cc.DB.PG.
			Insert("sources").
			Rows(models.SourceModel{
				Ip: remoteAddr,
			}).
			Returning("*").
			ToSQL()

		if err != nil {
			return err
		}

		source, err = db.QueryOne[models.SourceModel](cc.DB, cc.ReqCtx(), query, params...)
		if err != nil {
			return err
		}
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

	query, params, err = cc.DB.PG.
		Insert("messages").
		Rows(
			models.MessageModel{
				SourceId: source.ID,
				Raw:      bodyStr,
			},
		).
		ToSQL()

	if err != nil {
		return err
	}

	message, err := db.QueryOne[models.MessageModel](cc.DB, cc.ReqCtx(), query, params...)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, message)
}

func GetMessages(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	bind := binds.PaginationBind{}
	if err := c.Bind(&bind); err != nil {
		return err
	}

	paginationDto := dtos.PaginationDTO[models.MessageModel]{}.FromBind(bind)

	query, params, err := cc.DB.PG.
		From("messages").
		Limit(paginationDto.Limit()).
		Offset(paginationDto.Offset()).
		Order(goqu.C("id").Desc()).
		ToSQL()

	if err != nil {
		return err
	}

	messages, err := db.QueryMany[models.MessageModel](cc.DB, cc.ReqCtx(), query, params...)
	if err != nil {
		return err
	}

	paginationDto.SetData(messages)

	return c.JSON(http.StatusOK, paginationDto)
}
