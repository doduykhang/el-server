package entity

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"el.com/m/dto"
	"el.com/m/models"
	"el.com/m/util"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TestBo struct {
	db *sql.DB
}

func NewTestBo(db *sql.DB) *TestBo {
	return &TestBo{db: db}
}

func (test *TestBo) CreateTest(ctx context.Context, request dto.CreateTestRequest) (*models.Test, error) {
	err := validate.Struct(request)

	if err != nil {
		return nil, err
	}

	tx, err := test.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	var models models.Test
	models.TestName = request.TestName
	models.Time = int(request.Time)
	models.Level = int(request.Level)
	models.LessonID = request.LessonID
	models.ManagerID = request.ManagerID

	err = models.Insert(ctx, test.db, boil.Infer())

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &models, nil
}

func (test *TestBo) UpdateTest(ctx context.Context, request dto.UpdateTestRequest) (*models.Test, error) {
	err := validate.Struct(request)

	if err != nil {
		return nil, err
	}

	tx, err := test.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	models, err := models.Tests(
		models.TestWhere.ID.EQ(request.ID),
	).One(ctx, test.db)

	if err != nil {
		return nil, err
	}

	if models.Published == 1 {
		return nil, errors.New("Can not update published test")
	}

	models.TestName = request.TestName
	models.Time = int(request.Time)
	models.Level = int(request.Level)

	_, err = models.Update(ctx, test.db, boil.Infer())

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return models, nil
}

func (test *TestBo) Delete(ctx context.Context, ID uint) (*models.Test, error) {

	tx, err := test.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	models, err := models.Tests(
		models.TestWhere.ID.EQ(ID),
	).One(ctx, test.db)

	if err != nil {
		return nil, err
	}

	_, err = models.Delete(ctx, test.db)

	if err != nil {
		return nil, err
	}

	if models.Published == 1 {
		return nil, errors.New("Can not delete published test")
	}

	tx.Commit()
	return models, nil
}

func (test *TestBo) FindTest(ctx context.Context, ID uint) (*models.Test, error) {

	tx, err := test.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	models, err := models.Tests(
		models.TestWhere.ID.EQ(ID),
	).One(ctx, test.db)

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return models, nil
}

func (test *TestBo) FindTests(ctx context.Context, request dto.FindTestRequest) (*dto.FindTestResponse, error) {

	tx, err := test.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	count, err := models.Tests(
		qm.Offset(int(request.PageNum*request.PageSize)),
		qm.Limit(int(request.PageSize)),
	).Count(ctx, test.db)

	if err != nil {
		return nil, err
	}

	models, err := models.Tests(
		qm.Offset(int(request.PageNum*request.PageSize)),
		qm.Limit(int(request.PageSize)),
	).All(ctx, test.db)

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &dto.FindTestResponse{Total: uint(count), Data: models}, nil
}

func (test *TestBo) GetQuestions(ctx context.Context, ID uint) (*models.QuestionSlice, error) {

	tx, err := test.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	questions, err := models.Questions(
		models.QuestionWhere.TestID.EQ(ID),
	).All(ctx, test.db)

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &questions, nil
}

func (test *TestBo) PublishTest(ctx context.Context, ID uint) (*models.Test, error) {

	tx, err := test.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	testModel, err := models.Tests(
		models.TestWhere.ID.EQ(ID),
	).One(ctx, test.db)

	if err != nil {
		return nil, err
	}

	if testModel.Published == 1 {
		return nil, errors.New("Test already published")
	}

	questionCount, err := models.Questions(
		models.QuestionWhere.TestID.EQ(ID),
	).Count(ctx, test.db)

	if questionCount < 3 {
		return nil, errors.New("Test has less then 3 questions")
	}

	questions, err := models.Questions(
		models.QuestionWhere.TestID.EQ(ID),
		models.QuestionWhere.QuestionType.EQ(util.ChoiceQuestion),
	).All(ctx, test.db)

	if err != nil {
		return nil, err
	}

	for _, question := range questions {
		options, err := models.Options(
			models.OptionWhere.QuestionID.EQ(question.ID),
		).All(ctx, test.db)

		if err != nil {
			return nil, err
		}

		if len(options) < 2 {
			return nil, errors.New(fmt.Sprintf("Question with id %d have less then 2 options", question.ID))
		}

		var haveAnswer bool = false
		for _, option := range options {
			if option.Content == question.Answer {
				haveAnswer = true
				break
			}
		}

		if haveAnswer == false {
			return nil, errors.New(fmt.Sprintf("Question with id %d does not have an answer", question.ID))
		}
	}

	testModel.Published = 1

	_, err = testModel.Update(ctx, test.db, boil.Infer())

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return testModel, nil
}
