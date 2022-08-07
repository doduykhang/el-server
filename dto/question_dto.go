package dto

type CreateQuestionRequest struct {
	Content      string `json:"content" validate:"required"`
	Answer       string `json:"answer" validate:"required"`
	QuestionType string `json:"questionType" validate:"required"`
	TestID       uint   `json:"testID" validate:"required"`
	ManagerID    uint
}

type UpdateQuestionRequest struct {
	ID      uint   `json:"id" validate:"required"`
	Content string `json:"content" validate:"required"`
	Answer  string `json:"answer" validate:"required"`
	TestID  uint   `json:"testID" validate:"required"`
}
