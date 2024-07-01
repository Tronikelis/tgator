package handlers

import (
	"errors"
	"io"
	"net/http"
	"slices"
	"strings"
	"tgator/binds"
	"tgator/builders"
	"tgator/db"
	"tgator/dtos"
	"tgator/middleware"
	"tgator/models"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func CreateSource(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	bind := binds.CreateSourceBind{}
	if err := c.Bind(&bind); err != nil {
		return err
	}

	if bind.Name == "" {
		return echo.ErrBadRequest
	}

	query, params, err := cc.DB.PG.
		Insert("sources").
		Rows(
			models.SourceModel{
				Name: bind.Name,
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

func GetSource(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	bind := binds.GetSourceBind{}
	if err := cc.Bind(&bind); err != nil {
		return err
	}

	query, params, err := cc.DB.PG.From("sources").Where(goqu.C("id").Eq(bind.Id)).ToSQL()

	source, err := db.QueryOne[models.SourceModel](cc.DB, cc.ReqCtx(), query, params...)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, source)
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

	paginationDto := new(dtos.PaginationDTO[models.MessageModel]).SetFromBind(bind.PaginationBind)

	builder := builders.NewGetSourceMessagesBuilder(cc.DB.PG, *paginationDto, bind.Id)

	if bind.Search != "" {
		builder.WhereSearch(bind.Search)
	}

	if bind.OrderBy != "" {
		builder.OrderBy(bind.OrderBy)
	}

	rowCount, err := db.CountRows(cc.DB, cc.ReqCtx(), builder.UnwrapSelectDataset())
	if err != nil {
		return err
	}

	paginationDto.SetPages(rowCount)

	query, params, err := builder.ToSQL()
	if err != nil {
		return err
	}

	messages, err := db.QueryMany[models.MessageModel](cc.DB, cc.ReqCtx(), query, params...)
	if err != nil {
		return err
	}

	paginationDto.Data = messages

	return c.JSON(http.StatusOK, paginationDto)
}

func CreateSourceMessage(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	bind := binds.CreateSourceMessageBind{}
	if err := (&echo.DefaultBinder{}).BindPathParams(c, &bind); err != nil {
		return err
	}

	query, params, err := cc.DB.PG.From("sources").Where(goqu.C("id").Eq(bind.Id)).ToSQL()
	if err != nil {
		return err
	}

	source, err := db.QueryOne[models.SourceModel](cc.DB, cc.ReqCtx(), query, params...)
	if errors.Is(err, pgx.ErrNoRows) {
		return echo.ErrBadRequest
	}
	if err != nil {
		return err
	}

	bodyBytes, err := io.ReadAll(cc.Request().Body)
	if err != nil {
		return err
	}

	bodyStr := string(bodyBytes)

	insertRows := []models.MessageModel{}
	for _, v := range strings.Split(bodyStr, "\n") {
		insertRows = append(insertRows, models.MessageModel{
			SourceId: source.ID,
			Raw:      v,
		})
	}

	insertRows = slices.DeleteFunc(insertRows, func(m models.MessageModel) bool {
		return strings.TrimSpace(m.Raw) == ""
	})

	query, params, err = cc.DB.PG.
		Insert("messages").
		Rows(insertRows).
		Returning("*").
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
