package converter

import (
	"measurements-api/internal/model"
	"measurements-api/pkg/asoiza"
)

func ToMeasurementFromAsoiza(measurement *asoiza.Measurement) *model.Measurement {
	if measurement == nil {
		return nil
	}

	return &model.Measurement{
		ID:                   measurement.ID,
		ObjectID:             measurement.ObjectID,
		ObjName:              measurement.ObjName,
		ObjNum:               measurement.ObjNum,
		Changed:              &measurement.Changed.Time,
		Created:              measurement.Created.Time,
		RealDateTime:         measurement.RealDateTime,
		DateTime:             measurement.DateTime,
		Temp:                 measurement.Temp,
		Pressure:             measurement.Pressure,
		WindDir:              measurement.WindDir,
		WindDirStr:           measurement.WindDirStr,
		WindSpeed:            measurement.WindSpeed,
		Humid:                measurement.Humid,
		WaterVaporElasticity: measurement.WaterVaporElasticity,
		AtmPhenom:            measurement.AtmPhenom,
		HumidInt:             measurement.HumidInt,
		TempInt:              measurement.TempInt,
		V202917:              measurement.V202917,
		M202917:              measurement.M202917,
		Q202917:              measurement.Q202917,
		V202918:              measurement.V202918,
		M202918:              measurement.M202918,
		Q202918:              measurement.Q202918,
		V202919:              measurement.V202919,
		M202919:              measurement.M202919,
		Q202919:              measurement.Q202919,
		V202920:              measurement.V202920,
		M202920:              measurement.M202920,
		Q202920:              measurement.Q202920,
		V202921:              measurement.V202921,
		M202921:              measurement.M202921,
		Q202921:              measurement.Q202921,
		V202932:              measurement.V202932,
		M202932:              measurement.M202932,
		Q202932:              measurement.Q202932,
		V202935:              measurement.V202935,
		M202935:              measurement.M202935,
		Q202935:              measurement.Q202935,
		V202924:              measurement.V202924,
		M202924:              measurement.M202924,
		Q202924:              measurement.Q202924,
		V202925:              measurement.V202925,
		M202925:              measurement.M202925,
		Q202925:              measurement.Q202925,
		V203565:              measurement.V203565,
		M203565:              measurement.M203565,
		Q203565:              measurement.Q203565,
		V209190:              measurement.V209190,
		M209190:              measurement.M209190,
		Q209190:              measurement.Q209190,
		V203570:              measurement.V203570,
		M203570:              measurement.M203570,
		Q203570:              measurement.Q203570,
		V203551:              measurement.V203551,
		M203551:              measurement.M203551,
		Q203551:              measurement.Q203551,
		V202936:              measurement.V202936,
		M202936:              measurement.M202936,
		Q202936:              measurement.Q202936,
		V203569:              measurement.V203569,
		M203569:              measurement.M203569,
		Q203569:              measurement.Q203569,
		V203557:              measurement.V203557,
		M203557:              measurement.M203557,
		Q203557:              measurement.Q203557,
		V203568:              measurement.V203568,
		M203568:              measurement.M203568,
		Q203568:              measurement.Q203568,
		V203559:              measurement.V203559,
		M203559:              measurement.M203559,
		Q203559:              measurement.Q203559,
		V203577:              measurement.V203577,
		M203577:              measurement.M203577,
		Q203577:              measurement.Q203577,
		V211082:              measurement.V211082,
		M211082:              measurement.M211082,
		Q211082:              measurement.Q211082,
		V202931:              measurement.V202931,
		M202931:              measurement.M202931,
		Q202931:              measurement.Q202931,
	}
}

func ToMeasurementsFromAsoiza(measurements []*asoiza.Measurement) []*model.Measurement {
	models := make([]*model.Measurement, 0, len(measurements))
	for i := 0; i < len(measurements); i++ {
		models = append(models, ToMeasurementFromAsoiza(measurements[i]))
	}
	return models
}
