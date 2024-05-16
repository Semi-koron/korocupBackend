package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"websocket/utils/config"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID   int64 `bun:",pk,autoincrement"`
	Name string
}

func main() {
	// Open a PostgreSQL database.
	config.LoadEnv()
	dsn := "postgres://" + config.PostgreSQLUser + ":" + config.PostgreSQLPassword + "@" + config.PostgreSQLConfig
	fmt.Println(dsn)
	// dsn := "unix://user:pass@dbname/var/run/postgresql/.s.PGSQL.5432"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	defer sqldb.Close()
	// 使用するデータベースに合わせて、第二引数を変える。
	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))
	// クエリを標準出力するコードです。
	// 動作が分かりやすいため、入れておくことをオススメします。
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))
	if _, err := db.NewCreateTable().Model((*User)(nil)).Exec(context.TODO()); err != nil {
		log.Fatalln(err)
	}
	user := &User{Name: "鈴木太郎"}
	if _, err := db.NewInsert().Model(user).Exec(context.TODO()); err != nil {
		log.Fatalln(err)
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "change files")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
