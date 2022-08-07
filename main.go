package main

import (
	"net/http"

	"el.com/m/server"
	"github.com/go-chi/chi/v5"
)

func main() {
	server.Api.Route("/", server.UserServer)
	server.Api.Route("/api", func(r chi.Router) {
		r.Route("/word", server.WordRoute)
		r.Route("/lesson", server.LessonRoute)
		r.Route("/test", server.TestRoute)
		r.Route("/question", server.QuestionRoute)
	})
	http.ListenAndServe(":3000", server.Api)
}
