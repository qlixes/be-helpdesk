package bootstrap

import (
	"net/http"

	"github.com/qlixes/be-helpdesk/web/handlers"
)

type Http struct {
	mux *http.ServeMux
}

type HttpManager interface {
	ListenAndServe() error
	Routes()
}

func NewHttp() *Http {

	mux := http.NewServeMux()

	return &Http{
		mux: mux,
	}
}

func (h *Http) ListenAndServe(port string) error {

	return http.ListenAndServe(port, h.mux)
}

func (h *Http) Routes() {
	http.HandleFunc("POST /", handlers.GetHelloWorld)
}
