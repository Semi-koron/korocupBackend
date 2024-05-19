package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PostgreSQLUser     string
	PostgreSQLPassword string
	PostgreSQLDBName   string
	PostgreSQLPort     string
	PostgreSQLConfig   string
)

// .envを呼び出します。
func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("読み込み出来ませんでした: %v", err)
	}
	PostgreSQLUser = os.Getenv("POSTGRESQLUSER")
	PostgreSQLPassword = os.Getenv("POSTGRESQLPASSWORD")
	PostgreSQLDBName = os.Getenv("POSTGRESQLDBNAME")
	PostgreSQLPort = os.Getenv("POSTGRESQLPORT")
	PostgreSQLConfig = os.Getenv("POSTGRESQLCONFIG")
}
