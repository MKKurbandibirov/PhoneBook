package handlers

import (
	"context"
	"github.com/sirupsen/logrus"
	"phonebook/internal/domain"
	"phonebook/pkg/parsing"
)

type NumbersService interface {
	Numbers(ctx context.Context, filter domain.NumbersFilter, operation domain.Operation) (any, error)
}

type Numbers struct {
	logger  logrus.FieldLogger
	service NumbersService
}

func NewNumbers(logger logrus.FieldLogger, service NumbersService) *Numbers {
	return &Numbers{
		logger:  logger,
		service: service,
	}
}

func (n *Numbers) GetQuery() error {
	for {
		operation, err := parsing.ParsQuery()
		if err != nil {
			return err
		}

		if operation == "" {
			continue
		} else if operation == "exit" {
			break
		}

		filter, err := parsing.ParsFilter()
		if err != nil {
			return err
		}

		n.service.Numbers(context.Background(), *filter, operation)
	}
	return nil
}
