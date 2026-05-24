package deps

import (
	"github.com/merulis/shuffle/internal/config"
	"github.com/merulis/shuffle/internal/provider/factory"
	"github.com/merulis/shuffle/internal/source"
)

type Deps struct {
	ConfigService  *config.Service
	SourceResolver *source.Resolver
}

func NewDeps() Deps {
	repo := config.NewFileRepository("")
	configService := config.NewService(repo)

	providerFactory := factory.NewFactory()
	sourceResolver := source.NewResolver(
		configService,
		providerFactory,
	)

	return Deps{
		ConfigService:  configService,
		SourceResolver: sourceResolver,
	}
}
