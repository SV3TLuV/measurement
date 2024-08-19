package exporter

import (
	"bytes"
	"encoding/csv"
	"github.com/pkg/errors"
)

type csvExporter struct{}

func (e *csvExporter) Export(data [][]string) ([]byte, error) {
	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)
	for i := 0; i < len(data); i++ {
		if err := writer.Write(data[i]); err != nil {
			return nil, errors.Wrap(err, "failed to write to csv")
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, errors.Wrap(err, "error flushing csv writer")
	}

	return buffer.Bytes(), nil
}
