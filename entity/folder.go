package entity

import "database/sql"

type FolderBo struct {
	db *sql.DB
}

func NewFolderBo(db *sql.DB) *FolderBo {
	return &FolderBo{db: db}
}
