package repository

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/lucas-code42/rinha-backend/internal/domain"
)

type PersonRepository struct {
	sqlClient *sql.DB
}

func New(sqlClient *sql.DB) *PersonRepository {
	return &PersonRepository{sqlClient: sqlClient}
}

// TODO: pegar o último id inserido, precisa retornar ele no header http com location
func (p *PersonRepository) CreatePerson(person *domain.PessoaDto) error {
	stmt, err := p.sqlClient.Prepare("INSERT INTO pessoa(id, apelido, nome, nascimento, stack) values (?, ?, ?, ? , ?)")
	if err != nil {
		slog.Error("error preapare statement", err)
		return err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			slog.Error("error statement close", err)
		}
	}()

	response, err := stmt.Exec(
		uuid.NewString(),
		person.Apelido,
		person.Nome,
		person.Nascimento,
		person.Stack,
	)
	if err != nil {
		slog.Error("error statement execute", err)
		return err
	}

	lastID, _ := response.LastInsertId()
	slog.Debug(fmt.Sprintf("%d", lastID))

	n, err := response.RowsAffected()
	if n != 1 || err != nil {
		slog.Error("error rows affected", err)
		return err
	}

	return nil
}
