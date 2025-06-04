package handlers

import (
	"net/http"
)

type HelloWorldHandler interface {
	Helloworld(w http.ResponseWriter, r *http.Request)
}

func GetHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello world"))
}
