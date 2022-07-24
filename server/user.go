package server

import (
	"encoding/json"
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
	router.Post("/register", func(w http.ResponseWriter, r *http.Request) {
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
	})

	router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		var request dto.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			util.SadResp(err, 500, w)
			return
		}

		result, err := userBo.Login(r.Context(), request)

		if err != nil {
			util.SadResp(err, 500, w)
			return
		}

		//set cookie
		expiresAt := time.Now().Add(120 * time.Second)
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   result.FirstName,
			Expires: expiresAt,
		})

		util.JSONResp(result, 200, w)
	})

	router.Route("/test", func(r chi.Router) {
		r.Use(middlewares.EnsureAuthenticatedJwtMw())
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			util.JSONResp("You are in", 200, w)
		})
	})

}
