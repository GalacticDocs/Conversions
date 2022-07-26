package reader

import "strconv"

func ReadSingularDecimal(value float64) string {
	var float string = strconv.FormatFloat(value, 'f', 1, 64)

	return float
}
