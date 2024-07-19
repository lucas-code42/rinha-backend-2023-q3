package sql

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/lucas-code42/rinha-backend/internal/configs"
)

type sqlDb struct {
	SqlClient *sql.DB
}

func New() *sqlDb {
	addr := "db:3306"
	switch configs.Environment {
	case "test":
		addr = "localhost:3306"
	}

	cfg := mysql.Config{
		User:   os.Getenv("DBUser"),
		Passwd: os.Getenv("DBPassword"),
		Net:    "tcp",
		Addr:   addr,
		DBName: os.Getenv("DBName"),
	}

	client, err := sql.Open(os.Getenv("DBDriver"), cfg.FormatDSN())
	if err != nil {
		slog.Error("error connect with database", err.Error(), err)
		panic(err)
	}

	if err = client.Ping(); err != nil {
		slog.Error("error ping database", err.Error(), err)
		panic(err)
	}

	return &sqlDb{SqlClient: client}
}
