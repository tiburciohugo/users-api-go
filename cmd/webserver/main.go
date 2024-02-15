package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tiburciohugo/users-api-go/config/env"
	"github.com/tiburciohugo/users-api-go/config/logger"
	"github.com/tiburciohugo/users-api-go/internal/database"
	"github.com/tiburciohugo/users-api-go/internal/database/sqlc"
	"github.com/tiburciohugo/users-api-go/internal/handler/routes"
	"github.com/tiburciohugo/users-api-go/internal/handler/userhandler"
	"github.com/tiburciohugo/users-api-go/internal/repository/userrepository"
	"github.com/tiburciohugo/users-api-go/internal/service/userservice"
)

func main() {
	logger.InitLogger()
	slog.Info("Starting API...")

	_, err := env.LoadingConfig(".")
	if err != nil {
		slog.Error("Error loading environment variables", slog.String("package", "main"))
		return
	}

	dbConnection, err := database.NewDBConnection()
	if err != nil {
		slog.Error("Error establishing database connection", "err", err, slog.String("package", "main"))
		return
	}

	router := chi.NewRouter()
	queries := sqlc.New(dbConnection)

	userRepo := userrepository.NewUserRepository(dbConnection, queries)
	newUserService := userservice.NewUserService(userRepo)
	newUserHandler := userhandler.NewUserHandler(newUserService)

	routes.InitUserRoutes(router, newUserHandler)

	port := fmt.Sprintf(":%s", env.Env.GoPort)
	slog.Info(fmt.Sprintf("Server running on port %s", port), slog.String("package", "main"))
	err = http.ListenAndServe(port, router)
	if err != nil {
		slog.Error("Error starting server", slog.String("package", "main"))
	}

}
