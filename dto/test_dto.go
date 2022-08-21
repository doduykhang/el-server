package dto

import (
	"time"

	"el.com/m/models"
)

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
	UserID  uint
}

type GetUserTest struct {
	ID         uint      `json:"id" boil:"id"`
	StartTime  time.Time `json:"startTime" boil:"start_time"`
	Time       uint      `json:"time" boil:"time"`
	Score      float32   `json:"score" boil:"score"`
	TestID     uint      `json:"testId" boil:"test_id"`
	TestName   string    `json:"testName" boil:"test_name"`
	Level      uint      `json:"level" boil:"level"`
	LessonID   uint      `json:"lessonId" boil:"lesson_id"`
	LessonName string    `json:"lessonName" boil:"lesson_name"`
}

type GetUserTestDetail struct {
	TestID       uint   `json:"testId" boil:"test_id"`
	Content      string `json:"content" boil:"content"`
	QuestionType string `json:"questionType" boil:"question_type"`
	UserAnswer   string `json:"userAnswer" boil:"user_answer"`
	Answer       string `json:"answer" boil:"answer"`
}
