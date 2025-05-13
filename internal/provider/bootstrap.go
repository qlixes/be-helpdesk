package provider

import "sync"

var bootstrap *Bootstrap
var mutex sync.Mutex
var once sync.Once

type Bootstrap struct {
	Config *ConfigMethod
	Db *DatabaseMethod
}

type BootstrapMethod interface {

}

func NewBootstrap() *Bootstrap {

	once.Do(func() {
		bootstrap = &Bootstrap{}
	})

	return bootstrap
}