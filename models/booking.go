package models

type Booking struct {
	Id              string  `json:"id"`
	FlightId        string  `json:"flight_id"`
	NumberPassenger int64   `json:"number_passenger"`
	TotalPrice      float64 `json:"total_price"`
	//BookingDetail   []BookingDetail `json:"booking_detail"`
}

type BookingDetail struct {
	Id                string `json:"id"`
	BookingId         string `json:"booking_id"`
	UserId            string `json:"user_id"`
	NamePassenger     string `json:"name_passenger"`
	LastNamePassenger string `json:"lastname_passenger"`
	DocPassenger      string `json:"doc_passenger"`
	SeatId            string `json:"seat_id"`
	BaggageId         string `json:"baggage_id"`
}

type BookingList struct {
	Id                string  `json:"id"`
	TotalPrice        float64 `json:"total_price"`
	NamePassenger     string  `json:"name_passenger"`
	LastNamePassenger string  `json:"lastname_passenger"`
	DocPassenger      string  `json:"doc_passenger"`
}
