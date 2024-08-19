package converter

import (
	"github.com/doug-martin/goqu/v9"
	"measurements-api/internal/model"
	"measurements-api/internal/repository/entities"
)

func ToConfigurationFromRepo(configuration *entities.Configuration) *model.Configuration {
	return &model.Configuration{
		ID:                 configuration.ID,
		AsoizaLogin:        configuration.AsoizaLogin,
		AsoizaPassword:     configuration.AsoizaPassword,
		CollectingInterval: configuration.CollectingInterval,
		DeletingInterval:   configuration.DeletingInterval,
		DeletingThreshold:  configuration.DeletingThreshold,
		DisablingInterval:  configuration.DisablingInterval,
		DisablingThreshold: configuration.DisablingThreshold,
	}
}

func ToConfigurationRecordFromService(configuration *model.Configuration) *goqu.Record {
	return &goqu.Record{
		"configuration_id":    configuration.ID,
		"asoiza_login":        configuration.AsoizaLogin,
		"asoiza_password":     configuration.AsoizaPassword,
		"collecting_interval": configuration.CollectingInterval,
		"deleting_interval":   configuration.DeletingInterval,
		"deleting_threshold":  configuration.DeletingThreshold,
		"disabling_interval":  configuration.DisablingInterval,
		"disabling_threshold": configuration.DisablingThreshold,
	}
}
