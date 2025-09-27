package models

import "time"

type Salon struct {
	ID          int
	Name        string
	Description string
	Address     string
	City        string
	State       string
	Latitude    float64
	Longitude   float64
	CreatedAt   time.Time
	UpdatedAt time.Time
}
type Service struct {
	ID          int
	SalonID 	int
	Name string
	DurationMinutes time.Time
	CreatedAt   time.Time
	UpdatedAt time.Time
}
type Price struct {
	ID        int
	ServiceID int
	Amount 	  float64
	Currency  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type User struct {
	ID     int
	Name   string
	Email  string
}