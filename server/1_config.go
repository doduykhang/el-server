package server

import (
	"database/sql"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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
	Api.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))
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
