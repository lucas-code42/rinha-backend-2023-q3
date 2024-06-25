package sql

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/go-sql-driver/mysql"
)

type sqlDb struct {
	SqlClient *sql.DB
}

func New() *sqlDb {
	cfg := mysql.Config{
		User:   os.Getenv("DBUser"),
		Passwd: os.Getenv("DBPassword"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DBName"),
	}

	client, err := sql.Open(os.Getenv("DBDriver"), cfg.FormatDSN())
	if err != nil {
		slog.Error("error connect with database", err)
		panic(err)
	}

	if err = client.Ping(); err != nil {
		slog.Error("error ping database", err)
		panic(err)
	}

	return &sqlDb{SqlClient: client}
}
