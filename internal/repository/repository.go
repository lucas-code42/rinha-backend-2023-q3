package repository

import (
	"database/sql"
	"log/slog"

	"github.com/lucas-code42/rinha-backend/internal/domain"
)

type PersonRepository struct {
	sqlClient *sql.DB
}

func New(sqlClient *sql.DB) *PersonRepository {
	return &PersonRepository{sqlClient: sqlClient}
}

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
		person.Id,
		person.Apelido,
		person.Nome,
		person.Nascimento,
		person.Stack,
	)
	if err != nil {
		slog.Error("error statement execute", err)
		return err
	}

	n, err := response.RowsAffected()
	if n != 1 || err != nil {
		slog.Error("error rows affected", err)
		return err
	}

	return nil
}

func (p *PersonRepository) GetPersonById(personId string) (*domain.PessoaDto, error) {
	stmt, err := p.sqlClient.Prepare("SELECT * FROM pessoa WHERE id=?")
	if err != nil {
		slog.Error("error query get person by id", err)
		return &domain.PessoaDto{}, err
	}

	var personDto domain.PessoaDto
	if err := stmt.QueryRow(personId).Scan(
		&personDto.Id,
		&personDto.Apelido,
		&personDto.Nome,
		&personDto.Nascimento,
		&personDto.Stack,
	); err != nil {
		slog.Error("error scan query row", err)
		return &domain.PessoaDto{}, err
	}

	return &personDto, nil
}
