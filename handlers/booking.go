package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/rodolfopenia/apuestatotal-eval/models"
	"github.com/rodolfopenia/apuestatotal-eval/repository"
	"github.com/rodolfopenia/apuestatotal-eval/server"
	"github.com/segmentio/ksuid"
)

type BookingRequest struct {
	Id              string                 `json:"id"`
	FlightId        string                 `json:"flight_id"`
	NumberPassenger int64                  `json:"number_passenger"`
	TotalPrice      float64                `json:"total_price"`
	BookingDetail   []BookingDetailRequest `json:"booking_detail"`
}

type BookingDetailRequest struct {
	Id                string `json:"id"`
	BookingId         string `json:"booking_id"`
	UserId            string `json:"user_id"`
	NamePassenger     string `json:"name_passenger"`
	LastNamePassenger string `json:"lastname_passenger"`
	DocPassenger      string `json:"doc_passenger"`
	SeatId            string `json:"seat_id"`
	BaggageId         string `json:"baggage_id"`
}

func InsertBooking(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			var bookingRequest = BookingRequest{}
			if err := json.NewDecoder(r.Body).Decode(&bookingRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			id, err := ksuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			booking := models.Booking{
				Id:              id.String(),
				FlightId:        bookingRequest.FlightId,
				NumberPassenger: bookingRequest.NumberPassenger,
				TotalPrice:      bookingRequest.TotalPrice,
			}

			err = repository.InsertBooking(r.Context(), &booking)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			for _, value := range bookingRequest.BookingDetail {
				idDet, err := ksuid.NewRandom()
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				bookingDetail := models.BookingDetail{
					Id:                idDet.String(),
					BookingId:         booking.Id,
					UserId:            claims.UserId,
					NamePassenger:     value.NamePassenger,
					LastNamePassenger: value.LastNamePassenger,
					DocPassenger:      value.DocPassenger,
					SeatId:            value.SeatId,
					BaggageId:         value.BaggageId,
				}

				err = repository.InsertBookingDetail(r.Context(), &bookingDetail)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(InsertFlightResponse{
				Id:      booking.Id,
				Message: "Insert booking and detail successfully",
			})

		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func ListBookingHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		bookings, err := repository.ListBooking(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(bookings)
	}
}

func ListBookingByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var err error
		bookings, err := repository.ListBookingById(r.Context(), params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(bookings)
	}
}
