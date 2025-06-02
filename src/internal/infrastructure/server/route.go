package server

import (
	"net/http"

	"github.com/qlixes/be-helpdesk/web/handler"
)

type Handler struct {
	mux *http.ServeMux
}

func New(s *http.ServeMux) *Handler {
	h := Handler{s}
	h.routes()

	return &Handler{}
}

func (h *Handler) routes() {
	h.mux.HandleFunc("/", handler.HelloWorld)
}
