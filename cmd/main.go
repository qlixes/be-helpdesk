package main

import (
	"net/http"

	"github.com/qlixes/helpdesk/internal/infrastructure/server"
)

func main() {
	mux := http.NewServeMux()

	server.New(mux)

	http.ListenAndServe(":8000", mux)
}
