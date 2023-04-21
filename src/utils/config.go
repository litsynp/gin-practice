package utils

import (
	"os"
)

const PORT = "8080"

var DB_URL = getEnv("DB_URL", "postgres://gostgres:go_pw@localhost:5436/gin_practice?sslmode=disable")

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
