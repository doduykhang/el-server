package entity

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"el.com/m/dto"
	"el.com/m/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type WordBo struct {
	db *sql.DB
}

func NewWordBo(db *sql.DB) *WordBo {
	return &WordBo{db: db}
}

func (word *WordBo) CreateWord(ctx context.Context, request dto.CreateWordRequest) (*models.Word, error) {
	err := validate.Struct(request)

	if err != nil {
		return nil, err
	}

	existedWord, err := models.Words(qm.Where("word=?", request.Word)).One(ctx, word.db)

	if existedWord != nil {
		return nil, errors.New("Word already exist")
	}

	tx, err := word.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	var wordModels models.Word
	wordModels.Word = request.Word
	wordModels.Definition = request.Definition
	wordModels.Example = request.Example
	wordModels.Pronounciation = request.Pronounciation
	wordModels.Type = request.Type
	wordModels.ManagerID = request.ManagerID

	err = wordModels.Insert(ctx, word.db, boil.Infer())

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &wordModels, nil
}

func (word *WordBo) UpdateWord(ctx context.Context, request dto.UpdateWordRequest) (*models.Word, error) {
	err := validate.Struct(request)

	if err != nil {
		return nil, err
	}

	tx, err := word.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	wordModels, err := models.Words(
		models.WordWhere.ID.EQ(request.ID),
	).One(ctx, word.db)

	if err != nil {
		return nil, err
	}

	wordModels.Definition = request.Definition
	wordModels.Example = request.Example
	wordModels.Pronounciation = request.Pronounciation
	wordModels.Type = request.Type

	_, err = wordModels.Update(ctx, word.db, boil.Infer())

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return wordModels, nil
}

func (word *WordBo) DeleteWord(ctx context.Context, ID uint) (*models.Word, error) {

	tx, err := word.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	wordModels, err := models.Words(
		models.WordWhere.ID.EQ(ID),
	).One(ctx, word.db)

	if err != nil {
		return nil, err
	}

	_, err = wordModels.Delete(ctx, word.db)

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return wordModels, nil
}

func (word *WordBo) FindWords(ctx context.Context, request dto.FindWordsRequest) (*dto.FindWordsResponse, error) {

	var wordsResponse []dto.FindWordsWithSaved

		rawQuery := fmt.Sprintf(`call el.sp_GetWords(%d, "%s", %d, %d)`, request.UserID, request.Word, request.PageNum*request.PageSize, request.PageSize)
		err := queries.Raw(rawQuery).Bind(ctx, word.db, &wordsResponse)
		if err != nil {
			return nil, err
		}

	count, err := models.Words(
		qm.Where("word like ?", "%"+request.Word+"%"),
	).Count(ctx, word.db)

	if err != nil {
		return nil, err
	}

	return &dto.FindWordsResponse{Total: uint(count), Data: &wordsResponse}, nil
}

func (word *WordBo) FindWordsWithSave(ctx context.Context, request dto.FindWordsRequest) (*dto.FindWordsWithSavedReponse, error) {
	rawQuery := fmt.Sprintf(`call el.sp_GetWords(%d, "%s", %d, %d)`, request.UserID, request.Word, request.PageNum*request.PageSize, request.PageSize)
	var words []dto.FindWordsWithSaved
	err := queries.Raw(rawQuery).Bind(ctx, word.db, &words)

	if err != nil {
		return nil, err
	}

	count, err := models.Words(
		qm.Where("word like ?", "%"+request.Word+"%"),
	).Count(ctx, word.db)

	if err != nil {
		return nil, err
	}

	// return &dto.FindWordsResponse{Total: uint(count), Data: &words}, nil
	return &dto.FindWordsWithSavedReponse{Total: uint(count), Data: &words}, nil
}

func (word *WordBo) FindWord(ctx context.Context, ID uint) (*models.Word, error) {

	wordModels, err := models.Words(
		models.WordWhere.ID.EQ(ID),
	).One(ctx, word.db)

	if err != nil {
		return nil, err
	}

	return wordModels, nil
}

func (word *WordBo) AddWordToUser(ctx context.Context, request dto.AddWordToUser) (string, error) {
	err := validate.Struct(request)

	if err != nil {
		return "", err
	}

	tx, err := word.db.BeginTx(ctx, nil)

	if err != nil {
		return "", err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	userModel, err := models.Users(
		models.UserWhere.ID.EQ(request.UserID),
	).One(ctx, tx)

	if err != nil {
		return "", err
	}

	wordModel, err := models.Words(
		models.WordWhere.ID.EQ(request.WordID),
	).One(ctx, tx)

	if err != nil {
		return "", err
	}

	err = userModel.AddWords(ctx, tx, false, wordModel)
	if err != nil {
		return "", err
	}

	tx.Commit()
	return "Added", nil
}

func (word *WordBo) RemoveWordFromUser(ctx context.Context, request dto.AddWordToUser) (string, error) {
	err := validate.Struct(request)

	if err != nil {
		return "", err
	}

	tx, err := word.db.BeginTx(ctx, nil)

	if err != nil {
		return "", err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	userModel, err := models.Users(
		models.UserWhere.ID.EQ(request.UserID),
	).One(ctx, tx)

	if err != nil {
		return "", err
	}

	wordModel, err := models.Words(
		models.WordWhere.ID.EQ(request.WordID),
	).One(ctx, tx)

	if err != nil {
		return "", err
	}

	err = userModel.RemoveWords(ctx, tx, wordModel)
	if err != nil {
		return "", err
	}

	tx.Commit()
	return "Remove", nil
}

func (word *WordBo) GetWordsOfUser(ctx context.Context, ID uint) (*models.WordSlice, error) {
	tx, err := word.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	userModel, err := models.Users(
		models.UserWhere.ID.EQ(ID),
	).One(ctx, tx)

	if err != nil {
		return nil, err
	}

	wordsQuery := userModel.Words()

	words, err := wordsQuery.All(ctx, tx)

	tx.Commit()
	return &words, nil
}
