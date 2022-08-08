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
	testBo *entity.TestBo
)

func init() {
	testBo = entity.NewTestBo(db)
}

func TestRoute(router chi.Router) {
	router.Route("/", func(r chi.Router) {
		r.Use(middlewares.EnsureAuthenticatedJwtMw(db, util.AdminRole))
		r.Post("/", CreateTest)
		r.Put("/", UpdateTest)
		r.Put("/publish/{ID}", PublishTest)
		r.Put("/un-publish/{ID}", UnPublishTest)
		r.Delete("/{ID}", DeleteTest)
	})

	router.Route("/user", func(r chi.Router) {
		r.Use(middlewares.EnsureAuthenticatedJwtMw(db, util.UserRole))
		r.Post("/submit-test", SubmitTest)
	})

	router.Get("/{ID}", FindTest)
	router.Get("/questions/{ID}", GetQuestions)
	router.Get("/all", FindTests)
}

func CreateTest(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateTestRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	ID := util.UserIDFromContext(r.Context())
	request.ManagerID = ID

	result, err := testBo.CreateTest(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func UpdateTest(w http.ResponseWriter, r *http.Request) {
	var request dto.UpdateTestRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	result, err := testBo.UpdateTest(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func DeleteTest(w http.ResponseWriter, r *http.Request) {

	IDString := chi.URLParam(r, "ID")
	ID, err := util.IDFromStr(IDString)
	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	result, err := testBo.Delete(r.Context(), ID)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func FindTest(w http.ResponseWriter, r *http.Request) {

	IDString := chi.URLParam(r, "ID")
	ID, err := util.IDFromStr(IDString)
	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	result, err := testBo.FindTest(r.Context(), ID)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func FindTests(w http.ResponseWriter, r *http.Request) {
	var request dto.FindTestRequest
	decoder.Decode(&request, r.URL.Query())

	result, err := testBo.FindTests(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	IDString := chi.URLParam(r, "ID")
	ID, err := util.IDFromStr(IDString)

	result, err := testBo.GetQuestions(r.Context(), ID)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func PublishTest(w http.ResponseWriter, r *http.Request) {

	IDString := chi.URLParam(r, "ID")
	ID, err := util.IDFromStr(IDString)
	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	result, err := testBo.PublishTest(r.Context(), ID)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func UnPublishTest(w http.ResponseWriter, r *http.Request) {
	IDString := chi.URLParam(r, "ID")
	ID, err := util.IDFromStr(IDString)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	result, err := testBo.UnPublishTest(r.Context(), ID)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func SubmitTest(w http.ResponseWriter, r *http.Request) {
	var request dto.SubmitTestRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	ID := util.UserIDFromContext(r.Context())
	request.UserID = ID

	result, err := testBo.SubmitTest(r.Context(), request)


	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}
