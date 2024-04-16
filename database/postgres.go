package database

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/rodolfopenia/apuestatotal-eval/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}

func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (id,name, email, password) VALUES ($1, $2, $3, $4)", user.Id, user.Name, user.Email, user.Password)
	return err
}

func (repo *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, email FROM users WHERE id = $1", id)
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Name, &user.Email); err == nil {
			return &user, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, email, password FROM users WHERE email = $1", email)
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err == nil {
			return &user, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *PostgresRepository) InsertFlight(ctx context.Context, flight *models.Flight) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO flights (id, name, airplane_id, departure_date, departure_time, arrival_date, arrival_time, origin_city, destination_city,price) VALUES ($1, $2, $3, to_date($4, 'DD/MM/YYYY'), $5, to_date($6, 'DD/MM/YYYY'), $7, $8, $9, $10)", flight.Id, flight.Name, flight.AirplaneId, flight.DepartureDate, flight.DepartureTime, flight.ArrivalDate, flight.ArrivalTime, flight.OriginCity, flight.DestinationCity, flight.Price)
	return err
}

func (repo *PostgresRepository) ListFlight(ctx context.Context, inputDate string, originCity string) ([]*models.Flight, error) {
	var commandSql string = ""
	if inputDate == "" && originCity == "" {
		commandSql = "SELECT id, name, departure_date, departure_time, arrival_date, arrival_time, origin_city, destination_city, price FROM flights"
	} else {
		commandSql = "SELECT id, name, departure_date, departure_time, arrival_date, arrival_time, origin_city, destination_city, price FROM flights WHERE departure_date = to_date('" + inputDate + "', 'DDMMYYYY') and origin_city ='" + originCity + "'"
	}

	rows, err := repo.db.QueryContext(ctx, commandSql)

	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var flights []*models.Flight
	for rows.Next() {
		var flight = models.Flight{}
		if err = rows.Scan(&flight.Id, &flight.Name, &flight.DepartureDate, &flight.DepartureTime, &flight.ArrivalDate, &flight.ArrivalTime, &flight.OriginCity, &flight.DestinationCity, &flight.Price); err == nil {
			flights = append(flights, &flight)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return flights, nil
}

func (repo *PostgresRepository) InsertBooking(ctx context.Context, booking *models.Booking) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO bookings (id, flight_id, number_passenger, total_price) VALUES ($1, $2, $3, $4)", booking.Id, booking.FlightId, booking.NumberPassenger, booking.TotalPrice)
	return err
}

func (repo *PostgresRepository) InsertBookingDetail(ctx context.Context, bookingDetail *models.BookingDetail) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO booking_detail (id, booking_id, user_id, name_passenger, lastname_passenger, doc_passenger, seat_id, baggage_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", bookingDetail.Id, bookingDetail.BookingId, bookingDetail.UserId, bookingDetail.NamePassenger, bookingDetail.LastNamePassenger, bookingDetail.DocPassenger, bookingDetail.SeatId, bookingDetail.BaggageId)
	return err
}

func (repo *PostgresRepository) ListBooking(ctx context.Context) ([]*models.BookingList, error) {

	rows, err := repo.db.QueryContext(ctx, "SELECT a.id, a.total_price, b.name_passenger, b.lastname_passenger, b.doc_passenger FROM bookings a INNER JOIN booking_detail b ON a.id = b.booking_id")

	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var bookings []*models.BookingList
	for rows.Next() {
		var booking = models.BookingList{}
		if err = rows.Scan(&booking.Id, &booking.TotalPrice, &booking.NamePassenger, &booking.LastNamePassenger, &booking.DocPassenger); err == nil {
			bookings = append(bookings, &booking)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bookings, nil
}

func (repo *PostgresRepository) ListBookingById(ctx context.Context, id string) ([]*models.BookingList, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT a.id, a.total_price, b.name_passenger, b.lastname_passenger, b.doc_passenger FROM bookings a INNER JOIN booking_detail b ON a.id = b.booking_id WHERE a.id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var bookings []*models.BookingList
	for rows.Next() {
		var booking = models.BookingList{}
		if err = rows.Scan(&booking.Id, &booking.TotalPrice, &booking.NamePassenger, &booking.LastNamePassenger, &booking.DocPassenger); err == nil {
			bookings = append(bookings, &booking)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bookings, nil
}
