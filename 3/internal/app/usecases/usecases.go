package usecases

import (
	"test-3/config"
	persistentUser "test-3/internal/user/repo/persistent"
	"test-3/internal/user/usecase/user"
	"test-3/pkg/postgres"
)

// UseCases holds all the use cases for DI

type UseCases struct {
	User *user.UseCase
	// Add other services here
}

// New -.
func New(pg *postgres.Postgres, cfg *config.Config) *UseCases {
	return &UseCases{
		User: user.New(
			persistentUser.New(pg),
			cfg,
		),
		// Add other services here
	}
}
