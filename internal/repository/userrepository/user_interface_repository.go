package userrepository

import "database/sql"

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
