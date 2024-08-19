package exporter

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
)

type xslxExporter struct{}

func (e *xslxExporter) Export(data [][]string) ([]byte, error) {
	f := excelize.NewFile()
	sheetName := "Sheet1"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return nil, errors.Wrap(err, "create sheet")
	}

	for i, row := range data {
		for j, cell := range row {
			cellName, err := excelize.CoordinatesToCellName(j+1, i+1)
			if err != nil {
				return nil, errors.Wrap(err, "failed to convert to cell name")
			}

			err = f.SetCellValue(sheetName, cellName, cell)
			if err != nil {
				return nil, errors.Wrap(err, "failed to set cell value")
			}
		}
	}

	f.SetActiveSheet(index)
	var buffer bytes.Buffer

	if err = f.Write(&buffer); err != nil {
		return nil, errors.Wrap(err, "failed to write")
	}

	return buffer.Bytes(), nil
}
