package bootstrap

import (
	"sync"
)

var (
	app  Application
	once sync.Once
)

type Application struct {
	http     *Http
	config   *Config
	provider *Provider
	facade   *Facade
}

type ApplicationManager interface {
	Run() error
}

func NewApplication() *Application {

	once.Do(func() {
		app = Application{
			http:     NewHttp(),
			config:   NewConfig(),
			provider: NewProvider(),
			facade:   NewFacade(),
		}
	})

	return &app
}

func (a *Application) Run() error {

	return a.http.ListenAndServe(a.config.GetAppPort())
}
