package bootstrap

import (
	"net/http"

	"github.com/qlixes/be-helpdesk/web/handlers"
)

type Application struct {
	mux *http.ServeMux
}

func NewApplication() *Application {

	mux := http.NewServeMux()

	// routing
	mux.HandleFunc("/", handlers.GetHelloWorld)

	return &Application{
		mux: mux,
	}
}
