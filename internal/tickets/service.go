package tickets

import (
	"context"

	"fmt"
)

type Service interface {
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

func (s *ServiceImpl) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	tickets, err := s.repository.GetAll(ctx)
	cont := 0
	if err != nil {
		return 0, err
	}
	for _, ticket := range tickets {
		if ticket.Country == destination {
			cont++
		}
	}

	return cont, nil
}

func (s *ServiceImpl) AverageDestination(ctx context.Context, destination string) (float64, error) {
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	countPassangers := len(tickets)
	countPassangersByDestination, _ := s.GetTotalTickets(ctx, destination)

	if countPassangers == 0 {
		return 0.0, fmt.Errorf("Error: No existen tickets")
	}

	average := float64(countPassangersByDestination) / float64(countPassangers)
	return average, nil
}
