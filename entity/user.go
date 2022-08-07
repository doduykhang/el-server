package entity

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"el.com/m/dto"
	"el.com/m/models"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

var (
	jwtKey = []byte("my_secret_key")
)

type Claims struct {
	ID uint
	jwt.StandardClaims
}

func init() {
	validate = validator.New()
}

type UserBo struct {
	db *sql.DB
}

func NewUserBo(db *sql.DB) *UserBo {
	return &UserBo{db: db}
}

func (user *UserBo) RegisterUser(ctx context.Context, request dto.RegisterRequest) (*models.User, error) {
	err := validate.Struct(request)
	if err != nil {
		return nil, err
	}

	tx, err := user.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	existedAccount, err := models.Accounts(qm.Where("email=?", request.Email)).One(ctx, user.db)

	if existedAccount != nil {
		return nil, errors.New("Email already in use")
	}

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	var account models.Account
	account.Email = request.Email
	account.Password = string(hashedPassword)
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

func (user *UserBo) Login(ctx context.Context, request dto.LoginRequest) (string, error) {
	account, err := models.Accounts(
		models.AccountWhere.Email.EQ(request.Email),
	).One(ctx, user.db)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(request.Password))
	if err != nil {
		return "", errors.New("Wrong username or password")
	}

	expirationTime := time.Now().Add(100 * time.Minute)
	claims := &Claims{
		ID: account.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
