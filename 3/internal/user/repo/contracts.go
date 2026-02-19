// Package repo implements application outer layer logic. Each logic group in own file.
package repo

import (
	"context"
	"test-3/internal/user/entity"
)

type (
	// UserRepo -.
	UserRepo interface {
		GetAllData(ctx context.Context) ([]entity.User, error)
		GetMostSpendCountryData(ctx context.Context) ([]entity.CountrySpend, error)
		GetMostCreditCardTypeData(ctx context.Context) ([]entity.CreditCardTypeCount, error)
		CreateData(ctx context.Context, req entity.User) (*entity.User, error)
	}
)
