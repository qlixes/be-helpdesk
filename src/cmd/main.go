package main

import (
	"net/http"

	"github.com/qlixes/be-helpdesk/internal/interfaces/bootstrap"
	"github.com/spf13/viper"
)

func init() {

}

func main() {

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	conf := viper.ReadInConfig()

	if conf != nil {

	}

	app := bootstrap.NewApplication()

	err := http.ListenAndServe(":8000", app.Mux)

	if err != nil {

	}
}
