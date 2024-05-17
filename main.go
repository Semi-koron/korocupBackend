package main

import (
	"github.com/semikoron/korocupbackend/crud"
	"github.com/semikoron/korocupbackend/database"
	"github.com/semikoron/korocupbackend/services"
	"github.com/semikoron/korocupbackend/utils/config"

	"github.com/labstack/echo/v4"
)

func main() {
	// Open a PostgreSQL database.
	config.LoadEnv()
	database.Connect()
	defer database.DB.Close()
	crud.Test()
	e := echo.New()
	e.GET("/", services.Hello)
	e.Logger.Fatal(e.Start(":8080"))
}
