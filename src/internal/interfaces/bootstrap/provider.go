package bootstrap

import (
	"github.com/qlixes/be-helpdesk/pkg/shared/facades"
)

type Provider struct {
	Cache      *facades.Cache
	Queue      *facades.Queue
	Mail       *facades.Mail
	Database   *facades.Database
	Validation *facades.Validation
}

type ProviderManager interface {
}

func NewProvider() *Provider {
	return &Provider{}
}
