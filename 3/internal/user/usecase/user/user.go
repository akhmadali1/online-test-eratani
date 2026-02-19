package user

import (
	"context"
	"fmt"

	"test-3/config"
	"test-3/internal/user/entity"
	userRepo "test-3/internal/user/repo"
	sharedDto "test-3/shared/dto"
)

// UseCase -.
type UseCase struct {
	repo userRepo.UserRepo
	cfg  *config.Config
}

// New -.
func New(r userRepo.UserRepo, cfg *config.Config) *UseCase {
	return &UseCase{
		repo: r,
		cfg:  cfg,
	}
}

func (uc *UseCase) GetAllData(ctx context.Context) ([]entity.User, *sharedDto.APIError) {
	datas, err := uc.repo.GetAllData(ctx)
	if err != nil {
		return nil, &sharedDto.APIError{
			Code:         500,
			Message:      "Internal server error",
			DebugMessage: fmt.Errorf("UserUseCase - GetAllData - repo.GetAllData: %v", err),
		}
	}

	return datas, nil
}

func (uc *UseCase) GetMostSpendCountryData(ctx context.Context) ([]entity.CountrySpend, *sharedDto.APIError) {
	datas, err := uc.repo.GetMostSpendCountryData(ctx)
	if err != nil {
		return nil, &sharedDto.APIError{
			Code:         500,
			Message:      "Internal server error",
			DebugMessage: fmt.Errorf("UserUseCase - GetMostSpendCountryData - repo.GetMostSpendCountryData: %v", err),
		}
	}

	return datas, nil
}

func (uc *UseCase) GetMostCreditCardTypeData(ctx context.Context) ([]entity.CreditCardTypeCount, *sharedDto.APIError) {
	datas, err := uc.repo.GetMostCreditCardTypeData(ctx)
	if err != nil {
		return nil, &sharedDto.APIError{
			Code:         500,
			Message:      "Internal server error",
			DebugMessage: fmt.Errorf("UserUseCase - GetMostCreditCardTypeData - repo.GetMostCreditCardTypeData: %v", err),
		}
	}

	return datas, nil
}

func (uc *UseCase) CreateData(ctx context.Context, req entity.User) (*entity.User, *sharedDto.APIError) {
	data, err := uc.repo.CreateData(ctx, req)
	if err != nil {
		return nil, &sharedDto.APIError{
			Code:         500,
			Message:      "Internal server error",
			DebugMessage: fmt.Errorf("UserUseCase - CreateData - repo.CreateData: %v", err),
		}
	}

	return data, nil
}
