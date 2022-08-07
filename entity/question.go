package entity

import (
	"context"
	"database/sql"
	"errors"

	"el.com/m/dto"
	"el.com/m/models"
	"el.com/m/util"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type QuestionBo struct {
	db *sql.DB
}

func NewQuestionBo(db *sql.DB) *QuestionBo {
	return &QuestionBo{db: db}
}

func (question *QuestionBo) CreateQuestion(ctx context.Context, request dto.CreateQuestionRequest) (*models.Question, error) {
	err := validate.Struct(request)

	if err != nil {
		return nil, err
	}

	if request.QuestionType != util.AudioQuestion && request.QuestionType != util.ChoiceQuestion && request.QuestionType != util.FillQuestion {
		return nil, errors.New("Not valid quesiton type")
	}

	tx, err := question.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	test, err := models.Tests(
		models.TestWhere.ID.EQ(request.TestID),
	).One(ctx, question.db)

	if err != nil {
		return nil, err
	}

	if test.Published == 1 {
		return nil, errors.New("Can not add quetion to a published test")
	}

	var models models.Question
	models.Content = request.Content
	models.Answer = request.Answer
	models.QuestionType = request.QuestionType
	models.TestID = request.TestID

	err = models.Insert(ctx, question.db, boil.Infer())

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &models, nil
}

func (question *QuestionBo) UpdateQuestion(ctx context.Context, request dto.UpdateQuestionRequest) (*models.Question, error) {
	err := validate.Struct(request)

	if err != nil {
		return nil, err
	}

	tx, err := question.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	test, err := models.Tests(
		models.TestWhere.ID.EQ(request.TestID),
	).One(ctx, question.db)

	if err != nil {
		return nil, err
	}

	if test.Published == 1 {
		return nil, errors.New("Can not update quetion of a published test")
	}

	models, err := models.Questions(
		models.QuestionWhere.ID.EQ(request.ID),
	).One(ctx, question.db)

	if err != nil {
		return nil, err
	}

	models.Content = request.Content
	models.Answer = request.Answer

	_, err = models.Update(ctx, question.db, boil.Infer())

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return models, nil
}

func (question *QuestionBo) DeleteQuestion(ctx context.Context, ID uint) (*models.Question, error) {

	tx, err := question.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	questionModel, err := models.Questions(
		models.QuestionWhere.ID.EQ(ID),
	).One(ctx, question.db)

	if err != nil {
		return nil, err
	}

	test, err := models.Tests(
		models.TestWhere.ID.EQ(questionModel.TestID),
	).One(ctx, question.db)

	if test.Published == 1 {
		return nil, errors.New("Can not delete quetion of a published test")
	}

	_, err = questionModel.Delete(ctx, question.db)

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return questionModel, nil
}
