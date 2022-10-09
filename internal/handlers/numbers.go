package handlers

import (
	"context"
	"github.com/sirupsen/logrus"
	"phonebook/internal/domain"
	"phonebook/pkg/parsing"
)

type NumbersService interface {
	Numbers(ctx context.Context, operation domain.Operation) (any, error)
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
		operation := parsing.ParsQuery()

		if operation == "" {
			continue
		} else if operation == "exit" {
			break
		}

		n.service.Numbers(context.Background(), operation)
	}
	return nil
}
