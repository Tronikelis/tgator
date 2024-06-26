package handlers

import (
	"errors"
	"net/http"
	"tgator/binds"
	"tgator/db"
	"tgator/dtos"
	"tgator/middleware"
	"tgator/models"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func CreateMessage(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	bind := binds.CreateMessageBind{}
	if err := cc.Bind(&bind); err != nil {
		return err
	}

	if bind.Raw == "" || bind.SourceId == 0 {
		return echo.ErrBadRequest
	}

	query, params, err := cc.DB.PG.From("sources").Where(goqu.C("id").Eq(bind.SourceId)).ToSQL()
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

	query, params, err = cc.DB.PG.
		Insert("messages").
		Rows(
			models.MessageModel{
				SourceId: source.ID,
				Raw:      bind.Raw,
			},
		).
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

func GetMessages(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	bind := binds.GetMessagesBind{}
	if err := c.Bind(&bind); err != nil {
		return err
	}

	paginationDto := new(dtos.PaginationDTO[models.MessageModel]).SetFromBind(bind.PaginationBind)

	builder := cc.DB.PG.
		From("messages").
		Select(
			goqu.I("messages.*"),
			goqu.I("sources.*"),
		).
		LeftJoin(goqu.T("sources"), goqu.On(goqu.I("sources.id").Eq(goqu.I("messages.source_id")))).
		Limit(uint(paginationDto.Limit)).
		Offset(uint(paginationDto.Offset)).
		Order(goqu.I("messages.id").Desc())

	if bind.Search != "" {
		builder = builder.Where(goqu.C("raw").Like("%" + bind.Search + "%"))
	}

	if bind.OrderBy != "" {
		builder = builder.ClearOrder()

		var order exp.OrderedExpression

		switch bind.OrderBy {
		case "asc":
			order = goqu.C("id").Asc()
		default:
			order = goqu.C("id").Desc()
		}

		builder = builder.Order(order)
	}

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
