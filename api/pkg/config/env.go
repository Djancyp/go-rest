package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Env struct {
	Database *DB
}
type DB struct {
	MysqlDatabase     string
	Host              string
	MysqlPort         string
	MysqlUser         string
	MysqlPassword     string
	MysqlRootPassword string
}

func GetConfig(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
