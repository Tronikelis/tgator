package handlers

import (
	"fmt"
	"net/http"
	"tgator/binds"
	"tgator/db"
	"tgator/dtos"
	"tgator/middleware"
	"tgator/models"

	"github.com/doug-martin/goqu/v9"
	"github.com/labstack/echo/v4"
)

func CreateSource(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	bind := binds.CreateSourceBind{}
	if err := c.Bind(&bind); err != nil {
		return err
	}

	if bind.Ip == "" {
		return echo.ErrBadRequest
	}

	query, params, err := cc.DB.PG.
		Insert("sources").
		Rows(
			models.SourceModel{
				Ip: bind.Ip,
			},
		).
		Returning("*").
		ToSQL()

	if err != nil {
		return err
	}

	source, err := db.QueryOne[models.SourceModel](cc.DB, cc.ReqCtx(), query, params...)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, source)
}

func GetSources(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	query, params, err := cc.DB.PG.From("sources").ToSQL()
	if err != nil {
		return err
	}

	sources, err := db.QueryMany[models.SourceModel](cc.DB, cc.ReqCtx(), query, params...)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, sources)
}

func GetSourceMessages(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	bind := binds.GetSourceMessagesBind{}
	if err := c.Bind(&bind); err != nil {
		return err
	}

	if bind.Id == 0 {
		return echo.ErrBadRequest
	}

	paginationDto := dtos.PaginationDTO[models.MessageModel]{}.FromBind(bind.PaginationBind)

	query, params, err := cc.DB.PG.
		From("messages").
		Where(
			goqu.C("source_id").Eq(bind.Id),
		).
		Limit(paginationDto.Limit()).
		Offset(paginationDto.Offset()).
		Order(goqu.C("id").Desc()).
		ToSQL()

	messages, err := db.QueryMany[models.MessageModel](cc.DB, cc.ReqCtx(), query, params...)
	if err != nil {
		return err
	}

	paginationDto.SetData(messages)

	return c.JSON(http.StatusOK, paginationDto)
}
