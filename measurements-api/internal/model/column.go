package model

import (
	"strings"
	"unicode"
)

type Column struct {
	ID         uint64
	Title      string
	ShortTitle string
	Formula    *string
	ObjField   string
	Code       *string
}

func (c *Column) GetFormattedObjectField() string {
	var result strings.Builder
	parts := strings.Split(c.ObjField, "_")
	for _, part := range parts {
		if len(part) > 0 {
			result.WriteRune(unicode.ToUpper(rune(part[0])))
			if len(part) > 1 {
				result.WriteString(part[1:])
			}
		}
	}
	return result.String()
}
