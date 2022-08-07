package dto

type AddWordToLesson struct {
	WordID   uint `json:"wordID"`
	LessonID uint `json:"lessonID"`
}
