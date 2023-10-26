package entities

type Transport struct {
	Id            int     `json:"id"`
	CanBeRented   bool    `json:"can_be_rented"`
	TransportType string  `json:"transport_type"`
	Model         string  `json:"model"`
	Color         string  `json:"color"`
	Identifier    string  `json:"identifier"`
	Description   string  `json:"description"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	MinutePrice   float64 `json:"minute_price"`
	DayPrice      float64 `json:"day_price"`
}
