package converter

import (
	"strconv"

	"github.com/iVitaliya/logger-go"
)

func StringToInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		logger.Error(err.Error())
	}

	return val
}
