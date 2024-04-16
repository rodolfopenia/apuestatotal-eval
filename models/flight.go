package models

type Flight struct {
	Id              string  `json:"id"`
	Name            string  `json:"name"`
	AirplaneId      string  `json:"airplane_id"`
	DepartureDate   string  `json:"departure_date"`
	DepartureTime   string  `json:"departure_time"`
	ArrivalDate     string  `json:"arrival_date"`
	ArrivalTime     string  `json:"arrival_time"`
	OriginCity      string  `json:"origin_city"`
	DestinationCity string  `json:"destination_city"`
	Price           float64 `json:"price"`
}
