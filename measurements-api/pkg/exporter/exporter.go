package exporter

import "measurements-api/internal/model"

type Exporter interface {
	Export([][]string) ([]byte, error)
}

func NewExporter(format model.ExportFormat) Exporter {
	switch format {
	case model.Xlsx:
		return &xslxExporter{}
	case model.Csv:
		return &csvExporter{}
	default:
		return nil
	}
}
