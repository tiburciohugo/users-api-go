package userhandler

import (
	"net/http"

	"github.com/tiburciohugo/users-api-go/internal/service/userservice"
)

func NewUserHandler(service userservice.UserService) UserHandler {
	return &handler{
		service,
	}
}

type handler struct {
	service userservice.UserService
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request) error
}
