package repository

import (
	"context"

	"github.com/rodolfopenia/apuestatotal-eval/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	InsertFlight(ctx context.Context, flight *models.Flight) error
	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func Close() error {
	return implementation.Close()
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func InsertFlight(ctx context.Context, flight *models.Flight) error {
	return implementation.InsertFlight(ctx, flight)
}
