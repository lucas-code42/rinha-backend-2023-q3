package repository

import (
	"database/sql"
	"fmt"
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

func (p *PersonRepository) SearchPerson(searchTerm string) ([]*domain.PessoaDto, error) {
	stmt, err := p.sqlClient.Prepare(`
		SELECT * FROM pessoa WHERE apelido LIKE ? OR nome LIKE ? OR stack LIKE ?
	`)
	if err != nil {
		slog.Error("error query search person by term", err)
		return []*domain.PessoaDto{}, err
	}

	rows, err := stmt.Query("%"+searchTerm+"%", "%"+searchTerm+"%", "%"+searchTerm+"%")
	if err != nil {
		slog.Error("error to exec query search person by term", err)
		return []*domain.PessoaDto{}, err
	}

	var paginationPerson []*domain.PessoaDto
	for rows.Next() {
		var person domain.PessoaDto
		if err := rows.Scan(
			&person.Id,
			&person.Apelido,
			&person.Nome,
			&person.Nascimento,
			&person.Stack,
		); err != nil {
			slog.Error("error scan query rows", err)
			return []*domain.PessoaDto{}, err
		}

		paginationPerson = append(paginationPerson, &person)
	}

	return paginationPerson, nil
}

func (p *PersonRepository) Count() (int, error) {
	var total int
	if err := p.sqlClient.QueryRow("SELECT COUNT(id) FROM pessoa").Scan(&total); err != nil {
		slog.Error("error cannot count id from data base", err)
		return 0, err
	}

	if total <= 0 {
		e := fmt.Errorf("count db error")
		slog.Error("error data base has no data to count", e)
		return 0, e
	}

	return total, nil
}
