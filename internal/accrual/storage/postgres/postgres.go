package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gophermart/internal/accrual/config"
	"gophermart/internal/accrual/entity"

	"github.com/jackc/pgx/v5"
)

type postgresStorage struct {
	db *pgx.Conn
}

var (
	ErrOrderNotFound  = errors.New("order not found")
	ErrQueryExecution = errors.New("query execution failed")
	ErrConnectTimeout = errors.New("connect timeout")
)

// NewPostgresStorage create new data storage in PostgreSQL.
func NewPostgresStorage(cfg *config.PostgresConfig) (*postgresStorage, error) {
	connCtx, cancel := context.WithTimeoutCause(context.Background(), cfg.ConnectTimeout, ErrConnectTimeout)
	defer cancel()
	conn, err := pgx.Connect(connCtx, cfg.DatabaseDSN)
	if err != nil {
		return nil, fmt.Errorf("postgres connect: %w", err)
	}

	return &postgresStorage{
		db: conn,
	}, nil
}

func (s *postgresStorage) CreateOrder(ctx context.Context, o entity.Order) error {
	query := "INSERT INTO orders (number, status, accrual) VALUES ($1, $2, $3)"
	_, err := s.db.Exec(ctx, query, o.Number, o.Status, o.Accrual)
	if err != nil {
		return ErrQueryExecution
	}

	return nil
}

func (s *postgresStorage) GetOrderByID(ctx context.Context, id entity.ID) (entity.Order, error) {
	query := "SELECT number, status, accrual FROM orders WHERE id = $1"
	row := s.db.QueryRow(ctx, query, id)

	var o entity.Order
	err := row.Scan(&o.Number, &o.Status, &o.Accrual)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.Order{}, ErrOrderNotFound
	}
	if err != nil {
		return entity.Order{}, ErrQueryExecution
	}

	return o, nil
}

func (s *postgresStorage) Close() error {
	err := s.db.Close(context.Background())

	if err != nil {
		return fmt.Errorf("postgres close: %w", err)
	}

	return nil
}
