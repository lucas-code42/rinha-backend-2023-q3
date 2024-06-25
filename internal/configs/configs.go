package configs

import (
	"log/slog"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		slog.Error("error load environments variables", err)
		panic(err)
	}
}
