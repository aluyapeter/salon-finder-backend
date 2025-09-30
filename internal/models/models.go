package models

import "time"

type UserRole string

const (
	RoleCustomer UserRole = "customer"
	RoleProvider UserRole = "provider"
)

type User struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Role      UserRole  `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ProviderProfile struct {
	ID           int64     `db:"id"`
	UserID       int64     `db:"user_id"`
	BusinessNmae string    `db:"business_name"`
	Address      string    `db:"address"`
	City         string    `db:"city"`
	State        string    `db:"state"`
	Latitude     float64   `db:"latitude"`
	Longitude    float64   `db:"longitude"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type Service struct {
	ID              int64     `db:"id"`
	ProviderID      int64     `db:"provider_id"`
	Name            string    `db:"name"`
	DurationMinutes int       `db:"duration_minutes"`
	PriceAmount     int64     `db:"price_amount"`
	Currency        string    `db:"currency"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type BookingStatus string

const (
	StatusPending   BookingStatus = "pending"
	StatusConfirmed BookingStatus = "confirmed"
	StatusCanceled  BookingStatus = "canceled"
)

type Booking struct {
	ID        int64         `db:"id"`
	UserID    int64         `db:"user_id"`
	ServiceID int64         `db:"service_id"`
	BookedAt  time.Time     `db:"booked_at"`
	Status    BookingStatus `db:"status"`
	CreatedAt time.Time     `db:"created_at"`
	UpdatedAt time.Time     `db:"updated_at"`
}

type Review struct {
	ID        int64     `db:"id"`
	BookingID int64     `db:"booking_id"`
	Rating    int       `db:"rating"`
	Comment   string    `db:"comment"`
	CreatedAt time.Time `db:"created_at"`
}

// type Salon struct {
// 	ID          int
// 	Name        string
// 	Description string
// 	Address     string
// 	City        string
// 	State       string
// 	Latitude    float64
// 	Longitude   float64
// 	CreatedAt   time.Time
// 	UpdatedAt   time.Time
// }

// type Price struct {
// 	ID        int
// 	ServiceID int
// 	Amount 	  float64
// 	Currency  string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }
