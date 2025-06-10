package bootstrap

import (
	"net/http"
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

	return http.ListenAndServe(port, h.mux)
}
