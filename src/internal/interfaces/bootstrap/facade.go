package bootstrap

import "github.com/qlixes/be-helpdesk/pkg/shared/facades"

type Facade struct {
	Db    *facades.Db
	Cache *facades.Cache
	Queue *facades.Queue
	Mail  *facades.Mail
}

type FacadeManager interface {
	SetConfig(config *Config)
}

func NewFacade() *Facade {
	return &Facade{}
}

func (f *Facade) SetConfig(config *Config) {

}
