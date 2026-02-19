package middleware

import (
	"strconv"
	"strings"

	"test-3/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func buildRequestMessage(ctx *fiber.Ctx) string {
	var result strings.Builder

	result.WriteString(ctx.IP())
	result.WriteString(" - ")
	result.WriteString(ctx.Method())
	result.WriteString(" ")
	result.WriteString(ctx.OriginalURL())
	result.WriteString(" - ")
	result.WriteString(strconv.Itoa(ctx.Response().StatusCode()))
	result.WriteString(" ")
	result.WriteString(strconv.Itoa(len(ctx.Response().Body())))

	return result.String()
}

func Logger(l logger.Interface) func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Next()

		l.Info(buildRequestMessage(ctx))

		reqID := ctx.Get("X-Request-ID")
		if reqID == "" {
			uuidV7, err := uuid.NewV7()
			if err != nil {
				return err
			}

			reqID = uuidV7.String()
		}

		ctx.Set("X-Request-ID", reqID)
		ctx.Locals("request_id", reqID)

		return err
	}
}
