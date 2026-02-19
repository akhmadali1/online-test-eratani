package v1

import (
	"net/http"

	"test-3/internal/user/entity"
	sharedDto "test-3/shared/dto"
	sharedHelper "test-3/shared/helper"

	"github.com/gofiber/fiber/v2"
)

type (
	APIResponse[T any] sharedDto.APIResponse[T]
	APIRequest[T any]  sharedDto.APIRequest[T]
)

const (
	UnknownRequestID = "unknown"
	// ... other constants.
)

func (r *V1) getAllData(ctx *fiber.Ctx) error {
	reqID, ok := ctx.Locals("request_id").(string)
	if !ok {
		reqID = UnknownRequestID
	}

	datas, err := r.t.GetAllData(ctx.UserContext())
	if err != nil {
		r.l.Error(err.DebugMessage, "[req_id=%s] http - v1 - user - getAllData", reqID)
		return sharedHelper.ErrorResponse[any](ctx, err.Code, err)
	}

	return sharedHelper.SuccessResponse(ctx, http.StatusOK, datas)
}

func (r *V1) getMostSpendCountryData(ctx *fiber.Ctx) error {
	reqID, ok := ctx.Locals("request_id").(string)
	if !ok {
		reqID = UnknownRequestID
	}

	datas, err := r.t.GetMostSpendCountryData(ctx.UserContext())
	if err != nil {
		r.l.Error(err.DebugMessage, "[req_id=%s] http - v1 - user - getAllData", reqID)
		return sharedHelper.ErrorResponse[any](ctx, err.Code, err)
	}

	return sharedHelper.SuccessResponse(ctx, http.StatusOK, datas)
}

func (r *V1) getMostCreditCardTypeData(ctx *fiber.Ctx) error {
	reqID, ok := ctx.Locals("request_id").(string)
	if !ok {
		reqID = UnknownRequestID
	}

	datas, err := r.t.GetMostCreditCardTypeData(ctx.UserContext())
	if err != nil {
		r.l.Error(err.DebugMessage, "[req_id=%s] http - v1 - user - getAllData", reqID)
		return sharedHelper.ErrorResponse[any](ctx, err.Code, err)
	}

	return sharedHelper.SuccessResponse(ctx, http.StatusOK, datas)
}

func (r *V1) createData(ctx *fiber.Ctx) error {
	reqID, ok := ctx.Locals("request_id").(string)
	if !ok {
		reqID = UnknownRequestID
	}

	var body APIRequest[entity.User]
	if err := ctx.BodyParser(&body); err != nil {
		r.l.Error(err, "[req_id=%s] http - v1 - user - createData", reqID)
		return sharedHelper.ErrorResponse[any](ctx, 400, &sharedDto.APIError{
			Code:         400,
			DebugMessage: err,
			Message:      err.Error(),
		})
	}

	if err := r.v.Struct(body.Data); err != nil {
		r.l.Error(err, "[req_id=%s] http - v1 - user - createData", reqID)
		return sharedHelper.ErrorResponse[any](ctx, 400, &sharedDto.APIError{
			Code:         400,
			DebugMessage: err,
			Message:      err.Error(),
		})
	}

	data, err := r.t.CreateData(ctx.UserContext(), body.Data)
	if err != nil {
		r.l.Error(err.DebugMessage, "[req_id=%s] http - v1 - user - createData", reqID)
		return sharedHelper.ErrorResponse[any](ctx, err.Code, err)
	}

	return sharedHelper.SuccessResponse(ctx, http.StatusOK, data)
}
