package dto

import "time"

type RegisterRequest struct {
	Email       string    `json:"email" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	LastName    string    `json:"lastName" validate:"required"`
	FirtstName  string    `json:"firstName" validate:"required"`
	Gender      bool      `json:"gender" validate:"required"`
	DateOfBirth time.Time `json:"dateOfBirth" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}