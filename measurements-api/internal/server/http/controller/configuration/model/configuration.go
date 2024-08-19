package model

type ConfigurationView struct {
	AsoizaLogin        *string `json:"asoizaLogin"`
	AsoizaPassword     *string `json:"asoizaPassword"`
	CollectingInterval uint64  `json:"collectingInterval"`
	DeletingInterval   uint64  `json:"deletingInterval"`
	DeletingThreshold  uint64  `json:"deletingThreshold"`
	DisablingInterval  uint64  `json:"disablingInterval"`
	DisablingThreshold uint64  `json:"disablingThreshold"`
}

type UpdateConfigurationRequest struct {
	AsoizaLogin        *string `form:"asoizaLogin"`
	AsoizaPassword     *string `form:"asoizaPassword"`
	CollectingInterval uint64  `form:"collectingInterval"`
	DeletingInterval   uint64  `form:"deletingInterval"`
	DeletingThreshold  uint64  `form:"deletingThreshold"`
	DisablingInterval  uint64  `form:"disablingInterval"`
	DisablingThreshold uint64  `form:"disablingThreshold"`
}
