package reader

import (
	"fmt"
	"regexp"
	
)

func ReadOnlyNumbers(str string) string {
	re := regexp.MustCompile(`^(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)$`)
	if re.FindString(str) == "" {

	}
}
