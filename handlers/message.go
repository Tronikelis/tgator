package handlers

import (
	"net/http"
	"tgator/binds"
	"tgator/builders"
	"tgator/db"
	"tgator/dtos"
	"tgator/middleware"
	"tgator/models"

	"github.com/labstack/echo/v4"
)

func GetMessages(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	bind := binds.GetMessagesBind{}
	if err := c.Bind(&bind); err != nil {
		return err
	}

	paginationDto := new(dtos.PaginationDTO[models.MessageModel]).SetFromBind(bind.PaginationBind)

	builder := builders.NewGetMessagesBuilder(cc.DB.PG, *paginationDto)

	builder.JoinSources()

	if bind.Search != "" {
		builder.WhereSearch(bind.Search)
	}

	if bind.OrderBy != "" {
		builder.OrderBy(bind.OrderBy)
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
