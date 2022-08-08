package entity

import (
	"context"
	"database/sql"
	"errors"

	"el.com/m/dto"
	"el.com/m/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type FolderBo struct {
	db *sql.DB
}

func NewFolderBo(db *sql.DB) *FolderBo {
	return &FolderBo{db: db}
}

func (folder *FolderBo) CreateFolder(ctx context.Context, request dto.CreateFolderRequest) (*models.Folder, error) {
	err := validate.Struct(request)

	if err != nil {
		return nil, err
	}

	tx, err := folder.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	var models models.Folder
	models.Name = request.Name
	models.UserID = request.UserId

	err = models.Insert(ctx, tx, boil.Infer())

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &models, nil
}

func (folder *FolderBo) UpdateFolder(ctx context.Context, request dto.UpdateFolderRequest) (*models.Folder, error) {
	err := validate.Struct(request)

	if err != nil {
		return nil, err
	}

	tx, err := folder.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	models, err := models.Folders(
		models.FolderWhere.ID.EQ(request.ID),
	).One(ctx, tx)

	if err != nil {
		return nil, err
	}

	if models.UserID != request.UserId {
		return nil, errors.New("Not your folder")
	}

	models.Name = request.Name

	_, err = models.Update(ctx, tx, boil.Infer())

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return models, nil
}

func (folder *FolderBo) DeleteFolder(ctx context.Context, request dto.DeleteFolderRequest) (*models.Folder, error) {
	err := validate.Struct(request)

	if err != nil {
		return nil, err
	}

	tx, err := folder.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	models, err := models.Folders(
		models.FolderWhere.ID.EQ(request.ID),
	).One(ctx, tx)

	if err != nil {
		return nil, err
	}

	if models.UserID != request.UserId {
		return nil, errors.New("Not your folder")
	}

	_, err = models.Delete(ctx, tx)

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return models, nil
}

func (folder *FolderBo) FindFolder(ctx context.Context, request dto.DeleteFolderRequest) (*models.Folder, error) {
	err := validate.Struct(request)

	if err != nil {
		return nil, err
	}

	tx, err := folder.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	models, err := models.Folders(
		models.FolderWhere.ID.EQ(request.ID),
		models.FolderWhere.UserID.EQ(request.UserId),
	).One(ctx, tx)

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return models, nil
}

func (folder *FolderBo) FindFolders(ctx context.Context, request dto.FindFoldersRequest) (*dto.FindFoldersResponse, error) {
	err := validate.Struct(request)

	if err != nil {
		return nil, err
	}

	tx, err := folder.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	count, err := models.Folders().Count(ctx, tx)

	modelsFolder, err := models.Folders(
		models.FolderWhere.UserID.EQ(request.UserId),
		qm.Offset(int(request.PageNum*request.PageSize)),
		qm.Limit(int(request.PageSize)),
	).All(ctx, tx)

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &dto.FindFoldersResponse{Total: uint(count), Data: &modelsFolder}, nil
}
