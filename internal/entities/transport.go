package entities

type Transport struct {
	Id            int     `json:"id" db:"id"`
	OwnerId       int     `json:"owner_id" db:"owner_id"`
	CanBeRented   bool    `json:"can_be_rented" db:"can_be_rented"`
	TransportType string  `json:"transport_type" db:"transport_type"`
	Model         string  `json:"model" db:"model"`
	Color         string  `json:"color" db:"color"`
	Identifier    string  `json:"identifier" db:"identifier"`
	Description   string  `json:"description" db:"description"`
	Latitude      float64 `json:"latitude" db:"latitude"`
	Longitude     float64 `json:"longitude" db:"longitude"`
	MinutePrice   float64 `json:"minute_price" db:"minute_price"`
	DayPrice      float64 `json:"day_price" db:"day_price"`
}

type Rent struct {
	Id           int     `json:"id"`
	TransportId  int     `json:"transport_id" db:"transport_id"`
	UserId       int     `json:"user_id" db:"user_id"`
	RentType     string  `json:"rent_type" db:"rent_type"`
	RentingEnded bool    `json:"renting_ended" db:"renting_ended"`
	StartedAt    string  `json:"started_at" db:"started_at"`
	EndedAt      *string `json:"ended_at" db:"ended_at"`
}
