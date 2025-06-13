package config

import (
	"net/http"

	"github.com/qlixes/be-helpdesk/web/handlers"
)

func Routes(h *http.ServeMux) {
	h.HandleFunc("POST /", handlers.GetHelloWorld)
	h.HandleFunc("POST /index", handlers.GetHelloWorld)
}
