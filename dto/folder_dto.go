package dto

import "el.com/m/models"

type CreateFolderRequest struct {
	Name   string `json:"name" validate:"required"`
	UserId uint
}

type UpdateFolderRequest struct {
	ID     uint   `json:"id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	UserId uint
}

type DeleteFolderRequest struct {
	ID     uint
	UserId uint
}

type FindFoldersRequest struct {
	PaginationRequest
	UserId uint
}

type FindFoldersResponse struct {
	Total uint                `json:"total"`
	Data  *models.FolderSlice `json:"data"`
}

type FindFolderWithSaveRequest struct {
  WordID uint `form:"wordID"`
  UserID uint 
}

type FolderWithSave struct {
	ID    string `json:"id" boil:"id"`
	Name  string `json:"name" boil:"name"`
	Saved bool   `json:"saved" boil:"saved"`
}

type AddWordToFolder struct {
	WordID   uint `json:"wordID" validate:"required"`
	FolderId uint `json:"folderId" validate:"required"`
	UserId   uint
}

type GetWordRequest struct {
	FolderId uint `json:"folderId" validate:"required"`
	UserId   uint
}
