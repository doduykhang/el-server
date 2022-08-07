package dto

type CreateLessonRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageURL    string `json:"imageURL" validate:"required"`
	Content     string `json:"content" validate:"required"`
	ManagerID   uint
}
