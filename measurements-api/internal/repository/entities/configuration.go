package entities

type Configuration struct {
	ID                 uint64  `db:"configuration_id"`
	AsoizaLogin        *string `db:"asoiza_login"`
	AsoizaPassword     *string `db:"asoiza_password"`
	CollectingInterval uint64  `db:"collecting_interval"`
	DeletingInterval   uint64  `db:"deleting_interval"`
	DeletingThreshold  uint64  `db:"deleting_threshold"`
	DisablingInterval  uint64  `db:"disabling_interval"`
	DisablingThreshold uint64  `db:"disabling_threshold"`
}
