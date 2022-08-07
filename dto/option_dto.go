package dto

type CreateOptionRequest struct {
	Content    string `json:"content" validate:"required"`
	Position   uint   `json:"position" validate:"required"`
	QuestionID uint   `json:"questionID" validate:"required"`
}

type UpdateOptionRequest struct {
	ID         uint   `json:"id" validate:"required"`
	Content    string `json:"content" validate:"required"`
	Position   uint   `json:"position" validate:"required"`
	QuestionID uint   `json:"questionID" validate:"required"`
}
