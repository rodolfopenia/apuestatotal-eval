package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rodolfopenia/apuestatotal-eval/handlers"
	"github.com/rodolfopenia/apuestatotal-eval/middleware"
	"github.com/rodolfopenia/apuestatotal-eval/server"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JWT_SECRET,
		Port:        PORT,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRouter)
}

func BindRouter(s server.Server, r *mux.Router) {
	r.Use(middleware.CheckAuthMiddleware(s))
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/flights", handlers.InsertFlight(s)).Methods(http.MethodPost)
	r.HandleFunc("/flights", handlers.ListFlightHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/bookings", handlers.InsertBooking(s)).Methods(http.MethodPost)
	r.HandleFunc("/bookings", handlers.ListBookingHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/bookings/{id}", handlers.ListBookingByIdHandler(s)).Methods(http.MethodGet)
}
