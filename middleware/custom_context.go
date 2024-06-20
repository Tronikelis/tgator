package middleware

import (
	"context"
	"tgator/db"
	"tgator/db/sqlc"

	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Queries *sqlc.Queries
	DB      *db.DB
}

func GetCustomContextMiddleware(queries *sqlc.Queries, db *db.DB) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := CustomContext{c, queries, db}
			return next(&cc)
		}
	}
}

// wrapper for c.Request().Context()
func (cc *CustomContext) ReqCtx() context.Context {
	return cc.Request().Context()
}
