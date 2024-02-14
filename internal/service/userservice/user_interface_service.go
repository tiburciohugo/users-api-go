package userservice

import "github.com/tiburciohugo/users-api-go/internal/repository/userrepository"

func NewUserService(repo userrepository.UserRepository) UserService {
	return &service{
		repo,
	}
}

type service struct {
	repo userrepository.UserRepository
}

// CreateUser implements UserService.
func (s *service) CreateUser() error {
	return nil
}

type UserService interface {
	CreateUser() error
}
