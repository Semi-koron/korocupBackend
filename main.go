package main

import (
	"context"
	"log"
	"net/http"

	"github.com/semikoron/korocupbackend/database"
	"github.com/semikoron/korocupbackend/schema"
	"github.com/semikoron/korocupbackend/utils/config"

	"github.com/labstack/echo/v4"
)

func main() {
	// Open a PostgreSQL database.
	config.LoadEnv()
	database.Connect()
	if _, err := database.DB.NewCreateTable().Model((*schema.User)(nil)).Exec(context.TODO()); err != nil {
		log.Fatalln(err)
	}
	user := &schema.User{Name: "鈴木太郎"}
	if _, err := database.DB.NewInsert().Model(user).Exec(context.TODO()); err != nil {
		log.Fatalln(err)
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "change files")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
