package bootstrap

import (
	"sync"
)

var (
	once sync.Once
	app  Application
)

type Application struct {
	http *Http
	// db     *Database
	// cache  *Cache
	// mail   *Mail
	config *Config
	// queue  *Queue
}

type ApplicationManager interface {
	Run() error
}

func NewApplication() *Application {

	once.Do(func() {
		app = Application{
			http: NewHttp(),
			// db:     NewDatabase(),
			// cache:  NewCache(),
			// mail:   NewMail(),
			config: NewConfig(),
			// queue:  NewQueue(),
		}
	})

	return &app
}

func (a *Application) Run() error {
	return a.http.ListenAndServe(a.config.GetAppPort())
}
