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
	questionBo *entity.QuestionBo
)

func init() {
	questionBo = entity.NewQuestionBo(db)
}

func QuestionRoute(router chi.Router) {
	router.Route("/", func(r chi.Router) {
		r.Use(middlewares.EnsureAuthenticatedJwtMw(db, util.AdminRole))
		r.Post("/", CreateQuestion)
		r.Put("/", UpdateQuestion)
		r.Delete("/{ID}", DeleteQuestion)

	})

	router.Get("/options/{ID}", GetOptions)
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateQuestionRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	ID := util.UserIDFromContext(r.Context())
	request.ManagerID = ID

	result, err := questionBo.CreateQuestion(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	var request dto.UpdateQuestionRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	result, err := questionBo.UpdateQuestion(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	IDString := chi.URLParam(r, "ID")
	ID, err := util.IDFromStr(IDString)
	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	result, err := questionBo.DeleteQuestion(r.Context(), ID)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func GetOptions(w http.ResponseWriter, r *http.Request) {
	IDString := chi.URLParam(r, "ID")
	ID, err := util.IDFromStr(IDString)
	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	result, err := questionBo.GetOptions(r.Context(), ID)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}
