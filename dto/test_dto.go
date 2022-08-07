package dto

import "el.com/m/models"

type CreateTestRequest struct {
	TestName  string `json:"testName" validate:"required"`
	Time      uint   `json:"time" validate:"required"`
	Level     uint   `json:"level" validate:"required"`
	LessonID  uint   `json:"lessonID" validate:"required"`
	ManagerID uint
}

type UpdateTestRequest struct {
	ID       uint   `json:"id" validate:"required"`
	TestName string `json:"testName" validate:"required"`
	Time     uint   `json:"time" validate:"required"`
	Level    uint   `json:"level" validate:"required"`
	LessonID uint   `json:"lessonID" validate:"required"`
}

type FindTestRequest struct {
	PaginationRequest
}

type FindTestResponse struct {
	Total uint             `json:"total"`
	Data  models.TestSlice `json:"data"`
}

type SubmitTestRequest struct {
	TestID  uint            `json:"testID" validate:"required"`
	Answers map[uint]string `json:"answers" validate:"required"`
  UserID uint 
}

