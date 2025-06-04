package handlers

import "net/http"

type UserHandler interface {
	Signupr(w http.ResponseWriter, r *http.Request)
	Signin(w http.ResponseWriter, r *http.Request)
	Forget(w http.ResponseWriter, r *http.Request)
	Password(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Activate(w http.ResponseWriter, r *http.Request)
}

func Signupr(w http.ResponseWriter, r *http.Request) {

}

func Signin(w http.ResponseWriter, r *http.Request) {

}

func Forget(w http.ResponseWriter, r *http.Request) {

}

func Password(w http.ResponseWriter, r *http.Request) {

}

func Update(w http.ResponseWriter, r *http.Request) {

}

func Activate(w http.ResponseWriter, r *http.Request) {

}
