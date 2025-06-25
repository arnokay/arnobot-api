package api

import (
	"log/slog"
	"net/http"

	"github.com/arnokay/arnobot-shared/appctx"
	"github.com/arnokay/arnobot-shared/apperror"
	"github.com/arnokay/arnobot-shared/applog"
	"github.com/arnokay/arnobot-shared/data"
	sharedServices "github.com/arnokay/arnobot-shared/service"
	"github.com/labstack/echo/v4"

	"github.com/arnokay/arnobot-api/internal/api/middleware"
)

type PlatformController struct {
	platformModule *sharedServices.PlatformModuleIn
	middlewares    *middleware.Middlewares

	logger *slog.Logger
}

func NewPlatformController(
	platformModule *sharedServices.PlatformModuleIn,
	middlewares *middleware.Middlewares,
) *PlatformController {
	logger := applog.NewServiceLogger("platform-controller")

	return &PlatformController{
		platformModule: platformModule,
		middlewares:    middlewares,

		logger: logger,
	}
}

func (c *PlatformController) Routes(parentGroup *echo.Group) {
	g := parentGroup.Group("/platform", c.middlewares.AuthMiddlewares.UserSessionGuard)
	g.POST("/:platform/bot", c.StartBot)
	g.DELETE("/:platform/bot", c.StopBot)
}

func (c *PlatformController) StartBot(ctx echo.Context) error {
	var payload data.PlatformToggleBot

	err := ctx.Bind(&payload)
	if err != nil {
		c.logger.DebugContext(ctx.Request().Context(), "cannot bind body", "err", err)
		return apperror.ErrInvalidInput
	}

	err = ctx.Validate(payload)
	if err != nil {
		c.logger.DebugContext(ctx.Request().Context(), "failed validation", "err", err)
		return err
	}

	user := appctx.GetUser(ctx.Request().Context())
	payload.UserID = user.ID

	err = c.platformModule.StartBot(ctx.Request().Context(), payload)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *PlatformController) StopBot(ctx echo.Context) error {
	var payload data.PlatformToggleBot

	err := ctx.Bind(&payload)
	if err != nil {
		c.logger.DebugContext(ctx.Request().Context(), "cannot bind body", "err", err)
		return apperror.ErrInvalidInput
	}

	err = ctx.Validate(payload)
	if err != nil {
		c.logger.DebugContext(ctx.Request().Context(), "failed validation", "err", err)
		return err
	}

	user := appctx.GetUser(ctx.Request().Context())
	payload.UserID = user.ID

	err = c.platformModule.StopBot(ctx.Request().Context(), payload)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
