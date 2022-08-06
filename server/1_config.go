package server

import (
	"database/sql"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/form"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	Api     *chi.Mux
	decoder *form.Decoder
	db      *sql.DB
)

func init() {
	Api = chi.NewRouter()
	var err error
	db, err = sql.Open("mysql", "khang:bao123@tcp(127.0.0.1:3306)/el?parseTime=true")
	if err != nil {
		fmt.Println(err)
		panic("can not connect to db")
	} else {
		boil.SetDB(db)
	}
	decoder = form.NewDecoder()
}
