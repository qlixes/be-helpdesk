package main

import (
	"log"
	"net/http"

	"github.com/qlixes/be-helpdesk/internal/interfaces/bootstrap"
)

func main() {
	app := bootstrap.NewApplication()

	err := http.ListenAndServe(":8000")

	if err != nil {
		log.Fatalln(err.Error())
	}
}
