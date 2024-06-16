package handlers

import (
	"net/http"
	"tgator/binds"
	"tgator/middleware"

	"github.com/labstack/echo/v4"
)

func CreateSource(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	sourceBind := binds.SourceBind{}
	if err := c.Bind(&sourceBind); err != nil {
		return err
	}

	source, err := cc.Queries.CreateSource(
		c.Request().Context(),
		sourceBind.Ip,
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
