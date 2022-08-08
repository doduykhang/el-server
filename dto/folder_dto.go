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
