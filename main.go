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
	database.ConnectDB()
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.GET("/", services.Hello)
	// User
	e.POST("/create/user", services.NewUser)                 // ユーザーを作成
	e.GET("/fetch/users", services.FetchUsers)               // すべてのユーザーを取得
	e.PUT("/update/user/:firebase_uid", services.UpdateUser) // ユーザーのプロフィールを更新
	// Post
	e.POST("/create/post", services.NewPost)
	e.Logger.Fatal(e.Start(":8080"))
}
