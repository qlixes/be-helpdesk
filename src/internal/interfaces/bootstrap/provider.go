package bootstrap

import "github.com/qlixes/be-helpdesk/pkg/shared"

type Provider struct {
	Config *shared.Config
}

func NewProvider() *Provider {
	return &Provider{}
}
