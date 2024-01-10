package main

import (
	"log/slog"

	"github.com/tiburciohugo/users-api-go/config/logger"
)

type user struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

func (u user) LogUser() slog.Value {
	return slog.GroupValue(
		slog.String("name", u.Name),
		slog.Int("age", u.Age),
		slog.String("password", "********"),
	)
}

func main() {
	logger.InitLogger()

	user := user{
		Name:     "Hugo",
		Age:      32,
		Password: "123456",
	}

	slog.Info("Starting API...")
	slog.Info("User", user.LogUser())
}
