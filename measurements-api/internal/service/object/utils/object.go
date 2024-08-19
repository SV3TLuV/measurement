package utils

import "measurements-api/internal/model"

func SortObjectByAscID(a, b *model.Object) int {
	if a.ID < b.ID {
		return -1
	} else if a.ID > b.ID {
		return 1
	}
	return 0
}
