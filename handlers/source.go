package handlers

import (
	"net/http"
	"net/netip"
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

	ip, err := netip.ParseAddr(sourceBind.Ip)
	if err != nil {
		return err
	}

	err = cc.Queries.CreateSource(
		c.Request().Context(),
		ip,
	)
	if err != nil {
		return err
	}

	return nil
}

func GetSources(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	sources, err := cc.Queries.GetSources(cc.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, sources)
}
