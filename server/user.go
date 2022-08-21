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
	router.Post("/logout", Logout)
	router.Route("/test-admin", TestMwAdmin)
	router.Route("/test-user", TestMwUser)
	router.Route("/", func(r chi.Router) {
		r.Use(middlewares.EnsureAuthenticatedJwtMw(db, util.AllRole))
		r.Get("/profile", GetProfile)
	})
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
	expiresAt := time.Now().Add(100 * time.Minute)
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   token,
		Expires: expiresAt,
	})

	util.JSONResp("Logged in", 200, w)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Unix(0, 0),
	})
	util.JSONResp("Logged in", 200, w)
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	var request dto.GetProfileRequest

	ID := util.UserIDFromContext(r.Context())
	RoleID := util.RoleIDFromContext(r.Context())
	request.ID = ID
	request.RoleID = RoleID

	token, err := userBo.GetProfile(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(token, 200, w)
}

func TestMwAdmin(r chi.Router) {

	r.Use(middlewares.EnsureAuthenticatedJwtMw(db, util.AdminRole))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		accountID := util.UserIDFromContext(r.Context())
		util.JSONResp(fmt.Sprintf("you are in admin %d", accountID), 200, w)
	})
}

func TestMwUser(r chi.Router) {

	r.Use(middlewares.EnsureAuthenticatedJwtMw(db, util.UserRole))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		accountID := util.UserIDFromContext(r.Context())
		util.JSONResp(fmt.Sprintf("you are in user %d", accountID), 200, w)
	})
}
