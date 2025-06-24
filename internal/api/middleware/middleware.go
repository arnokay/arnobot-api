package middleware

import (
	"log/slog"

	"github.com/arnokay/arnobot-shared/applog"
	"github.com/arnokay/arnobot-shared/middlewares"
)

type Middlewares struct {
	AuthMiddlewares *middlewares.AuthMiddlewares

	logger *slog.Logger
}

func New(
	authMiddlewares *middlewares.AuthMiddlewares,
) *Middlewares {
	logger := applog.NewServiceLogger("api-middleware")
	return &Middlewares{
		AuthMiddlewares: authMiddlewares,

		logger: logger,
	}
}
