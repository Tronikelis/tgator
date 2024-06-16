package handlers

import (
	"net/http"
	"tgator/binds"
	"tgator/db/sqlc"
	"tgator/dtos"
	"tgator/middleware"

	"github.com/labstack/echo/v4"
)

func CreateSource(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	bind := binds.CreateSourceBind{}
	if err := c.Bind(&bind); err != nil {
		return err
	}

	source, err := cc.Queries.CreateSource(
		c.Request().Context(),
		bind.Ip,
	)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, source)
}

func GetSources(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	sources, err := cc.Queries.GetSources(cc.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, sources)
}

func GetSourceMessages(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	ctx := c.Request().Context()

	bind := binds.GetSourceMessagesBind{}
	if err := c.Bind(&bind); err != nil {
		return err
	}

	if bind.Id == 0 {
		return echo.ErrBadRequest
	}

	paginationDto := dtos.PaginationDTO[sqlc.GetMessagesWhereSourceIdRow]{
		Offset: bind.Offset(),
		Limit:  bind.Limit(),
	}

	messages, err := cc.Queries.GetMessagesWhereSourceId(
		ctx,
		sqlc.GetMessagesWhereSourceIdParams{
			SourceID: bind.Id,
			Limit:    paginationDto.Limit,
			Offset:   paginationDto.Offset,
		},
	)
	if err != nil {
		return err
	}

	paginationDto.Data = messages

	return c.JSON(http.StatusOK, paginationDto)
}
