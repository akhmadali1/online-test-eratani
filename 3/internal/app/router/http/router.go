//nolint:godot,revive // Disable for swagger.
package httprouter

import (
	"test-3/config"
	"test-3/internal/app/router/http/middleware"
	"test-3/internal/app/usecases"
	userV1 "test-3/internal/user/controller/http/v1"
	"test-3/pkg/logger"
	"test-3/pkg/postgres"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, pg *postgres.Postgres, cfg *config.Config, l logger.Interface) {
	// Options
	app.Use(middleware.Logger(l))
	app.Use(middleware.Recovery(l))

	// Usecases
	allUsecases := usecases.New(pg, cfg)

	// Routers
	apiV1Group := app.Group("/v1")
	{
		userV1.NewUserRoutes(apiV1Group, cfg, allUsecases.User, l)
	}
}
