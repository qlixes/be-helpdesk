package provider

import (
	"fmt"
	"sync"
)

var bootstrap *Bootstrap

var mutex sync.Mutex
var once sync.Once

type Bootstrap struct {
	Config *Config
	Database *Database
}

type BootstrapMethod interface {
}

func NewBootstrap() *Bootstrap {

//	once.Do(func() {
		bootstrap = &Bootstrap{}
		bootstrap.Config = NewConfig()
		bootstrap.Database = NewPgsql() // default database used
//	})

	return bootstrap
}