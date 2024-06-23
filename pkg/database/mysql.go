package database

import "database/sql"

type db struct {
	ClientDb *sql.DB
}

func New() *db {
	return &db{}
}
