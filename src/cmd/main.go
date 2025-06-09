package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/qlixes/be-helpdesk/internal/interfaces/bootstrap"
)

func main() {

	conf := godotenv.Load()

	if conf != nil {
		log.Fatal(conf.Error())
	}

	app := bootstrap.NewApplication()

	err := http.ListenAndServe(":8000", app.GetInstance())

	if err != nil {
		log.Fatal(err.Error())
	}
}
