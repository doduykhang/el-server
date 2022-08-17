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

type FindLessonsRequest struct {
	PaginationRequest
	Name string `form:"name"`
}

type FindLessonsResponse struct {
	Total uint                `json:"total"`
	Data  *models.LessonSlice `json:"data"`
}

type FindLessonResponse struct {
	Lesson  *models.Lesson      `json:"lesson"`
	Words   *[]FindWordsWithSaved `json:"words"`
	Manager *models.Manager     `json:"manager"`
	Tets    *models.TestSlice   `json:"tests"`
}
