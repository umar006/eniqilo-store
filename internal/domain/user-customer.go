package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserCustomer struct {
	ID          string    `db:"id"`
	CreatedAt   time.Time `db:"created_at"`
	Name        string    `db:"name"`
	PhoneNumber string    `db:"phone_number"`
	Password    string    `db:"password"`
}

type RegisterUserCustomerRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterUserCustomerResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}

type UserCustomerResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}

func (cr *RegisterUserCustomerRequest) NewUserCustomer() UserCustomer {
	id := uuid.New()
	rawCreatedAt := time.Now().Format(time.RFC3339)
	createdAt, _ := time.Parse(time.RFC3339, rawCreatedAt)

	return UserCustomer{
		ID:          id.String(),
		CreatedAt:   createdAt,
		Name:        cr.Name,
		PhoneNumber: cr.PhoneNumber,
		Password:    cr.Password,
	}
}