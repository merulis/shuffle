package source

import (
	"fmt"

	"github.com/merulis/shuffle/internal/domain"
)

type ConfigProvider interface {
	Get(name string) (Source, error)
}

type ProviderFactory interface {
	NewProvider(source Source) (domain.Source, error)
}

type Resolver struct {
	configProvider  ConfigProvider
	providerFactory ProviderFactory
}

func NewResolver(provider ConfigProvider, factory ProviderFactory) *Resolver {
	return &Resolver{
		configProvider:  provider,
		providerFactory: factory,
	}
}

func (r *Resolver) Resolve(sourceRef SourceRef) (domain.Source, Source, error) {
	source, err := r.configProvider.Get(sourceRef.Name)
	if err != nil {
		return nil, Source{}, fmt.Errorf("get source config: %w", err)
	}

	src, err := r.providerFactory.NewProvider(source)
	if err != nil {
		return nil, Source{}, fmt.Errorf("create source: %w", err)
	}

	return src, source, nil
}
