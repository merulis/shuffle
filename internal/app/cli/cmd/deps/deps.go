package cdeps

import (
	rconfiguration "github.com/merulis/shuffle/internal/app/repository/configuration"
	"github.com/merulis/shuffle/internal/app/service"
	sconfiguration "github.com/merulis/shuffle/internal/app/service/configuration"
	sprovider "github.com/merulis/shuffle/internal/app/service/provider"
)

type Deps struct {
	ConfigService service.Configuration
	Factory       *sprovider.Factory
}

func NewDeps() Deps {
	repo := rconfiguration.NewFileRepository("")
	configService := sconfiguration.NewService(repo)

	factory := sprovider.NewFactory()

	return Deps{
		ConfigService: configService,
		Factory:       factory,
	}
}
