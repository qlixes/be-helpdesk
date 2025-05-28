package handlers

import (
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello world"))
}
