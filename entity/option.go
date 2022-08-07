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

type OptionBo struct {
	db *sql.DB
}

func NewOptionBo(db *sql.DB) *OptionBo {
	return &OptionBo{db: db}
}

func (option *OptionBo) CreateOption(ctx context.Context, request dto.CreateOptionRequest) (*models.Option, error) {
	err := validate.Struct(request)

	if err != nil {
		return nil, err
	}

	tx, err := option.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	question, err := models.Questions(
		models.QuestionWhere.ID.EQ(request.QuestionID),
	).One(ctx, option.db)

	if err != nil {
		return nil, err
	}

	if question.QuestionType != util.ChoiceQuestion {
		return nil, errors.New("Can only add option to choice question")
	}

	test, err := models.Tests(
		models.TestWhere.ID.EQ(question.TestID),
	).One(ctx, option.db)

	if err != nil {
		return nil, err
	}

	if test.Published == 1 {
		return nil, errors.New("Can not update question of published test")
	}

	checkOption, err := models.Options(
		models.OptionWhere.QuestionID.EQ(request.QuestionID),
		models.OptionWhere.Position.EQ(request.Position),
	).One(ctx, option.db)

	if checkOption != nil {
		return nil, errors.New("Can not use this position")
	}

	var models models.Option
	models.Content = request.Content
	models.Position = request.Position
	models.QuestionID = request.QuestionID

	err = models.Insert(ctx, option.db, boil.Infer())

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &models, nil
}

func (option *OptionBo) UpdateOption(ctx context.Context, request dto.UpdateOptionRequest) (*models.Option, error) {
	err := validate.Struct(request)

	if err != nil {
		return nil, err
	}

	tx, err := option.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	question, err := models.Questions(
		models.QuestionWhere.ID.EQ(request.QuestionID),
	).One(ctx, option.db)

	if err != nil {
		return nil, err
	}

	if question.QuestionType != util.ChoiceQuestion {
		return nil, errors.New("Can only add option to choice question")
	}

	test, err := models.Tests(
		models.TestWhere.ID.EQ(question.TestID),
	).One(ctx, option.db)

	if err != nil {
		return nil, err
	}

	if test.Published == 1 {
		return nil, errors.New("Can not update question of published test")
	}

	checkOption, err := models.Options(
		models.OptionWhere.QuestionID.EQ(request.QuestionID),
		models.OptionWhere.Position.EQ(request.Position),
		models.OptionWhere.ID.NEQ(request.ID),
	).One(ctx, option.db)

	if checkOption != nil {
		return nil, errors.New("Can not use this position")
	}

	optionModels, err := models.Options(
		models.OptionWhere.ID.EQ(request.ID),
	).One(ctx, option.db)

	if err != nil {
		return nil, err
	}

	optionModels.Content = request.Content
	optionModels.Position = request.Position

	_, err = optionModels.Update(ctx, option.db, boil.Infer())

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return optionModels, nil
}

func (option *OptionBo) DeleteOption(ctx context.Context, id uint) (*models.Option, error) {

	tx, err := option.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	optionModels, err := models.Options(
		models.OptionWhere.ID.EQ(id),
	).One(ctx, option.db)

	if err != nil {
		return nil, err
	}

	question, err := models.Questions(
		models.QuestionWhere.ID.EQ(optionModels.QuestionID),
	).One(ctx, option.db)

	if err != nil {
		return nil, err
	}

	test, err := models.Tests(
		models.TestWhere.ID.EQ(question.TestID),
	).One(ctx, option.db)

	if err != nil {
		return nil, err
	}

	if test.Published == 1 {
		return nil, errors.New("Can not update question of published test")
	}

	_, err = optionModels.Delete(ctx, option.db)

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return optionModels, nil
}
