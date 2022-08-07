package entity

import (
	"context"
	"database/sql"

	"el.com/m/dto"
	"el.com/m/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type LessonBo struct {
	db *sql.DB
}

func NewLessonBo(db *sql.DB) *LessonBo {
	return &LessonBo{db: db}
}

func (lesson *LessonBo) CreateLesson(ctx context.Context, request dto.CreateLessonRequest) (*models.Lesson, error) {
	err := validate.Struct(request)

	if err != nil {
		return nil, err
	}

	tx, err := lesson.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	var models models.Lesson
	models.LessonName = request.Name
	models.ImageURL = request.ImageURL
	models.Description = request.Description
	models.Content = request.Content
	models.ManagerID = request.ManagerID

	err = models.Insert(ctx, lesson.db, boil.Infer())

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &models, nil
}
