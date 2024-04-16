package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/rodolfopenia/apuestatotal-eval/models"
	"github.com/rodolfopenia/apuestatotal-eval/repository"
	"github.com/rodolfopenia/apuestatotal-eval/server"
	"github.com/segmentio/ksuid"
)

type InsertFlightResponse struct {
	Id      string `json:"id"`
	Message string `json:"name"`
}

func InsertFlight(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if _, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			var flightRequest = models.Flight{}
			if err := json.NewDecoder(r.Body).Decode(&flightRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			id, err := ksuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			flight := models.Flight{
				Id:              id.String(),
				Name:            flightRequest.Name,
				AirplaneId:      flightRequest.AirplaneId,
				DepartureDate:   flightRequest.DepartureDate,
				DepartureTime:   flightRequest.DepartureTime,
				ArrivalDate:     flightRequest.ArrivalDate,
				ArrivalTime:     flightRequest.ArrivalTime,
				OriginCity:      flightRequest.OriginCity,
				DestinationCity: flightRequest.DestinationCity,
				Price:           flightRequest.Price,
			}

			err = repository.InsertFlight(r.Context(), &flight)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(InsertFlightResponse{
				Id:      flight.Id,
				Message: "Insert flight successfully",
			})

		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
