// Package postgres implements postgres connection.
package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	_defaultMaxPoolSize      = 1
	_defaultConnAttempts     = 10
	_defaultConnTimeout      = time.Second
	_defaultTimezone         = "UTC"
	_defaultVerificationTime = 5 * time.Second
)

// Postgres -.
type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration
	timeZone     string

	Builder squirrel.StatementBuilderType
	Pool    *pgxpool.Pool
}

// New -.
func New(url string, opts ...Option) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:  _defaultMaxPoolSize,
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
		timeZone:     _defaultTimezone,
	}

	// Custom options
	for _, opt := range opts {
		opt(pg)
	}

	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - pgxpool.ParseConfig: %w", err)
	}

	poolConfig.MaxConns = int32(pg.maxPoolSize) //nolint:gosec // skip integer overflow conversion int -> int32

	// Set timezone for the connection
	poolConfig.ConnConfig.RuntimeParams["timezone"] = pg.timeZone

	for pg.connAttempts > 0 {
		pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err != nil {
			log.Printf("Postgres is trying to connect, attempts left: %d", pg.connAttempts)
			time.Sleep(pg.connTimeout)
			pg.connAttempts--

			continue
		}

		// Verify timezone is set correctly.
		verifyCtx, verifyCancel := context.WithTimeout(context.Background(), _defaultVerificationTime)

		var actualTimezone string

		verifyErr := pg.Pool.QueryRow(verifyCtx, "SHOW timezone").Scan(&actualTimezone)

		verifyCancel()

		if verifyErr != nil {
			log.Printf("Warning: Could not verify timezone: %v", verifyErr)
		} else if actualTimezone != pg.timeZone {
			log.Printf("Info: Database timezone is %s (configured: %s)", actualTimezone, pg.timeZone)
		}

		break
	}

	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
	}

	return pg, nil
}

// Close -.
func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
