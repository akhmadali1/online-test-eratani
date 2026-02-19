package helper

import (
	"time"

	"test-3/shared/dto"

	"github.com/gofiber/fiber/v2"
)

const (
	UnknownRequestID = "unknown"
	// ... other constants.
)

func ErrorResponse[T any](ctx *fiber.Ctx, status int, apiErr *dto.APIError) error {
	reqID, ok := ctx.Locals("request_id").(string)
	if !ok {
		reqID = UnknownRequestID
	}

	return ctx.Status(status).JSON(dto.APIResponse[T]{
		Success: false,
		Error:   apiErr,
		Meta: &dto.Meta{
			Timestamp: time.Now().Unix(),
			RequestID: reqID,
		},
	})
}

func SuccessResponse[T any](ctx *fiber.Ctx, status int, data T) error {
	reqID, ok := ctx.Locals("request_id").(string)
	if !ok {
		reqID = UnknownRequestID
	}

	return ctx.Status(status).JSON(dto.APIResponse[T]{
		Success: true,
		Data:    data,
		Meta: &dto.Meta{
			Timestamp: time.Now().Unix(),
			RequestID: reqID,
		},
	})
}
