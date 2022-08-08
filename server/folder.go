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
	folderBo *entity.FolderBo
)

func init() {
	folderBo = entity.NewFolderBo(db)
}

func FolderRoute(router chi.Router) {
	router.Route("/", func(r chi.Router) {
		r.Use(middlewares.EnsureAuthenticatedJwtMw(db, util.UserRole))
		r.Post("/", CreateFolder)
		r.Put("/", UpdateFolder)
		r.Delete("/{ID}", DeleteFolder)
		r.Get("/{ID}", FindFolder)
		r.Get("/all", FindFolders)

		r.Post("/add", AddWordToFolder)
		r.Delete("/remove/{folderID}/{wordID}", RemoveWordFromFolder)

		r.Get("/words/{ID}", GetWordOfFolder)
	})
}

func CreateFolder(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateFolderRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	ID := util.UserIDFromContext(r.Context())
	request.UserId = ID

	result, err := folderBo.CreateFolder(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func UpdateFolder(w http.ResponseWriter, r *http.Request) {
	var request dto.UpdateFolderRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	ID := util.UserIDFromContext(r.Context())
	request.UserId = ID

	result, err := folderBo.UpdateFolder(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func DeleteFolder(w http.ResponseWriter, r *http.Request) {
	var request dto.DeleteFolderRequest

	ID := util.UserIDFromContext(r.Context())
	request.UserId = ID

	IDString := chi.URLParam(r, "ID")
	folderID, err := util.IDFromStr(IDString)
	request.ID = folderID

	result, err := folderBo.DeleteFolder(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func FindFolder(w http.ResponseWriter, r *http.Request) {
	var request dto.DeleteFolderRequest

	ID := util.UserIDFromContext(r.Context())
	request.UserId = ID

	IDString := chi.URLParam(r, "ID")
	folderID, err := util.IDFromStr(IDString)
	request.ID = folderID

	result, err := folderBo.FindFolder(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func FindFolders(w http.ResponseWriter, r *http.Request) {

	var request dto.FindFoldersRequest
	decoder.Decode(&request, r.URL.Query())

	ID := util.UserIDFromContext(r.Context())

	request.UserId = ID

	result, err := folderBo.FindFolders(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func AddWordToFolder(w http.ResponseWriter, r *http.Request) {

	var request dto.AddWordToFolder
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	ID := util.UserIDFromContext(r.Context())
	request.UserId = ID

	result, err := folderBo.AddWordToFolder(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func RemoveWordFromFolder(w http.ResponseWriter, r *http.Request) {

	var request dto.AddWordToFolder

	IDString := chi.URLParam(r, "folderID")
	folderID, err := util.IDFromStr(IDString)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	IDString = chi.URLParam(r, "wordID")
	wordID, err := util.IDFromStr(IDString)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	ID := util.UserIDFromContext(r.Context())
	request.UserId = ID
	request.FolderId = folderID
	request.WordID = wordID

	result, err := folderBo.RemoveToFolder(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}

func GetWordOfFolder(w http.ResponseWriter, r *http.Request) {

	var request dto.GetWordRequest

	IDString := chi.URLParam(r, "ID")
	folderID, err := util.IDFromStr(IDString)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	ID := util.UserIDFromContext(r.Context())
	request.UserId = ID
	request.FolderId = folderID

	result, err := folderBo.GetWordsOfFolder(r.Context(), request)

	if err != nil {
		util.SadResp(err, 500, w)
		return
	}

	util.JSONResp(result, 200, w)
}
