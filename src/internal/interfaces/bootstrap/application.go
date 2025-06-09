package bootstrap

import (
	"net/http"
	"sync"

	"github.com/qlixes/be-helpdesk/web/handlers"
)

var (
	once sync.Once
	mux  *http.ServeMux
)

type ApplicationManager interface {
	GetInstance() *Application
}

type Application struct {
	mux *http.ServeMux
}

func NewApplication() *Application {

	once.Do(func() {
		mux = http.NewServeMux()
	})

	// routing
	mux.HandleFunc("/", handlers.GetHelloWorld)

	return &Application{
		mux: mux,
	}
}

func (m *Application) GetInstance() *http.ServeMux {
	return m.mux
}
