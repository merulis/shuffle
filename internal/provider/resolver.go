package provider

import (
	"fmt"

	s "github.com/merulis/shuffle/internal/source"
)

type ConfigProvider interface {
	Get(name string) (s.Source, error)
}

type Resolver struct {
	configProvider  ConfigProvider
	providerFactory Factory
}

func NewResolver(provider ConfigProvider, factory ProviderFactory) *Resolver {
	return &Resolver{
		configProvider:  provider,
		providerFactory: factory,
	}
}

func (r *Resolver) Resolve(sourceRef s.SourceRef) (Provider, s.Source, error) {
	source, err := r.configProvider.Get(sourceRef.Name)
	if err != nil {
		return nil, s.Source{}, fmt.Errorf("get source config: %w", err)
	}

	src, err := r.providerFactory.NewProvider(source)
	if err != nil {
		return nil, s.Source{}, fmt.Errorf("create source: %w", err)
	}

	return src, source, nil
}
