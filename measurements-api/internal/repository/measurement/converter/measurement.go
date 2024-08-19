package converter

import (
	"github.com/doug-martin/goqu/v9"
	"measurements-api/internal/model"
	"measurements-api/internal/repository/entities"
)

func ToMeasurementRecordFromService(measurement *model.Measurement) *goqu.Record {
	if measurement == nil {
		return nil
	}

	return &goqu.Record{
		"measurement_id":         measurement.ID,
		"object_id":              measurement.ObjectID,
		"obj_name":               measurement.ObjName,
		"obj_num":                measurement.ObjNum,
		"created":                measurement.Created,
		"changed":                measurement.Changed,
		"date_time":              measurement.DateTime,
		"real_date_time":         measurement.RealDateTime,
		"temp":                   measurement.Temp,
		"pressure":               measurement.Pressure,
		"wind_dir":               measurement.WindDir,
		"wind_dir_str":           measurement.WindDirStr,
		"wind_speed":             measurement.WindSpeed,
		"humid":                  measurement.Humid,
		"water_vapor_elasticity": measurement.WaterVaporElasticity,
		"atm_phenom":             measurement.AtmPhenom,
		"humid_int":              measurement.HumidInt,
		"temp_int":               measurement.TempInt,
		"v_202917":               measurement.V202917,
		"m_202917":               measurement.M202917,
		"q_202917":               measurement.Q202917,
		"v_202918":               measurement.V202918,
		"m_202918":               measurement.M202918,
		"q_202918":               measurement.Q202918,
		"v_202919":               measurement.V202919,
		"m_202919":               measurement.M202919,
		"q_202919":               measurement.Q202919,
		"v_202920":               measurement.V202920,
		"m_202920":               measurement.M202920,
		"q_202920":               measurement.Q202920,
		"v_202921":               measurement.V202921,
		"m_202921":               measurement.M202921,
		"q_202921":               measurement.Q202921,
		"v_202932":               measurement.V202932,
		"m_202932":               measurement.M202932,
		"q_202932":               measurement.Q202932,
		"v_202935":               measurement.V202935,
		"m_202935":               measurement.M202935,
		"q_202935":               measurement.Q202935,
		"v_202924":               measurement.V202924,
		"m_202924":               measurement.M202924,
		"q_202924":               measurement.Q202924,
		"v_202925":               measurement.V202925,
		"m_202925":               measurement.M202925,
		"q_202925":               measurement.Q202925,
		"v_203565":               measurement.V203565,
		"m_203565":               measurement.M203565,
		"q_203565":               measurement.Q203565,
		"v_209190":               measurement.V209190,
		"m_209190":               measurement.M209190,
		"q_209190":               measurement.Q209190,
		"v_203570":               measurement.V203570,
		"m_203570":               measurement.M203570,
		"q_203570":               measurement.Q203570,
		"v_203551":               measurement.V203551,
		"m_203551":               measurement.M203551,
		"q_203551":               measurement.Q203551,
		"v_202936":               measurement.V202936,
		"m_202936":               measurement.M202936,
		"q_202936":               measurement.Q202936,
		"v_203569":               measurement.V203569,
		"m_203569":               measurement.M203569,
		"q_203569":               measurement.Q203569,
		"v_203557":               measurement.V203557,
		"m_203557":               measurement.M203557,
		"q_203557":               measurement.Q203557,
		"v_203568":               measurement.V203568,
		"m_203568":               measurement.M203568,
		"q_203568":               measurement.Q203568,
		"v_203559":               measurement.V203559,
		"m_203559":               measurement.M203559,
		"q_203559":               measurement.Q203559,
		"v_203577":               measurement.V203577,
		"m_203577":               measurement.M203577,
		"q_203577":               measurement.Q203577,
		"v_211082":               measurement.V211082,
		"m_211082":               measurement.M211082,
		"q_211082":               measurement.Q211082,
		"v_202931":               measurement.V202931,
		"m_202931":               measurement.M202931,
		"q_202931":               measurement.Q202931,
	}
}

func ToMeasurementRecordsFromService(measurements []*model.Measurement) []*goqu.Record {
	records := make([]*goqu.Record, 0, len(measurements))
	for i := 0; i < len(measurements); i++ {
		records = append(records, ToMeasurementRecordFromService(measurements[i]))
	}
	return records
}

func ToMeasurementFromRepo(measurement *entities.Measurement) *model.Measurement {
	if measurement == nil {
		return nil
	}

	return &model.Measurement{
		ID:                   measurement.ID,
		ObjectID:             measurement.ObjectID,
		ObjName:              measurement.ObjName,
		ObjNum:               measurement.ObjNum,
		Changed:              measurement.Changed,
		Created:              measurement.Created,
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

func ToMeasurementsFromRepo(measurements []*entities.Measurement) []*model.Measurement {
	models := make([]*model.Measurement, 0, len(measurements))
	for i := 0; i < len(measurements); i++ {
		models = append(models, ToMeasurementFromRepo(measurements[i]))
	}
	return models
}
