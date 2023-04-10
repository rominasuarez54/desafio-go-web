package tickets

import (
	"context"
	"fmt"
	"testing"

	"desafio-go-web/internal/domain"

	"github.com/stretchr/testify/assert"
)

var cxt = context.Background()

var tickets = []domain.Ticket{
	{
		Id:      "1",
		Name:    "Tait Mc Caughan",
		Email:   "tmc0@scribd.com",
		Country: "Finland",
		Time:    "17:11",
		Price:   785.00,
	},
	{
		Id:      "2",
		Name:    "Padget McKee",
		Email:   "pmckee1@hexun.com",
		Country: "China",
		Time:    "20:19",
		Price:   537.00,
	},
	{
		Id:      "3",
		Name:    "Yalonda Jermyn",
		Email:   "yjermyn2@omniture.com",
		Country: "China",
		Time:    "18:11",
		Price:   579.00,
	},
}

var ticketsByDestination = []domain.Ticket{
	{
		Id:      "2",
		Name:    "Padget McKee",
		Email:   "pmckee1@hexun.com",
		Country: "China",
		Time:    "20:19",
		Price:   537.00,
	},
	{
		Id:      "3",
		Name:    "Yalonda Jermyn",
		Email:   "yjermyn2@omniture.com",
		Country: "China",
		Time:    "18:11",
		Price:   579.00,
	},
}

type stubRepo struct {
	db *DbMock
}

type DbMock struct {
	db  []domain.Ticket
	spy bool
	err error
}

func NewRepositoryTest(dbM *DbMock) Repository {
	return &stubRepo{dbM}
}

func (r *stubRepo) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	r.db.spy = true
	if r.db.err != nil {
		return []domain.Ticket{}, r.db.err
	}
	return tickets, nil
}

func (r *stubRepo) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	r.db.spy = true
	cont := 0
	for _, ticket := range r.db.db {
		if ticket.Country == destination {
			cont++
		}
	}
	return cont, nil
}

func (r *stubRepo) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {

	var tkts []domain.Ticket

	r.db.spy = true
	if r.db.err != nil {
		return []domain.Ticket{}, r.db.err
	}

	for _, t := range r.db.db {
		if t.Country == destination {
			tkts = append(tkts, t)
		}
	}

	return tkts, nil
}
func (r *stubRepo) AverageDestination(ctx context.Context, destination string) (float64, error) {
	r.db.spy = true
	tickets, _ := r.GetTicketByDestination(ctx, destination)

	countPassangers := len(tickets)
	countPassangersByDestination, _ := r.GetTotalTickets(ctx, destination)

	if countPassangers == 0 {
		return 0.0, fmt.Errorf("Error: No existen tickets")
	}

	average := float64(countPassangersByDestination) / float64(countPassangers)
	return average, nil
}

func TestGetTicketByDestination(t *testing.T) {

	dbMock := &DbMock{
		db:  tickets,
		spy: false,
		err: nil,
	}
	repo := NewRepositoryTest(dbMock)
	service := NewService(repo)

	tkts, err := service.GetTotalTickets(cxt, "China")

	assert.Nil(t, err)
	assert.True(t, dbMock.spy)
	assert.Equal(t, len(ticketsByDestination), tkts)
}

func TestGetTotalTickets(t *testing.T) {

	dbMock := &DbMock{
		db:  tickets,
		spy: false,
		err: nil,
	}
	repo := NewRepositoryTest(dbMock)
	service := NewService(repo)

	avr, err := service.AverageDestination(cxt, "China")

	assert.Nil(t, err)
	assert.NotNil(t, avr)
	assert.True(t, dbMock.spy)
}
