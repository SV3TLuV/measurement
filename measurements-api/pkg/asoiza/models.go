package asoiza

import "time"

type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02 15:04:05.999999"

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := string(b[1 : len(b)-1])
	t, err := time.Parse(ctLayout, s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

type Configuration struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type BaseResponse struct {
	Success bool `json:"success"`
}

type NavStoreResponse struct {
	BaseResponse
	Total uint    `json:"total,string"`
	Data  []*Node `json:"data"`
}

type MeasurementResponse struct {
	BaseResponse
	Total uint           `json:"total,string"`
	Data  []*Measurement `json:"data"`
}

type Node struct {
	ID         string      `json:"id,string"`
	Created    CustomTime  `json:"created"`
	Changed    *CustomTime `json:"changed"`
	Title      string      `json:"title"`
	ObjectType string      `json:"object_type"` // labs, sys_localities, objects
	Address    *string     `json:"address"`
	LabID      *string     `json:"lab"`
	CityID     *string     `json:"locality"`
	Lat        *float64    `json:"lat,string"`
	Lon        *float64    `json:"lon,string"`
}

type Measurement struct {
	ID                   uint64      `json:"id,string"`
	ObjectID             uint64      `json:"obj,string"`
	ObjName              *string     `json:"obj_name"`
	ObjNum               *string     `json:"obj_num"`
	Created              CustomTime  `json:"created"`
	Changed              *CustomTime `json:"changed"`
	DateTime             *string     `json:"date_time"`
	RealDateTime         *time.Time  `json:"real_date_time"`
	Temp                 *float32    `json:"temp,string"`
	TempInt              *float32    `json:"temp_int,string"`
	Pressure             *float32    `json:"pressure,string"`
	WindDir              *int        `json:"wind_dir,string"`
	WindDirStr           *string     `json:"wind_dir_str,string"`
	WindSpeed            *float32    `json:"wind_speed,string"`
	Humid                *float32    `json:"humid,string"`
	HumidInt             *float32    `json:"humid_int,string"`
	WaterVaporElasticity *float32    `json:"water_vapor_elasticity,string"`
	AtmPhenom            *float32    `json:"atm_phenom,string"`

	V202917 *float32 `json:"v_202917,string"`
	Q202917 *string  `json:"q_202917"`
	M202917 *float32 `json:"m_202917,string"`

	V202918 *float32 `json:"v_202918,string"`
	Q202918 *string  `json:"q_202918"`
	M202918 *float32 `json:"m_202918,string"`

	V202919 *float32 `json:"v_202919,string"`
	Q202919 *string  `json:"q_202919"`
	M202919 *float32 `json:"m_202919,string"`

	V202920 *float32 `json:"v_202920,string"`
	Q202920 *string  `json:"q_202920"`
	M202920 *float32 `json:"m_202920,string"`

	V202921 *float32 `json:"v_202921,string"`
	Q202921 *string  `json:"q_202921"`
	M202921 *float32 `json:"m_202921,string"`

	V202932 *float32 `json:"v_202932,string"`
	Q202932 *string  `json:"q_202932"`
	M202932 *float32 `json:"m_202932,string"`

	V202935 *float32 `json:"v_202935,string"`
	Q202935 *string  `json:"q_202935"`
	M202935 *float32 `json:"m_202935,string"`

	V202924 *float32 `json:"v_202924,string"`
	Q202924 *string  `json:"q_202924"`
	M202924 *float32 `json:"m_202924,string"`

	V202925 *float32 `json:"v_202925,string"`
	Q202925 *string  `json:"q_202925"`
	M202925 *float32 `json:"m_202925,string"`

	V203565 *float32 `json:"v_203565,string"`
	Q203565 *string  `json:"q_203565"`
	M203565 *float32 `json:"m_203565,string"`

	V209190 *float32 `json:"v_209190,string"`
	Q209190 *string  `json:"q_209190"`
	M209190 *float32 `json:"m_209190,string"`

	V203570 *float32 `json:"v_203570,string"`
	Q203570 *string  `json:"q_203570"`
	M203570 *float32 `json:"m_203570,string"`

	V203551 *float32 `json:"v_203551,string"`
	Q203551 *string  `json:"q_203551"`
	M203551 *float32 `json:"m_203551,string"`

	V202936 *float32 `json:"v_202936,string"`
	Q202936 *string  `json:"q_202936"`
	M202936 *float32 `json:"m_202936,string"`

	V203569 *float32 `json:"v_203569,string"`
	Q203569 *string  `json:"q_203569"`
	M203569 *float32 `json:"m_203569,string"`

	V203557 *float32 `json:"v_203557,string"`
	Q203557 *string  `json:"q_203557"`
	M203557 *float32 `json:"m_203557,string"`

	V203568 *float32 `json:"v_203568,string"`
	Q203568 *string  `json:"q_203568"`
	M203568 *float32 `json:"m_203568,string"`

	V203559 *float32 `json:"v_203559,string"`
	Q203559 *string  `json:"q_203559"`
	M203559 *float32 `json:"m_203559,string"`

	V203577 *float32 `json:"v_203577,string"`
	Q203577 *string  `json:"q_203577"`
	M203577 *float32 `json:"m_203577,string"`

	V211082 *float32 `json:"v_211082,string"`
	Q211082 *string  `json:"q_211082"`
	M211082 *float32 `json:"m_211082,string"`

	V202931 *float32 `json:"v_202931,string"`
	Q202931 *string  `json:"q_202931"`
	M202931 *float32 `json:"m_202931,string"`
}
