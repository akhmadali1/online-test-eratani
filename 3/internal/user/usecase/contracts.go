// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"test-3/internal/user/entity"
	sharedDto "test-3/shared/dto"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_usecase_test.go -package=usecase_test

type (
	// User -.
	User interface {
		GetAllData(ctx context.Context) ([]entity.User, *sharedDto.APIError)
		GetMostSpendCountryData(ctx context.Context) ([]entity.CountrySpend, *sharedDto.APIError)
		GetMostCreditCardTypeData(ctx context.Context) ([]entity.CreditCardTypeCount, *sharedDto.APIError)
		CreateData(ctx context.Context, req entity.User) (*entity.User, *sharedDto.APIError)
	}
)
