package userrepository

import (
	"database/sql"

	"github.com/tiburciohugo/users-api-go/internal/database/sqlc"
)

func NewUserRepository(db *sql.DB, q *sqlc.Queries) UserRepository {
	return &repository{
		db,
		q,
	}
}

type repository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

// CreateUser implements UserRepository.
func (r *repository) CreateUser() error {
	return nil
}

type UserRepository interface {
	CreateUser() error
}
