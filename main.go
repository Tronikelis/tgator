package main

import (
	"os"
	"tgator/db"
	"tgator/db/sqlc"
	"tgator/middleware"
	"tgator/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := db.New(os.Getenv("PG_URL"))
	if err != nil {
		panic(err)
	}

	if err := db.CreateSchema("./db/schema.sql"); err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(middleware.GetCustomContextMiddleware(sqlc.New(db.Pool)))

	routes.AddV1(e)

	e.Start("localhost:3000")
}
