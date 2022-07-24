package server

import (
	"database/sql"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/form"
	_ "github.com/lib/pq"
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
	db, err = sql.Open("postgres", "dbname=HOC_TU_VUNG_TIENG_ANH user=khang password=123")
	if err != nil {
		fmt.Println(err)
		panic("can not connect to db")
	} else {
		boil.SetDB(db)
	}
	decoder = form.NewDecoder()
}
