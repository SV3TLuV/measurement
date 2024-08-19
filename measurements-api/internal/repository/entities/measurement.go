package entities

import (
	"time"
)

type Measurement struct {
	ID                   uint64 `db:"measurement_id"`
	ObjectID             uint64 `db:"object_id"`
	Object               *Object
	ObjName              *string    `db:"obj_name"`
	ObjNum               *string    `db:"obj_num"`
	Changed              *time.Time `db:"changed"`
	Created              time.Time  `db:"created"`
	RealDateTime         *time.Time `db:"real_date_time"`
	DateTime             *string    `db:"date_time"`
	Temp                 *float32   `db:"temp"`
	Pressure             *float32   `db:"pressure"`
	WindDir              *int       `db:"wind_dir"`
	WindDirStr           *string    `db:"wind_dir_str"`
	WindSpeed            *float32   `db:"wind_speed"`
	Humid                *float32   `db:"humid"`
	WaterVaporElasticity *float32   `db:"water_vapor_elasticity"`
	AtmPhenom            *float32   `db:"atm_phenom"`
	HumidInt             *float32   `db:"humid_int"`
	TempInt              *float32   `db:"temp_int"`

	V202917 *float32 `db:"v_202917"`
	M202917 *float32 `db:"m_202917"`
	Q202917 *string  `db:"q_202917"`

	V202918 *float32 `db:"v_202918"`
	M202918 *float32 `db:"m_202918"`
	Q202918 *string  `db:"q_202918"`

	V202919 *float32 `db:"v_202919"`
	M202919 *float32 `db:"m_202919"`
	Q202919 *string  `db:"q_202919"`

	V202920 *float32 `db:"v_202920"`
	M202920 *float32 `db:"m_202920"`
	Q202920 *string  `db:"q_202920"`

	V202921 *float32 `db:"v_202921"`
	M202921 *float32 `db:"m_202921"`
	Q202921 *string  `db:"q_202921"`

	V202932 *float32 `db:"v_202932"`
	M202932 *float32 `db:"m_202932"`
	Q202932 *string  `db:"q_202932"`

	V202935 *float32 `db:"v_202935"`
	M202935 *float32 `db:"m_202935"`
	Q202935 *string  `db:"q_202935"`

	V202924 *float32 `db:"v_202924"`
	M202924 *float32 `db:"m_202924"`
	Q202924 *string  `db:"q_202924"`

	V202925 *float32 `db:"v_202925"`
	M202925 *float32 `db:"m_202925"`
	Q202925 *string  `db:"q_202925"`

	V203565 *float32 `db:"v_203565"`
	M203565 *float32 `db:"m_203565"`
	Q203565 *string  `db:"q_203565"`

	V209190 *float32 `db:"v_209190"`
	M209190 *float32 `db:"m_209190"`
	Q209190 *string  `db:"q_209190"`

	V203570 *float32 `db:"v_203570"`
	M203570 *float32 `db:"m_203570"`
	Q203570 *string  `db:"q_203570"`

	V203551 *float32 `db:"v_203551"`
	M203551 *float32 `db:"m_203551"`
	Q203551 *string  `db:"q_203551"`

	V202936 *float32 `db:"v_202936"`
	M202936 *float32 `db:"m_202936"`
	Q202936 *string  `db:"q_202936"`

	V203569 *float32 `db:"v_203569"`
	M203569 *float32 `db:"m_203569"`
	Q203569 *string  `db:"q_203569"`

	V203557 *float32 `db:"v_203557"`
	M203557 *float32 `db:"m_203557"`
	Q203557 *string  `db:"q_203557"`

	V203568 *float32 `db:"v_203568"`
	M203568 *float32 `db:"m_203568"`
	Q203568 *string  `db:"q_203568"`

	V203559 *float32 `db:"v_203559"`
	M203559 *float32 `db:"m_203559"`
	Q203559 *string  `db:"q_203559"`

	V203577 *float32 `db:"v_203577"`
	M203577 *float32 `db:"m_203577"`
	Q203577 *string  `db:"q_203577"`

	V211082 *float32 `db:"v_211082"`
	M211082 *float32 `db:"m_211082"`
	Q211082 *string  `db:"q_211082"`

	V202931 *float32 `db:"v_202931"`
	M202931 *float32 `db:"m_202931"`
	Q202931 *string  `db:"q_202931"`
}
