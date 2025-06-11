package bootstrap

import (
	"net/http"

	"github.com/qlixes/be-helpdesk/internal/config"
)

type Http struct {
	mux *http.ServeMux
}

type HttpManager interface {
	ListenAndServe() error
}

func NewHttp() *Http {

	mux := http.NewServeMux()

	return &Http{
		mux: mux,
	}
}

func (h *Http) ListenAndServe(port string) error {

	config.Routes(h.mux)

	return http.ListenAndServe(port, h.mux)
}
