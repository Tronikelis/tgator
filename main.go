package main

import (
	"tgator/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.AddV1(e)

	e.Start("localhost:3000")
}
