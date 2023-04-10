package tickets

import (
	"context"
	"desafio-go-web/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTotalTickets(ctx context.Context, destination string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type ServiceImpl struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &ServiceImpl{
		repository: repository,
	}
}

func (s *ServiceImpl) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	return s.repository.GetAll(ctx)
}
func (s *ServiceImpl) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	tickets, err := s.repository.GetTotalTickets(ctx, destination)
	if err != nil {
		return 0, err
	}
	return tickets, nil
}

func (s *ServiceImpl) AverageDestination(ctx context.Context, destination string) (float64, error) {
	average, err := s.repository.AverageDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	return average, nil
}
