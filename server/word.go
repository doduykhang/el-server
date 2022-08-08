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
	wordBo *entity.WordBo
)

func init() {
	wordBo = entity.NewWordBo(db)
}

func WordRoute(router chi.Router) {
	router.Route("/", func(r chi.Router) {
		r.Use(middlewares.EnsureAuthenticatedJwtMw(db, util.AdminRole))
		r.Post("/", CreateWord)
		r.Put("/", UpdateWord)
		r.Delete("/{ID}", DeleteWord)
	})
	router.Route("/user", func(r chi.Router) {
		r.Use(middlewares.EnsureAuthenticatedJwtMw(db, util.UserRole))
		r.Post("/", AddWordToUser)
		r.Get("/", GetWordsOfUser)
		r.Delete("/", RemoveWordFromUser)
	})
	router.Get("/{ID}", FindWord)
	router.Get("/all", FindWords)
}

func CreateWord(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateWordRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	ID := util.UserIDFromContext(r.Context())
	request.ManagerID = ID

	result, err := wordBo.CreateWord(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func UpdateWord(w http.ResponseWriter, r *http.Request) {
	var request dto.UpdateWordRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	result, err := wordBo.UpdateWord(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func DeleteWord(w http.ResponseWriter, r *http.Request) {
	IDString := chi.URLParam(r, "ID")
	ID, err := util.IDFromStr(IDString)
	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	result, err := wordBo.DeleteWord(r.Context(), ID)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func FindWord(w http.ResponseWriter, r *http.Request) {
	IDString := chi.URLParam(r, "ID")
	ID, err := util.IDFromStr(IDString)
	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	result, err := wordBo.FindWord(r.Context(), ID)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func FindWords(w http.ResponseWriter, r *http.Request) {
	var request dto.PaginationRequest
	decoder.Decode(&request, r.URL.Query())

	result, err := wordBo.FindWords(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func AddWordToUser(w http.ResponseWriter, r *http.Request) {

	var request dto.AddWordToUser
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	ID := util.UserIDFromContext(r.Context())
	request.UserID = ID

	result, err := wordBo.AddWordToUser(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func RemoveWordFromUser(w http.ResponseWriter, r *http.Request) {

	var request dto.AddWordToUser
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	ID := util.UserIDFromContext(r.Context())
	request.UserID = ID

	result, err := wordBo.RemoveWordFromUser(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func GetWordsOfUser(w http.ResponseWriter, r *http.Request) {

	ID := util.UserIDFromContext(r.Context())

	result, err := wordBo.GetWordsOfUser(r.Context(), ID)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}
