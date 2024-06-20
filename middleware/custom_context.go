package middleware

import (
	"context"
	"tgator/db"

	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	DB *db.DB
}

func GetCustomContextMiddleware(db *db.DB) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := CustomContext{c, db}
			return next(&cc)
		}
	}
}

// wrapper for c.Request().Context()
func (cc *CustomContext) ReqCtx() context.Context {
	return cc.Request().Context()
}
