package model

import "time"

type Period uint16

const (
	All Period = iota + 1
	Year
	Month
	Week
	Day
)

func (p *Period) ToDuration() time.Duration {
	switch *p {
	case Year:
		return 365 * 24 * time.Hour
	case Month:
		return 30 * 24 * time.Hour
	case Week:
		return 7 * 24 * time.Hour
	case Day:
		return 24 * time.Hour
	default:
		return 0
	}
}
