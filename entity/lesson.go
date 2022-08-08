package entity

import (
	"context"
	"database/sql"

	"el.com/m/dto"
	"el.com/m/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

func (lesson *LessonBo) UpdateLesson(ctx context.Context, request dto.UpdateLessonRequest) (*models.Lesson, error) {
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

	models, err := models.Lessons(
		models.LessonWhere.ID.EQ(request.ID),
	).One(ctx, lesson.db)

	if err != nil {
		return nil, err
	}

	models.LessonName = request.Name
	models.ImageURL = request.ImageURL
	models.Description = request.Description
	models.Content = request.Content

	_, err = models.Update(ctx, lesson.db, boil.Infer())

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return models, nil
}

func (lesson *LessonBo) DeleteLesson(ctx context.Context, ID uint) (*models.Lesson, error) {

	tx, err := lesson.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	models, err := models.Lessons(
		models.LessonWhere.ID.EQ(ID),
	).One(ctx, lesson.db)

	if err != nil {
		return nil, err
	}

	_, err = models.Delete(ctx, lesson.db)

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return models, nil
}

func (lesson *LessonBo) FindLesson(ctx context.Context, ID uint) (*models.Lesson, error) {

	tx, err := lesson.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	models, err := models.Lessons(
		models.LessonWhere.ID.EQ(ID),
	).One(ctx, lesson.db)

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return models, nil
}

func (lesson *LessonBo) FindLessons(ctx context.Context, request dto.PaginationRequest) (*dto.FindLessonsResponse, error) {

	tx, err := lesson.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	count, err := models.Lessons().Count(ctx, lesson.db)

	if err != nil {
		return nil, err
	}

	models, err := models.Lessons(
		qm.Offset(int(request.PageNum*request.PageSize)),
		qm.Limit(int(request.PageSize)),
	).All(ctx, lesson.db)

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &dto.FindLessonsResponse{Total: uint(count), Data: &models}, nil
}

func (lesson *LessonBo) AddWord(ctx context.Context, request dto.AddWordToLesson) (string, error) {
	err := validate.Struct(request)

	if err != nil {
		return "", err
	}

	tx, err := lesson.db.BeginTx(ctx, nil)

	if err != nil {
		return "", err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	lessonModel, err := models.Lessons(
		models.LessonWhere.ID.EQ(request.LessonID),
	).One(ctx, lesson.db)

	if err != nil {
		return "", err
	}

	word, err := models.Words(
		models.WordWhere.ID.EQ(request.WordID),
	).One(ctx, lesson.db)

	if err != nil {
		return "", err
	}

	err = lessonModel.AddWords(ctx, lesson.db, false, word)
	if err != nil {
		return "", err
	}

	tx.Commit()
	return "Added", nil
}

func (lesson *LessonBo) RemoveWord(ctx context.Context, request dto.AddWordToLesson) (string, error) {
	err := validate.Struct(request)

	if err != nil {
		return "", err
	}

	tx, err := lesson.db.BeginTx(ctx, nil)

	if err != nil {
		return "", err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	lessonModel, err := models.Lessons(
		models.LessonWhere.ID.EQ(request.LessonID),
	).One(ctx, lesson.db)

	if err != nil {
		return "", err
	}

	word, err := models.Words(
		models.WordWhere.ID.EQ(request.WordID),
	).One(ctx, lesson.db)

	if err != nil {
		return "", err
	}

	err = lessonModel.RemoveWords(ctx, lesson.db, word)
	if err != nil {
		return "", err
	}

	tx.Commit()
	return "Removed", nil
}

func (lesson *LessonBo) GetWords(ctx context.Context, ID uint) (*models.WordSlice, error) {
	tx, err := lesson.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	lessonModel, err := models.Lessons(
		models.LessonWhere.ID.EQ(ID),
	).One(ctx, lesson.db)

	if err != nil {
		return nil, err
	}

	wordsQuery := lessonModel.Words()

	words, err := wordsQuery.All(ctx, lesson.db)

	tx.Commit()
	return &words, nil
}

func (lesson *LessonBo) GetTests(ctx context.Context, ID uint) (*models.TestSlice, error) {
	tx, err := lesson.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	tests, err := models.Tests(
		models.TestWhere.LessonID.EQ(ID),
	).All(ctx, lesson.db)

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &tests, nil
}
