// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package sqlc

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (id, name, email, password, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6)
`

type CreateUserParams struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT u.id, u.name, u.email FROM users u WHERE u.email = $1
`

type FindUserByEmailRow struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (FindUserByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, findUserByEmail, email)
	var i FindUserByEmailRow
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}
