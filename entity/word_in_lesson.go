package entity

import (
	"database/sql"

)

type WordInLessonBo struct {
	db *sql.DB
}

func NewWordInLessonBo(db *sql.DB) *WordInLessonBo {
	return &WordInLessonBo{db: db}
}

