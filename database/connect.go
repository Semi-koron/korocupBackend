package database

import (
	"database/sql"

	"github.com/semikoron/korocupbackend/utils/config"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var DB *bun.DB

func Connect() {
	dsn := "postgres://" + config.PostgreSQLUser + ":" + config.PostgreSQLPassword + "@" + config.PostgreSQLConfig
	Sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	DB = bun.NewDB(Sqldb, pgdialect.New())
	DB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))
}
