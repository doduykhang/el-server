package dto

import "el.com/m/models"

type CreateWordRequest struct {
	Word           string `json:"word" validate:"required"`
	Definition     string `json:"definition" validate:"required"`
	Example        string `json:"example" validate:"required"`
	Pronounciation string `json:"pronounciation" validate:"required"`
	Type           string `json:"type" validate:"required"`
	ManagerID      uint
}

type UpdateWordRequest struct {
	ID             uint   `json:"id" validate:"required"`
	Definition     string `json:"definition" validate:"required"`
	Example        string `json:"example" validate:"required"`
	Pronounciation string `json:"pronounciation" validate:"required"`
	Type           string `json:"type" validate:"required"`
}

type FindWordsResponse struct {
	Total uint              `json:"total"`
	Data  *models.WordSlice `json:"data"`
}
