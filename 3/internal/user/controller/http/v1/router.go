package v1

import (
	"test-3/config"
	"test-3/internal/user/usecase/user"
	"test-3/pkg/logger"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// NewUserRoutes -.
func NewUserRoutes(apiV1Group fiber.Router, cfg *config.Config, t *user.UseCase, l logger.Interface) {
	r := &V1{t: t, l: l, v: validator.New(validator.WithRequiredStructEnabled())}
	userGroup := apiV1Group.Group("/user")
	{
		userGroup.Get("", func(c *fiber.Ctx) error {
			return r.getAllData(c)
		})
		userGroup.Get("/a", func(c *fiber.Ctx) error {
			return r.getMostSpendCountryData(c)
		})
		userGroup.Get("/b", func(c *fiber.Ctx) error {
			return r.getMostCreditCardTypeData(c)
		})
		userGroup.Post("", func(c *fiber.Ctx) error {
			return r.createData(c)
		})
	}
}
