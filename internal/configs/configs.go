package configs

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

var Environment string

func Init() {
	if err := godotenv.Load(); err != nil {
		slog.Error("error load environments variables", err.Error(), err)
		panic(err)
	}
	Environment = os.Getenv("ENVIRONMENT")
}
