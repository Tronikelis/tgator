package main

import (
	"fmt"
	"os"
	"strconv"
	"tgator/db"
	"tgator/middleware"
	"tgator/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
)

func intEnv(key string) (int, error) {
	value := os.Getenv(key)
	return strconv.Atoi(value)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		// dont panic on err as .env file does not exist in prod
		fmt.Println(err)
	}

	db, err := db.New(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "dev"
	}

	port, err := intEnv("PORT")
	if err != nil {
		panic(err)
	}

	if err := db.Migrate(); err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(echo_middleware.Static("./client/dist/"))

	e.Use(middleware.GetCustomContextMiddleware(db))

	e.Use(echo_middleware.Logger())

	routes.AddV1(e)

	addr := ""
	if appEnv == "dev" {
		addr += "localhost"
	} else {
		addr += "0.0.0.0"
	}

	addr += ":" + fmt.Sprintf("%v", port)

	if err := e.Start(addr); err != nil {
		panic(err)
	}
}
