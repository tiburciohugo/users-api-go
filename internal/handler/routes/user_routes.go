package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tiburciohugo/users-api-go/internal/handler/userhandler"
)

func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
	router.Route("/user", func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			h.CreateUser(w, r)
		})
	})
}
