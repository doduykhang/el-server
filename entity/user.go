package entity

import (
	"context"
	"database/sql"
	"errors"

	"el.com/m/dto"
	"el.com/m/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"
)

type UserBo struct {
	db *sql.DB
}

func NewUserBo(db *sql.DB) *UserBo {
	return &UserBo{db: db}
}

func (user *UserBo) RegisterUser(ctx context.Context, request dto.RegisterRequest) (*models.User, error) {
	tx, err := user.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var account models.Account
	account.Email = request.Email
	account.Password = hashedPassword
	account.RoleID = 2

	err = account.Insert(ctx, tx, boil.Infer())

	if err != nil {
		return nil, err
	}

	var userModel models.User
	userModel.LastName = request.LastName
	userModel.FirstName = request.FirtstName
	userModel.Gender = request.Gender
	userModel.DateOfBirth = request.DateOfBirth
	userModel.AccountID = account.ID

	err = userModel.Insert(ctx, tx, boil.Infer())

	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &userModel, nil
}

func (user *UserBo) Login(ctx context.Context, request dto.LoginRequest) (*models.User, error) {
	account, err := models.Accounts(
		models.AccountWhere.Email.EQ(request.Email),
	).One(ctx, user.db)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(account.Password, []byte(request.Password))
	if err != nil {
		return nil, errors.New("Wrong username or password")
	}

	userModel, err := models.Users(
		models.UserWhere.AccountID.EQ(account.ID),
	).One(ctx, user.db)

	if err != nil {
		return nil, err
	}
	return userModel, nil
}
