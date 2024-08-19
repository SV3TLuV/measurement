package asoiza

func WindDirToRmb16(direction int) string {
	dir := float64(direction * 10)

	switch {
	case dir <= 11.25 || dir > 348.75:
		return "С"
	case dir <= 33.75 && dir > 11.25:
		return "ССВ"
	case dir <= 56.25 && dir > 33.75:
		return "СВ"
	case dir <= 78.75 && dir > 56.25:
		return "ВСВ"
	case dir <= 101.25 && dir > 78.75:
		return "В"
	case dir <= 123.75 && dir > 101.25:
		return "ВЮВ"
	case dir <= 146.25 && dir > 123.75:
		return "ЮВ"
	case dir <= 168.75 && dir > 146.25:
		return "ЮЮВ"
	case dir <= 191.25 && dir > 168.75:
		return "Ю"
	case dir <= 213.75 && dir > 191.25:
		return "ЮЮЗ"
	case dir <= 236.25 && dir > 213.75:
		return "ЮЗ"
	case dir <= 258.75 && dir > 236.25:
		return "ЗЮЗ"
	case dir <= 281.25 && dir > 258.75:
		return "З"
	case dir <= 303.75 && dir > 281.25:
		return "ЗСЗ"
	case dir <= 326.25 && dir > 303.75:
		return "СЗ"
	case dir <= 348.75 && dir > 326.25:
		return "ССЗ"
	default:
		return ""
	}
}
