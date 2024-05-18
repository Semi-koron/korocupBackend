package main

import (
	"github.com/semikoron/korocupbackend/database"
	"github.com/semikoron/korocupbackend/services"
	"github.com/semikoron/korocupbackend/utils/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Open a PostgreSQL database.
	config.LoadEnv()
	database.Connect()
	defer database.DB.Close()
	database.InitDB()
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.GET("/", services.Hello)
	e.Logger.Fatal(e.Start(":8080"))
}
