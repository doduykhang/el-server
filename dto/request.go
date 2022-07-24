package dto

import "time"

type RegisterRequest struct {
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	LastName    string    `json:"lastName"`
	FirtstName  string    `json:"firstName"`
	Gender      bool      `json:"gender"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
