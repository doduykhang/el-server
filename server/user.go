package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"el.com/m/dto"
	"el.com/m/entity"
	"el.com/m/middlewares"
	"el.com/m/util"
	"github.com/go-chi/chi/v5"
)

var (
	userBo *entity.UserBo
)

func init() {
	userBo = entity.NewUserBo(db)
}

func UserServer(router chi.Router) {
	router.Post("/register", Register)
	router.Post("/login", Login)
	router.Route("/test", TestMw)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var request dto.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	result, err := userBo.RegisterUser(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var request dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	token, err := userBo.Login(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	//set cookie
	expiresAt := time.Now().Add(120 * time.Second)
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   token,
		Expires: expiresAt,
	})

	util.JSONResp("Logged in", 200, w)
}

func TestMw(r chi.Router) {
	r.Use(middlewares.EnsureAuthenticatedJwtMw(db, util.UserRole))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		accountID := util.UserIDFromContext(r.Context())
		util.JSONResp(fmt.Sprintf("you are in %d", accountID), 200, w)
	})
}
