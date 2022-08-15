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

	count, err := models.Tests().Count(ctx, test.db)

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

func (test *TestBo) UnPublishTest(ctx context.Context, ID uint) (*models.Test, error) {

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

	if testModel.Published == 0 {
		return nil, errors.New("Test has not been published")
	}

	userTest, err := models.UserTests(
		models.UserTestWhere.TestID.EQ(ID),
	).One(ctx, test.db)

	if userTest != nil {
		return nil, errors.New("User has already taken this test")
	}

	testModel.Published = 0
	_, err = testModel.Update(ctx, test.db, boil.Infer())

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return testModel, nil
}

func (test *TestBo) SubmitTest(ctx context.Context, request dto.SubmitTestRequest) (string, error) {

	tx, err := test.db.BeginTx(ctx, nil)

	if err != nil {
		return "", err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	testModel, err := models.Tests(
		models.TestWhere.ID.EQ(request.TestID),
	).One(ctx, tx)

	if err != nil {
		return "", err
	}

	if testModel.Published == 0 {
		return "", errors.New("Test has not been published")
	}
	time := 1
	previousTest, err := models.UserTests(
		models.UserTestWhere.UserID.EQ(request.UserID),
		models.UserTestWhere.TestID.EQ(request.TestID),
		qm.OrderBy("time desc"),
	).One(ctx, tx)

	if previousTest != nil {
		time = previousTest.Time + 1
	}

	var userTest models.UserTest
	userTest.UserID = request.UserID
	userTest.TestID = request.TestID
	userTest.Score = 0
	userTest.Time = time

	err = userTest.Insert(ctx, tx, boil.Infer())

	if err != nil {
		return "", err
	}

	questions, err := models.Questions(
		models.QuestionWhere.TestID.EQ(testModel.ID),
	).All(ctx, tx)

	score := 0.0

	for _, question := range questions {
		answer := ""
		if val, ok := request.Answers[question.ID]; ok {
			if val == question.Answer {
				score++
			}
			answer = val
		}

		var userTestDetail models.UserTestDetail
		userTestDetail.TestID = userTest.ID
		userTestDetail.QuestionID = question.ID
		userTestDetail.Answer = answer
		err = userTestDetail.Insert(ctx, tx, boil.Infer())

		if err != nil {
			return "", err
		}
	}

	score = score / float64(len(questions))

	userTest.Score = float32(score)
	_, err = userTest.Update(ctx, tx, boil.Infer())

	if err != nil {
		return "", err
	}

	tx.Commit()

	return "Ok", nil
}
