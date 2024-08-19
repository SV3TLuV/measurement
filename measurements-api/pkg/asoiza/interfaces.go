package asoiza

import "context"

type Client interface {
	GetMeasurements(ctx context.Context, objectID uint64, limit int, offset int) ([]*Measurement, error)
	GetNewestMeasurement(ctx context.Context, objectID uint64) (*Measurement, error)
	GetNavTree(ctx context.Context, node string) ([]*Node, error)
}

type ConfigurationFactory interface {
	Get() *Configuration
}
