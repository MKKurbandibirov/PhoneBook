package services

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"phonebook/internal/domain"
)

type numbersRepo interface {
	Numbers(ctx context.Context) ([]domain.Number, error)
}

type NumbersService struct {
	repo   numbersRepo
	logger log.FieldLogger
}

func NewNumberService(repo numbersRepo, logger log.FieldLogger) *NumbersService {
	return &NumbersService{
		repo:   repo,
		logger: logger,
	}
}

func (s *NumbersService) Numbers(ctx context.Context, filter domain.NumbersFilter,
	operation domain.Operation) (any, error) {
	switch operation {
	case domain.Select:
		numbers, err := s.repo.Numbers(ctx)
		if err != nil {
			return nil, err
		}

		if len(numbers) > filter.Limit && filter.Limit > 0 {
			numbers = numbers[:filter.Limit]
		}

		return numbers, nil
	case domain.Insert:
		return nil, nil
	case domain.Update:
		return nil, nil
	case domain.Delete:
		return nil, nil
	default:
		return nil, errors.New("service invalid operation")
	}
}
