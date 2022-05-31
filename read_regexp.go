package conversions

import (
	"fmt"
	"regexp"

	CLRS "github.com/iVitaliya/colors-go"
)

func logErr(exception string, msg string) {
	fmt.Println(
		
	)
}

func ReadOnlyNumbers(str string) string {
	re := regexp.MustCompile(`^(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)$`)
	if re.FindString(str) == "" {
		
	}
}