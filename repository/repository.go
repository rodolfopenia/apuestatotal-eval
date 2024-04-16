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
	ListFlight(ctx context.Context, inputDate string, originCity string) ([]*models.Flight, error)
	InsertBooking(ctx context.Context, booking *models.Booking) error
	InsertBookingDetail(ctx context.Context, bookingDetail *models.BookingDetail) error
	ListBooking(ctx context.Context) ([]*models.BookingList, error)
	ListBookingById(ctx context.Context, id string) ([]*models.BookingList, error)
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

func ListFlight(ctx context.Context, inputDate string, originCity string) ([]*models.Flight, error) {
	return implementation.ListFlight(ctx, inputDate, originCity)
}

func InsertBooking(ctx context.Context, booking *models.Booking) error {
	return implementation.InsertBooking(ctx, booking)
}

func InsertBookingDetail(ctx context.Context, bookingDetail *models.BookingDetail) error {
	return implementation.InsertBookingDetail(ctx, bookingDetail)
}

func ListBooking(ctx context.Context) ([]*models.BookingList, error) {
	return implementation.ListBooking(ctx)
}

func ListBookingById(ctx context.Context, id string) ([]*models.BookingList, error) {
	return implementation.ListBookingById(ctx, id)
}
