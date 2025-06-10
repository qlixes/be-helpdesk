package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/qlixes/be-helpdesk/internal/interfaces/bootstrap"
)

func main() {

	conf := godotenv.Load()

	if conf != nil {
		log.Fatalln(conf.Error())
	}

	app := bootstrap.NewApplication()

	err := app.Run()

	if err != nil {
		log.Fatalln(err.Error())
	}
}
