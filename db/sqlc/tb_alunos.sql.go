// Code generated by sqlc. DO NOT EDIT.
// source: tb_alunos.sql

package db

import (
	"context"
	"database/sql"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO tb_alunos  (
    desnome,
    email,
    telefone,
    endereco,
    turma) VALUES ($1, $2, $3, $4, $5)
RETURNING idaluno, desnome, email, telefone, endereco, matricula, turma, dt_cadastro, modificado_em
`

type CreateAccountParams struct {
	Desnome  string         `json:"desnome"`
	Email    string         `json:"email"`
	Telefone int64          `json:"telefone"`
	Endereco string         `json:"endereco"`
	Turma    sql.NullString `json:"turma"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (TbAluno, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.Desnome,
		arg.Email,
		arg.Telefone,
		arg.Endereco,
		arg.Turma,
	)
	var i TbAluno
	err := row.Scan(
		&i.Idaluno,
		&i.Desnome,
		&i.Email,
		&i.Telefone,
		&i.Endereco,
		&i.Matricula,
		&i.Turma,
		&i.DtCadastro,
		&i.ModificadoEm,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM tb_alunos
WHERE idaluno = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, idaluno int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, idaluno)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT idaluno, desnome, email, telefone, endereco, matricula, turma, dt_cadastro, modificado_em FROM tb_alunos
WHERE idaluno = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, idaluno int64) (TbAluno, error) {
	row := q.db.QueryRowContext(ctx, getAccount, idaluno)
	var i TbAluno
	err := row.Scan(
		&i.Idaluno,
		&i.Desnome,
		&i.Email,
		&i.Telefone,
		&i.Endereco,
		&i.Matricula,
		&i.Turma,
		&i.DtCadastro,
		&i.ModificadoEm,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT idaluno, desnome, email, telefone, endereco, matricula, turma, dt_cadastro, modificado_em FROM tb_alunos
ORDER BY id
LIMIT $1
    OFFSET $2
`

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]TbAluno, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TbAluno
	for rows.Next() {
		var i TbAluno
		if err := rows.Scan(
			&i.Idaluno,
			&i.Desnome,
			&i.Email,
			&i.Telefone,
			&i.Endereco,
			&i.Matricula,
			&i.Turma,
			&i.DtCadastro,
			&i.ModificadoEm,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE tb_alunos
SET desnome = $2,
email = $3,
telefone = $4,
endereco = $5,
turma = $6

WHERE idaluno = $1
RETURNING idaluno, desnome, email, telefone, endereco, matricula, turma, dt_cadastro, modificado_em
`

type UpdateAccountParams struct {
	Idaluno  int64          `json:"idaluno"`
	Desnome  string         `json:"desnome"`
	Email    string         `json:"email"`
	Telefone int64          `json:"telefone"`
	Endereco string         `json:"endereco"`
	Turma    sql.NullString `json:"turma"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (TbAluno, error) {
	row := q.db.QueryRowContext(ctx, updateAccount,
		arg.Idaluno,
		arg.Desnome,
		arg.Email,
		arg.Telefone,
		arg.Endereco,
		arg.Turma,
	)
	var i TbAluno
	err := row.Scan(
		&i.Idaluno,
		&i.Desnome,
		&i.Email,
		&i.Telefone,
		&i.Endereco,
		&i.Matricula,
		&i.Turma,
		&i.DtCadastro,
		&i.ModificadoEm,
	)
	return i, err
}
