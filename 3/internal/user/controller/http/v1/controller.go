package v1

import (
	"test-3/internal/user/usecase"
	"test-3/pkg/logger"

	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	t usecase.User
	l logger.Interface
	v *validator.Validate
}
