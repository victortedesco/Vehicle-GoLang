package models

type Vehicle struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Brand   string  `json:"brand"`
	Price   float64 `json:"price"`
	Mileage uint32  `json:"mileage"`
	Year    uint16  `json:"year"`
}
