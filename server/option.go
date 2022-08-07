package server

import (
	"encoding/json"
	"net/http"

	"el.com/m/dto"
	"el.com/m/entity"
	"el.com/m/middlewares"
	"el.com/m/util"
	"github.com/go-chi/chi/v5"
)

var (
	optionBo *entity.OptionBo
)

func init() {
	optionBo = entity.NewOptionBo(db)
}

func OptionRoute(router chi.Router) {
	router.Route("/", func(r chi.Router) {
		r.Use(middlewares.EnsureAuthenticatedJwtMw(db, util.AdminRole))
		r.Post("/", CreateOption)
		r.Put("/", UpdateOption)
		r.Delete("/{ID}", DeleteOption)

		r.Post("/add-word", AddWord)
		r.Delete("/remove-word", RemoveWord)
	})

	router.Get("/{ID}", FindLesson)
	router.Get("/all", FindLessons)
	router.Get("/get-words/{ID}", GetWords)
	router.Get("/get-tests/{ID}", GetWords)
}

func CreateOption(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateOptionRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	result, err := optionBo.CreateOption(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func UpdateOption(w http.ResponseWriter, r *http.Request) {
	var request dto.UpdateOptionRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	result, err := optionBo.UpdateOption(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func DeleteOption(w http.ResponseWriter, r *http.Request) {
	IDString := chi.URLParam(r, "ID")
	ID, err := util.IDFromStr(IDString)

	result, err := optionBo.DeleteOption(r.Context(), ID)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}
