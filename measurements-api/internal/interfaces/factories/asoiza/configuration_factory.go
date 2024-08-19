package asoiza

import (
	"context"
	"measurements-api/internal/service"
	"measurements-api/pkg/asoiza"
)

var _ asoiza.ConfigurationFactory = (*configurationFactory)(nil)

type configurationFactory struct {
	configurationService service.ConfigurationService
}

func NewConfigurationFactory(service service.ConfigurationService) *configurationFactory {
	return &configurationFactory{
		configurationService: service,
	}
}

func (f *configurationFactory) Get() *asoiza.Configuration {
	configuration, err := f.configurationService.Get(context.Background())
	if err != nil {
		return nil
	}
	if configuration.AsoizaPassword == nil {
		return nil
	}
	if configuration.AsoizaPassword == nil {
		return nil
	}

	return &asoiza.Configuration{
		Username: *configuration.AsoizaLogin,
		Password: *configuration.AsoizaPassword,
	}
}
