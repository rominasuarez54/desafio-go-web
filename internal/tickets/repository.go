package tickets

import (
	"context"
	"fmt"

	"desafio-go-web/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTotalTickets(ctx context.Context, destination string) (int, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type repository struct {
	db []domain.Ticket
}

func NewRepository(db []domain.Ticket) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Ticket, error) {

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	return r.db, nil
}

func (r *repository) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	cont := 0
	for _, ticket := range r.db {
		if ticket.Country == destination {
			cont++
		}
	}
	return cont, nil
}

func (r *repository) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {

	var ticketsDest []domain.Ticket

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}

func (r *repository) AverageDestination(ctx context.Context, destination string) (float64, error) {
	tickets, _ := r.GetTicketByDestination(ctx, destination)

	countPassangers := len(tickets)
	countPassangersByDestination, _ := r.GetTotalTickets(ctx, destination)

	if countPassangers == 0 {
		return 0.0, fmt.Errorf("Error: No existen tickets")
	}

	average := float64(countPassangersByDestination) / float64(countPassangers)
	return average, nil
}
