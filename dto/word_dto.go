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

type FindWordsRequest struct {
	PaginationRequest
	Word   string `form:"word"`
	UserID uint
}

type FindWordsResponse struct {
	Total uint              `json:"total"`
	Data  *models.WordSlice `json:"data"`
}

type AddWordToUser struct {
	WordID uint `json:"wordID"`
	UserID uint
}

type FindWordsWithSaved struct {
	ID             string `json:"id" boil:"id"`
	Word           string `json:"word" boil:"word"`
	Definition     string `json:"definition" boil:"definition"`
	Example        string `json:"example" boil:"example"`
	Type           string `json:"type" boil:"type"`
	Pronounciation string `json:"pronounciation" boil:"pronounciation"`
	Saved          bool `json:"saved" boil:"saved" `
}

type FindWordsWithSavedReponse struct {
	Total uint                  `json:"total"`
	Data  *[]FindWordsWithSaved `json:"data"`
}
