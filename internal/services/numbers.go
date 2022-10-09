package services

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"phonebook/internal/domain"
)

type numbersRepo interface {
	Numbers(ctx context.Context) ([]domain.Number, error)
	AddNumber(ctx context.Context) (int, error)
	ChangeNumber(ctx context.Context) error
	DeleteNumber(ctx context.Context) error
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

func (s *NumbersService) Numbers(ctx context.Context,
	operation domain.Operation) (any, error) {
	switch operation {
	case domain.Select:
		numbers, err := s.repo.Numbers(ctx)
		if err != nil {
			return nil, err
		}

		fmt.Println(numbers)

		return numbers, nil
	case domain.Insert:
		id, err := s.repo.AddNumber(ctx)
		if err != nil {
			return nil, err
		}

		fmt.Println(id)

		return id, nil
	case domain.Update:
		err := s.repo.ChangeNumber(ctx)
		if err != nil {
			return nil, err
		}
		return nil, nil
	case domain.Delete:
		err := s.repo.DeleteNumber(ctx)
		if err != nil {
			return nil, err
		}
		return nil, nil
	default:
		return nil, errors.New("service invalid operation")
	}
}
