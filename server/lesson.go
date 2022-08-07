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
	lessonBo *entity.LessonBo
)

func init() {
	lessonBo = entity.NewLessonBo(db)
}

func LessonRoute(router chi.Router) {
	router.Route("/", func(r chi.Router) {
		r.Use(middlewares.EnsureAuthenticatedJwtMw(db, util.AdminRole))
		r.Post("/", CreateLesson)
		r.Put("/", UpdateLesson)
		r.Delete("/{ID}", DeleteLesson)

		r.Post("/add-word", AddWord)
		r.Delete("/remove-word", RemoveWord)
	})

	router.Get("/{ID}", FindLesson)
	router.Get("/all", FindLessons)
	router.Get("/get-words/{ID}", GetWords)
}

func CreateLesson(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateLessonRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	ID := util.UserIDFromContext(r.Context())
	request.ManagerID = ID

	result, err := lessonBo.CreateLesson(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func UpdateLesson(w http.ResponseWriter, r *http.Request) {
	var request dto.UpdateLessonRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	result, err := lessonBo.UpdateLesson(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func DeleteLesson(w http.ResponseWriter, r *http.Request) {

	IDString := chi.URLParam(r, "ID")
	ID, err := util.IDFromStr(IDString)

	result, err := lessonBo.DeleteLesson(r.Context(), ID)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func FindLesson(w http.ResponseWriter, r *http.Request) {

	IDString := chi.URLParam(r, "ID")
	ID, err := util.IDFromStr(IDString)

	result, err := lessonBo.FindLesson(r.Context(), ID)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func FindLessons(w http.ResponseWriter, r *http.Request) {

	var request dto.PaginationRequest
	decoder.Decode(&request, r.URL.Query())

	result, err := lessonBo.FindLessons(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func AddWord(w http.ResponseWriter, r *http.Request) {

	var request dto.AddWordToLesson
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	result, err := lessonBo.AddWord(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func RemoveWord(w http.ResponseWriter, r *http.Request) {

	var request dto.AddWordToLesson
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	result, err := lessonBo.RemoveWord(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func GetWords(w http.ResponseWriter, r *http.Request) {

	IDString := chi.URLParam(r, "ID")
	ID, err := util.IDFromStr(IDString)

	result, err := lessonBo.GetWords(r.Context(), ID)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}
