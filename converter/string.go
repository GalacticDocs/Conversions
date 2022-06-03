package converter

import (
	"strconv"
)

func IntToString(integer int) string {
	val := strconv.Itoa(integer)

	return val
}