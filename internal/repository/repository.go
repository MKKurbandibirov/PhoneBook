package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"phonebook/internal/domain"
	"phonebook/internal/repository/queries"
)

type repo struct {
	*queries.Queries
	pool   *pgxpool.Pool
	logger logrus.FieldLogger
}

func NewRepository(pool *pgxpool.Pool, logger logrus.FieldLogger) Repository {
	return &repo{
		Queries: queries.New(pool),
		pool:    pool,
		logger:  logger,
	}
}

type Repository interface {
	Numbers(ctx context.Context) ([]domain.Number, error)
	AddNumber(ctx context.Context) (int, error)
	ChangeNumber(ctx context.Context) error
	DeleteNumber(ctx context.Context) error
}
