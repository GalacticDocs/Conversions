package converter

import (
	"strconv"

	"github.com/iVitaliya/logger-go"
)

func StringToFloat64(value string) float64 {
	var float, err = strconv.ParseFloat(value, 64)
	if err != nil {
		logger.Fatal("Couldn't translate given string value to a float64")
		return 0.0
	}

	return float
}
