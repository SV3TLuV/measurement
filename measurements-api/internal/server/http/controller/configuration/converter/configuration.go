package converter

import (
	"measurements-api/internal/model"
	model2 "measurements-api/internal/server/http/controller/configuration/model"
)

func ToConfigurationViewFromService(configuration *model.Configuration) *model2.ConfigurationView {
	return &model2.ConfigurationView{
		AsoizaLogin:        configuration.AsoizaLogin,
		AsoizaPassword:     configuration.AsoizaPassword,
		CollectingInterval: configuration.CollectingInterval,
		DeletingInterval:   configuration.DeletingInterval,
		DeletingThreshold:  configuration.DeletingThreshold,
		DisablingInterval:  configuration.DisablingInterval,
		DisablingThreshold: configuration.DisablingThreshold,
	}
}

func ToConfigurationFromRequest(request *model2.UpdateConfigurationRequest) *model.Configuration {
	return &model.Configuration{
		AsoizaLogin:        request.AsoizaLogin,
		AsoizaPassword:     request.AsoizaPassword,
		CollectingInterval: request.CollectingInterval,
		DeletingInterval:   request.DeletingInterval,
		DeletingThreshold:  request.DeletingThreshold,
		DisablingInterval:  request.DisablingInterval,
		DisablingThreshold: request.DisablingThreshold,
	}
}
