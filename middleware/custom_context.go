package middleware

import (
	"tgator/db/sqlc"

	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Queries *sqlc.Queries
}

func GetCustomContextMiddleware(queries *sqlc.Queries) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := CustomContext{c, queries}
			return next(&cc)
		}
	}
}
