package provider

import "sync"

var bootstrap *Bootstrap

var mutex sync.Mutex
var once sync.Once

type Bootstrap struct {
	Config *Config
	Db *Database
}

type BootstrapMethod interface {

}

func NewBootstrap() *Bootstrap {

	once.Do(func() {
		bootstrap = &Bootstrap{
			Db: NewDatabase(),
			Config: NewConfig(),
		}
	})

	return bootstrap
}