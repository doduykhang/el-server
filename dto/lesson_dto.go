package dto

import "el.com/m/models"

type CreateLessonRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageURL    string `json:"imageURL" validate:"required"`
	Content     string `json:"content" validate:"required"`
	ManagerID   uint
}

type UpdateLessonRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageURL    string `json:"imageURL" validate:"required"`
	Content     string `json:"content" validate:"required"`
}

type FindLessonsResponse struct {
	Total uint              `json:"total"`
	Data  *models.LessonSlice `json:"data"`
}
