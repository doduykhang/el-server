package main

import (
	"net/http"

	"el.com/m/server"
)

func main() {
	server.Api.Route("/user", server.UserServer)
	http.ListenAndServe(":3000", server.Api)
}
