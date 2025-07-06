package middleware

import (
	

	"github.com/arnokay/arnobot-shared/applog"
	"github.com/arnokay/arnobot-shared/middlewares"
)

type Middlewares struct {
	AuthMiddlewares *middlewares.AuthMiddlewares

	logger applog.Logger
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
