package model

type Configuration struct {
	ID                 uint64
	AsoizaLogin        *string
	AsoizaPassword     *string
	CollectingInterval uint64
	DeletingInterval   uint64
	DeletingThreshold  uint64
	DisablingInterval  uint64
	DisablingThreshold uint64
}
