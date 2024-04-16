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

/*

func (repo *PostgresRepository) InsertPost(ctx context.Context, post *models.Post) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO posts (id,post_content, user_id) VALUES ($1, $2, $3)", post.Id, post.PostContent, post.UserId)
	return err
}

func (repo *PostgresRepository) GetPostById(ctx context.Context, id string) (*models.Post, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, post_content, created_at, user_id FROM posts WHERE id = $1", id)
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var post = models.Post{}
	for rows.Next() {
		if err = rows.Scan(&post.Id, &post.PostContent, &post.CreatedAt, &post.UserId); err == nil {
			return &post, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &post, nil
}

func (repo *PostgresRepository) UpdatePost(ctx context.Context, post *models.Post, userId string) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE posts SET post_content = $1 WHERE id = $2 and user_id = $3", post.PostContent, post.Id, userId)
	return err
}

func (repo *PostgresRepository) DeletePost(ctx context.Context, id string, userId string) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM posts WHERE id = $1 and user_id = $2", id, userId)
	return err
}

func (repo *PostgresRepository) ListPost(ctx context.Context, page uint64) ([]*models.Post, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, post_content, user_id, created_at FROM posts LIMIT $1 OFFSET $2", 2, page*2)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var posts []*models.Post
	for rows.Next() {
		var post = models.Post{}
		if err = rows.Scan(&post.Id, &post.PostContent, &post.UserId, &post.CreatedAt); err == nil {
			posts = append(posts, &post)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}
*/
