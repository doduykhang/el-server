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
		r.Put("/", UpdateWord)
		r.Delete("/{ID}", DeleteWord)
	})
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
