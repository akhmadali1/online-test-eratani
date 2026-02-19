package persistent

import (
	"context"
	"fmt"

	"test-3/internal/user/entity"
	"test-3/pkg/postgres"
)

// UserRepo -.
type UserRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (r *UserRepo) GetAllData(ctx context.Context) ([]entity.User, error) {
	sql, _, err := r.Builder.
		Select("id, country, credit_card_type, credit_card_number, first_name, last_name").
		From("users").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UserRepo - GetAllData - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("UserRepo - GetAllData - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	datas := make([]entity.User, 0)
	for rows.Next() {
		e := entity.User{}
		err = rows.Scan(&e.ID, &e.Country, &e.CreditCardType, &e.CreditCard, &e.FirstName, &e.LastName)
		if err != nil {
			return nil, fmt.Errorf("UserRepo - GetAllData - rows.Scan: %w", err)
		}

		datas = append(datas, e)
	}

	return datas, nil
}

func (r *UserRepo) GetMostSpendCountryData(ctx context.Context) ([]entity.CountrySpend, error) {
	sql, args, err := r.Builder.
		Select("u.country", "SUM(t.total_buy) AS total_spend").
		From("users u").
		Join("transactions t ON u.id = t.user_id").
		GroupBy("u.country").
		OrderBy("total_spend DESC").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UserRepo - GetMostSpendCountryData - Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("UserRepo - GetMostSpendCountryData - Query: %w", err)
	}
	defer rows.Close()

	var datas []entity.CountrySpend

	for rows.Next() {
		var e entity.CountrySpend
		if err := rows.Scan(&e.Country, &e.TotalSpend); err != nil {
			return nil, fmt.Errorf("UserRepo - GetMostSpendCountryData - Scan: %w", err)
		}
		datas = append(datas, e)
	}

	return datas, nil
}

func (r *UserRepo) GetMostCreditCardTypeData(ctx context.Context) ([]entity.CreditCardTypeCount, error) {
	sql, args, err := r.Builder.
		Select("credit_card_type", "COUNT(*) AS total").
		From("users").
		GroupBy("credit_card_type").
		OrderBy("total DESC").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UserRepo - GetMostCreditCardTypeData - Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("UserRepo - GetMostCreditCardTypeData - Query: %w", err)
	}
	defer rows.Close()

	var datas []entity.CreditCardTypeCount

	for rows.Next() {
		var e entity.CreditCardTypeCount
		if err := rows.Scan(&e.CreditCardType, &e.Total); err != nil {
			return nil, fmt.Errorf("UserRepo - GetMostCreditCardTypeData - Scan: %w", err)
		}
		datas = append(datas, e)
	}

	return datas, nil
}

func (r *UserRepo) CreateData(ctx context.Context, req entity.User) (*entity.User, error) {
	sql, args, err := r.Builder.
		Insert("users").
		Columns(
			"country",
			"credit_card_type",
			"credit_card_number",
			"first_name",
			"last_name",
		).
		Values(
			req.Country,
			req.CreditCardType,
			req.CreditCard,
			req.FirstName,
			req.LastName,
		).
		Suffix("RETURNING country, credit_card_type, credit_card_number, first_name, last_name").
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("UserRepo - CreateData - ToSql: %w", err)
	}

	var data entity.User
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&data.Country, &data.CreditCardType, &data.CreditCard, &data.FirstName, &data.LastName)
	if err != nil {
		return nil, fmt.Errorf("UserRepo - CreateData - QueryRow: %w", err)
	}

	return &data, nil
}
