package bootstrap

type Provider struct {
}

type ProviderManager interface {
}

func NewProvider() *Provider {
	return &Provider{}
}
